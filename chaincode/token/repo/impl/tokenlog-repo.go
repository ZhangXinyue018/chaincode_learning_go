package impl

import (
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/chaincode_learning_go/chaincode/token/domain/obj"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type _TokenLogRepo struct {
	common.BaseRepo
}

func GenTokenLogRepo() *_TokenLogRepo {
	logger := shim.NewLogger("token-log-repo")
	return &_TokenLogRepo{common.BaseRepo{
		Logger: logger,
		IndexNamePackage: common.IndexNamePackage{
			{Indicator: "ByFromUserAndToUser", Names: []string{"from_user_name", "to_user_name", "token_name", "token_log_id"}},
			{Indicator: "ByFromUserAndToken", Names: []string{"from_user_name", "token_name", "to_user_name", "token_log_id"}},
			{Indicator: "ByToUserAndFromUser", Names: []string{"to_user_name", "from_user_name", "token_name", "token_log_id"}},
			{Indicator: "ByToUserAndToken", Names: []string{"to_user_name", "token_name", "from_user_name", "token_log_id"}},
		},
		Factory:       &obj.TokenLogDataFactory{},
		BaseKeyPrefix: "tokenlog",
	}}
}

func (repo *_TokenLogRepo) CreateTokenLog(stub shim.ChaincodeStubInterface, requestId, userName, tokenName string, tokenAmount int64) error {
	tokenLog := &obj.TokenLog{
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
