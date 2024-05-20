package exception

import "strings"

// ValidationError represents a custom error type that contains a list of validation errors.
type ValidationError struct {
	Errs []string
}

// NewValidationError creates a new instance of ValidationError.
// This function is used to initialize a ValidationError with a given list of error messages.
//
// Parameters:
//   - errs: A slice of strings representing the validation error messages.
//
// Returns:
//   - *ValidationError: A pointer to the newly created ValidationError instance.
func NewValidationError(errs []string) *ValidationError {
	return &ValidationError{errs}
}

// Error implements the error interface for ValidationError.
// This method returns a string representation of the validation errors,
// concatenating all error messages in the Errs slice with a comma separator.
//
// Returns:
//   - string: A single string containing all validation error messages, separated by commas.
func (v *ValidationError) Error() string {
	return strings.Join(v.Errs, ",")
}
