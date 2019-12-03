package main

import (
	"fmt"
	"github.com/chaincode_learning_go/chaincode/token/domain/protos"
	"github.com/chaincode_learning_go/chaincode/token/domain/resp"
	_ "github.com/chaincode_learning_go/chaincode/token/repo"
	"github.com/chaincode_learning_go/chaincode/token/service"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"strings"
)

var mainLogger = shim.NewLogger("main")

type SmartContract struct {
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	mainLogger.Debug("Hello request")
	function, args := APIstub.GetFunctionAndParameters()
	if len(args) <= 0 {
		return shim.Error("arg not enough")
	}
	inputs := strings.Split(args[0], ",")
	requestBytes := make([]byte, 0)
	for _, input := range inputs{
		inputv, _ := strconv.Atoi(input)
		requestBytes = append(requestBytes, byte(inputv))
	}

	switch function {
	case "Ping":
		request := &protos.PingRequest{}
		err := proto.Unmarshal(requestBytes, request)
		if err != nil {
			return shim.Error("non known")
		}
		response := service.TokenService.Ping(APIstub, request)
		return shim.Success(resp.ToResponseBytes(response))
	case "CreateToken":
		request := &protos.CreateTokenRequest{}
		err := proto.Unmarshal(requestBytes, request)
		if err != nil {
			mainLogger.Error(err.Error())
			return shim.Error(err.Error())
		}
		if request.MaxAmount == 0 {
			return shim.Error("not valid protobuf")
		}
		response := service.TokenService.CreateToken(APIstub, request)
		return shim.Success(resp.ToResponseBytes(response))
	case "IssueToken":
		request := &protos.IssueTokenRequest{}
		err := proto.Unmarshal(requestBytes, request)
		if err != nil {
			return shim.Error("non known")
		}
		response := service.TokenService.IssueToken(APIstub, request)
		return shim.Success(resp.ToResponseBytes(response))
	case "GetToken":
		request := &protos.GetTokenRequest{}
		err := proto.Unmarshal(requestBytes, request)
		if err != nil {
			return shim.Error("non known")
		}
		response := service.TokenService.GetToken(APIstub, request)
		return shim.Success(resp.ToResponseBytes(response))
	case "TransferToken":
		request := &protos.TransferTokenRequest{}
		err := proto.Unmarshal(requestBytes, request)
		if err != nil {
			return shim.Error("non known")
		}
		response := service.TokenService.TransferToken(APIstub, request)
		return shim.Success(resp.ToResponseBytes(response))
	case "PaginateTokenBalanceByUser":
		request := &protos.PaginateTokenByUserNameRequest{}
		err := proto.Unmarshal(requestBytes, request)
		if err != nil {
			return shim.Error("non known")
		}
		response := service.TokenService.PaginateTokenByUserName(APIstub, request)
		return shim.Success(resp.ToResponseBytes(response))
	case "PaginateTokenBalanceByToken":
		request := &protos.PaginateTokenByTokenNameRequest{}
		err := proto.Unmarshal(requestBytes, request)
		if err != nil {
			return shim.Error("non known")
		}
		response := service.TokenService.PaginateTokenByTokenName(APIstub, request)
		return shim.Success(resp.ToResponseBytes(response))
	}

	return shim.Error("not supported")
}

func main() {
	mainLogger.SetLevel(shim.LogDebug)
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
