package main

import (
	"bufio"
	"fmt"
	"io"
	"learningGo/application/cmd/poker"
	"strings"
)

type CLI struct {
	store poker.PlayerStore
	in    io.Reader
}

func (cli CLI) PlayPoker() {
	scanner := bufio.NewScanner(cli.in)
	scanner.Scan()
	cli.store.RecordWin(extractWinner(scanner.Text()))
}

func main() {
	fmt.Println("Let's play poker")
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
