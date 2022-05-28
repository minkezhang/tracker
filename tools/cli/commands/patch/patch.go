package patch

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/minkezhang/tracker/api/go/database/utils"
	"github.com/minkezhang/tracker/database"
	"github.com/minkezhang/tracker/tools/cli/commands/get/common"
	"google.golang.org/protobuf/proto"

	dpb "github.com/minkezhang/tracker/api/go/database"
	ce "github.com/minkezhang/tracker/formats/cli"
	se "github.com/minkezhang/tracker/formats/cli/struct"
)

type C struct {
	db *database.DB

	title *se.Title
	id    *se.ID

	body *se.Body
}

func New(db *database.DB) *C {
	return &C{
		db: db,

		title: &se.Title{},
		body:  &se.Body{},
		id:    &se.ID{},
	}
}

func (c *C) Name() string     { return "patch" }
func (c *C) Synopsis() string { return "patch entry with matching query parameters" }
func (c *C) Usage() string    { return fmt.Sprintf("%v\n", c.Synopsis()) }

func (c *C) SetFlags(f *flag.FlagSet) {
	c.title.SetFlags(f)
	c.body.SetFlags(f)
	c.id.SetFlags(f)
}

func (c *C) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	epb, err := common.Get(common.O{
		DB:     c.db,
		Title:  c.title.Title,
		ID:     c.id.ID,
		Corpus: c.body.GetCorpus(),
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return subcommands.ExitFailure
	}

	c.body.SetCorpus(epb.GetCorpus())
	s, err := c.body.Load()
	if err != nil {
		fmt.Printf("%v\n", err)
		return subcommands.ExitFailure
	}

	fpb := s.(*dpb.Entry)

	proto.Merge(epb, fpb)
	// Ensure we actually replace lists, not the default append as per
	// default Merge behavior.
	switch utils.AuxDataL[epb.GetCorpus()] {
	case utils.AuxDataVideo:
		if len(fpb.GetAuxDataVideo().GetStudios()) > 0 {
			epb.GetAuxDataVideo().Studios = fpb.GetAuxDataVideo().GetStudios()
		}
		if len(fpb.GetAuxDataVideo().GetWriters()) > 0 {
			epb.GetAuxDataVideo().Writers = fpb.GetAuxDataVideo().GetWriters()
		}
		if len(fpb.GetAuxDataVideo().GetDirectors()) > 0 {
			epb.GetAuxDataVideo().Directors = fpb.GetAuxDataVideo().GetDirectors()
		}
	case utils.AuxDataAudio:
		if len(fpb.GetAuxDataAudio().GetComposers()) > 0 {
			epb.GetAuxDataAudio().Composers = fpb.GetAuxDataAudio().GetComposers()
		}
	case utils.AuxDataBook:
		if len(fpb.GetAuxDataBook().GetAuthors()) > 0 {
			epb.GetAuxDataBook().Authors = fpb.GetAuxDataBook().GetAuthors()
		}
	case utils.AuxDataGame:
		if len(fpb.GetAuxDataGame().GetStudios()) > 0 {
			epb.GetAuxDataGame().Studios = fpb.GetAuxDataGame().GetStudios()
		}
		if len(fpb.GetAuxDataGame().GetStudios()) > 0 {
			epb.GetAuxDataGame().Studios = fpb.GetAuxDataGame().GetStudios()
		}
		if len(fpb.GetAuxDataGame().GetWriters()) > 0 {
			epb.GetAuxDataGame().Writers = fpb.GetAuxDataGame().GetWriters()
		}
	}

	gpb, err := c.db.PutEntry(epb)
	if err != nil {
		fmt.Printf("%v\n", err)
		return subcommands.ExitFailure
	}

	e := &ce.E{}
	e.Dump(gpb)
	fmt.Printf(string(e.Data))

	return subcommands.ExitSuccess
}
