"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var client_1 = require("@apollo/client");
var client = new client_1.ApolloClient({
    uri: 'http://localhost:8080/query',
    cache: new client_1.InMemoryCache(),
});
var _Q = (0, client_1.gql)("\n  query {\n    list(input: {\n      title: \"Evangelion\"\n      corpora: [\n        CORPUS_ANIME\n      ]\n      apis: [\n        API_MAL\n      ]\n    }) {\n      metadata {\n        sources {\n          titles {\n            title\n          }\n        }\n      }\n    }\n  }\n");
client.query({
    query: _Q,
}).then(function (result) { return console.log(result); });
