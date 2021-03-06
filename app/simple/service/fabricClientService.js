'use strict';

const { simpleContractOperator } = require("../common");

async function Ping() {
    try {
        const result = await simpleContractOperator.EvaluateTransaction("Ping");
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        return result.toString();
    } catch (err) {
        console.log(err);
    }
}

async function Create(key, message) {
    try {
        // Submit the specified transaction.
        await simpleContractOperator.SubmitTransaction('Create', key, message);
        console.log('Transaction has been submitted');
    } catch (err) {
        console.log(err);
    }
}

async function Conflict(message) {
    try {
        await simpleContractOperator.SubmitTransaction('Conflict', message);
        console.log('Transaction of conflict has been submitted');
    } catch (err) {
        console.log(err);
    }
}

async function Get(key) {
    try {
        const result = await simpleContractOperator.EvaluateTransaction('Get', key);
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        return result.toString();
    } catch (err) {
        console.log(err);
    }
}

module.exports = {
    Ping: Ping,
    Get: Get,
    Create: Create,
    Conflict: Conflict
};