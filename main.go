package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	const (
		defaultPath = "NONSET"
	)

	path := flag.String("p", defaultPath, "path to file")
	flag.Parse()

	if len(os.Args) < 2 {
		exitWithHelp()
	}
	var lines []string
	if *path == defaultPath {
		lines = readStdin()
	} else {
		lines = readFileLines(*path)
	}
	search(lines, flag.Args())
	os.Exit(0)
}

func exitWithHelp() {
	fmt.Println("****************  USE  *******************")
	fmt.Println("")
	fmt.Println("greg -p <path to file> <text to search>")
	fmt.Println("")
	fmt.Println("****************  OR   *******************")
	fmt.Println("")
	fmt.Println("some command | greg <text to search>")
	fmt.Println("")
	fmt.Println("******************************************")
	os.Exit(1)
}

func readFileLines(filepath string) (lines []string) {
	readFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	if err := fileScanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading file:", err)
	}
	readFile.Close()
	return lines
}

func readStdin() (lines []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return lines
}

func search(lines []string, searchTerms []string) {
	atleastOneMatch := false
	for i, line := range lines {
		for _, term := range searchTerms {
			if strings.Contains(line, term) {
				s := strings.Replace(line, term, fmt.Sprintf("\033[31m%s\033[0m", term), -1)
				fin := fmt.Sprintf("%d: %s", i, s)
				fmt.Println(fin)
				atleastOneMatch = true
			}
		}
	}
	if !atleastOneMatch {
		fmt.Println("No matches found")
	}
}
