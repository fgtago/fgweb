package form

type FormError struct {
	FieldName string
	Message   string
}

type errors map[string][]FormError

// Add adds a new FormError to the errors map.
//
// Parameters:
// - fieldName: the name of the field associated with the error.
// - message: the error message.
//
// Return type: none.
func (e errors) Add(fieldName string, message string) {
	e[fieldName] = append(e[fieldName], FormError{
		FieldName: fieldName,
		Message:   message,
	})
}

// Get gets the error message for the specified field name.
//
// Parameters:
// - fieldName: the name of the field associated with the error.
// Return type: string.
func (e errors) Get(fieldName string) string {
	errorMessages := e[fieldName]
	if len(errorMessages) == 0 {
		return ""
	}
	return errorMessages[0].Message
}
