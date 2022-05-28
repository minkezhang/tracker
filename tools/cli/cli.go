package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/subcommands"
	"github.com/minkezhang/tracker/database"
	"github.com/minkezhang/tracker/tools/cli/commands/add"
	"github.com/minkezhang/tracker/tools/cli/commands/bump"
	"github.com/minkezhang/tracker/tools/cli/commands/get"
	"github.com/minkezhang/tracker/tools/cli/commands/patch"
	"github.com/minkezhang/tracker/tools/cli/commands/search"
)

const (
	fn = "data/database.textproto"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	fp, err := os.Open(fn)
	if err != nil {
		log.Fatalf("could not open file %v: %v", fp, err)
	}
	defer fp.Close()

	data, err := ioutil.ReadAll(fp)
	if err != nil {
		log.Fatalf("could not read file %v: %v", fp, err)
	}

	db, err := database.Unmarshal(data)
	if err != nil {
		log.Fatalf("could not read database: %v", err)
	}

	subcommands.Register(get.New(db), "")
	subcommands.Register(add.New(db), "")
	subcommands.Register(search.New(db), "")
	subcommands.Register(patch.New(db), "")
	subcommands.Register(bump.New(db), "")

	flag.Parse()
	ctx := context.Background()

	status := subcommands.Execute(ctx)
	os.Exit(int(status))
}
