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
			{Indicator: "ByFromUser", Names: []string{"from_user_name", "timestamp", "to_user_name", "token_name", "token_log_id"}},
			{Indicator: "ByFromUserAndToken", Names: []string{"from_user_name", "token_name", "timestamp", "to_user_name", "token_log_id"}},
			{Indicator: "ByToUser", Names: []string{"to_user_name", "timestamp", "from_user_name", "token_name", "token_log_id"}},
			{Indicator: "ByToUserAndToken", Names: []string{"to_user_name", "token_name", "timestamp", "from_user_name", "token_log_id"}},
		},
		Factory:       &obj.TokenLogDataFactory{},
		BaseKeyPrefix: "tokenlog",
	}}
}

func (repo *_TokenLogRepo) CreateTokenLog(stub shim.ChaincodeStubInterface, requestId, fromUserName, toUserName,
	tokenName string, tokenAmount int64) error {
	currTime, err := stub.GetTxTimestamp()
	if err != nil {
		repo.Logger.Error(err.Error())
		return err
	}
	tokenLog := &obj.TokenLog{
		TokenLogId:   requestId,
		FromUserName: fromUserName,
		ToUserName:   toUserName,
		TokenName:    tokenName,
		TokenAmount:  tokenAmount,
		Comment:      "issue token",
		Timestamp:    currTime.Seconds,
	}
	err = repo.Create(stub, tokenLog)
	if err != nil {
		repo.Logger.Error(err.Error())
		return err
	}
	return nil
}

func (repo *_TokenLogRepo) PaginateTokenLogByFromUserName(stub shim.ChaincodeStubInterface, query []string,
	pageSize int32, bookMark string) ([]*obj.TokenLog, string, error) {
	indexSearchKey := common.IndexSearchKey{Indicator: "ByFromUser", Keys: query}
	resultList, newBookMark, err := repo.PaginateByIndexKey(stub, indexSearchKey, pageSize, bookMark)
	if err != nil {
		return nil, "", err
	}
	finalResultList := make([]*obj.TokenLog, 0)
	for _, result := range resultList {
		finalResultList = append(finalResultList, result.(*obj.TokenLog))
	}
	return finalResultList, newBookMark, nil
}

func (repo *_TokenLogRepo) PaginateTokenLogByToUserName(stub shim.ChaincodeStubInterface, query []string,
	pageSize int32, bookMark string) ([]*obj.TokenLog, string, error) {
	indexSearchKey := common.IndexSearchKey{Indicator: "ByToUser", Keys: query}
	resultList, newBookMark, err := repo.PaginateByIndexKey(stub, indexSearchKey, pageSize, bookMark)
	if err != nil {
		return nil, "", err
	}
	finalResultList := make([]*obj.TokenLog, 0)
	for _, result := range resultList {
		finalResultList = append(finalResultList, result.(*obj.TokenLog))
	}
	return finalResultList, newBookMark, nil
}
