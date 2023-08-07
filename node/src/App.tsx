import React from 'react';
import logo from './logo.svg';
import './App.css';

import { useQuery, gql } from '@apollo/client';
import * as types from './graphql/graphql';

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

function RenderEntry(e: types.Entry) {
  let t = ""
  if (e.metadata.truffle && e.metadata.truffle.titles) {
    t = e.metadata.truffle.titles[0].title
  } else if (e.metadata.sources) {
    if (e.metadata.sources[0].titles) {
      t = e.metadata.sources[0].titles[0].title
    }
  }
  return (
    <div key='{ e.id }'>{ t }</div>
  )
}

function F() {
  const { loading, error, data } = useQuery(_Q);
  if (loading) {
    return <p>Loading...</p>
  }
  return (
    <div>{ data.list.map(RenderEntry) }</div>
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
