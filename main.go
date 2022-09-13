package main

import (
	"fmt"

	"github.com/tb0hdan/aws-k8s/pkg/commands"
	"github.com/tb0hdan/aws-k8s/pkg/handler"
)

type XCmd struct {
	Print bool `cmd:"x" help:"Example command"`
}

func (p *XCmd) Run(ctx *commands.CLIContext) error {
	if p.Print {
		fmt.Printf("Hello, %s. Debug mode enabled: %v\n", ctx.User.Name, ctx.Debug)
	}
	return nil
}

func main() {
	cli := &handler.CLI{}
	commandHandler := handler.New()
	commandHandler.ClearEnvironment()
	commandHandler.AddCommand("x", "do stuff", "", &XCmd{})
	ctx := commandHandler.Parse(cli)
	commandHandler.Run(ctx, cli.Debug)
}
