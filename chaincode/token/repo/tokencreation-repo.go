package repo

import (
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/chaincode_learning_go/chaincode/token/domain"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var TokenCreationRepository *_TokenCreationRepo

func init() {
	TokenCreationRepository = genTokenCreationRepo()
}

type _TokenCreationRepo struct {
	common.BaseRepo
}

func genTokenCreationRepo() *_TokenCreationRepo {
	logger := shim.NewLogger("token-creation-repo")
	logger.SetLevel(shim.LogDebug)
	return &_TokenCreationRepo{common.BaseRepo{
		Logger: logger,
		IndexNames: []common.IndexName{
			{"creator", "token"},
		},
		Factory: &domain.TokenCreationDataFactory{},
		BaseKeyPrefix: "tokencreation",
	}}
}
