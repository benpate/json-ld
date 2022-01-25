package reader

import "github.com/benpate/convert"

type Map map[string]interface{}

// Property returns a property of this document as a JSON-LD object
func (m Map) Property(property string) JSONLD {
	return New(m[property])
}

// AsObject returns this document as an object
func (m Map) AsObject(property string) JSONLD {
	return m
}

// AsString returns this document as a string
func (m Map) AsString(property string) string {
	return convert.String(m[property])
}

// AsSliceOfString returns this document as a slice of strings
func (m Map) AsSliceOfString(property string) []string {
	return []string{m.AsString(property)}
}

// AsSliceOfObject returns this document as a slice of Objects
func (m Map) AsSliceOfObject(property string) []JSONLD {
	return []JSONLD{m}
}

// GetStringMap returns a property of this JSONLD document from a multi-language mapped value
func (m Map) GetStringMap(property string, languages ...string) string {
	return ""
}

func (m Map) Len() int {
	return 1
}

// Head returns the first JSONLD document in a series
func (m Map) Head() JSONLD {
	return m
}

func (m Map) Tail() JSONLD {
	return nil
}
