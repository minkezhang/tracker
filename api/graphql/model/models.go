// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Aux interface {
	IsAux()
}

type Tracker interface {
	IsTracker()
}

type APIData struct {
	API       APIType        `json:"api"`
	ID        string         `json:"id"`
	Corpus    CorpusType     `json:"corpus"`
	Queued    bool           `json:"queued"`
	Cached    bool           `json:"cached"`
	Completed bool           `json:"completed"`
	Score     *float64       `json:"score,omitempty"`
	Titles    []*Title       `json:"titles,omitempty"`
	Providers []ProviderType `json:"providers,omitempty"`
	Tags      []string       `json:"tags,omitempty"`
	Aux       Aux            `json:"aux,omitempty"`
	Tracker   Tracker        `json:"tracker,omitempty"`
}

type AuxAlbum struct {
	Studios   []string `json:"studios,omitempty"`
	Composers []string `json:"composers,omitempty"`
	Labels    []string `json:"labels,omitempty"`
	Producers []string `json:"producers,omitempty"`
}

func (AuxAlbum) IsAux() {}

type AuxAnime struct {
	Composers []string `json:"composers,omitempty"`
	Directors []string `json:"directors,omitempty"`
	Studios   []string `json:"studios,omitempty"`
	Writers   []string `json:"writers,omitempty"`
}

func (AuxAnime) IsAux() {}

type AuxAnimeFilm struct {
	Composers []string `json:"composers,omitempty"`
	Directors []string `json:"directors,omitempty"`
	Studios   []string `json:"studios,omitempty"`
	Writers   []string `json:"writers,omitempty"`
}

func (AuxAnimeFilm) IsAux() {}

type AuxBook struct {
	Authors []string `json:"authors,omitempty"`
}

func (AuxBook) IsAux() {}

type AuxFilm struct {
	Cinematographers []string `json:"cinematographers,omitempty"`
	Composers        []string `json:"composers,omitempty"`
	Directors        []string `json:"directors,omitempty"`
	Editors          []string `json:"editors,omitempty"`
	Writers          []string `json:"writers,omitempty"`
}

func (AuxFilm) IsAux() {}

type AuxGame struct {
	Artists     []string `json:"artists,omitempty"`
	Composers   []string `json:"composers,omitempty"`
	Designers   []string `json:"designers,omitempty"`
	Developers  []string `json:"developers,omitempty"`
	Directors   []string `json:"directors,omitempty"`
	Programmers []string `json:"programmers,omitempty"`
	Writers     []string `json:"writers,omitempty"`
}

func (AuxGame) IsAux() {}

type AuxManga struct {
	Artists []string `json:"artists,omitempty"`
	Authors []string `json:"authors,omitempty"`
}

func (AuxManga) IsAux() {}

type AuxShortStory struct {
	Authors []string `json:"authors,omitempty"`
}

func (AuxShortStory) IsAux() {}

type AuxTv struct {
	Cinematographers []string `json:"cinematographers,omitempty"`
	Creators         []string `json:"creators,omitempty"`
	Composers        []string `json:"composers,omitempty"`
	Editors          []string `json:"editors,omitempty"`
	Writers          []string `json:"writers,omitempty"`
}

func (AuxTv) IsAux() {}

type Entry struct {
	ID       string    `json:"id"`
	Metadata *Metadata `json:"metadata"`
}

type ListInput struct {
	ID      *string      `json:"id,omitempty"`
	Corpus  *CorpusType  `json:"corpus,omitempty"`
	Title   *string      `json:"title,omitempty"`
	Corpora []CorpusType `json:"corpora,omitempty"`
	Apis    []APIType    `json:"apis,omitempty"`
	Nsfw    *bool        `json:"nsfw,omitempty"`
}

type Metadata struct {
	Truffle *APIData   `json:"truffle"`
	Sources []*APIData `json:"sources,omitempty"`
}

type PatchInput struct {
	ID        *string                `json:"id,omitempty"`
	Corpus    *CorpusType            `json:"corpus,omitempty"`
	Queued    *bool                  `json:"queued,omitempty"`
	Titles    []*PatchInputTitle     `json:"titles,omitempty"`
	Score     *float64               `json:"score,omitempty"`
	Providers []ProviderType         `json:"providers,omitempty"`
	Tags      []string               `json:"tags,omitempty"`
	Aux       *PatchInputAux         `json:"aux,omitempty"`
	Tracker   *PatchInputTracker     `json:"tracker,omitempty"`
	Sources   []*PatchInputAPISource `json:"sources,omitempty"`
}

type PatchInputAPISource struct {
	API APIType `json:"api"`
	ID  string  `json:"id"`
}

type PatchInputAux struct {
	Studios    []string `json:"studios,omitempty"`
	Authors    []string `json:"authors,omitempty"`
	Composers  []string `json:"composers,omitempty"`
	Directors  []string `json:"directors,omitempty"`
	Developers []string `json:"developers,omitempty"`
}

type PatchInputTitle struct {
	Locale string `json:"locale"`
	Title  string `json:"title"`
}

type PatchInputTracker struct {
	Season  *string `json:"season,omitempty"`
	Episode *string `json:"episode,omitempty"`
	Volume  *string `json:"volume,omitempty"`
	Chapter *string `json:"chapter,omitempty"`
}

type Title struct {
	Locale string `json:"locale"`
	Title  string `json:"title"`
}

