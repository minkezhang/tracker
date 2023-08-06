"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var client_1 = require("@apollo/client");
var client = new client_1.ApolloClient({
    uri: "http://localhost:8080/query",
    cache: new client_1.InMemoryCache(),
});
console.log("HELLO WORLD");
