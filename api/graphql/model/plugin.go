// Package model adds supplementary plugin behavior for the auto-generated
// GraphQL code. This must reside in the same directory as the auto-generated
// code, since it contains logic which adds custom instance methods.
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

	// TrackerLookup are Tracker type constructors. There should be one per
	// Tracker type. This list needs to be manually maintained.
	TrackerLookup = map[string]func() Tracker{
		reflect.TypeOf(TrackerAnime{}).Name(): func() Tracker { return &TrackerAnime{} },
		reflect.TypeOf(TrackerBook{}).Name():  func() Tracker { return &TrackerBook{} },
		reflect.TypeOf(TrackerManga{}).Name(): func() Tracker { return &TrackerManga{} },
		reflect.TypeOf(TrackerTv{}).Name():    func() Tracker { return &TrackerTv{} },
	}
)

// APIDataUnion embeds the concrete Aux type into the APIData JSON output. We are
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
type APIDataUnion struct {
	AuxTypename     string `json:"__truffle_aux_typename"`
	TrackerTypename string `json:"__truffle_tracker_typename"`
}

type PartialAPIData struct {
	API       APIType        `json:"api"`
	ID        string         `json:"id"`
	Queued    bool           `json:"queued"`
	Cached    bool           `json:"cached"`
	Corpus    CorpusType     `json:"corpus"`
	Completed bool           `json:"completed"`
	Score     *float64       `json:"score,omitempty"`
	Titles    []*Title       `json:"titles,omitempty"`
	Providers []ProviderType `json:"providers,omitempty"`
	Tags      []string       `json:"tags,omitempty"`
}

// MarshalJSON is a custom truffle encoder. This function needs to reside in the
// same directory as the dynamically generated model.
func (a APIData) MarshalJSON() ([]byte, error) {
	type FakeAPIData APIData
	type MarshalData struct {
		FakeAPIData
		APIDataUnion
	}

	f := &MarshalData{
		FakeAPIData: FakeAPIData(a),
	}

	// Add AuxTypename annotation.
	if a.Aux != nil {
		var t string
		// The underlying name of pointer types must be explicitly
		// grokked.
		if k := reflect.ValueOf(a.Aux).Kind(); k == reflect.Ptr || k == reflect.Interface {
			t = reflect.TypeOf(a.Aux).Elem().Name()
		} else {
			t = reflect.TypeOf(a.Aux).Name()
		}
		f.APIDataUnion.AuxTypename = t
	}

	// Add TrackerTypename annotation.
	if a.Tracker != nil {
		var t string
		if k := reflect.ValueOf(a.Tracker).Kind(); k == reflect.Ptr || k == reflect.Interface {
			t = reflect.TypeOf(a.Tracker).Elem().Name()
		} else {
			t = reflect.TypeOf(a.Tracker).Name()
		}
		f.APIDataUnion.TrackerTypename = t
	}

	return json.Marshal(f)
}

func (a *APIData) UnmarshalJSON(data []byte) error {
	p := PartialAPIData{}
	if err := json.Unmarshal(data, &p); err != nil {
		return err
	}

	for _, f := range reflect.VisibleFields(reflect.TypeOf(PartialAPIData{})) {
		reflect.ValueOf(a).Elem().FieldByName(f.Name).Set(reflect.ValueOf(p).FieldByName(f.Name))
	}

	t := APIDataUnion{}
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	// Look for the concrete Aux type.
	if f, ok := AuxLookup[t.AuxTypename]; ok {
		a.Aux = f()
		return json.Unmarshal(data, a.Aux)
	}
	if t.AuxTypename != "" {
		return fmt.Errorf("Invalid __truffle_aux_typename: %s", t.AuxTypename)
	}

	// Look for the concrete Tracker type.
	if f, ok := TrackerLookup[t.TrackerTypename]; ok {
		a.Tracker = f()
		return json.Unmarshal(data, a.Tracker)
	}
	if t.TrackerTypename != "" {
		return fmt.Errorf("Invalid __truffle_tracker_typename: %s", t.TrackerTypename)
	}

	return nil
}
