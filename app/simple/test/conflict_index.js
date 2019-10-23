'use strict';

const http = require('http');

for (let i = 0; i < 5; i++) {
    http.get({
        hostname: 'localhost',
        port: 3000,
        path: `/conflict?message=123`,
        agent: false  // Create a new agent just for this one request
    }, (res) => {
        console.log(`sent: get`);
        // Do stuff with response
    });
}