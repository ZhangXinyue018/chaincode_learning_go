package iterf

import (
	"github.com/chaincode_learning_go/chaincode/token/domain/protos"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type ITokenService interface {
	Ping(stub shim.ChaincodeStubInterface, request *protos.PingRequest) *protos.PingResponse

	CreateToken(stub shim.ChaincodeStubInterface, request *protos.CreateTokenRequest) *protos.CreateTokenResponse

	IssueToken(stub shim.ChaincodeStubInterface, request *protos.IssueTokenRequest) *protos.IssueTokenResponse

	TransferToken(stub shim.ChaincodeStubInterface, request *protos.TransferTokenRequest) *protos.TransferTokenResponse

	GetToken(stub shim.ChaincodeStubInterface, request *protos.GetTokenRequest) *protos.GetTokenResponse

	PaginateTokenByUserName(stub shim.ChaincodeStubInterface, request *protos.PaginateTokenByUserNameRequest) *protos.PaginateTokenByUserNameResponse

	PaginateTokenByTokenName(stub shim.ChaincodeStubInterface, request *protos.PaginateTokenByTokenNameRequest) *protos.PaginateTokenByTokenNameResponse
}
