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

func IssueToken(stub shim.ChaincodeStubInterface, requestId string, userName string, tokenName string, tokenAmount int64) peer.Response {
	err := repo.TokenCreationRepository.UpdateTokenIssueAmount(stub, userName, tokenName, tokenAmount)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = repo.TokenBalanceRepository.AddBalance(stub, userName, tokenName, tokenAmount)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = repo.TokenLogRepository.CreateTokenLog(stub, requestId, userName, tokenName, tokenAmount)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("ok"))
}

func TransferToken(stub shim.ChaincodeStubInterface, requestId, fromUserName, toUserName, tokenName string,
	tokenAmount int64) peer.Response {
	err := repo.TokenBalanceRepository.DeductBalance(stub, fromUserName, tokenName, tokenAmount)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = repo.TokenBalanceRepository.AddBalance(stub, toUserName, tokenName, tokenAmount)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = repo.TokenLogRepository.CreateTokenLog(stub, requestId, fromUserName, tokenName, -tokenAmount)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = repo.TokenLogRepository.CreateTokenLog(stub, requestId, toUserName, tokenName, tokenAmount)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("ok"))
}

func GetToken(stub shim.ChaincodeStubInterface, userName, tokenName string) peer.Response {
	userBalancePrimaryKey := domain.GetTokenBalancePrimaryKey(userName, tokenName)
	userBalance, err := repo.TokenBalanceRepository.GetByPrimaryKey(stub, userBalancePrimaryKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	if userBalance == nil {
		userBalance = domain.GetDefaultTokenBalance(userName, tokenName)
	}
	result, err := userBalance.(*domain.TokenBalance).ToBytes()
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(result)
}
