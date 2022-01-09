package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Preparing dictionary")
	input := os.Args[1]
	output := os.Args[2]

	fmt.Println("Preparing dictionary:", input, "->", output)
	// open input
	f, err := os.Open(input)
	check(err)

	// Create output file
	out, err := os.Create(output)
	check(err)

	// remember to close the input at the end of the program
	defer f.Close()

	// read the input line by line using scanner
	scanner := bufio.NewScanner(f)

	var IsLetters = regexp.MustCompile(`^[a-z]+$`).MatchString

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		line := scanner.Text()
		line = strings.ToLower(line)
		if len(line) == 5 && IsLetters(line) {
			_, err := out.WriteString(line)
			check(err)
			_, err = out.WriteString("\n")
			check(err)
		}
	}

}
