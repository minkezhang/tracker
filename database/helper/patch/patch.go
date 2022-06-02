package patch

import (
	"context"

	"github.com/minkezhang/truffle/api/go/database/utils"
	"github.com/minkezhang/truffle/database"
	"github.com/minkezhang/truffle/database/helper/get"
	"google.golang.org/protobuf/proto"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

func Patch(ctx context.Context, db *database.DB, epb *dpb.Entry) (*dpb.Entry, error) {
	fpb, err := get.Get(ctx, db, epb, []dpb.API{
		dpb.API_API_TRUFFLE,
	})
	if err != nil {
		return nil, err
	}

	// Use the user-supplied ETag as a concurrency check.
	if len(epb.GetEtag()) > 0 {
		fpb.Etag = nil
	}

	// Entry titles are not directly changeable.
	epb.Titles = nil

	// Assume the user-supplied queue tag is the source of truth, as there
	// is no way to unset a bool field using proto.Merge.
	fpb.Queued = false

	proto.Merge(fpb, epb)

	// Ensure we actually replace lists, not the default append as per
	// default Merge behavior.
	switch utils.AuxDataL[fpb.GetCorpus()] {
	case utils.AuxDataVideo:
		if len(epb.GetAuxDataVideo().GetStudios()) > 0 {
			fpb.GetAuxDataVideo().Studios = epb.GetAuxDataVideo().GetStudios()
		}
		if len(epb.GetAuxDataVideo().GetWriters()) > 0 {
			fpb.GetAuxDataVideo().Writers = epb.GetAuxDataVideo().GetWriters()
		}
		if len(epb.GetAuxDataVideo().GetDirectors()) > 0 {
			fpb.GetAuxDataVideo().Directors = epb.GetAuxDataVideo().GetDirectors()
		}
	case utils.AuxDataAudio:
		if len(epb.GetAuxDataAudio().GetComposers()) > 0 {
			fpb.GetAuxDataAudio().Composers = epb.GetAuxDataAudio().GetComposers()
		}
	case utils.AuxDataBook:
		if len(epb.GetAuxDataBook().GetAuthors()) > 0 {
			fpb.GetAuxDataBook().Authors = epb.GetAuxDataBook().GetAuthors()
		}
	case utils.AuxDataGame:
		if len(epb.GetAuxDataGame().GetStudios()) > 0 {
			fpb.GetAuxDataGame().Studios = epb.GetAuxDataGame().GetStudios()
		}
		if len(epb.GetAuxDataGame().GetStudios()) > 0 {
			fpb.GetAuxDataGame().Studios = epb.GetAuxDataGame().GetStudios()
		}
		if len(epb.GetAuxDataGame().GetWriters()) > 0 {
			fpb.GetAuxDataGame().Writers = epb.GetAuxDataGame().GetWriters()
		}
	}
	if len(epb.GetLinkedIds()) > 0 {
		fpb.LinkedIds = epb.GetLinkedIds()
	}
	if len(epb.GetProviders()) > 0 {
		fpb.Providers = epb.GetProviders()
	}

	return db.Put(ctx, fpb)
}
