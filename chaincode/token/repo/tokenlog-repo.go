package repo

import (
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/chaincode_learning_go/chaincode/token/domain"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var TokenLogRepository *_TokenLogRepo

func init() {
	TokenLogRepository = genTokenLogRepo()
}

type _TokenLogRepo struct {
	common.BaseRepo
}

func genTokenLogRepo() *_TokenLogRepo {
	logger := shim.NewLogger("token-log-repo")
	logger.SetLevel(shim.LogDebug)
	return &_TokenLogRepo{common.BaseRepo{
		Logger: logger,
		IndexNames: []common.IndexName{
			{"token", "user"}},
		Factory: &domain.TokenLogDataFactory{},
		BaseKeyPrefix: "tokenlog",
	}}
}
