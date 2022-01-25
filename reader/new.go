package reader

func New(value interface{}) JSONLD {

	switch value := value.(type) {

	case JSONLD:
		return value

	case string:

		if value != "" {
			return String(value)
		}

	case map[string]interface{}:

		if len(value) > 0 {
			return Map(value)
		}

	case []interface{}:

		if len(value) > 0 {
			return Slice(value)
		}
	}

	// All other cases return zero value
	return Zero{}
}
