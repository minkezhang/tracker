# graphql API

## Example

### List

```
query {
  list(input: {
    title: "Evangelion"
    corpora: [CORPUS_ANIME]
    apis: [API_MAL]
  }) {
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
  corpus
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

### Patch

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
  corpus
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
