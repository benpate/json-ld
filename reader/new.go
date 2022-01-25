package reader

func New(value interface{}) JSONLD {

	switch value := value.(type) {
	case string:
		return String(value)
	case map[string]interface{}:
		return Map(value)

	case []interface{}:
		return Slice(value)
	}

	return nil
}
