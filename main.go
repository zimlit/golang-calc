package main

import (
	"bufio"
	"calc/scanner"
	"fmt"
	"os"
)

func main() {
	runPrompt()
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')

		if line == "" {
			break
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		run(line)
	}
}

func run(line string) {
	sc := scanner.NewScanner(line)

	tokens, err := sc.ScanTokens()
	if err != nil {
		fmt.Println(err)
		os.Exit(65)
	}

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		fmt.Println(token)
	}
}
