import React from 'react';
import logo from './logo.svg';
import './App.css';

import { useQuery, gql } from '@apollo/client';
import * as types from './graphql/graphql';
import { Entry } from './Entry';

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

function F() {
  const { loading, error, data } = useQuery(_Q);
  if (loading) {
    return <p>Loading...</p>
  }
  return (
    <div>{ data.list.map((x: types.Entry) => Entry({e: x})) }</div>
  )
}



function App() {
  return (
    <div className="App">
      <header className="App-header">
        <F />
      </header>
    </div>
  );
}

export default App;
