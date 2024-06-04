package formAttributes

import "time"

type FieldAttribute struct {
	Name   string
	Code   string
	Errors []interface{}
}

type FieldAttributeInterface interface {
	AddError(message interface{})
	GetCode() string
	GetErrors() []interface{}
	Time() *time.Time
	IsClean() bool

	// Validators
	ValidateRequired()
	ValidateMin(min interface{})
	ValidateMax(max interface{})
	ValidateFormat(formatter string, formatterRemind string)
}

// ValidationMessage returns a formatted validation message.
func ValidationMessage(column string, message interface{}) interface{} {
	return message
}
