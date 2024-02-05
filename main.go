package main

import (
	"regexp"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var printBefore = flag.Int("b", 0, "print lines before the match")
var printAfter = flag.Int("a", 0, "print lines after the match")

func main() {
    
	if len(os.Args) < 2 {
		exitWithHelp()
	}
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, `
   __ _ _ __ ___  __ _ 
  / _  | '__/ _ \/ _  |
 | (_| | | |  __/ (_| |
  \__, |_|  \___|\__, |
   __/ |          __/ |
  |___/          |___/ 
		
  `)
		fmt.Fprintf(os.Stdout, "<some command> | greg [OPTIONS] PATTERN\n")
		fmt.Fprintf(os.Stdout, "\nOptions:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	input := readStdin()
	if len(input) == 0 {
		exitWithHelp()
	}
	search(readStdin(), flag.Args())
	os.Exit(0)
}

func exitWithHelp() {
	fmt.Println("type 'greg --help' for help")
	os.Exit(1)
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
			//change Contains to Regex match
			if regexp.MustCompile(term).MatchString(line) {
				s := strings.Replace(line, term, fmt.Sprintf("\033[31m%s\033[0m", term), -1)
				fin := fmt.Sprintf("%d: %s", i, s)
				fmt.Println("***MATCH***")
				if *printBefore > 0 {
					for j := i - *printBefore; j < i; j++ {
						if j >= 0 {
							fmt.Println(fmt.Sprintf("%d: %s", j, lines[j]))
						}
					}
				}
				fmt.Println(fin)
				if *printAfter > 0 {
					for j := i + 1; j <= i+*printAfter; j++ {
						if j < len(lines) {
							fmt.Println(fmt.Sprintf("%d: %s", j, lines[j]))
						}
					}
				}
				atleastOneMatch = true
			}
		}
	}
	if !atleastOneMatch {
		fmt.Println("No matches found")
	}
}
