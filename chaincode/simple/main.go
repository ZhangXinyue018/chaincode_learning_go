package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("main")

type SmartContract struct {
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	switch function {
	case "Ping":
		logger.Warning("Ping")
		return shim.Success([]byte("Pong"))
	case "Create":
		input1 := args[0]
		input2 := args[1]
		key, _ := APIstub.CreateCompositeKey("temp", []string{"hello", input1})
		logger.Warning(key)
		err := APIstub.PutState(key, []byte(input2))
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(nil)
	case "Conflict":
		// This is a dummy test to generate mvcc conflict
		input := []byte(args[0])
		err := APIstub.PutState("temp", input)
		if err != nil {
			return shim.Error(err.Error())
		}

		result, err := APIstub.GetState("temp")

		err = APIstub.PutState("temp", []byte("testconflict"))
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(result)
	case "Get":
		input1 := args[0]
		key, _ := APIstub.CreateCompositeKey("temp", []string{"hello", input1})
		logger.Warning(key)
		result, err := APIstub.GetState(key)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(result)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func main() {
	logger.SetLevel(shim.LogInfo)
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
