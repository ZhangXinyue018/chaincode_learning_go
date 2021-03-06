'use strict';

const http = require('http');

// This is another test to generate MVCC READ CONFLICT, 
// but in service, the Get method needs to be submitted to orderer
// meaning that use SubmitTransaction instead of EvaluateTransaction

http.get({
    hostname: 'localhost',
    port: 3000,
    path: `/create?message=test1`,
    agent: false  // Create a new agent just for this one request
}, (res) => {
    console.log(`sent1: create test1`);
    // Do stuff with response
});

http.get({
    hostname: 'localhost',
    port: 3000,
    path: `/get`,
    agent: false  // Create a new agent just for this one request
}, (res) => {
    console.log(`sent2: get`);
    // Do stuff with response
});

http.get({
    hostname: 'localhost',
    port: 3000,
    path: `/create?message=test3`,
    agent: false  // Create a new agent just for this one request
}, (res) => {
    console.log(`sent3: create test3`);
    // Do stuff with response
});

http.get({
    hostname: 'localhost',
    port: 3000,
    path: `/get`,
    agent: false  // Create a new agent just for this one request
}, (res) => {
    console.log(`sent4: get`);
    // Do stuff with response
});

http.get({
    hostname: 'localhost',
    port: 3000,
    path: `/create?message=test5`,
    agent: false  // Create a new agent just for this one request
}, (res) => {
    console.log(`sent5: create test5`);
    // Do stuff with response
});

http.get({
    hostname: 'localhost',
    port: 3000,
    path: `/get`,
    agent: false  // Create a new agent just for this one request
}, (res) => {
    console.log(`sent6: get`);
    // Do stuff with response
});
