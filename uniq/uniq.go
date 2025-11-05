package uniq

import (
	"strings"
	"fmt"
)

// Options содержит параметры для обработки строк
type Options struct {
	Count          bool // -c: показывать количество
	DuplicatesOnly bool // -d: только дубликаты
	UniqueOnly     bool // -u: только уникальные
	IgnoreCase     bool // -i: игнорировать регистр
	SkipFields     int  // -f: пропустить N полей
	SkipChars      int  // -s: пропустить N символов
}

// preprocess обрабатывает строку согласно опциям
func preprocess(s string, opts Options) string {
	// Пропускаем первые N полей (слов)
	fields := strings.Fields(s)
	if opts.SkipFields > 0 {
		if len(fields) > opts.SkipFields {
			fields = fields[opts.SkipFields:]
		} else {
			fields = []string{}
		}
	}
	s = strings.Join(fields, " ")

	// Пропускаем первые N символов
	if opts.SkipChars > 0 {
		if len(s) > opts.SkipChars {
			s = s[opts.SkipChars:]
		} else {
			s = ""
		}
	}

	// Игнорируем регистр
	if opts.IgnoreCase {
		s = strings.ToLower(s)
	}

	return s
}

// Uniq обрабатывает слайс строк и возвращает результат
func Uniq(lines []string, opts Options) []string {
	if len(lines) == 0 {
		return []string{}
	}

	counts := make(map[string]int)
	originals := make(map[string]string)
	order := []string{} // для сохранения порядка

	for _, line := range lines {
		key := preprocess(line, opts)
		
		// Сохраняем первое вхождение оригинальной строки
		if _, exists := originals[key]; !exists {
			originals[key] = line
			order = append(order, key)
		}
		counts[key]++
	}

	result := []string{}
	for _, key := range order {
		count := counts[key]
		
		// Фильтруем по типу вывода
		if opts.DuplicatesOnly && count < 2 {
			continue
		}
		if opts.UniqueOnly && count > 1 {
			continue
		}

		line := originals[key]
		if opts.Count {
			line = fmt.Sprintf("%d %s", count, line)
		}
		result = append(result, line)
	}

	return result
}
