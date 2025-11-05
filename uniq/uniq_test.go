package uniq

import (
	"testing"
)

func TestUniqBasic(t *testing.T) {
	input := []string{"foo", "bar", "foo", "baz"}
	opts := Options{}
	result := Uniq(input, opts)

	expected := map[string]bool{"foo": true, "bar": true, "baz": true}
	if len(result) != 3 {
		t.Errorf("Expected 3 unique lines, got %d", len(result))
	}

	for _, line := range result {
		if !expected[line] {
			t.Errorf("Unexpected line: %s", line)
		}
	}
}

func TestUniqCount(t *testing.T) {
	input := []string{"foo", "foo", "bar"}
	opts := Options{Count: true}
	result := Uniq(input, opts)

	expectedMap := map[string]bool{"2 foo": true, "1 bar": true}
	if len(result) != 2 {
		t.Errorf("Expected 2 lines, got %d", len(result))
	}

	for _, line := range result {
		if !expectedMap[line] {
			t.Errorf("Unexpected line: %s", line)
		}
	}
}

func TestUniqDuplicatesOnly(t *testing.T) {
	input := []string{"a", "b", "a", "c"}
	opts := Options{DuplicatesOnly: true}
	result := Uniq(input, opts)

	if len(result) != 1 || result[0] != "a" {
		t.Errorf("Expected only 'a', got %v", result)
	}
}

func TestUniqUniqueOnly(t *testing.T) {
	input := []string{"a", "b", "a", "c"}
	opts := Options{UniqueOnly: true}
	result := Uniq(input, opts)

	expected := map[string]bool{"b": true, "c": true}
	if len(result) != 2 {
		t.Errorf("Expected 2 lines, got %d", len(result))
	}

	for _, line := range result {
		if !expected[line] {
			t.Errorf("Unexpected line: %s", line)
		}
	}
}

func TestUniqIgnoreCase(t *testing.T) {
	input := []string{"Hello", "hello", "HELLO", "world"}
	opts := Options{IgnoreCase: true}
	result := Uniq(input, opts)

	if len(result) != 2 {
		t.Errorf("Expected 2 unique lines (case-insensitive), got %d", len(result))
	}
}

func TestUniqSkipFields(t *testing.T) {
	input := []string{"a b c", "x b c", "a b d"}
	opts := Options{SkipFields: 1}
	result := Uniq(input, opts)

	if len(result) != 2 {
		t.Errorf("Expected 2 unique lines, got %d", len(result))
	}
}

func TestUniqSkipChars(t *testing.T) {
	input := []string{"abcdef", "xycdef", "abcxyz"}
	opts := Options{SkipChars: 2}
	result := Uniq(input, opts)

	if len(result) != 2 {
		t.Errorf("Expected 2 unique lines, got %d", len(result))
	}
}
