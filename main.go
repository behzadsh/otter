package main

import (
	"context"

	"github.com/behzadsh/otter/commands"
)

func main() {
	_ = commands.Execute(context.Background())
}

