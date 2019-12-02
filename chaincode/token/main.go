package main

import (
	"fmt"
	_ "github.com/chaincode_learning_go/chaincode/token/repo"
	"github.com/chaincode_learning_go/chaincode/token/service"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

var mainLogger = shim.NewLogger("main")

type SmartContract struct {
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	switch function {
	case "Ping":
		mainLogger.Debugf("ping")
		return service.TokenService.Ping(APIstub)
	case "CreateToken":
		//requestId := args[0]
		maxAmount, _ := strconv.Atoi(args[2])
		tokenName := args[1]
		creator := args[3]
		issuer := args[4]
		return service.TokenService.CreateToken(APIstub, tokenName, int64(maxAmount), creator, issuer)
	case "IssueToken":
		requestId := args[0]
		userName := args[1]
		tokenName := args[2]
		tokenAmount, _ := strconv.Atoi(args[3])
		return service.TokenService.IssueToken(APIstub, requestId, userName, tokenName, int64(tokenAmount))
	case "GetToken":
		//requestId := args[0]
		userName := args[1]
		tokenName := args[2]
		return service.TokenService.GetToken(APIstub, userName, tokenName)
	case "TransferToken":
		requestId := args[0]
		fromUserName := args[1]
		toUserName := args[2]
		tokenName := args[3]
		tokenAmount, _ := strconv.Atoi(args[4])
		return service.TokenService.TransferToken(APIstub, requestId, fromUserName, toUserName, tokenName, int64(tokenAmount))
	case "PaginateTokenBalanceByUser":
		//requestId := args[0]
		userName := args[1]
		pageSize, _ := strconv.Atoi(args[2])
		bookMark := args[3]
		return service.TokenService.PaginateTokenByUserName(APIstub, []string{userName}, int32(pageSize), bookMark)
	case "PaginateTokenBalanceByToken":
		//requestId := args[0]
		tokenName := args[1]
		pageSize, _ := strconv.Atoi(args[2])
		bookMark := args[3]
		return service.TokenService.PaginateTokenByTokenName(APIstub, []string{tokenName}, int32(pageSize), bookMark)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func main() {
	mainLogger.SetLevel(shim.LogDebug)
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
