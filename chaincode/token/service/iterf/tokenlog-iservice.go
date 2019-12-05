package iterf

import (
	"github.com/chaincode_learning_go/chaincode/token/domain/protos"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type ITokenLogService interface {
	PaginateTokenLogByFromUserName(stub shim.ChaincodeStubInterface, request *protos.PaginateTokenLogByFromUserNameRequest) *protos.PaginateTokenLogByFromUserNameResponse

	PaginateTokenLogByToUserName(stub shim.ChaincodeStubInterface, request *protos.PaginateTokenLogByToUserNameRequest) *protos.PaginateTokenLogByToUserNameResponse
}
