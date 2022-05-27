package entry

import (
	"fmt"
	"strings"

	"github.com/minkezhang/tracker/api/go/database/utils"
	"google.golang.org/protobuf/proto"

	dpb "github.com/minkezhang/tracker/api/go/database"
)

type E struct {
	Data []byte
}

func (e *E) Dump(m proto.Message) error {
	epb := m.(*dpb.Entry)

	lines := append([]string{},
		func() string {
			if len(epb.GetTitles()) > 0 {
				return fmt.Sprintf("Title: %v", epb.GetTitles()[0])
			}
			return ""
		}(),

		func() string {
			if l := len(epb.GetTitles()); l > 1 {
				return fmt.Sprintf("Alternate Titles: %v", strings.Join(epb.GetTitles()[1:l-1], ", "))
			}
			return ""
		}(),

		func() string {
			if epb.GetCorpus() != dpb.Corpus_CORPUS_UNKNOWN {
				return fmt.Sprintf("Category: %v", epb.GetCorpus().String())
			}
			return ""
		}(),

		func() string {
			if s := epb.GetScore(); s > 0 {
				return fmt.Sprintf("Score: %.1f", s)
			}
			return ""
		}(),
	)
	lines = append(lines,
		func() []string {
			var data []string

			switch utils.AuxDataL[epb.GetCorpus()] {
			case utils.AuxDataBook:
				if len(epb.GetAuxDataBook().GetAuthors()) > 0 {
					data = append(data, fmt.Sprintf("Authors: %v", strings.Join(epb.GetAuxDataBook().GetAuthors(), ", ")))
				}
			case utils.AuxDataVideo:
				if len(epb.GetAuxDataVideo().GetDirectors()) > 0 {
					data = append(data, fmt.Sprintf("Directors: %v", strings.Join(epb.GetAuxDataVideo().GetDirectors(), ", ")))
				}
				if len(epb.GetAuxDataVideo().GetWriters()) > 0 {
					data = append(data, fmt.Sprintf("Writers: %v", strings.Join(epb.GetAuxDataVideo().GetWriters(), ", ")))
				}
				if len(epb.GetAuxDataVideo().GetStudios()) > 0 {
					data = append(data, fmt.Sprintf("Studios: %v", strings.Join(epb.GetAuxDataVideo().GetStudios(), ", ")))
				}
			case utils.AuxDataAudio:
				if len(epb.GetAuxDataAudio().GetComposers()) > 0 {
					data = append(data, fmt.Sprintf("Composers: %v", strings.Join(epb.GetAuxDataAudio().GetComposers(), ", ")))
				}
			case utils.AuxDataGame:
				if len(epb.GetAuxDataGame().GetDirectors()) > 0 {
					data = append(data, fmt.Sprintf("Directors: %v", strings.Join(epb.GetAuxDataGame().GetDirectors(), ", ")))
				}
				if len(epb.GetAuxDataGame().GetWriters()) > 0 {
					data = append(data, fmt.Sprintf("Writers: %v", strings.Join(epb.GetAuxDataGame().GetWriters(), ", ")))
				}
				if len(epb.GetAuxDataGame().GetStudios()) > 0 {
					data = append(data, fmt.Sprintf("Studios: %v", strings.Join(epb.GetAuxDataGame().GetStudios(), ", ")))
				}
			}

			return data

		}()...,
	)

	lines = append(lines,
		func() []string {
			var data []string

			switch utils.TrackerL[epb.GetCorpus()] {
			case utils.TrackerVideo:
				if epb.GetTrackerVideo().GetSeason() != 0 {
					data = append(data, fmt.Sprintf("Season: %v", epb.GetTrackerVideo().GetSeason()))
				}
				if epb.GetTrackerVideo().GetEpisode() != 0 {
					data = append(data, fmt.Sprintf("Episode: %v", epb.GetTrackerVideo().GetEpisode()))
				}
			case utils.TrackerBook:
				if epb.GetTrackerBook().GetVolume() != 0 {
					data = append(data, fmt.Sprintf("Volume: %v", epb.GetTrackerBook().GetVolume()))
				}
				if epb.GetTrackerBook().GetChapter() != 0 {
					data = append(data, fmt.Sprintf("Chapter: %v", epb.GetTrackerBook().GetChapter()))
				}
			}

			return data
		}()...,
	)

	lines = append(lines,
		func() string {
			if len(epb.GetProviders()) > 0 {
				var providers []string
				for _, p := range epb.GetProviders() {
					providers = append(providers, p.String())
				}
				return fmt.Sprintf("Providers: %v", strings.Join(providers, ", "))
			}
			return ""
		}(),

		fmt.Sprintf("ID: %s", epb.GetId()),
		fmt.Sprintf("ETag: %s", epb.GetEtag()),
	)

	var data []string
	for _, l := range lines {
		if l != "" {
			data = append(data, l)
		}
	}

	e.Data = []byte(strings.Join(data, "\n"))
	return nil
}
