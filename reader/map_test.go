package reader

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func unmarshalMap(value string) (JSONLD, error) {

	result := Map{}

	if err := json.Unmarshal([]byte(value), &result); err != nil {
		return nil, err
	}

	return result, nil
}
func TestID(t *testing.T) {

	object, err := unmarshalMap(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"name": "Foo",
		"id": "http://example.org/foo",
		"links": [{
			"type":"Link",
			"href":"https://emissary.dev"
		}, {
			"type":"Link",
			"href":"https://test.com"
		}]
	  }`)

	require.Nil(t, err)
	require.Equal(t, "Foo", object.Property("name").AsString(""))
	require.Equal(t, "http://example.org/foo", object.Property("id").AsString(""))
	require.Equal(t, []string{"Foo"}, object.AsSliceOfString("name"))

	require.Equal(t, 2, object.Property("links").Len())
	require.Equal(t, "Link", object.Property("links").Head().Property("type").AsString(""))
	require.Equal(t, 1, object.Property("links").Tail().Len())
	require.Equal(t, Map{"type": "Link", "href": "https://emissary.dev"}, object.Property("links").Head().AsObject("href"))
	require.Equal(t, []string{"https://emissary.dev", "https://test.com"}, object.Property("links").AsSliceOfString("href"))
	require.Equal(t, Map{"type": "Link", "href": "https://test.com"}, object.Property("links").AsSliceOfObject("href")[1])
}

/*
func TestType1(t *testing.T) {

	object, err := unmarshal(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"summary": "A foo",
		"type": "http://example.org/Foo"
	  }`)

	require.Nil(t, err)
	require.Equal(t, "http://example.org/Foo", object.Type())
	require.Equal(t, []string{"http://example.org/Foo"}, object.TypeArray())
}

func TestType2(t *testing.T) {

	object, err := unmarshal(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"summary": "A foo",
		"type": ["first", "second", "third"]
	  }`)

	require.Nil(t, err)
	require.Equal(t, "first", object.Type())
	require.Equal(t, []string{"first", "second", "third"}, object.TypeArray())
}

func TestActor1(t *testing.T) {

	object, err := unmarshal(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"summary": "Sally offered the Foo object",
		"type": "Offer",
		"actor": "http://sally.example.org",
		"object": "http://example.org/foo"
	  }`)

	require.Nil(t, err)
	require.Equal(t, object.Actor(), "http://sally.example.org")

	t.Log(spew.Sdump(object))

	actor, ok := object.ActorObject()

	if !ok {
		t.Fail()
		return
	}

	require.Equal(t, actor.Type(), "")
	require.Equal(t, actor.ID(), "http://sally.example.org")
	require.Equal(t, actor.Summary(), "")
	require.Equal(t, actor.Summary("es"), "")
}

func TestActor2(t *testing.T) {

	object, err := unmarshal(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"summary": "Sally offered the Foo object",
		"type": "Offer",
		"actor": {
		  "type": "Person",
		  "id": "http://sally.example.org",
		  "summary": "Sally"
		},
		"object": "http://example.org/foo"
	  }`)

	require.Nil(t, err)
	require.Equal(t, object.Actor(), "http://sally.example.org")

	t.Log(spew.Sdump(object))

	actor, ok := object.ActorObject()

	if !ok {
		t.Fail()
		return
	}

	require.Equal(t, actor.Type(), "Person")
	require.Equal(t, actor.ID(), "http://sally.example.org")
	require.Equal(t, actor.Summary(), "Sally")
	require.Equal(t, actor.Summary("es"), "Sally")
}

func TestAttachment(t *testing.T) {

	object, err := unmarshal(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"type": "Note",
		"name": "Have you seen my cat?",
		"attachment": [
		  {
			"type": "Image",
			"content": "This is what he looks like.",
			"url": "http://example.org/cat.jpeg"
		  }
		]
	  }`)

	attachment, ok := object.AttachmentObject()
	t.Log(spew.Sdump(object))
	t.Log(spew.Sdump(attachment))

	require.Nil(t, err)
	require.Equal(t, "Note", object.Type())
	require.Equal(t, "Have you seen my cat?", object.Name())
	require.Equal(t, "http://example.org/cat.jpeg", object.Attachment())

	require.True(t, ok)
	require.Equal(t, "Image", attachment.Type())
	require.Equal(t, "This is what he looks like.", attachment.Content())
	require.Equal(t, "http://example.org/cat.jpeg", attachment.URL())
}

func TestName1(t *testing.T) {

	object, err := unmarshal(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"type": "Note",
		"name": "A simple note"
		}`)

	require.Nil(t, err)
	require.Equal(t, "Note", object.Type())
	require.Equal(t, "A simple note", object.Name())
	require.Equal(t, "A simple note", object.Name("en"))
	require.Equal(t, "A simple note", object.Name("es"))
}

func TestName2(t *testing.T) {

	object, err := unmarshal(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"type": "Note",
		"nameMap": {
			"en": "A simple note",
			"es": "Una nota sencilla",
			"zh-Hans": "一段简单的笔记"
		}
		}`)

	require.Nil(t, err)
	require.Equal(t, object.Type(), "Note")
	require.Equal(t, object.Name(), "A simple note")
	require.Equal(t, object.Name("en"), "A simple note")
	require.Equal(t, object.Name("es"), "Una nota sencilla")
	require.Equal(t, object.Name("zh-Hans"), "一段简单的笔记")

}
*/
/*
func TestDuration1(t *testing.T) {

	object, err := unmarshal(`{
		"@context": "https://www.w3.org/ns/activitystreams",
		"type": "Video",
		"name": "Birds Flying",
		"url": "http://example.org/video.mkv",
		"duration": "PT2H"
	  }`)

	require.Nil(t, err)
	require.Equal(t, object.Type(), "Video")
	require.Equal(t, object.Duration(), time.Duration(5*time.Second))
}
*/
