package git

import (
	"context"
	"flag"
	"fmt"
	"os/exec"
	"strings"

	"github.com/google/subcommands"
	"github.com/minkezhang/truffle/truffle/commands/common"
)

type C struct {
	common    common.O
	directory string

	args []string
}

func New(directory string, common common.O) *C {
	return &C{
		common:    common,
		directory: directory,
	}
}

func (c *C) Name() string             { return "git" }
func (c *C) Synopsis() string         { return "execute git commands on the user DB repository" }
func (c *C) Usage() string            { return fmt.Sprintf("%v\n", c.Synopsis()) }
func (c *C) SetFlags(f *flag.FlagSet) {}

func (c *C) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	c.args = f.Args()

	b := fmt.Sprintf(
		`
		#/bin/bash
		cd %v
		git %v
		`, c.directory, strings.Join(c.args, " "),
	)

	cmd := exec.Command("bash", "-c", b)
	cmd.Stdout = c.common.Output
	cmd.Stderr = c.common.Error

	if err := cmd.Run(); err != nil {
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
