package iterf

import (
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type ITokenLogRepository interface {
	common.IBaseRepo

	CreateTokenLog(stub shim.ChaincodeStubInterface, requestId, userName, tokenName string, tokenAmount int64) error
}
