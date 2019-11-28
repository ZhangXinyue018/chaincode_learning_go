'use strict';

const ContractOperator = require("./contractOperator");

var simpleContractOperator = new ContractOperator(process.cwd(), "mychannel", "token");

async function InitSimpleContractOperator() {
    await simpleContractOperator.Init();
}

module.exports = { InitSimpleContractOperator, simpleContractOperator };