'use strict';

const {simpleContractOperator} = require("../common");

async function Ping() {
    try {
        const result = await simpleContractOperator.EvaluateTransaction("Ping");
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        return result.toString();
    } catch (err) {
        console.log(err);
    }
}

async function CreateToken(requestId, tokenName, maxAmount, creator, issuer) {
    try {
        // Submit the specified transaction.
        await simpleContractOperator.SubmitTransaction(
            'CreateToken', requestId, tokenName, maxAmount, creator, issuer);
        console.log('Transaction has been submitted');
    } catch (err) {
        console.log(err);
    }
}

async function IssueToken(requestId, userName, tokenType, tokenAmount) {
    try {
        await simpleContractOperator.SubmitTransaction(
            'IssueToken', requestId, userName, tokenType, tokenAmount);
        console.log('Transaction of conflict has been submitted');
    } catch (err) {
        console.log(err);
    }
}

module.exports = {
    Ping: Ping,
    CreateToken: CreateToken,
    IssueToken: IssueToken
};