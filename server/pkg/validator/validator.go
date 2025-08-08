package validator

import (
    "errors"
    "fmt"
    "reflect"
    "strings"

    "github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
    validate = validator.New()
    
    // Register custom tag name function to use json tag names
    validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
        name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
        if name == "-" {
            return ""
        }
        return name
    })

    // Register custom validators
    validate.RegisterValidation("password", validatePassword)
}

func Validate(s interface{}) error {
    if err := validate.Struct(s); err != nil {
        var validationErrors []string
        
        for _, err := range err.(validator.ValidationErrors) {
            validationErrors = append(validationErrors, formatValidationError(err))
        }
        
        return errors.New(strings.Join(validationErrors, ", "))
    }
    return nil
}

func formatValidationError(err validator.FieldError) string {
    field := err.Field()
    
    switch err.Tag() {
    case "required":
        return fmt.Sprintf("%s is required", field)
    case "email":
        return fmt.Sprintf("%s must be a valid email address", field)
    case "min":
        return fmt.Sprintf("%s must be at least %s characters long", field, err.Param())
    case "max":
        return fmt.Sprintf("%s must be at most %s characters long", field, err.Param())
    case "password":
        return fmt.Sprintf("%s must contain at least 8 characters with uppercase, lowercase, number and special character", field)
    case "gte":
        return fmt.Sprintf("%s must be greater than or equal to %s", field, err.Param())
    case "lte":
        return fmt.Sprintf("%s must be less than or equal to %s", field, err.Param())
    case "oneof":
        return fmt.Sprintf("%s must be one of: %s", field, err.Param())
    default:
        return fmt.Sprintf("%s is invalid", field)
    }
}

// Custom password validator
func validatePassword(fl validator.FieldLevel) bool {
    password := fl.Field().String()
    if len(password) < 8 {
        return false
    }
    
    hasUpper := false
    hasLower := false
    hasNumber := false
    hasSpecial := false
    
    for _, char := range password {
        switch {
        case 'A' <= char && char <= 'Z':
            hasUpper = true
        case 'a' <= char && char <= 'z':
            hasLower = true
        case '0' <= char && char <= '9':
            hasNumber = true
        case strings.ContainsRune("!@#$%^&*()_+-=[]{}|;:,.<>?", char):
            hasSpecial = true
        }
    }
    
    return hasUpper && hasLower && hasNumber && hasSpecial
}

// ValidateVar validates a single variable
func ValidateVar(field interface{}, tag string) error {
    return validate.Var(field, tag)
}