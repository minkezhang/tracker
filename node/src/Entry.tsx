import React from 'react';

import { useQuery, gql } from '@apollo/client';
import * as types from './graphql/graphql';

function Title(props: any) {
  const { e } = props;

  let sources: types.Title[] = [];
  for (var t of e.metadata.truffle == null ? [] : e.metadata.truffle.titles) {
    sources.push(t);
  }
  for (var s of e.metadata.sources == null ? [] : e.metadata.sources) {
    for (var t of s.titles == null ? [] : s.titles) {
      sources.push(t);
    }
  }

  let titles: string[] = sources.map((x) => x.title)
  if (titles.length > 0) {
    return (
      <div key='{e.id}'>
        <h2>{ titles[0] }</h2>
      </div>
    );
  }
  return (
    <></>
  );
}

export function Entry(props: any) {
  const { e } = props;

  return (
    <div key='{e.id}'>
      <Title e={e}></Title>
    </div>
  );
}


