package model

import (
	"encoding/json"
	"fmt"
	"reflect"
)

var (
	// AuxLookup are Aux type constructors. There should be one per Aux
	// type (corresponding to a Corpus). This list needs to be manually
	// maintained.
	AuxLookup = map[string]func() Aux{
		reflect.TypeOf(AuxAlbum{}).Name():      func() Aux { return &AuxAlbum{} },
		reflect.TypeOf(AuxAnime{}).Name():      func() Aux { return &AuxAnime{} },
		reflect.TypeOf(AuxAnimeFilm{}).Name():  func() Aux { return &AuxAnimeFilm{} },
		reflect.TypeOf(AuxBook{}).Name():       func() Aux { return &AuxBook{} },
		reflect.TypeOf(AuxFilm{}).Name():       func() Aux { return &AuxFilm{} },
		reflect.TypeOf(AuxGame{}).Name():       func() Aux { return &AuxGame{} },
		reflect.TypeOf(AuxManga{}).Name():      func() Aux { return &AuxManga{} },
		reflect.TypeOf(AuxShortStory{}).Name(): func() Aux { return &AuxShortStory{} },
		reflect.TypeOf(AuxGame{}).Name():       func() Aux { return &AuxGame{} },
	}
)

// APIDataAux embeds the concrete Aux type into the APIData JSON output. We are
// adding an annotation to the parent class instead of adding a __typename field
// to the Aux concrete type because of human scalability.
//
// Example:
//
//	{
//	  "api": "API_MAL",
//	  "id": "manga/698,
//	  "aux": {
//	  },
//	  "__truffle_aux_typename": "AuxManga"
//	}
type APIDataAux struct {
	Typename string `json:"__truffle_aux_typename"`
}

type PartialAPIData struct {
	API       APIType        `json:"api"`
	ID        string         `json:"id"`
	Queued    bool           `json:"queued"`
	Cached    bool           `json:"cached"`
	Completed bool           `json:"completed"`
	Score     *float64       `json:"score,omitempty"`
	Titles    []*Title       `json:"titles,omitempty"`
	Providers []ProviderType `json:"providers,omitempty"`
	Tags      []string       `json:"tags,omitempty"`
	Tracker   Tracker        `json:"tracker,omitempty"`
}

// MarshalJSON is a custom truffle encoder. This function needs to reside in the
// same directory as the dynamically generated model.
func (a APIData) MarshalJSON() ([]byte, error) {
	type FakeAPIData APIData

	if a.Aux != nil {
		var t string
		// The underlying name of pointer types must be explicitly
		// grokked.
		if k := reflect.ValueOf(a.Aux).Kind(); k == reflect.Ptr || k == reflect.Interface {
			t = reflect.TypeOf(a.Aux).Elem().Name()
		} else {
			t = reflect.TypeOf(a.Aux).Name()
		}
		return json.Marshal(struct {
			FakeAPIData
			APIDataAux
		}{
			FakeAPIData: FakeAPIData(a),
			APIDataAux: APIDataAux{
				Typename: t,
			},
		})
	} else {
		return json.Marshal(struct {
			FakeAPIData
		}{
			FakeAPIData: FakeAPIData(a),
		})
	}
}

func (a *APIData) UnmarshalJSON(data []byte) error {
	p := PartialAPIData{}
	if err := json.Unmarshal(data, &p); err != nil {
		return err
	}

	for _, f := range reflect.VisibleFields(reflect.TypeOf(PartialAPIData{})) {
		reflect.ValueOf(a).Elem().FieldByName(f.Name).Set(reflect.ValueOf(p).FieldByName(f.Name))
	}

	// Look for the concrete Aux type.
	t := APIDataAux{}
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	if f, ok := AuxLookup[t.Typename]; ok {
		a.Aux = f()
		return json.Unmarshal(data, a.Aux)
	}

	if t.Typename != "" {
		return fmt.Errorf("Invalid __truffle_aux_typename: %s", t.Typename)
	}

	return nil
}
