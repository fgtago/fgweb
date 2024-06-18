package appsmodel

import "net/url"

type FormError struct {
	FieldName string
	Message   string
}

type formerrors map[string][]FormError

type Form struct {
	url.Values
	Errors formerrors
}

// New creates a new Form with the provided values.
//
// values: The form values to initialize the Form.
// Returns a pointer to the newly created Form.
func New(values url.Values) *Form {
	return &Form{
		Values: values,
		Errors: make(formerrors),
	}
}

// Has checks if a form field exists and returns a boolean value indicating whether the field is present or not.
//
// Parameters:
// - field: a string representing the name of the form field to check.
//
// Returns:
// - bool: true if the field is present in the form, false otherwise.
func (f *Form) Has(field string) bool {
	x := f.Get(field)
	return x != ""
}

// Add adds a new FormError to the errors map.
//
// Parameters:
// - fieldName: the name of the field associated with the error.
// - message: the error message.
//
// Return type: none.
func (e formerrors) Add(fieldName string, message string) {
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
func (e formerrors) Get(fieldName string) string {
	errorMessages := e[fieldName]
	if len(errorMessages) == 0 {
		return ""
	}
	return errorMessages[0].Message
}
