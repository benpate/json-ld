package reader

import (
	"time"
)

type Slice []interface{}

// AsObject returns this document as an Object
func (s Slice) AsObject(property string) JSONLD {
	return s.Head().AsObject(property)
}

// AsString returns this document as a string
func (s Slice) AsString(property string) string {
	return s.Head().AsString(property)
}

// AsInt returns this document as an int
func (s Slice) AsInt(property string) int {
	return s.Head().AsInt(property)
}

// AsTime returns this document as an int
func (s Slice) AsTime(property string) time.Time {
	return s.Head().AsTime(property)
}

// AsSliceOfString returns this document as a slice of strings
func (s Slice) AsSliceOfString(property string) []string {

	result := make([]string, len(s))

	for index, value := range s {
		result[index] = New(value).AsString(property)
	}

	return result
}

// AsSliceOfObject returns this document as a slice of Objects
func (s Slice) AsSliceOfObject(property string) []JSONLD {

	result := make([]JSONLD, len(s))

	for index, value := range s {
		result[index] = New(value).AsObject(property)
	}

	return result
}

// Property returns a property of this document
func (s Slice) Property(property string) JSONLD {
	return s.Head().Property(property)
}

// PropertyMap returns a property of this document
func (s Slice) PropertyMap(property string, languages ...string) string {
	return s.Head().PropertyMap(property, languages...)
}

// Len returns the length of a series
func (s Slice) Len() int {
	return len(s)
}

// Head returns the first JSONLD document in a series
func (s Slice) Head() JSONLD {

	// If there is at least one value in this slice, return the first one
	if len(s) > 0 {
		return New(s[0])
	}

	// Otherwise, return empty Object
	return New("")
}

// Tail returns the series without the head document
func (s Slice) Tail() JSONLD {

	// If there are items remaining in the tail, return them
	if result := (s[1:]); len(result) > 0 {
		return Slice(result)
	}

	// Otherwise, return "empty" Object
	return New("")
}
