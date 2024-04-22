// Package server contains the main server functionaliry.
package main

import (
	"context"
	"fmt"

	"github.com/anicolaspp/battleship/maps"
	"github.com/anicolaspp/battleship/server"
)

func main() {
	fmt.Println("Hello")

	m := server.Module{}
	board, err := maps.Build(context.Background(), 10, 2)
	fmt.Println(board, err)

	fmt.Println(m)
}
