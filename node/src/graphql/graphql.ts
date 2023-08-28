/* eslint-disable */
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Time: { input: any; output: any; }
};

export type ApiData = {
  __typename?: 'APIData';
  api: ApiType;
  aux?: Maybe<Aux>;
  cached: Scalars['Boolean']['output'];
  completed: Scalars['Boolean']['output'];
  corpus: CorpusType;
  id: Scalars['ID']['output'];
  providers?: Maybe<Array<ProviderType>>;
  queued: Scalars['Boolean']['output'];
  score?: Maybe<Scalars['Float']['output']>;
  tags?: Maybe<Array<Scalars['String']['output']>>;
  titles?: Maybe<Array<Title>>;
  tracker?: Maybe<Tracker>;
};

export enum ApiType {
  ApiKitsu = 'API_KITSU',
  ApiMal = 'API_MAL',
  ApiNone = 'API_NONE',
  ApiSpotify = 'API_SPOTIFY',
  ApiSteam = 'API_STEAM',
  ApiTruffle = 'API_TRUFFLE'
}

export type Aux = AuxAlbum | AuxAnime | AuxAnimeFilm | AuxBook | AuxFilm | AuxGame | AuxManga | AuxShortStory | AuxTv;

export type AuxAlbum = {
  __typename?: 'AuxAlbum';
  composers?: Maybe<Array<Scalars['String']['output']>>;
  labels?: Maybe<Array<Scalars['String']['output']>>;
  producers?: Maybe<Array<Scalars['String']['output']>>;
  studios?: Maybe<Array<Scalars['String']['output']>>;
};

export type AuxAnime = {
  __typename?: 'AuxAnime';
  composers?: Maybe<Array<Scalars['String']['output']>>;
  directors?: Maybe<Array<Scalars['String']['output']>>;
  studios?: Maybe<Array<Scalars['String']['output']>>;
  writers?: Maybe<Array<Scalars['String']['output']>>;
};

export type AuxAnimeFilm = {
  __typename?: 'AuxAnimeFilm';
  composers?: Maybe<Array<Scalars['String']['output']>>;
  directors?: Maybe<Array<Scalars['String']['output']>>;
  studios?: Maybe<Array<Scalars['String']['output']>>;
  writers?: Maybe<Array<Scalars['String']['output']>>;
};

export type AuxBook = {
  __typename?: 'AuxBook';
  authors?: Maybe<Array<Scalars['String']['output']>>;
};

export type AuxFilm = {
  __typename?: 'AuxFilm';
  cinematographers?: Maybe<Array<Scalars['String']['output']>>;
  composers?: Maybe<Array<Scalars['String']['output']>>;
  directors?: Maybe<Array<Scalars['String']['output']>>;
  editors?: Maybe<Array<Scalars['String']['output']>>;
  writers?: Maybe<Array<Scalars['String']['output']>>;
};

export type AuxGame = {
  __typename?: 'AuxGame';
  artists?: Maybe<Array<Scalars['String']['output']>>;
  composers?: Maybe<Array<Scalars['String']['output']>>;
  designers?: Maybe<Array<Scalars['String']['output']>>;
  developers?: Maybe<Array<Scalars['String']['output']>>;
  directors?: Maybe<Array<Scalars['String']['output']>>;
  programmers?: Maybe<Array<Scalars['String']['output']>>;
  writers?: Maybe<Array<Scalars['String']['output']>>;
};

export type AuxManga = {
  __typename?: 'AuxManga';
  artists?: Maybe<Array<Scalars['String']['output']>>;
  authors?: Maybe<Array<Scalars['String']['output']>>;
};

export type AuxShortStory = {
  __typename?: 'AuxShortStory';
  authors?: Maybe<Array<Scalars['String']['output']>>;
};

export type AuxTv = {
  __typename?: 'AuxTV';
  cinematographers?: Maybe<Array<Scalars['String']['output']>>;
  composers?: Maybe<Array<Scalars['String']['output']>>;
  creators?: Maybe<Array<Scalars['String']['output']>>;
  editors?: Maybe<Array<Scalars['String']['output']>>;
  writers?: Maybe<Array<Scalars['String']['output']>>;
};

