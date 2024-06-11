package helpers

// GetStringOrDefault returns the value of the string pointer if not nil, otherwise returns "".
func GetStringOrDefault(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

// GetInt32OrDefault returns the value of the int32 pointer if not nil, otherwise returns 0.
func GetInt32OrDefault(num *int32) int32 {
	if num == nil {
		return 0
	}
	return *num
}
