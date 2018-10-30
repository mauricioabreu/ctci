package challenge1

import "testing"

func TestIsUnique(t *testing.T) {
	result := isUnique("foo")
	if result != false {
		t.Error("foo has duplicated characters")
	}
	result = isUnique("bar")
	if result != true {
		t.Error("bar has no duplicated characters")
	}
}

func TestIsUniqueWithSort(t *testing.T) {
	result := isUniqueWithSort("foo")
	if result != false {
		t.Error("foo has duplicated characters")
	}
	result = isUniqueWithSort("bar")
	if result != true {
		t.Error("bar has no duplicated characters")
	}
}
