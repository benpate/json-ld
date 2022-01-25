package reader

type String string

// AsObject returns this document as an Object
func (s String) AsObject(property string) JSONLD {
	return Map{
		property: s,
	}
}

// AsString returns this document as a string
func (s String) AsString(property string) string {
	return string(s)
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
