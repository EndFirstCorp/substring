package helpers

import (
	"testing"
)

func TestSubstringAfter(t *testing.T) {
	actual := SubstringAfter("https://basecamp.com/2753312/api/v1/calendars.json", "/v1/")
	if "calendars.json" != actual {
		t.Fatalf("expected didn't match actual: ", actual)
	}
}

func TestSubstringAfterNotFound(t *testing.T) {
	const url string = "https://basecamp.com/2753312/api/v1/calendars.json"
	actual := SubstringAfter(url, "/v2/")
	if url != actual {
		t.Fatalf("expected didn't match actual: ", actual)
	}
}

func TestSubstringBetween(t *testing.T) {
	if SubstringBetween("hello, my friend: there", ", my", ":") != " friend" {
		t.Fatal("didn't work")
	}
}
