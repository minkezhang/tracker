package model

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMarshalAPIData(t *testing.T) {
	season := "1"
	configs := []struct {
		name string
		data APIData
		want []byte
	}{
		{
			name: "NilAux",
			data: APIData{
				ID: "hello-world",
			},
			want: []byte(`
			{
				"api": "",
				"id": "hello-world",
				"corpus": "",
				"queued": false,
				"cached": false,
				"completed": false,
				"__truffle_aux_typename": "",
				"__truffle_tracker_typename": ""
			}`),
		},
		{
			name: "AuxGame",
			data: APIData{
				Aux: AuxGame{
					Developers: []string{"Westwood"},
				},
			},
			want: []byte(`
			{
				"api": "",
				"id": "",
				"corpus": "",
				"queued": false,
				"cached": false,
				"completed": false,
				"aux": {
					"developers":["Westwood"]
				},
				"__truffle_aux_typename": "AuxGame",
				"__truffle_tracker_typename": ""
			}`),
		},
		{
			name: "AuxGamePointer",
			data: APIData{
				Aux: &AuxGame{
					Developers: []string{"Westwood"},
				},
			},
			want: []byte(`
			{
				"api": "",
				"id": "",
				"corpus": "",
				"queued": false,
				"cached": false,
				"completed": false,
				"aux": {
					"developers":["Westwood"]
				},
				"__truffle_aux_typename": "AuxGame",
				"__truffle_tracker_typename": ""
			}`),
		},
		{
			name: "TrackerAnime",
			data: APIData{
				Tracker: TrackerAnime{
					Season: &season,
				},
			},
			want: []byte(`
			{
				"api": "",
				"id": "",
				"corpus": "",
				"queued": false,
				"cached": false,
				"completed": false,
				"tracker": {
					"season": "1"
				},
				"__truffle_aux_typename": "",
				"__truffle_tracker_typename": "TrackerAnime"
			}`),
		},
		{
			name: "TrackerAnimePointer",
			data: APIData{
				Tracker: &TrackerAnime{
					Season: &season,
				},
			},
			want: []byte(`
			{
				"api": "",
				"id": "",
				"corpus": "",
				"queued": false,
				"cached": false,
				"completed": false,
				"tracker": {
					"season": "1"
				},
				"__truffle_aux_typename": "",
				"__truffle_tracker_typename": "TrackerAnime"
			}`),
		},
	}

	for _, c := range configs {
		t.Run(c.name, func(t *testing.T) {
			// Strip all whitespace, as Marshal does not
			// pretty-print the output.
			want := strings.Join(strings.Fields(string(c.want)), "")
			got, err := json.Marshal(c.data)
			if err != nil {
				t.Fatalf("Marshal() returned unexpected error: %s", err)
			}
			if diff := cmp.Diff(want, string(got)); diff != "" {
				t.Errorf("Marshal() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUnmarshalAPIData(t *testing.T) {
	configs := []struct {
		name string
		data APIData
		want APIData
	}{
		{
			name: "Trivial",
			data: APIData{},
			want: APIData{},
		},
		{
			name: "AuxBook",
			data: APIData{
				Aux: &AuxBook{},
			},
			want: APIData{
				Aux: &AuxBook{},
			},
		},
		{
			name: "TrackerBook",
			data: APIData{
				Tracker: &TrackerBook{},
			},
			want: APIData{
				Tracker: &TrackerBook{},
			},
		},
	}

	for _, c := range configs {
		t.Run(c.name, func(t *testing.T) {
			data, err := json.Marshal(&c.data)
			if err != nil {
				t.Fatalf("Marshal() returned unexected error: %s", err)
			}
			got := APIData{}
			if err := json.Unmarshal(data, &got); err != nil {
				t.Fatalf("Unmarshal() returned unexpected error: %s", err)
			}
			if diff := cmp.Diff(c.want, got); diff != "" {
				t.Errorf("Unmarshal() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestPartialAPIDataConformance(t *testing.T) {
	UnionFields := map[string]bool{
		"Aux":     true,
		"Tracker": true,
	}

	for _, f := range reflect.VisibleFields(reflect.TypeOf(APIData{})) {
		if _, ok := UnionFields[f.Name]; !ok && reflect.ValueOf(PartialAPIData{}).FieldByName(f.Name) == (reflect.Value{}) {
			t.Fatalf(
				"VisibleFields() = %v, want = %v",
				reflect.VisibleFields(reflect.TypeOf(PartialAPIData{})),
				reflect.VisibleFields(reflect.TypeOf(APIData{})),
			)
		}
	}
}
