package repo

import (
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/chaincode_learning_go/chaincode/token/domain"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var TokenRepository *_TokenRepo

func init() {
	TokenRepository = genTokenRepo()
}

type _TokenRepo struct {
	common.BaseRepo
}

func genTokenRepo() *_TokenRepo {
	logger := shim.NewLogger("token-creation-repo")
	logger.SetLevel(shim.LogDebug)
	return &_TokenRepo{common.BaseRepo{
		Logger: logger,
		IndexNames: []common.IndexName{
			{"creator", "token"},
		},
		Factory: &domain.UserTokenDataFactory{},
	}}
}
