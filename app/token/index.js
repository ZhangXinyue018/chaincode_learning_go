'use strict';

const express = require("express");
const service = require("./service/fabricClientService");
const app = express();
const {InitSimpleContractOperator} = require("./common");
const swaggerUi = require('swagger-ui-express');
const swaggerDocument = require('./swagger.json');


app.get('/create-token', async (req, res) => {
    await service.CreateToken(
        req.query.requestId, req.query.tokenName, req.query.maxAmount,
        req.query.creator, req.query.issuer);
    res.send("ok");
});

app.get('/ping', async (req, res) => {
    var result = await service.Ping();
    res.send(result);
});

// This is to generate mvcc read conflict
app.get('/issue-token', async (req, res) => {
    await service.IssueToken(
        req.query.requestId, req.query.userName, req.query.tokenType,
        req.query.tokenAmount);
    res.send("ok");
});

app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerDocument));

(async () => {
    await InitSimpleContractOperator();
    app.listen(3000);
})();