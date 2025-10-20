package openguardrails

import "fmt"

// OpenGuardrailsError OpenGuardrails base error class
type OpenGuardrailsError struct {
	Message string
	Cause   error
}

func (e *OpenGuardrailsError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func (e *OpenGuardrailsError) Unwrap() error {
	return e.Cause
}

// NewOpenGuardrailsError Create new OpenGuardrails error
func NewOpenGuardrailsError(message string, cause error) *OpenGuardrailsError {
	return &OpenGuardrailsError{
		Message: message,
		Cause:   cause,
	}
}

// AuthenticationError Authentication error
type AuthenticationError struct {
	*OpenGuardrailsError
}

// NewAuthenticationError Create authentication error
func NewAuthenticationError(message string) *AuthenticationError {
	return &AuthenticationError{
		OpenGuardrailsError: &OpenGuardrailsError{Message: message},
	}
}

// RateLimitError Rate limit error
type RateLimitError struct {
	*OpenGuardrailsError
}

// NewRateLimitError Create rate limit error
func NewRateLimitError(message string) *RateLimitError {
	return &RateLimitError{
		OpenGuardrailsError: &OpenGuardrailsError{Message: message},
	}
}

// ValidationError Input validation error
type ValidationError struct {
	*OpenGuardrailsError
}

// NewValidationError Create validation error
func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		OpenGuardrailsError: &OpenGuardrailsError{Message: message},
	}
}

// NetworkError Network error
type NetworkError struct {
	*OpenGuardrailsError
}

// NewNetworkError Create network error
func NewNetworkError(message string, cause error) *NetworkError {
	return &NetworkError{
		OpenGuardrailsError: &OpenGuardrailsError{Message: message, Cause: cause},
	}
}

// ServerError Server error
type ServerError struct {
	*OpenGuardrailsError
}

// NewServerError Create server error
func NewServerError(message string) *ServerError {
	return &ServerError{
		OpenGuardrailsError: &OpenGuardrailsError{Message: message},
	}
}