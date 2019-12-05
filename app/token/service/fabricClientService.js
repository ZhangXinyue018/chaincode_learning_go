'use strict';

const {simpleContractOperator} = require("../common");
const request = require("../protos/tokenservice-req_pb");
const response = require("../protos/tokenservice-resp_pb");

async function Ping(requestId) {
    try {
        var commonRequest = new request.CommonRequest();
        commonRequest.setRequestid(requestId);
        var pingRequest = new request.PingRequest();
        pingRequest.setCommonrequest(commonRequest);
        var byte = pingRequest.serializeBinary();
        var ccresp = await simpleContractOperator.EvaluateTransaction("Ping", byte.toString());
        console.log(ccresp);
        return response.PingResponse.deserializeBinary(ccresp);
    } catch (err) {
        console.log(err);
        return err;
    }
}

async function CreateToken(requestId, tokenName, maxAmount, creator, issuer) {
    try {
        var commonRequest = new request.CommonRequest();
        commonRequest.setRequestid(requestId);
        var createTokenReq = new request.CreateTokenRequest();
        createTokenReq.setCommonrequest(commonRequest);
        createTokenReq.setTokenname(tokenName);
        createTokenReq.setMaxamount(Number(maxAmount));
        createTokenReq.setCreator(creator);
        createTokenReq.setIssuer(issuer);
        // Submit the specified transaction.
        var byte = createTokenReq.serializeBinary();
        await simpleContractOperator.SubmitTransaction('CreateToken', byte.toString());
        console.log('Transaction has been submitted');
    } catch (err) {
        console.log(err);
    }
}

async function IssueToken(requestId, userName, tokenName, tokenAmount) {
    try {
        var issueTokenReq = new request.IssueTokenRequest();
        issueTokenReq.setCommonrequest(new request.CommonRequest([requestId]));
        issueTokenReq.setUsername(userName);
        issueTokenReq.setTokenname(tokenName);
        issueTokenReq.setTokenamount(tokenAmount);
        var byte = issueTokenReq.serializeBinary();
        await simpleContractOperator.SubmitTransaction('IssueToken', byte.toString());
        console.log('Transaction of issuetoken has been submitted');
    } catch (err) {
        console.log(err);
    }
}

async function GetToken(requestId, userName, tokenName) {
    try {
        var getTokeReq = new request.GetTokenRequest();
        getTokeReq.setCommonrequest(new request.CommonRequest([requestId]));
        getTokeReq.setUsername(userName);
        getTokeReq.setTokenname(tokenName)
        var byte = getTokeReq.serializeBinary();
        var ccresp = await simpleContractOperator.SubmitTransaction('GetToken', byte.toString());
        return response.GetTokenResponse.deserializeBinary(Uint8Array.from(ccresp));
    } catch (err) {
        console.log(err);
    }
}

async function TransferToken(requestId, fromUserName, toUserName, tokenName, tokenAmount) {
    try {
        var transferTokenReq = new request.TransferTokenRequest();
        transferTokenReq.setCommonrequest(new request.CommonRequest([requestId]));
        transferTokenReq.setFromusername(fromUserName);
        transferTokenReq.setTousername(toUserName);
        transferTokenReq.setTokenname(tokenName);
        transferTokenReq.setTokenamount(tokenAmount);
        var byte = transferTokenReq.serializeBinary();
        return await simpleContractOperator.SubmitTransaction('TransferToken', byte.toString());
    } catch (err) {
        console.log(err);
        return "error";
    }
}

async function PaginateTokenBalanceByUser(requestId, userName, pageSize, bookMark) {
    try {
        var paginateRequest = new request.PaginateTokenByUserNameRequest();
        paginateRequest.setCommonrequest(new request.CommonRequest([requestId]));
        paginateRequest.setUsername(userName);
        paginateRequest.setPagesize(pageSize);
        paginateRequest.setBookmark(bookMark);
        var byte = paginateRequest.serializeBinary();
        var ccresp = await simpleContractOperator.SubmitTransaction('PaginateTokenBalanceByUser', byte.toString());
        return response.PaginateTokenByUserNameResponse.deserializeBinary(Uint8Array.from(ccresp));
    } catch (err) {
        console.log(err);
    }
}

async function PaginateTokenBalanceByToken(requestId, tokenName, pageSize, bookMark) {
    try {
        var paginationRequest = new request.PaginateTokenByTokenNameRequest();
        paginationRequest.setCommonrequest(new request.CommonRequest([requestId]));
        paginationRequest.setTokenname(tokenName);
        paginationRequest.setPagesize(pageSize);
        paginationRequest.setBookmark(bookMark);
        var byte = paginationRequest.serializeBinary();
        var ccresp = await simpleContractOperator.SubmitTransaction('PaginateTokenBalanceByToken', byte.toString());
        return response.PaginateTokenByTokenNameResponse.deserializeBinary(Uint8Array.from(ccresp));
    } catch (err) {
        console.log(err);
    }
}

async function PaginateTokenLogByFromUser(requestId, fromUserName, pageSize, bookMark) {
    try {
        var paginationRequest = new request.PaginateTokenLogByFromUserNameRequest();
        paginationRequest.setCommonrequest(new request.CommonRequest([requestId]));
        paginationRequest.setFromusername(fromUserName);
        paginationRequest.setPagesize(pageSize);
        paginationRequest.setBookmark(bookMark);
        var byte = paginationRequest.serializeBinary();
        var ccresp = await simpleContractOperator.SubmitTransaction('PaginateTokenLogByFromUser', byte.toString());
        return response.PaginateTokenLogByFromUserNameResponse.deserializeBinary(Uint8Array.from(ccresp));
    } catch (err) {
        console.log(err);
    }
}

async function PaginateTokenLogByToUser(requestId, toUserName, pageSize, bookMark) {
    try {
        var paginationRequest = new request.PaginateTokenLogByToUserNameRequest();
        paginationRequest.setCommonrequest(new request.CommonRequest([requestId]));
        paginationRequest.setTousername(toUserName);
        paginationRequest.setPagesize(pageSize);
        paginationRequest.setBookmark(bookMark);
        var byte = paginationRequest.serializeBinary();
        var ccresp = await simpleContractOperator.SubmitTransaction('PaginateTokenLogByToUser', byte.toString());
        return response.PaginateTokenLogByToUserNameResponse.deserializeBinary(Uint8Array.from(ccresp));
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
    TransferToken: TransferToken,
    PaginateTokenLogByFromUser: PaginateTokenLogByFromUser,
    PaginateTokenLogByToUser: PaginateTokenLogByToUser
};