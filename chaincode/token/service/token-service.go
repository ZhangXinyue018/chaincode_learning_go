package service

import (
	"github.com/chaincode_learning_go/chaincode/token/domain"
	"github.com/chaincode_learning_go/chaincode/token/repo"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func Ping(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("Pong"))
}

func CreateToken(stub shim.ChaincodeStubInterface, requestId string, token *domain.Token) peer.Response {
	err := repo.UserTokenRepository.Create(stub, token)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("ok"))
}

func IssueToken(stub shim.ChaincodeStubInterface, requestId string, userName string, tokenType string, tokenAmount int64) peer.Response {
	token, err := repo.TokenRepository.GetByBaseKey(stub, tokenType)
	if err != nil {
		return shim.Error(err.Error())
	}
	if token == nil {
		return shim.Error("token not exist")
	}

	tokenEntity := token.(*domain.Token)
	if tokenEntity.Issuer != userName {
		return shim.Error("you are not allowed to issue token")
	}
	if tokenEntity.CurrentAmount+tokenAmount > tokenEntity.MaxAmount {
		return shim.Error("amount exceed")
	}

	tokenEntity.CurrentAmount = tokenEntity.CurrentAmount + tokenAmount
	err = repo.TokenRepository.Update(stub, tokenEntity.GetBaseKey(), tokenEntity)
	if err != nil {
		return shim.Error(err.Error())
	}

	tokenLog := &domain.TokenLog{
		TokenLogId:   requestId,
		FromUserName: "0",
		ToUserName:   tokenEntity.Issuer,
		TokenName:    tokenEntity.TokenName,
		TokenAmount:  tokenAmount,
		Comment:      "issue token",
	}
	err = repo.TokenLogRepository.Create(stub, tokenLog)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("ok"))
}

//func TransferToken(stub shim.ChaincodeStubInterface, requestId string,) {
//	tokenLog := &domain.TokenLog{
//		TokenLogId:   requestId,
//		FromUserName: "0",
//		ToUserName:   tokenEntity.Issuer,
//		TokenName:    tokenEntity.TokenName,
//		TokenAmount:  tokenAmount,
//		Comment:      "issue token",
//	}
//	err = repo.TokenLogRepository.Create(stub, tokenLog)
//	if err != nil {
//		return shim.Error(err.Error())
//	}
//
//	return shim.Success([]byte("ok"))
//
//}
//
//func GetToken(stub shim.ChaincodeStubInterface) {
//
//}
