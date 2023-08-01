package mal

import (
	"context"
	"net/http"
	"testing"

	"github.com/nstratos/go-myanimelist/mal"
)

func TestMangaGet(u *testing.T) {
	m := NewManga(
		mal.NewClient(&http.Client{
			Transport: &t{
				cid: CLIENT_ID,
			},
		}),
	)
	_, err := m.Get(context.Background(), "698")

	if err != nil {
		u.Errorf("Get() got unexpected error: %s", err)
	}
}
