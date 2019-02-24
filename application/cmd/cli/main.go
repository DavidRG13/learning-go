package main

import (
	"fmt"
	"io"
	"learningGo/application/cmd/poker"
)

type CLI struct {
	store poker.PlayerStore
	in    io.Reader
}

func (cli CLI) PlayPoker() {
	cli.store.RecordWin("Chris")
}

func main() {
	fmt.Println("Let's play poker")
}