export enum CorpusType {
  CorpusAlbum = 'CORPUS_ALBUM',
  CorpusAnime = 'CORPUS_ANIME',
  CorpusAnimeFilm = 'CORPUS_ANIME_FILM',
  CorpusBook = 'CORPUS_BOOK',
  CorpusFilm = 'CORPUS_FILM',
  CorpusGame = 'CORPUS_GAME',
  CorpusManga = 'CORPUS_MANGA',
  CorpusNone = 'CORPUS_NONE',
  CorpusShortStory = 'CORPUS_SHORT_STORY',
  CorpusTv = 'CORPUS_TV'
}

export type Entry = {
  __typename?: 'Entry';
  id: Scalars['ID']['output'];
  metadata: Metadata;
};

export type ListInput = {
  apis?: InputMaybe<Array<ApiType>>;
  corpora?: InputMaybe<Array<CorpusType>>;
  corpus?: InputMaybe<CorpusType>;
  id?: InputMaybe<Scalars['ID']['input']>;
  mal?: InputMaybe<ListInputMal>;
  title?: InputMaybe<Scalars['String']['input']>;
};

export type ListInputMal = {
  nsfw: Scalars['Boolean']['input'];
};

export type Metadata = {
  __typename?: 'Metadata';
  sources?: Maybe<Array<ApiData>>;
  truffle: ApiData;
};

export type Mutation = {
  __typename?: 'Mutation';
  delete?: Maybe<Entry>;
  patch?: Maybe<Entry>;
};


export type MutationDeleteArgs = {
  input: Scalars['ID']['input'];
};


export type MutationPatchArgs = {
  input?: InputMaybe<PatchInput>;
};

export type PatchInput = {
  aux?: InputMaybe<PatchInputAux>;
  corpus?: InputMaybe<CorpusType>;
  id?: InputMaybe<Scalars['ID']['input']>;
  providers?: InputMaybe<Array<ProviderType>>;
  queued?: InputMaybe<Scalars['Boolean']['input']>;
  score?: InputMaybe<Scalars['Float']['input']>;
  sources?: InputMaybe<Array<PatchInputApiSource>>;
  tags?: InputMaybe<Array<Scalars['String']['input']>>;
  titles?: InputMaybe<Array<PatchInputTitle>>;
  tracker?: InputMaybe<PatchInputTracker>;
};

export type PatchInputApiSource = {
  api: ApiType;
  id: Scalars['ID']['input'];
};

export type PatchInputAux = {
  authors?: InputMaybe<Array<Scalars['String']['input']>>;
  composers?: InputMaybe<Array<Scalars['String']['input']>>;
  developers?: InputMaybe<Array<Scalars['String']['input']>>;
  directors?: InputMaybe<Array<Scalars['String']['input']>>;
  studios?: InputMaybe<Array<Scalars['String']['input']>>;
};

export type PatchInputTitle = {
  locale: Scalars['String']['input'];
  title: Scalars['String']['input'];
};

export type PatchInputTracker = {
  chapter?: InputMaybe<Scalars['String']['input']>;
  episode?: InputMaybe<Scalars['String']['input']>;
  season?: InputMaybe<Scalars['String']['input']>;
  volume?: InputMaybe<Scalars['String']['input']>;
};

export enum ProviderType {
  ProviderAmazonStreaming = 'PROVIDER_AMAZON_STREAMING',
  ProviderAppleTv = 'PROVIDER_APPLE_TV',
  ProviderCrunchyroll = 'PROVIDER_CRUNCHYROLL',
  ProviderDisneyPlus = 'PROVIDER_DISNEY_PLUS',
  ProviderGooglePlay = 'PROVIDER_GOOGLE_PLAY',
  ProviderHboMax = 'PROVIDER_HBO_MAX',
  ProviderHulu = 'PROVIDER_HULU',
  ProviderNetflix = 'PROVIDER_NETFLIX',
  ProviderNone = 'PROVIDER_NONE',
  ProviderOther = 'PROVIDER_OTHER',
  ProviderSteam = 'PROVIDER_STEAM'
}

