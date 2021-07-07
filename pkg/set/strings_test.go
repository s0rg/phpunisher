package set

import "testing"

const needle = "needle"

func TestAdd(t *testing.T) {
	s := make(Strings)

	if s.Has(needle) {
		t.Fatal("found in empty")
	}

	s.Add("one")

	if s.Has(needle) {
		t.Fatal("found non-existent")
	}

	s.Add(needle)

	if !s.Has(needle) {
		t.Fatal("not found")
	}
}

func TestFromList(t *testing.T) {
	const needle = "needle"

	s := make(Strings)

	s.FromList([]string{"one", "two"})

	if s.Has(needle) {
		t.Fatal("found non-existent")
	}

	s.FromList([]string{"three", needle})

	if !s.Has(needle) {
		t.Fatal("not found")
	}
}
