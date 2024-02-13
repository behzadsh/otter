package main

import (
	"context"

	"github.com/behzadsh/otter/cmd/commands"
)

func main() {
	_ = commands.Execute(context.Background())
}
