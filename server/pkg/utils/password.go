package utils

import (
    "crypto/rand"
    "encoding/base64"
    "fmt"

    "golang.org/x/crypto/bcrypt"
)

const (
    // DefaultCost is the default cost for password hashing
    DefaultCost = bcrypt.DefaultCost
    // MinCost is the minimum cost for password hashing
    MinCost = bcrypt.MinCost
    // MaxCost is the maximum cost for password hashing
    MaxCost = bcrypt.MaxCost
)

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
    if len(password) == 0 {
        return "", fmt.Errorf("password cannot be empty")
    }
    
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
    if err != nil {
        return "", fmt.Errorf("failed to hash password: %w", err)
    }
    
    return string(bytes), nil
}

// CheckPasswordHash compares a password with its hash
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// GenerateRandomPassword generates a random password of specified length
func GenerateRandomPassword(length int) (string, error) {
    if length < 8 {
        return "", fmt.Errorf("password length must be at least 8 characters")
    }
    
    bytes := make([]byte, length)
    if _, err := rand.Read(bytes); err != nil {
        return "", fmt.Errorf("failed to generate random password: %w", err)
    }
    
    return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

// ValidatePasswordStrength checks if password meets strength requirements
func ValidatePasswordStrength(password string) error {
    if len(password) < 8 {
        return fmt.Errorf("password must be at least 8 characters long")
    }
    
    hasUpper := false
    hasLower := false
    hasNumber := false
    hasSpecial := false
    
    specialChars := "!@#$%^&*()_+-=[]{}|;:,.<>?"
    
    for _, char := range password {
        switch {
        case 'A' <= char && char <= 'Z':
            hasUpper = true
        case 'a' <= char && char <= 'z':
            hasLower = true
        case '0' <= char && char <= '9':
            hasNumber = true
        default:
            for _, special := range specialChars {
                if char == special {
                    hasSpecial = true
                    break
                }
            }
        }
    }
    
    if !hasUpper {
        return fmt.Errorf("password must contain at least one uppercase letter")
    }
    if !hasLower {
        return fmt.Errorf("password must contain at least one lowercase letter")
    }
    if !hasNumber {
        return fmt.Errorf("password must contain at least one number")
    }
    if !hasSpecial {
        return fmt.Errorf("password must contain at least one special character")
    }
    
    return nil
}