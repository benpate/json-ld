package reader

import (
	"time"

	"github.com/benpate/convert"
)

type String string

// AsObject returns this document as an Object
func (s String) AsObject(property string) JSONLD {
	return Map{property: s}
}

// AsString returns this document as a string
func (s String) AsString(property string) string {
	return string(s)
}

// AsInt returns this document as an int
func (s String) AsInt(property string) int {
	return convert.Int(string(s))
}

// AsTime returns this document as an int
func (s String) AsTime(property string) time.Time {
	result, _ := time.Parse(time.RFC3339, string(s))
	return result
}

// AsSliceOfString returns this document as a slice of strings
func (s String) AsSliceOfString(property string) []string {
	return []string{s.AsString(property)}
}

// AsSliceOfObject returns this document as a slice of Objects
func (s String) AsSliceOfObject(property string) []JSONLD {
	return []JSONLD{s.AsObject(property)}
}

// Property returns a property of this document
func (s String) Property(property string) JSONLD {
	return s
}

// PropertyMap returns a property of this document
func (s String) PropertyMap(property string, languages ...string) string {
	return string(s)
}

// Len returns the length of a series
func (s String) Len() int {
	return 1
}

// Head returns the first JSONLD document in a series
func (s String) Head() JSONLD {
	return s
}

// Tail returns the series without the head document
func (s String) Tail() JSONLD {
	return nil
}
