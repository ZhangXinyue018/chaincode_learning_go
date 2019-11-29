'use strict';

const {simpleContractOperator} = require("../common");

async function Ping() {
    try {
        const result = await simpleContractOperator.EvaluateTransaction("Ping");
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        return result;
    } catch (err) {
        console.log(err);
        return err;
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

async function IssueToken(requestId, userName, tokenName, tokenAmount) {
    try {
        await simpleContractOperator.SubmitTransaction(
            'IssueToken', requestId, userName, tokenName, tokenAmount);
        console.log('Transaction of issuetoken has been submitted');
    } catch (err) {
        console.log(err);
    }
}

async function GetToken(requestId, userName, tokenName) {
    try {
        return await simpleContractOperator.SubmitTransaction(
            'GetToken', requestId, userName, tokenName);
    } catch (err) {
        console.log(err);
    }
}

async function TransferToken(requestId, fromUserName, toUserName, tokenName, tokenAmount) {
    try {
        return await simpleContractOperator.SubmitTransaction(
            'TransferToken', requestId, fromUserName, toUserName, tokenName, tokenAmount);
    } catch (err) {
        console.log(err);
        return "error";
    }
}

async function PaginateTokenBalanceByUser(requestId, userName, pageSize, bookMark) {
    try {
        return await simpleContractOperator.SubmitTransaction(
            'PaginateTokenBalanceByUser', requestId, userName, pageSize, bookMark);
    } catch (err) {
        console.log(err);
    }
}

async function PaginateTokenBalanceByToken(requestId, tokenName, pageSize, bookMark) {
    try {
        return await simpleContractOperator.SubmitTransaction(
            'PaginateTokenBalanceByToken', requestId, tokenName, pageSize, bookMark);
    } catch (err) {
        console.log(err);
    }
}

module.exports = {
    Ping: Ping,
    CreateToken: CreateToken,
    IssueToken: IssueToken,
    GetToken: GetToken,
    PaginateTokenBalanceByUser: PaginateTokenBalanceByUser,
    PaginateTokenBalanceByToken: PaginateTokenBalanceByToken,
    TransferToken: TransferToken
};