package repo

import (
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/chaincode_learning_go/chaincode/token/domain"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var TokenBalanceRepository *_TokenBalanceRepo

func init() {
	TokenBalanceRepository = genTokenBalanceRepo()
}

type _TokenBalanceRepo struct {
	common.BaseRepo
}

func genTokenBalanceRepo() *_TokenBalanceRepo {
	logger := shim.NewLogger("user-token-repo")
	logger.SetLevel(shim.LogDebug)
	return &_TokenBalanceRepo{common.BaseRepo{
		Logger: logger,
		IndexNames: []common.IndexName{
			{"fromuser", "touser"},
			{"fromuser", "token"},
			{"touser", "fromuser"},
			{"touser", "token"},
		},
		Factory: &domain.TokenBalanceDataFactory{},
		BaseKeyPrefix: "tokenbalance",
	}}
}
