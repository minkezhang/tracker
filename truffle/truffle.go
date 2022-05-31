// Package main maintains a database of user media entries.
//
// To see how to use the tool, run
//
//   truffle help
//
// To see global flags, run
//
//   truffle flags
package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/google/subcommands"
	"github.com/minkezhang/truffle/database"
	"github.com/minkezhang/truffle/truffle/commands/add"
	"github.com/minkezhang/truffle/truffle/commands/bump"
	"github.com/minkezhang/truffle/truffle/commands/get"
	"github.com/minkezhang/truffle/truffle/commands/patch"
	"github.com/minkezhang/truffle/truffle/commands/search"

	del "github.com/minkezhang/truffle/truffle/commands/delete"
)

var (
	home, _         = os.UserHomeDir()
	defaultFilename = filepath.Join(home, ".truffle/database.textproto")
)

var (
	fn   = flag.String("database", defaultFilename, "user database location")
	mock = flag.Bool("dry_run", false, "do not commit changes to database")
)

func read(fn string) ([]byte, error) {
	if err := os.MkdirAll(filepath.Dir(fn), 0777); err != nil {
		return nil, err
	}
	fp, err := os.OpenFile(fn, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	return ioutil.ReadAll(fp)
}

func write(fn string, data []byte) error {
	return ioutil.WriteFile(fn, data, 0666)
}

func main() {
	flag.Parse()

	data, err := read(*fn)
	if err != nil {
		log.Fatalf("could not read file %v: %v", fn, err)
	}

	db, err := database.Unmarshal(data)
	if err != nil {
		log.Fatalf("could not read database: %v", err)
	}

	for _, c := range []subcommands.Command{
		subcommands.HelpCommand(),
		subcommands.FlagsCommand(),
		subcommands.CommandsCommand(),
		get.New(db),
		add.New(db),
		search.New(db),
		patch.New(db),
		bump.New(db),
		del.New(db),
	} {
		subcommands.Register(c, "")
	}

	ctx := context.Background()

	status := subcommands.Execute(ctx)

	if status == subcommands.ExitSuccess && !*mock {
		data, err := database.Marshal(db)
		if err != nil {
			log.Fatalf("could not marshal database: %v", err)
		}
		if err := write(*fn, data); err != nil {
			log.Fatalf("could not write to database: %v", err)
		}
	}

	os.Exit(int(status))
}
