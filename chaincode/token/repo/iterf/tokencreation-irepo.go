package iterf

import (
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type ITokenCreationRepository interface {
	common.IBaseRepo

	CreateToken(stub shim.ChaincodeStubInterface, tokenName string, maxAmount int64, creator, issuer string) error

	UpdateTokenIssueAmount(stub shim.ChaincodeStubInterface, userName, tokenName string, tokenAmount int64) error
}
