# graphql API

## Example

```
mutation {
  patch(input: {
    corpus: CORPUS_ANIME,
    id: "",
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
    ]
    aux: {
      studios: ["Gainax", "Tatsunoko"]
    }
    sources: [
      {
        api: API_MAL,
        id: "anime/30"
      }
    ]
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
  cached
  queued
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
}
```