type TrackerAnime struct {
	Season      *string    `json:"season,omitempty"`
	Episode     *string    `json:"episode,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

func (TrackerAnime) IsTracker() {}

type TrackerBook struct {
	Volume      *string    `json:"volume,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

func (TrackerBook) IsTracker() {}

type TrackerManga struct {
	Volume      *string    `json:"volume,omitempty"`
	Chapter     *string    `json:"chapter,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

func (TrackerManga) IsTracker() {}

type TrackerTv struct {
	Season      *string    `json:"season,omitempty"`
	Episode     *string    `json:"episode,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

func (TrackerTv) IsTracker() {}

type APIType string

const (
	APITypeAPINone    APIType = "API_NONE"
	APITypeAPITruffle APIType = "API_TRUFFLE"
	APITypeAPIMal     APIType = "API_MAL"
	APITypeAPIKitsu   APIType = "API_KITSU"
	APITypeAPISteam   APIType = "API_STEAM"
	APITypeAPISpotify APIType = "API_SPOTIFY"
)

var AllAPIType = []APIType{
	APITypeAPINone,
	APITypeAPITruffle,
	APITypeAPIMal,
	APITypeAPIKitsu,
	APITypeAPISteam,
	APITypeAPISpotify,
}

func (e APIType) IsValid() bool {
	switch e {
	case APITypeAPINone, APITypeAPITruffle, APITypeAPIMal, APITypeAPIKitsu, APITypeAPISteam, APITypeAPISpotify:
		return true
	}
	return false
}

func (e APIType) String() string {
	return string(e)
}

func (e *APIType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = APIType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid APIType", str)
	}
	return nil
}

func (e APIType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CorpusType string

const (
	CorpusTypeCorpusNone       CorpusType = "CORPUS_NONE"
	CorpusTypeCorpusAnime      CorpusType = "CORPUS_ANIME"
	CorpusTypeCorpusAnimeFilm  CorpusType = "CORPUS_ANIME_FILM"
	CorpusTypeCorpusManga      CorpusType = "CORPUS_MANGA"
	CorpusTypeCorpusBook       CorpusType = "CORPUS_BOOK"
	CorpusTypeCorpusShortStory CorpusType = "CORPUS_SHORT_STORY"
	CorpusTypeCorpusFilm       CorpusType = "CORPUS_FILM"
	CorpusTypeCorpusTv         CorpusType = "CORPUS_TV"
	CorpusTypeCorpusAlbum      CorpusType = "CORPUS_ALBUM"
	CorpusTypeCorpusGame       CorpusType = "CORPUS_GAME"
)

var AllCorpusType = []CorpusType{
	CorpusTypeCorpusNone,
	CorpusTypeCorpusAnime,
	CorpusTypeCorpusAnimeFilm,
	CorpusTypeCorpusManga,
	CorpusTypeCorpusBook,
	CorpusTypeCorpusShortStory,
	CorpusTypeCorpusFilm,
	CorpusTypeCorpusTv,
	CorpusTypeCorpusAlbum,
	CorpusTypeCorpusGame,
}

func (e CorpusType) IsValid() bool {
	switch e {
	case CorpusTypeCorpusNone, CorpusTypeCorpusAnime, CorpusTypeCorpusAnimeFilm, CorpusTypeCorpusManga, CorpusTypeCorpusBook, CorpusTypeCorpusShortStory, CorpusTypeCorpusFilm, CorpusTypeCorpusTv, CorpusTypeCorpusAlbum, CorpusTypeCorpusGame:
		return true
	}
	return false
}

func (e CorpusType) String() string {
	return string(e)
}

func (e *CorpusType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CorpusType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CorpusType", str)
	}
	return nil
}

func (e CorpusType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProviderType string

const (
	ProviderTypeProviderNone            ProviderType = "PROVIDER_NONE"
	ProviderTypeProviderOther           ProviderType = "PROVIDER_OTHER"
	ProviderTypeProviderCrunchyroll     ProviderType = "PROVIDER_CRUNCHYROLL"
	ProviderTypeProviderDisneyPlus      ProviderType = "PROVIDER_DISNEY_PLUS"
	ProviderTypeProviderHulu            ProviderType = "PROVIDER_HULU"
	ProviderTypeProviderHboMax          ProviderType = "PROVIDER_HBO_MAX"
	ProviderTypeProviderNetflix         ProviderType = "PROVIDER_NETFLIX"
	ProviderTypeProviderAmazonStreaming ProviderType = "PROVIDER_AMAZON_STREAMING"
	ProviderTypeProviderGooglePlay      ProviderType = "PROVIDER_GOOGLE_PLAY"
	ProviderTypeProviderAppleTv         ProviderType = "PROVIDER_APPLE_TV"
	ProviderTypeProviderSteam           ProviderType = "PROVIDER_STEAM"
)

var AllProviderType = []ProviderType{
	ProviderTypeProviderNone,
	ProviderTypeProviderOther,
	ProviderTypeProviderCrunchyroll,
	ProviderTypeProviderDisneyPlus,
	ProviderTypeProviderHulu,
	ProviderTypeProviderHboMax,
	ProviderTypeProviderNetflix,
	ProviderTypeProviderAmazonStreaming,
	ProviderTypeProviderGooglePlay,
	ProviderTypeProviderAppleTv,
	ProviderTypeProviderSteam,
}

func (e ProviderType) IsValid() bool {
	switch e {
	case ProviderTypeProviderNone, ProviderTypeProviderOther, ProviderTypeProviderCrunchyroll, ProviderTypeProviderDisneyPlus, ProviderTypeProviderHulu, ProviderTypeProviderHboMax, ProviderTypeProviderNetflix, ProviderTypeProviderAmazonStreaming, ProviderTypeProviderGooglePlay, ProviderTypeProviderAppleTv, ProviderTypeProviderSteam:
		return true
	}
	return false
}

func (e ProviderType) String() string {
	return string(e)
}

func (e *ProviderType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProviderType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProviderType", str)
	}
	return nil
}

func (e ProviderType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
