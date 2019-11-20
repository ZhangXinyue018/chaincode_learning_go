package school

import (
	"github.com/chaincode_learning_go/chaincode/coursera/common"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// SchoolRepo is repository for school
type SchoolRepo struct {
	common.Repository
}

func (repo *SchoolRepo) CreateSchool(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	name := args[0]
	school := School{
		Name:        args[0],
		Description: args[1],
	}
	repo.create(APIstub, name, school)
	return shim.Success(nil)
}

func (repo *SchoolRepo) GetSchool(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	name := args[0]
	result, err := repo.getBytes(APIstub, name)
	if err != nil {
		return shim.Error(err.Error())
	} else {
		return shim.Success(result)
	}
}
