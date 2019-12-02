package common

import "github.com/hyperledger/fabric/core/chaincode/shim"

type IBaseRepo interface {
	Create(stub shim.ChaincodeStubInterface, entity Data) error

	GetByPrimaryKey(stub shim.ChaincodeStubInterface, primaryKey string) (Data, error)

	ListByIndexKey(stub shim.ChaincodeStubInterface, indexKey IndexSearchKey) ([]Data, error)

	PaginateByIndexKey(stub shim.ChaincodeStubInterface, indexKey IndexSearchKey, pageSize int32,
		bookMark string) (*PaginationResponse, error)

	UpdateWithEntity(stub shim.ChaincodeStubInterface, originalEntity Data, entity Data) error

	UpdateWithPrimaryKey(stub shim.ChaincodeStubInterface, primaryKey string, entity Data) error

	Delete(stub shim.ChaincodeStubInterface, primaryKey string) (Data, error)
}
