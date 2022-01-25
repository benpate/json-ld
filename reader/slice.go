package reader

type Slice []interface{}

// AsObject returns this document as an Object
func (s Slice) AsObject(property string) JSONLD {
	return s.Head().AsObject(property)
}

// AsString returns this document as a string
func (s Slice) AsString(property string) string {
	return s.Head().AsString(property)
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

// Len returns the length of a series
func (s Slice) Len() int {
	return len(s)
}

// Head returns the first JSONLD document in a series
func (s Slice) Head() JSONLD {
	return New(s[0])
}

// Tail returns the series without the head document
func (s Slice) Tail() JSONLD {
	return Slice(s[1:])
}
