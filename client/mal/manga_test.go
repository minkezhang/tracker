package mal

import (
	"context"
	"testing"
)

func TestMangaGet(t *testing.T) {
	m := NewManga(O{
		ClientID: CLIENT_ID,
	})
	_, err := m.Get(context.Background(), "698")

	if err != nil {
		t.Errorf("Get() got unexpected error: %s", err)
	}
}
