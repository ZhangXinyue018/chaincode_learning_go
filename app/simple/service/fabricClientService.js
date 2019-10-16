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

async function Create(message) {
    try {
        // Submit the specified transaction.
        await simpleContractOperator.SubmitTransaction('Create', message);
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

async function Get() {
    try {
        const result = await simpleContractOperator.EvaluateTransaction('Get');
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