package main

import (
	"fmt"
	"github.com/chaincode_learning_go/chaincode/coursera/common"
	"github.com/chaincode_learning_go/chaincode/coursera/feature/school"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type CourseraContract struct {
	SchoolRepo school.SchoolRepo
}

func (cc *CourseraContract) Init(APIstub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (cc *CourseraContract) Invoke(APIstub shim.ChaincodeStubInterface) pb.Response {
	function, args := APIstub.GetFunctionAndParameters()
	switch function {
	case "CreateSchool":
		return cc.SchoolRepo.CreateSchool(APIstub, args)
	case "GetSchool":
		return cc.SchoolRepo.GetSchool(APIstub, args)
	}
	return shim.Error("Invalid Smart Contract function name.")
}

func main() {
	courseraContract := CourseraContract{
		SchoolRepo: school.SchoolRepo{
			Repository: common.Repository{Prefix: "school"},
		},
	}
	err := shim.Start(&courseraContract)
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
