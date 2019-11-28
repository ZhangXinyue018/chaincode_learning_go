package repo

import (
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/chaincode_learning_go/chaincode/token/domain"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var UserTokenRepository *_UserTokenRepo

func init() {
	UserTokenRepository = genUserTokenRepo()
}

type _UserTokenRepo struct {
	common.BaseRepo
}

func genUserTokenRepo() *_UserTokenRepo {
	logger := shim.NewLogger("user-token-repo")
	logger.SetLevel(shim.LogDebug)
	return &_UserTokenRepo{common.BaseRepo{
		Logger: logger,
		IndexNames: []common.IndexName{
			{"fromuser", "touser"},
			{"fromuser", "token"},
			{"touser", "fromuser"},
			{"touser", "token"},
		},
		Factory: &domain.UserTokenDataFactory{},
	}}
}
