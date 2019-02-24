package main

import (
	"fmt"
	"learningGo/application/cmd/poker"
)

type CLI struct {
	store poker.PlayerStore
}

func (cli CLI) PlayPoker() {
	cli.store.RecordWin("something")
}

func main() {
	fmt.Println("Let's play poker")
}
