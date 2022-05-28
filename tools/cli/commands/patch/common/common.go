package common

import (
	"github.com/minkezhang/truffle/api/go/database/utils"
	"github.com/minkezhang/truffle/database"
	"github.com/minkezhang/truffle/tools/cli/commands/get/common"
	"google.golang.org/protobuf/proto"

	dpb "github.com/minkezhang/truffle/api/go/database"
	se "github.com/minkezhang/truffle/formats/cli/struct"
)

type O struct {
	DB *database.DB

	Title string
	ID    string

	Body *se.Body
}

func Patch(opts O) (*dpb.Entry, error) {
	epb, err := common.Get(common.O{
		DB:     opts.DB,
		Title:  opts.Title,
		ID:     opts.ID,
		Corpus: opts.Body.GetCorpus(),
	})
	if err != nil {
		return nil, err
	}

	opts.Body.SetCorpus(epb.GetCorpus())
	s, err := opts.Body.Load()
	if err != nil {
		return nil, err
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

	return opts.DB.PutEntry(epb)
}
