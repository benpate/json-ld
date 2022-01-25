package reader

import "time"

type JSONLD interface {

	// AsObject returns this document as a JSON-LD Object.  The "property" argument designates
	// the property to populate if the document is not already an object
	AsObject(property string) JSONLD

	// AsString returns this document as a string.  The "property" argument designates
	// the property to use as a string value if the document is not already a string
	AsString(property string) string

	// AsInt returns this document as a time.Time.  The "property" argument designated
	// the property to use as a string value if the document is not already a string
	AsTime(property string) time.Time

	// AsInt returns this document as an int.  The "property" argument designated
	// the property to use as a string value if the document is not already a string
	AsInt(property string) int

	// AsSliceOfString returns this document as a slice of strings.  The "property" argument designates
	// the property to use as a string value if the document is not already a string
	AsSliceOfString(property string) []string

	// AsSliceOfObject returns this document as a slice of JSON-LD Objects.  The "property" argument designates
	// the property to populate if the document is not already an object
	AsSliceOfObject(property string) []JSONLD

	// Property returns a property of this document as its own JSON-LD document
	Property(property string) JSONLD

	// PropertyMap retrieves a string value from this document, trying first to use the
	// multi-lingual <propertyName>Map collection first.  If the requested language is
	// not found, then the default language (or the naked property string) value is returned.
	PropertyMap(property string, languages ...string) string

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
