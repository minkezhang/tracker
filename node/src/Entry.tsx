import React from 'react';

import { useQuery, gql } from '@apollo/client';
import * as types from './graphql/graphql';

export class E {
  sources: types.ApiData[]
  constructor(e: types.Entry) {
    let sources: types.ApiData[] = [];

    if (e.metadata.truffle != null) {
      sources.push(e.metadata.truffle);
    }
    for (var s of e.metadata.sources == null ? [] : e.metadata.sources) {
      sources.push(s);
    }

    this.sources = sources;
  }

  _aux(): types.Maybe<types.Aux> {
    for (var s of this.sources) {
      if (s.aux != null) {
        return s.aux
      }
    }
    return null;
  }

  _tracker(): types.Maybe<types.Tracker> {
    for (var s of this.sources) {
      if (s.tracker != null) {
        return s.tracker
      }
    }
    return null;
  }

  _providers(): types.Maybe<Array<types.ProviderType>> {
    let providers: types.ProviderType[] = [];
    for (var s of this.sources) {
      for (var p of s.providers == null ? [] : s.providers) {
        providers.push(p);
      }
    }
    return providers;
  }

  _tags(): types.Maybe<Array<types.Scalars['String']['output']>> {
    let tags: types.Scalars['String']['output'][] = [];
    for (var s of this.sources) {
      for (var t of s.tags == null ? [] : s.tags) {
        tags.push(t);
      }
    }
    return tags;
  }

  _titles(): types.Maybe<Array<types.Title>> {
    let titles: types.Maybe<Array<types.Title>> = [];
    for (var s of this.sources) {
      for (var t of s.titles == null ? [] : s.titles) {
        titles.push(t)
      }
    }
    return titles;
  }

  default(): types.ApiData {
    let d = {
      api: types.ApiType.ApiNone,
      aux: this._aux(),
      cached: false,
      completed: false,
      corpus: types.CorpusType.CorpusNone,
      id: "",  // ID will only display the truffle ID, if it is available.
      queued: false,
      providers: this._providers(),
      score: 0,
      tags: this._tags(),
      titles: this._titles(),
      tracker: this._tracker(),
    }

    for (var s of this.sources) {
      d.api = d.api == types.ApiType.ApiNone ? s.api : d.api;
      d.cached = d.cached || s.cached;
      d.completed = d.completed || s.completed;
      d.corpus = d.corpus == types.CorpusType.CorpusNone ? s.corpus : d.corpus;
      d.id = s.api == types.ApiType.ApiTruffle ? s.id : d.id;
      d.queued = d.queued || s.queued;
      d.score = d.score == 0 && s.score != null ? s.score : d.score;
    }

    return d;
  }
}

function Title(props: any) {
  const { data } = props;

  if (data.titles != null && data.titles.length > 0) {
    return (
      <div key='{data.id}'>
        <h2>{ data.titles[0].title }</h2>
      </div>
    );
  }
  return (
    <></>
  );
}

export function ApiData(props: any) {
  const { data } = props;

  return (
    <div key='{data.id}'>
      <Title data={data}></Title>
    </div>
  );
}


