package mal

import (
	"context"
	"testing"

	"github.com/minkezhang/truffle/graphql/generated/model"
)

func TestMangaGet(t *testing.T) {
	m := NewManga(O{
		ClientID: CLIENT_ID,
	})
	_, err := m.Get(context.Background(), &model.APIData{
		API: model.APITypeAPIMal,
		ID:  "manga/698",
	})

	if err != nil {
		t.Errorf("Get() got unexpected error: %s", err)
	}
}
