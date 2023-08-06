"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.graphql = void 0;
/* eslint-disable */
var types = require("./graphql");
/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
var documents = {
    "fragment APIDataParts on APIData {\n  api\n  id\n  corpus\n  titles {\n    title\n  }\n  providers\n  tags\n  aux {\n    ... on AuxAnime {\n      studios\n    }\n    ... on AuxManga {\n      authors\n      artists\n    }\n  }\n  tracker {\n    ... on TrackerAnime {\n      season\n      episode\n      last_updated\n    }\n    ... on TrackerManga {\n      volume\n      chapter\n      last_updated\n    }\n  }\n}": types.ApiDataPartsFragmentDoc,
};
function graphql(source) {
    var _a;
    return (_a = documents[source]) !== null && _a !== void 0 ? _a : {};
}
exports.graphql = graphql;
