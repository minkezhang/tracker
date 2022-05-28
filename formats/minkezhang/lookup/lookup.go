package lookup

import (
	dpb "github.com/minkezhang/truffle/api/go/database"
)

var (
	Corpus = map[string]dpb.Corpus{
		"Anime":       dpb.Corpus_CORPUS_ANIME,
		"Anime Film":  dpb.Corpus_CORPUS_ANIME_FILM,
		"Manga":       dpb.Corpus_CORPUS_MANGA,
		"Book":        dpb.Corpus_CORPUS_BOOK,
		"TV Show":     dpb.Corpus_CORPUS_TV,
		"Film":        dpb.Corpus_CORPUS_FILM,
		"Game":        dpb.Corpus_CORPUS_GAME,
		"Album":       dpb.Corpus_CORPUS_ALBUM,
		"Short Story": dpb.Corpus_CORPUS_SHORT_STORY,
	}

	Provider = map[string]dpb.Provider{
		"Crunchyroll":      dpb.Provider_PROVIDER_CRUNCHYROLL,
		"Netflix":          dpb.Provider_PROVIDER_NETFLIX,
		"Amazon Streaming": dpb.Provider_PROVIDER_AMAZON_STREAMING,
		"Battlenet":        dpb.Provider_PROVIDER_BATTLENET,
		"Epic":             dpb.Provider_PROVIDER_EPIC,
		"Google Play":      dpb.Provider_PROVIDER_GOOGLE_PLAY,
		"Hulu":             dpb.Provider_PROVIDER_HULU,
		"Origins":          dpb.Provider_PROVIDER_ORIGINS,
		"Steam":            dpb.Provider_PROVIDER_STEAM,
		"Funimation":       dpb.Provider_PROVIDER_FUNIMATION,
		"Disney+":          dpb.Provider_PROVIDER_DISNEY_PLUS,
		"HBO Max":          dpb.Provider_PROVIDER_HBO_MAX,
	}
)
