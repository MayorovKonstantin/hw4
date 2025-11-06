package uniq

import (
	"fmt"
	"strings"
)

type Options struct {
	Count          bool
	DuplicatesOnly bool
	UniqueOnly     bool
	IgnoreCase     bool
	SkipFields     int
	SkipChars      int
}

func Uniq(lines []string, opts Options) []string {
	if len(lines) == 0 {
		return []string{}
	}

	result := []string{}
	var prevKey string
	var prevLine string
	count := 1

	for i, line := range lines {
		key := getKey(line, opts)

		if i == 0 {
			prevKey = key
			prevLine = line
			continue
		}

		if key == prevKey {
			count++
		} else {
			// Output previous group
			if shouldOutput(count, opts) {
				if opts.Count {
					result = append(result, formatWithCount(count, prevLine))
				} else {
					result = append(result, prevLine)
				}
			}
			prevKey = key
			prevLine = line
			count = 1
		}
	}

	// Output last group
	if shouldOutput(count, opts) {
		if opts.Count {
			result = append(result, formatWithCount(count, prevLine))
		} else {
			result = append(result, prevLine)
		}
	}

	return result
}

func getKey(line string, opts Options) string {
	key := line

	// Skip fields
	if opts.SkipFields > 0 {
		fields := strings.Fields(key)
		skip := opts.SkipFields
		if skip < len(fields) {
			key = strings.Join(fields[skip:], " ")
		} else {
			key = ""
		}
	}

	// Skip characters
	if opts.SkipChars > 0 {
		skip := opts.SkipChars
		if skip < len(key) {
			key = key[skip:]
		} else {
			key = ""
		}
	}

	// Ignore case
	if opts.IgnoreCase {
		key = strings.ToLower(key)
	}

	return key
}

func shouldOutput(count int, opts Options) bool {
	// -d: only duplicates (count > 1)
	if opts.DuplicatesOnly {
		return count > 1
	}

	// -u: only unique (count == 1)
	if opts.UniqueOnly {
		return count == 1
	}

	return true
}

func formatWithCount(count int, line string) string {
	return fmt.Sprintf("%7d %s", count, line)
}
