package iterf

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type ITokenService interface {
	Ping(stub shim.ChaincodeStubInterface) peer.Response

	CreateToken(stub shim.ChaincodeStubInterface, tokenName string, maxAmount int64,
		creator, issuer string) peer.Response

	IssueToken(stub shim.ChaincodeStubInterface, requestId, userName, tokenName string, tokenAmount int64) peer.Response

	TransferToken(stub shim.ChaincodeStubInterface, requestId, fromUserName, toUserName, tokenName string,
		tokenAmount int64) peer.Response

	GetToken(stub shim.ChaincodeStubInterface, userName, tokenName string) peer.Response

	PaginateTokenByUserName(stub shim.ChaincodeStubInterface, query []string, pageSize int32,
		bookMark string) peer.Response

	PaginateTokenByTokenName(stub shim.ChaincodeStubInterface, query []string, pageSize int32,
		bookMark string) peer.Response
}
