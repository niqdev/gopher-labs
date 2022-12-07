package main

import (
	"fmt"

	"github.com/niqdev/gopher-labs/cmd"
	labs "github.com/niqdev/gopher-labs/internal"
)

func main() {
	fmt.Println(labs.Banner())

	cmd.Execute()
}
