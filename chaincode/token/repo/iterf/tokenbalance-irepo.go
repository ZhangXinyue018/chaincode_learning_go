package iterf

import (
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/chaincode_learning_go/chaincode/token/domain"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type ITokenBalanceRepository interface {
	common.IBaseRepo

	GetBalance(stub shim.ChaincodeStubInterface, userName, tokenName string) (*domain.TokenBalance, error)

	ListBalanceByUserName(stub shim.ChaincodeStubInterface, query []string) ([]*domain.TokenBalance, error)

	PaginateBalanceByUserName(stub shim.ChaincodeStubInterface, query []string, pageSize int32,
		bookMark string) (*common.PaginationResponse, error)

	ListBalanceByTokenName(stub shim.ChaincodeStubInterface, query []string) ([]*domain.TokenBalance, error)

	PaginateBalanceByTokenName(stub shim.ChaincodeStubInterface, query []string, pageSize int32,
		bookMark string) (*common.PaginationResponse, error)

	AddBalance(stub shim.ChaincodeStubInterface, userName, tokenName string, amount int64) error

	DeductBalance(stub shim.ChaincodeStubInterface, userName, tokenName string, amount int64) error
}
