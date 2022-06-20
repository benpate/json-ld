package reader

import (
	"time"

	"github.com/benpate/rosetta/convert"
)

type Map map[string]interface{}

// Property returns a property of this document as a JSON-LD object
func (m Map) Property(property string) JSONLD {

	// Try to find, then wrap the value as a new JSON-LD
	if value, ok := m[property]; ok {
		return New(value)
	}

	// If not found, return empty value
	return New("")
}

// PropertyMap returns a property of this document
func (m Map) PropertyMap(property string, languages ...string) string {

	// Add "English" as default, catch-all language
	languages = append(languages, "en")
	propertyMap := property + "Map"

	// If mapped values exist (and are the correct type)
	if mappedValues, ok := m[propertyMap]; ok {
		if mappedValues, ok := mappedValues.(map[string]interface{}); ok {

			// Search for values matching the requested langauge (in order)
			for _, language := range languages {
				if value, ok := mappedValues[language]; ok {
					// Success!
					return convert.String(value)
				}
			}
		}
	}

	// Fall through means that we don't have a mapped value.
	// So just return a string value instead.
	return convert.String(m[property])
}

// AsObject returns this document as an object
func (m Map) AsObject(property string) JSONLD {
	return m
}

// AsString returns this document as a string
func (m Map) AsString(property string) string {
	return convert.String(m[property])
}

// AsInt returns this document as an int
func (m Map) AsInt(property string) int {
	return convert.Int(m[property])
}

// AsTime returns this document as an int
func (m Map) AsTime(property string) time.Time {
	result, _ := time.Parse(time.RFC3339, convert.String(m[property]))
	return result
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
	return New("")
}
