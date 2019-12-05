package iterf

import (
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/chaincode_learning_go/chaincode/token/domain/obj"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type ITokenLogRepository interface {
	common.IBaseRepo

	CreateTokenLog(stub shim.ChaincodeStubInterface, requestId, fromUserName, toUserName, tokenName string, tokenAmount int64) error

	PaginateTokenLogByFromUserName(stub shim.ChaincodeStubInterface, query []string, pageSize int32, bookMark string) ([]*obj.TokenLog, string, error)

	PaginateTokenLogByToUserName(stub shim.ChaincodeStubInterface, query []string, pageSize int32, bookMark string) ([]*obj.TokenLog, string, error)
}
