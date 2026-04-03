package validation

func SetIfNotZero(m map[string]interface{}, key string, value any) {
	if IsNotDefault(value) {
		m[key] = value
	}
}
