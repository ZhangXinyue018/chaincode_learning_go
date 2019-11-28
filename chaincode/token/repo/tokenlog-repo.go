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
	return &_TokenLogRepo{common.BaseRepo{
		Logger: logger,
		IndexNames: []common.IndexName{
			{"from_user_name", "to_user_name", "token_name", "token_log_id"},
			{"from_user_name", "token_name", "to_user_name", "token_log_id"},
			{"to_user_name", "from_user_name", "token_name", "token_log_id"},
			{"to_user_name", "token_name", "from_user_name", "token_log_id"},
		},
		Factory:       &domain.TokenLogDataFactory{},
		BaseKeyPrefix: "tokenlog",
	}}
}

func (repo *_TokenLogRepo) CreateTokenLog(stub shim.ChaincodeStubInterface, requestId, userName, tokenName string, tokenAmount int64) error {
	tokenLog := &domain.TokenLog{
		TokenLogId:   requestId,
		FromUserName: "0",
		ToUserName:   userName,
		TokenName:    tokenName,
		TokenAmount:  tokenAmount,
		Comment:      "issue token",
	}
	err := repo.Create(stub, tokenLog)
	if err != nil {
		repo.Logger.Error(err.Error())
		return err
	}
	return nil
}
