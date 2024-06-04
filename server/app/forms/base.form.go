package forms

import (
	"fmt"
	"server/app/exceptions"
	"server/app/formAttributes"
	"slices"
)

type Form struct {
	Attributes []formAttributes.FieldAttributeInterface
	Errors     exceptions.ResourceModificationError
}

func (form *Form) IsValid() bool {
	return len(form.Errors) == 0
}

func (form *Form) AddAttributes(attributes ...formAttributes.FieldAttributeInterface) {
	form.Attributes = append(form.Attributes, attributes...)
}

func (form *Form) GetAttribute(code string) formAttributes.FieldAttributeInterface {
	idx := slices.IndexFunc(form.Attributes, func(a formAttributes.FieldAttributeInterface) bool { return a.GetCode() == code })

	if idx != -1 {
		return form.Attributes[idx]
	} else {
		return nil
	}
}

// SummaryErrors summarizes errors in the form.
func (form *Form) summaryErrors() {
	err := exceptions.NewUnprocessableContentError("", nil)

	if form.Errors != nil {
		err.Errors = form.Errors
	}

	for _, attribute := range form.Attributes {
		attributeErr := attribute.GetErrors()

		if len(attributeErr) > 0 {
			err.AddError(attribute.GetCode(), attributeErr)
		}
	}

	form.Errors = err.Errors
}

func (form *Form) AddNestedErrors(fieldKey string, index int, errors exceptions.ResourceModificationError) {
	for key, innerErr := range errors {
		form.AddErrorDirectlyToField(form.NestedFieldKey(fieldKey, index, key), innerErr)
	}
}

// NestedFieldKey output a key for response
// such as `projectIssueStatuses.1.issueStatusId`
// use for nested attributes
func (form *Form) NestedFieldKey(wrapperFieldKey string, index int, nestedFieldKey string) string {
	return fmt.Sprintf("%s.%d.%s", wrapperFieldKey, index, nestedFieldKey)
}

func (form *Form) NestedDirectItemFieldKey(wrapperFieldKey string, index int) string {
	return fmt.Sprintf("%s.%d", wrapperFieldKey, index)
}

func (form *Form) AddErrorDirectlyToField(field string, errors []interface{}) {
	if len(form.Errors) == 0 {
		form.Errors = exceptions.ResourceModificationError{}
	}

	if len(form.Errors[field]) == 0 {
		form.Errors[field] = []interface{}{}
	}

	form.Errors[field] = append(form.Errors[field], errors...)
}
