'use strict';

const express = require("express");
const service = require("./service/fabricClientService");
const app = express();
const { InitSimpleContractOperator } = require("./common");

app.get('/get', async (req, res) => {
    var result = await service.Get();
    res.send(result);
});

app.get('/create', async (req, res) => {
    await service.Create(req.query.message);
    res.send("ok");
});

app.get('/ping', async (req, res) => {
    var result = await service.Ping();
    res.send(result);
});

(async () => {
    await InitSimpleContractOperator();
    app.listen(3000);
})();