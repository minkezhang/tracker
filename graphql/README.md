# graphql API

```bash
go run github.com/99designs/gqlgen generate
```

## Example

```
mutation {
  put(input: {
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
        cached
        queued
        api
        id
        titles {
          locale
          title
        }
        aux {
          ... on AuxAnime {
            studios
          }
        },
        providers,
        tags,
      }
      sources {
      	id
        cached
      }
    }
  }
}
```
