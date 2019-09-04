'use strict';

const ContractOperator = require("./contractOperator");

var simpleContractOperator = new ContractOperator(process.cwd(), "mychannel", "simple");

async function InitSimpleContractOperator() {
    await simpleContractOperator.Init();
}

module.exports = { InitSimpleContractOperator, simpleContractOperator };