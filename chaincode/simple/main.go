package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	switch function {
	case "Ping":
		return shim.Success([]byte("Pong"))
	case "Create":
		err := APIstub.PutState("temp", []byte(args[0]))
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(nil)
	case "Get":
		result, err := APIstub.GetState("temp")
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(result)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
