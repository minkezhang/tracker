import { ApolloClient, InMemoryCache, ApolloProvider, gql } from '@apollo/client';

const client = new ApolloClient({
  uri: 'http://localhost:8080/query',
  cache: new InMemoryCache(),
});

const _Q = gql(`
  query {
    list(input: {
      title: "Evangelion"
      corpora: [
        CORPUS_ANIME
      ]
      apis: [
        API_MAL
      ]
    }) {
      metadata {
        sources {
          titles {
            title
          }
        }
      }
    }
  }
`)

client.query({
  query: _Q,
}).then((result) => console.log(result));
