package course

import (
	"encoding/json"

	"github.com/coursera/common"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type CourseRepo struct {
	common.Repository
}

func (repo *CourseRepo) CreateCourse(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	name := args[0]
	course := Course{
		Name:        args[0],
		Description: args[1],
		SchoolName:  args[2],
	}
	repo.Create(APIstub, name, course)
	return shim.Success(nil)
}

func (repo *CourseRepo) GetCourse(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	name := args[0]
	result, err := repo.GetBytes(APIstub, name)
	if err != nil {
		return shim.Error(err.Error())
	} else {
		return shim.Success(result)
	}
}

func (repo *CourseRepo) DataToObject(dataBytes []byte) (Course, error) {
	result := Course{}
	err := json.Unmarshal(dataBytes, &result)
	return result, err
}
