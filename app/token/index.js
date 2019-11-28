'use strict';

const express = require("express");
const service = require("./service/fabricClientService");
const app = express();
const {InitSimpleContractOperator} = require("./common");
const swaggerUi = require('swagger-ui-express');
const swaggerDocument = require('./swagger.json');
const uuidv1 = require('uuid/v1');

app.get('/create-token', async (req, res) => {
    await service.CreateToken(
        uuidv1(), req.query.tokenName, req.query.maxAmount,
        req.query.creator, req.query.issuer);
    res.send("ok");
});

app.get('/ping', async (req, res) => {
    var result = await service.Ping();
    res.send(result);
});

app.get('/get-token', async (req, res) => {
    var result = await service.GetToken(
        uuidv1(), req.query.userName, req.query.tokenName
    );
    var strResult = new Buffer(result).toString('ascii');
    res.send(strResult);
});

app.get('/transfer-token', async (req, res) => {
    var result = await service.TransferToken(
        uuidv1(), req.query.fromUserName, req.query.toUserName,
        req.query.tokenName, req.query.tokenAmount
    );
    res.send(result);
});

// This is to generate mvcc read conflict
app.get('/issue-token', async (req, res) => {
    await service.IssueToken(
        uuidv1(), req.query.userName, req.query.tokenName,
        req.query.tokenAmount);
    res.send("ok");
});

app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerDocument));

(async () => {
    await InitSimpleContractOperator();
    app.listen(3000);
})();