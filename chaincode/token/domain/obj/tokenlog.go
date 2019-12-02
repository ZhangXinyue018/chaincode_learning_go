package obj

import (
	"fmt"
	"github.com/chaincode_learning_go/chaincode/common"
	"github.com/chaincode_learning_go/chaincode/lib"
	"github.com/chaincode_learning_go/chaincode/token/domain/protos/obj"
	"strconv"
)

type TokenLog protos.TokenLogPB

func (tokenlog *TokenLog) ToStoreBytes() ([]byte, error) {
	return lib.ToJsonBytes(tokenlog)
}

func (tokenlog *TokenLog) ToString() string {
	return lib.ToJsonString(tokenlog)
}

func (tokenlog *TokenLog) ToMap() map[string]string {
	return map[string]string{
		"token_log_id":   tokenlog.TokenLogId,
		"from_user_name": tokenlog.FromUserName,
		"to_user_name":   tokenlog.ToUserName,
		"token_name":     tokenlog.TokenName,
		"token_amount":   strconv.FormatInt(tokenlog.TokenAmount, 10),
		"comment":        tokenlog.Comment,
	}
}

func (tokenlog *TokenLog) GetPrimaryKey() string {
	return GetTokenLogPrimaryKey(tokenlog.TokenLogId)
}

func GetTokenLogPrimaryKey(logId string) string {
	return fmt.Sprintf("%s", logId)
}

type TokenLogDataFactory struct {
}

func (factory *TokenLogDataFactory) ToDataEntityFromStoredBytes(data []byte) (common.Data, error) {
	if data == nil {
		return nil, nil
	}
	result := &TokenLog{}
	err := lib.FromJsonBytes(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
