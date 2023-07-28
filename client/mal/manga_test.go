package mal

import (
	"context"
	"net/http"
	"testing"

	"github.com/nstratos/go-myanimelist/mal"
)

func TestMangaGet(t *testing.T) {
	m := NewManga(
		mal.NewClient(
			&http.Client{
				Transport: &transport{
					ClientID: CLIENT_ID,
				},
			},
		),
	)
	_, err := m.Get(context.Background(), "698")

	if err != nil {
		t.Errorf("Get() got unexpected error: %s", err)
	}
}
