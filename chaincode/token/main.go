package main

import (
	"fmt"
	"github.com/chaincode_learning_go/chaincode/token/domain"
	"github.com/chaincode_learning_go/chaincode/token/repo"
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
		mainLogger.Info(fmt.Sprintf("level for UserTokenRepo is %v", repo.TokenBalanceRepository.Logger.IsEnabledFor(shim.LogDebug)))
		mainLogger.Info(fmt.Sprintf("level for TokenLogRepo is %v", repo.TokenLogRepository.Logger.IsEnabledFor(shim.LogDebug)))
		mainLogger.Info(fmt.Sprintf("level for TokenRepo is %v", repo.TokenCreationRepository.Logger.IsEnabledFor(shim.LogDebug)))
		return service.Ping(APIstub)
	case "CreateToken":
		requestId := args[0]
		maxAmount, _ := strconv.Atoi(args[2])
		token := &domain.TokenCreation{
			TokenName:     args[1],
			MaxAmount:     int64(maxAmount),
			CurrentAmount: 0,
			Creator:       args[3],
			Issuer:        args[4],
		}
		return service.CreateToken(APIstub, requestId, token)
	case "IssueToken":
		requestId := args[0]
		userName := args[1]
		tokenName := args[2]
		tokenAmount, _ := strconv.Atoi(args[3])
		return service.IssueToken(APIstub, requestId, userName, tokenName, int64(tokenAmount))
	case "GetToken":
		//requestId := args[0]
		userName := args[1]
		tokenName := args[2]
		return service.GetToken(APIstub, userName, tokenName)
	case "TransferToken":
		requestId := args[0]
		fromUserName := args[1]
		toUserName := args[2]
		tokenName := args[3]
		tokenAmount, _ := strconv.Atoi(args[4])
		return service.TransferToken(APIstub, requestId, fromUserName, toUserName, tokenName, int64(tokenAmount))
	case "PaginateTokenBalanceByUser":
		//requestId := args[0]
		userName := args[1]
		pageSize, _ := strconv.Atoi(args[2])
		bookMark := args[3]
		return service.PaginateTokenByUserName(APIstub, []string{userName}, int32(pageSize), bookMark)
	case "PaginateTokenBalanceByToken":
		//requestId := args[0]
		tokenName := args[1]
		pageSize, _ := strconv.Atoi(args[2])
		bookMark := args[3]
		return service.PaginateTokenByTokenName(APIstub, []string{tokenName}, int32(pageSize), bookMark)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func main() {
	mainLogger.SetLevel(shim.LogDebug)
	mainLogger.Debug("start")
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
