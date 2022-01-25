package reader

import (
	"time"
)

// Zero reader has no value
type Zero struct{}

// AsObject returns this document as an Object
func (z Zero) AsObject(property string) JSONLD {
	return z
}

// AsString returns this document as a string
func (z Zero) AsString(property string) string {
	return ""
}

// AsInt returns this document as an int
func (z Zero) AsInt(property string) int {
	return 0
}

// AsTime returns this document as an int
func (z Zero) AsTime(property string) time.Time {
	var result time.Time
	return result
}

// AsSliceOfString returns this document as a slice of strings
func (z Zero) AsSliceOfString(property string) []string {
	result := make([]string, 0)
	return result
}

// AsSliceOfObject returns this document as a slice of Objects
func (z Zero) AsSliceOfObject(property string) []JSONLD {
	result := make([]JSONLD, 0)
	return result
}

// Property returns a property of this document
func (z Zero) Property(property string) JSONLD {
	return z
}

// PropertyMap returns a property of this document
func (z Zero) PropertyMap(property string, languages ...string) string {
	return ""
}

// Len returns the length of a series
func (z Zero) Len() int {
	return 0
}

// Head returns the first JSONLD document in a series
func (z Zero) Head() JSONLD {
	return z
}

// Tail returns the series without the head document
func (z Zero) Tail() JSONLD {
	return z
}
