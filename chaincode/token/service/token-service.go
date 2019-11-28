package service

import (
	"github.com/chaincode_learning_go/chaincode/token/domain"
	"github.com/chaincode_learning_go/chaincode/token/repo"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func Ping(stub shim.ChaincodeStubInterface) peer.Response {
	repo.TokenCreationRepository.Logger.Debug("hello ping")
	return shim.Success([]byte("Pong"))
}

func CreateToken(stub shim.ChaincodeStubInterface, requestId string, token *domain.TokenCreation) peer.Response {
	err := repo.TokenCreationRepository.Create(stub, token)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("ok"))
}

func IssueToken(stub shim.ChaincodeStubInterface, requestId string, userName string, tokenType string, tokenAmount int64) peer.Response {
	tokenCreationPrimaryKey := domain.GetTokenCreationPrimaryKey(tokenType)
	token, err := repo.TokenCreationRepository.GetByPrimaryKey(stub, tokenCreationPrimaryKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	if token == nil {
		return shim.Error("token not exist")
	}

	tokenEntity := token.(*domain.TokenCreation)
	if tokenEntity.Issuer != userName {
		return shim.Error("you are not allowed to issue token")
	}
	if tokenEntity.CurrentAmount+tokenAmount > tokenEntity.MaxAmount {
		return shim.Error("amount exceed")
	}

	tokenEntity.CurrentAmount = tokenEntity.CurrentAmount + tokenAmount
	err = repo.TokenCreationRepository.Update(stub, tokenEntity.GetPrimaryKey(), tokenEntity)
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
