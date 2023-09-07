import React from 'react';
import logo from './logo.svg';

import { useQuery, gql } from '@apollo/client';
import * as types from './graphql/graphql';
import * as entry from './Entry';

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
          api
          corpus
          id
          score
          titles {
            title
          }
        }
      }
    }
  }
`)

function F() {
  const { loading, error, data } = useQuery(_Q);
  if (loading) {
    return <p>Loading...</p>
  }
  return (
    <div>{ data.list.map((x: types.Entry) => entry.ApiData({ data: (new entry.E(x)).default()})) }</div>
  )
}



function App() {
  return (
    <F />
  );
}

export default App;
