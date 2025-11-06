package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/MayorovKonstantin/hw4/uniq"
)

// parseFlags определяет и парсит флаги командной строки
func parseFlags() uniq.Options {
	countFlag := flag.Bool("c", false, "Count occurrences")
	duplicatesFlag := flag.Bool("d", false, "Show only duplicates")
	uniqueFlag := flag.Bool("u", false, "Show only unique lines")
	ignoreCaseFlag := flag.Bool("i", false, "Ignore case")
	skipFieldsFlag := flag.Int("f", 0, "Skip first N fields")
	skipCharsFlag := flag.Int("s", 0, "Skip first N characters")

	flag.Parse()

	// Проверка взаимоисключающих флагов
	validateExclusiveFlags(*countFlag, *duplicatesFlag, *uniqueFlag)

	return uniq.Options{
		Count:          *countFlag,
		DuplicatesOnly: *duplicatesFlag,
		UniqueOnly:     *uniqueFlag,
		IgnoreCase:     *ignoreCaseFlag,
		SkipFields:     *skipFieldsFlag,
		SkipChars:      *skipCharsFlag,
	}
}

// validateExclusiveFlags проверяет, что флаги -c, -d, -u не используются одновременно
func validateExclusiveFlags(count, duplicates, unique bool) {
	exclusiveCount := 0
	if count {
		exclusiveCount++
	}
	if duplicates {
		exclusiveCount++
	}
	if unique {
		exclusiveCount++
	}

	if exclusiveCount > 1 {
		fmt.Fprintln(os.Stderr, "Error: flags -c, -d, -u are mutually exclusive")
		os.Exit(1)
	}
}

// openInputOutput открывает файлы ввода и вывода на основе аргументов
func openInputOutput() (*os.File, *os.File, func()) {
	input := os.Stdin
	output := os.Stdout
	var closeInput, closeOutput func()

	args := flag.Args()

	// Открываем входной файл, если указан
	if len(args) > 0 {
		f, err := os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
			os.Exit(1)
		}
		input = f
		closeInput = func() { f.Close() }
	}

	// Открываем выходной файл, если указан
	if len(args) > 1 {
		f, err := os.Create(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		output = f
		closeOutput = func() { f.Close() }
	}

	// Функция для закрытия всех открытых файлов
	cleanup := func() {
		if closeInput != nil {
			closeInput()
		}
		if closeOutput != nil {
			closeOutput()
		}
	}

	return input, output, cleanup
}

// readLines читает все строки из reader
func readLines(input *os.File) []string {
	scanner := bufio.NewScanner(input)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	return lines
}

// writeLines записывает строки в writer
func writeLines(output *os.File, lines []string) {
	for _, line := range lines {
		fmt.Fprintln(output, line)
	}
}

func main() {
	opts := parseFlags()
	input, output, cleanup := openInputOutput()
	defer cleanup()

	lines := readLines(input)
	result := uniq.Uniq(lines, opts)
	writeLines(output, result)
}
