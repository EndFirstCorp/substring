package substring

import (
	"testing"
)

func TestAfter(t *testing.T) {
	actual := After("https://basecamp.com/2753312/api/v1/calendars.json", "/v1/")
	if "calendars.json" != actual {
		t.Error("expected didn't match actual: ", actual)
	}
}

func TestAfterNotFound(t *testing.T) {
	const url string = "https://basecamp.com/2753312/api/v1/calendars.json"
	actual := After(url, "/v2/")
	if actual != "" {
		t.Error("expected didn't match actual: ", actual)
	}
}

func TestBetween(t *testing.T) {
	if val := Between("hello, my friend: there", ", my", ":"); val != " friend" {
		t.Error("expected match", val)
	}

	if val := Between("hello, my friend: there", " my", "!"); val != "" {
		t.Error("expected empty on no match", val)
	}

	if val := Between("hello, my friend: there", "nomatch", "friend"); val != "" {
		t.Error("expected empty on no match", val)
	}
}