export type Query = {
  __typename?: 'Query';
  list?: Maybe<Array<Entry>>;
};


export type QueryListArgs = {
  input?: InputMaybe<ListInput>;
};

export type Title = {
  __typename?: 'Title';
  locale: Scalars['String']['output'];
  title: Scalars['String']['output'];
};

export type Tracker = TrackerAnime | TrackerBook | TrackerManga | TrackerTv;

export type TrackerAnime = {
  __typename?: 'TrackerAnime';
  episode?: Maybe<Scalars['String']['output']>;
  last_updated?: Maybe<Scalars['Time']['output']>;
  season?: Maybe<Scalars['String']['output']>;
};

export type TrackerBook = {
  __typename?: 'TrackerBook';
  last_updated?: Maybe<Scalars['Time']['output']>;
  volume?: Maybe<Scalars['String']['output']>;
};

export type TrackerManga = {
  __typename?: 'TrackerManga';
  chapter?: Maybe<Scalars['String']['output']>;
  last_updated?: Maybe<Scalars['Time']['output']>;
  volume?: Maybe<Scalars['String']['output']>;
};

export type TrackerTv = {
  __typename?: 'TrackerTV';
  episode?: Maybe<Scalars['String']['output']>;
  last_updated?: Maybe<Scalars['Time']['output']>;
  season?: Maybe<Scalars['String']['output']>;
};

export type ApiDataPartsFragment = { __typename?: 'APIData', api: ApiType, id: string, corpus: CorpusType, providers?: Array<ProviderType> | null, tags?: Array<string> | null, titles?: Array<{ __typename?: 'Title', title: string }> | null, aux?: { __typename?: 'AuxAlbum' } | { __typename?: 'AuxAnime', studios?: Array<string> | null } | { __typename?: 'AuxAnimeFilm' } | { __typename?: 'AuxBook' } | { __typename?: 'AuxFilm' } | { __typename?: 'AuxGame' } | { __typename?: 'AuxManga', authors?: Array<string> | null, artists?: Array<string> | null } | { __typename?: 'AuxShortStory' } | { __typename?: 'AuxTV' } | null, tracker?: { __typename?: 'TrackerAnime', season?: string | null, episode?: string | null, last_updated?: any | null } | { __typename?: 'TrackerBook' } | { __typename?: 'TrackerManga', volume?: string | null, chapter?: string | null, last_updated?: any | null } | { __typename?: 'TrackerTV' } | null } & { ' $fragmentName'?: 'ApiDataPartsFragment' };

export const ApiDataPartsFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"APIDataParts"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"APIData"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"api"}},{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"corpus"}},{"kind":"Field","name":{"kind":"Name","value":"titles"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"title"}}]}},{"kind":"Field","name":{"kind":"Name","value":"providers"}},{"kind":"Field","name":{"kind":"Name","value":"tags"}},{"kind":"Field","name":{"kind":"Name","value":"aux"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"InlineFragment","typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"AuxAnime"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"studios"}}]}},{"kind":"InlineFragment","typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"AuxManga"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"authors"}},{"kind":"Field","name":{"kind":"Name","value":"artists"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"tracker"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"InlineFragment","typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"TrackerAnime"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"season"}},{"kind":"Field","name":{"kind":"Name","value":"episode"}},{"kind":"Field","name":{"kind":"Name","value":"last_updated"}}]}},{"kind":"InlineFragment","typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"TrackerManga"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"volume"}},{"kind":"Field","name":{"kind":"Name","value":"chapter"}},{"kind":"Field","name":{"kind":"Name","value":"last_updated"}}]}}]}}]}}]} as unknown as DocumentNode<ApiDataPartsFragment, unknown>;
