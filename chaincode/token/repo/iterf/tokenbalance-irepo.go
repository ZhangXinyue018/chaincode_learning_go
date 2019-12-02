package iterf

import (
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/chaincode_learning_go/chaincode/token/domain/obj"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type ITokenBalanceRepository interface {
	common.IBaseRepo

	GetBalance(stub shim.ChaincodeStubInterface, userName, tokenName string) (*obj.TokenBalance, error)

	ListBalanceByUserName(stub shim.ChaincodeStubInterface, query []string) ([]*obj.TokenBalance, error)

	PaginateBalanceByUserName(stub shim.ChaincodeStubInterface, query []string, pageSize int32,
		bookMark string) ([]*obj.TokenBalance, string, error)

	ListBalanceByTokenName(stub shim.ChaincodeStubInterface, query []string) ([]*obj.TokenBalance, error)

	PaginateBalanceByTokenName(stub shim.ChaincodeStubInterface, query []string, pageSize int32,
		bookMark string) ([]*obj.TokenBalance, string, error)

	AddBalance(stub shim.ChaincodeStubInterface, userName, tokenName string, amount int64) error

	DeductBalance(stub shim.ChaincodeStubInterface, userName, tokenName string, amount int64) error
}
