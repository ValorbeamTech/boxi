#!/bin/bash

# Database backup script
set -e

# Configuration
BACKUP_DIR="/backups"
DATE=$(date +%Y%m%d_%H%M%S)
DB_NAME="${DB_NAME:-myapp}"
DB_USER="${DB_USER:-postgres}"
RETENTION_DAYS=7

# Create backup directory
mkdir -p "$BACKUP_DIR"

# Create backup
echo "Creating backup for database: $DB_NAME"
pg_dump -h db -U "$DB_USER" -d "$DB_NAME" | gzip > "$BACKUP_DIR/backup_${DB_NAME}_${DATE}.sql.gz"

# Verify backup
if [ -f "$BACKUP_DIR/backup_${DB_NAME}_${DATE}.sql.gz" ]; then
    echo "Backup created successfully: backup_${DB_NAME}_${DATE}.sql.gz"
    
    # Check backup size
    SIZE=$(stat -f%z "$BACKUP_DIR/backup_${DB_NAME}_${DATE}.sql.gz" 2>/dev/null || stat -c%s "$BACKUP_DIR/backup_${DB_NAME}_${DATE}.sql.gz")
    echo "Backup size: $(($SIZE / 1024 / 1024)) MB"
else
    echo "Backup failed!"
    exit 1
fi

# Clean up old backups
echo "Cleaning up backups older than $RETENTION_DAYS days"
find "$BACKUP_DIR" -name "backup_${DB_NAME}_*.sql.gz" -mtime +$RETENTION_DAYS -delete

echo "Backup completed successfully"