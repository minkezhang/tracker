package model

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMarshalAPIData(t *testing.T) {
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
			want: []byte(`{"api":"","id":"hello-world","queued":false,"cached":false,"completed":false}`),
		},
		{
			name: "AuxGame",
			data: APIData{
				Aux: AuxGame{
					Developers: []string{"Supergiant Games"},
				},
			},
			want: []byte(`{"api":"","id":"","queued":false,"cached":false,"completed":false,"aux":{"developers":["Supergiant Games"]},"__truffle_aux_typename":"AuxGame"}`),
		},
		{
			name: "AuxGamePointer",
			data: APIData{
				Aux: &AuxGame{
					Developers: []string{"Supergiant Games"},
				},
			},
			want: []byte(`{"api":"","id":"","queued":false,"cached":false,"completed":false,"aux":{"developers":["Supergiant Games"]},"__truffle_aux_typename":"AuxGame"}`),
		},
	}

	for _, c := range configs {
		t.Run(c.name, func(t *testing.T) {
			got, err := json.Marshal(c.data)
			if err != nil {
				t.Fatalf("Marshal() returned unexpected error: %s", err)
			}
			if diff := cmp.Diff(string(c.want), string(got)); diff != "" {
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
	for _, f := range reflect.VisibleFields(reflect.TypeOf(APIData{})) {
		if f.Name != "Aux" && reflect.ValueOf(PartialAPIData{}).FieldByName(f.Name) == (reflect.Value{}) {
			t.Fatalf(
				"VisibleFields() = %v, want = %v",
				reflect.VisibleFields(reflect.TypeOf(PartialAPIData{})),
				reflect.VisibleFields(reflect.TypeOf(APIData{})),
			)
		}
	}
}
