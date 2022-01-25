package reader

type JSONLD interface {

	// AsObject returns this document as a JSON-LD Object.  The "property" argument designates
	// the property to populate if the document is not already an object
	AsObject(property string) JSONLD

	// AsString returns this document as a string.  The "property" argument designates
	// the property to use as a string value if the document is not already a string
	AsString(property string) string

	// AsSliceOfString returns this document as a slice of strings.  The "property" argument designates
	// the property to use as a string value if the document is not already a string
	AsSliceOfString(property string) []string

	// AsSliceOfObject returns this document as a slice of JSON-LD Objects.  The "property" argument designates
	// the property to populate if the document is not already an object
	AsSliceOfObject(property string) []JSONLD

	// Property returns a property of this document as its own JSON-LD document
	Property(property string) JSONLD

	// Len returns the number of documents in a series.  If this is not a series,
	// then Len() returns 1.
	Len() int

	// Head returns the first JSONLD document in a series.  If this is not a series,
	// then Head() returns this document
	Head() JSONLD

	// Tail returns a series contining all but the first document in a series.  If this
	// is not a series, then Tail() returns nil
	Tail() JSONLD
}
