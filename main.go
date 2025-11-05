package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/MayorovKonstantin/hw4/uniq"
)

func main() {
	// Определяем флаги
	countFlag := flag.Bool("c", false, "Count occurrences")
	duplicatesFlag := flag.Bool("d", false, "Show only duplicates")
	uniqueFlag := flag.Bool("u", false, "Show only unique lines")
	ignoreCaseFlag := flag.Bool("i", false, "Ignore case")
	skipFieldsFlag := flag.Int("f", 0, "Skip first N fields")
	skipCharsFlag := flag.Int("s", 0, "Skip first N characters")

	flag.Parse()

	// Проверка взаимоисключающих флагов
	exclusiveCount := 0
	if *countFlag {
		exclusiveCount++
	}
	if *duplicatesFlag {
		exclusiveCount++
	}
	if *uniqueFlag {
		exclusiveCount++
	}

	if exclusiveCount > 1 {
		fmt.Fprintln(os.Stderr, "Error: flags -c, -d, -u are mutually exclusive")
		os.Exit(1)
	}

	// Настройки
	opts := uniq.Options{
		Count:          *countFlag,
		DuplicatesOnly: *duplicatesFlag,
		UniqueOnly:     *uniqueFlag,
		IgnoreCase:     *ignoreCaseFlag,
		SkipFields:     *skipFieldsFlag,
		SkipChars:      *skipCharsFlag,
	}

	// Определяем источник ввода и вывода
	input := os.Stdin
	output := os.Stdout
	args := flag.Args()

	if len(args) > 0 {
		f, err := os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		input = f
	}

	if len(args) > 1 {
		f, err := os.Create(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		output = f
	}

	// Читаем строки
	scanner := bufio.NewScanner(input)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Обрабатываем
	result := uniq.Uniq(lines, opts)

	// Выводим
	for _, line := range result {
		fmt.Fprintln(output, line)
	}
}
