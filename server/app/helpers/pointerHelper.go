package helpers

func GetStringOrDefault(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
