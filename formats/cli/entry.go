package entry

import (
	"fmt"
	"strings"

	"github.com/minkezhang/truffle/api/go/database/utils"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

type FormatT int

const (
	FormatFull FormatT = iota
	FormatShort
)

type E struct {
	Format FormatT
	Data   []byte
}

func (e *E) Dump(m proto.Message) error {
	if e.Format == FormatFull {
		data, err := prototext.MarshalOptions{Multiline: true}.Marshal(m)
		if err != nil {
			return err
		}
		e.Data = data
		return nil
	}

	epb := m.(*dpb.Entry)

	lines := append([]string{},
		func() string {
			if len(epb.GetTitles()) > 0 {
				return epb.GetTitles()[0]
			}
			return ""
		}(),

		epb.GetCorpus().String(),

		epb.GetId(),

		map[bool]string{
			true:  "Q",
			false: "",
		}[epb.GetQueued()],

		func() string {
			if s := epb.GetScore(); s > 0 {
				return fmt.Sprintf("%.1f", s)
			}
			return ""
		}(),
	)

	lines = append(lines,
		func() string {
			type D struct {
				SeasonLabel  string
				Season       int32
				EpisodeLabel string
				Episode      int32
			}
			var data D

			switch utils.TrackerL[epb.GetCorpus()] {
			case utils.TrackerVideo:
				tracker := epb.GetTrackerVideo()
				data = D{
					SeasonLabel:  "s",
					Season:       tracker.GetSeason(),
					EpisodeLabel: "e",
					Episode:      tracker.GetEpisode(),
				}
			case utils.TrackerBook:
				tracker := epb.GetTrackerBook()
				data = D{
					SeasonLabel:  "v",
					Season:       tracker.GetVolume(),
					EpisodeLabel: "c",
					Episode:      tracker.GetChapter(),
				}
			}
			var s string
			if data.Season > 0 {
				s = fmt.Sprintf("%v%v", data.SeasonLabel, data.Season)
			}
			if data.Episode > 0 {
				s = fmt.Sprintf("%v%v%v", s, data.EpisodeLabel, data.Episode)
			}
			return s
		}(),
	)

	lines = append(lines,
		func() []string {
			type D struct {
				Director string
				Writer   string
				Studio   string
			}
			data := &D{}

			switch utils.AuxDataL[epb.GetCorpus()] {
			case utils.AuxDataBook:
				aux := epb.GetAuxDataBook()
				if len(aux.GetAuthors()) > 0 {
					data.Writer = aux.GetAuthors()[0]
				}
			case utils.AuxDataVideo:
				aux := epb.GetAuxDataVideo()
				if len(aux.GetDirectors()) > 0 {
					data.Director = aux.GetDirectors()[0]
				}
				if len(aux.GetWriters()) > 0 {
					data.Writer = aux.GetWriters()[0]
				}
				if len(aux.GetStudios()) > 0 {
					data.Studio = aux.GetStudios()[0]
				}
			case utils.AuxDataAudio:
				aux := epb.GetAuxDataAudio()
				if len(aux.GetComposers()) > 0 {
					data.Writer = aux.GetComposers()[0]
				}
			case utils.AuxDataGame:
				aux := epb.GetAuxDataGame()
				if len(aux.GetDirectors()) > 0 {
					data.Director = aux.GetDirectors()[0]
				}
				if len(aux.GetWriters()) > 0 {
					data.Writer = aux.GetWriters()[0]
				}
				if len(aux.GetStudios()) > 0 {
					data.Studio = aux.GetStudios()[0]
				}
			}

			return []string{data.Director, data.Writer, data.Studio}

		}()...,
	)

	lines = append(lines,
		func() string {
			if len(epb.GetProviders()) > 0 {
				return epb.GetProviders()[0].String()
			}
			return ""
		}(),
	)

	e.Data = []byte(fmt.Sprintf("%v\n", strings.Join(lines, "\t")))
	return nil
}
