package main

import (
	"fmt"
	"github.com/chaincode_learning_go/chaincode/token/domain"
	"github.com/chaincode_learning_go/chaincode/token/service"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
	"time"

	_ "github.com/chaincode_learning_go/chaincode/token/repo"
	sc "github.com/hyperledger/fabric/protos/peer"
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
		mainLogger.Debug("start to ping")
		return service.Ping(APIstub)
	case "CreateToken":
		requestId := args[0]
		maxAmount, _ := strconv.Atoi(args[2])
		token := &domain.Token{
			TokenName:     args[1],
			MaxAmount:     int64(maxAmount),
			CurrentAmount: 0,
			Creator:       args[3],
			Issuer:        args[4],
			CreatedTime:   time.Now(),
		}
		return service.CreateToken(APIstub, requestId, token)
	case "IssueToken":
		requestId := args[0]
		userName := args[1]
		tokenType := args[2]
		tokenAmount, _ := strconv.Atoi(args[3])
		return service.IssueToken(APIstub, requestId, userName, tokenType, int64(tokenAmount))
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
