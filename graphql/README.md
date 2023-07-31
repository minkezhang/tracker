# graphql API

## Example

```
mutation {
  patch(input: {
    corpus: CORPUS_ANIME,
    queued: false,
    titles: [{
   		locale: "en",
    	title: "Neon Genesis Evangelion"
    }],
    providers: [
      PROVIDER_NETFLIX,
    ],
    tags: [
      "mechs", "psychological",
    ],
    aux: {
      studios: ["Gainax", "Tatsunoko"]
    },
    sources: [
      {
        api: API_MAL,
        id: "anime/30"
      }
    ],
    tracker: {
      season: "1"
      episode: "10"
    }
  }) {
    id
    corpus
    metadata {
      truffle {
        ...APIDataParts
      }
      sources {
      	...APIDataParts
      }
    }
  }
}

fragment APIDataParts on APIData {
  api
  id
  cached
  titles {
    title
  }
  providers
  tags
  aux {
    ... on AuxAnime {
      studios
    }
    ... on AuxManga {
      authors
      artists
    }
  }
  tracker {
    ... on TrackerAnime {
      season
      episode
      last_updated
    }
    ... on TrackerManga {
      volume
      chapter
      last_updated
    }
  }
}
```
