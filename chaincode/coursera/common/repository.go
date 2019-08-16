package common

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// ----------
// Basic Ledger actions
// ----------

type Repository struct {
	Prefix RepoPrefix
}

func (repo *Repository) Create(APIstub shim.ChaincodeStubInterface, name string, value interface{}) error {
	dataBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return APIstub.PutState(repo.genStateKey(name), dataBytes)
}

// todo: change value based on key name
func (repo *Repository) Update(APIstub shim.ChaincodeStubInterface, name, key string, value interface{}) error {
	// 	m := make(map[string]string)
	// err := json.Unmarshal(input, &m)

	return nil
}

func (repo *Repository) Delete(APIstub shim.ChaincodeStubInterface, name string) error {
	return APIstub.DelState(repo.genStateKey(name))
}

func (repo *Repository) GetBytes(APIstub shim.ChaincodeStubInterface, name string) ([]byte, error) {
	return APIstub.GetState(repo.genStateKey(name))
}

func (repo *Repository) genStateKey(name string) string {
	return string(repo.Prefix) + "_" + name
}
