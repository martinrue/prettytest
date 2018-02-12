package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

func main() {
	skip := flag.Bool("skip", false, "skip [no test files] in output")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		color := reset

		switch {
		// failure
		case strings.HasPrefix(trimmed, "--- FAIL"):
			fallthrough
		case strings.HasPrefix(trimmed, "FAIL"):
			color = red

		// success
		case strings.HasPrefix(trimmed, "ok"):
			fallthrough
		case strings.HasPrefix(trimmed, "--- PASS"):
			fallthrough
		case strings.HasPrefix(trimmed, "PASS"):
			color = green

		// no test files
		case strings.HasSuffix(trimmed, "[no test files]"):
			if *skip {
				continue
			}

			color = yellow
		}

		fmt.Fprintf(os.Stderr, color+line+"\033[0m\n")
	}
}
