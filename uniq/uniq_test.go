package uniq

import (
	"testing"
)

func TestUniqBasic(t *testing.T) {
	// uniq работает с последовательными дубликатами
	input := []string{"foo", "bar", "foo", "baz"}
	opts := Options{}
	result := Uniq(input, opts)
	// Нет последовательных дубликатов, все строки выводятся
	expected := []string{"foo", "bar", "foo", "baz"}
	if len(result) != 4 {
		t.Errorf("Expected 4 lines, got %d", len(result))
	}
	for i, line := range result {
		if line != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], line)
		}
	}
}

func TestUniqCount(t *testing.T) {
	input := []string{"foo", "foo", "bar"}
	opts := Options{Count: true}
	result := Uniq(input, opts)
	expected := []string{"      2 foo", "      1 bar"}
	if len(result) != 2 {
		t.Errorf("Expected 2 lines, got %d", len(result))
	}
	for i, line := range result {
		if line != expected[i] {
			t.Errorf("Expected '%s', got '%s'", expected[i], line)
		}
	}
}

func TestUniqDuplicatesOnly(t *testing.T) {
	// Последовательные дубликаты
	input := []string{"a", "a", "b", "c", "c", "c"}
	opts := Options{DuplicatesOnly: true}
	result := Uniq(input, opts)
	expected := []string{"a", "c"}
	if len(result) != 2 {
		t.Errorf("Expected 2 lines, got %d: %v", len(result), result)
	}
	for i, line := range result {
		if line != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], line)
		}
	}
}

func TestUniqUniqueOnly(t *testing.T) {
	// Только уникальные (не повторяющиеся последовательно)
	input := []string{"a", "a", "b", "c", "c", "c"}
	opts := Options{UniqueOnly: true}
	result := Uniq(input, opts)
	expected := []string{"b"}
	if len(result) != 1 {
		t.Errorf("Expected 1 line, got %d: %v", len(result), result)
	}
	if len(result) > 0 && result[0] != expected[0] {
		t.Errorf("Expected %s, got %s", expected[0], result[0])
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
