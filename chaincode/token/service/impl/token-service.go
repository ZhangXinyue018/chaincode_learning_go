package impl

import (
	"github.com/chaincode_learning_go/chaincode/token/repo/iterf"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type TokenService struct {
	TokenCreationRepository iterf.ITokenCreationRepository
	TokenLogRepository      iterf.ITokenLogRepository
	TokenBalanceRepository  iterf.ITokenBalanceRepository
}

func (service *TokenService) Ping(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("Pong"))
}

func (service *TokenService) CreateToken(stub shim.ChaincodeStubInterface, tokenName string, maxAmount int64, creator, issuer string) peer.Response {
	err := service.TokenCreationRepository.CreateToken(stub, tokenName, maxAmount, creator, issuer)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("ok"))
}

func (service *TokenService) IssueToken(stub shim.ChaincodeStubInterface, requestId string, userName string, tokenName string, tokenAmount int64) peer.Response {
	err := service.TokenCreationRepository.UpdateTokenIssueAmount(stub, userName, tokenName, tokenAmount)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = service.TokenBalanceRepository.AddBalance(stub, userName, tokenName, tokenAmount)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = service.TokenLogRepository.CreateTokenLog(stub, requestId, userName, tokenName, tokenAmount)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("ok"))
}

func (service *TokenService) TransferToken(stub shim.ChaincodeStubInterface, requestId, fromUserName, toUserName, tokenName string,
	tokenAmount int64) peer.Response {
	err := service.TokenBalanceRepository.DeductBalance(stub, fromUserName, tokenName, tokenAmount)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = service.TokenBalanceRepository.AddBalance(stub, toUserName, tokenName, tokenAmount)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = service.TokenLogRepository.CreateTokenLog(stub, requestId, fromUserName, tokenName, -tokenAmount)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = service.TokenLogRepository.CreateTokenLog(stub, requestId, toUserName, tokenName, tokenAmount)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("ok"))
}

func (service *TokenService) GetToken(stub shim.ChaincodeStubInterface, userName, tokenName string) peer.Response {
	userBalance, err := service.TokenBalanceRepository.GetBalance(stub, userName, tokenName)
	if err != nil {
		return shim.Error(err.Error())
	}
	result, err := userBalance.ToBytes()
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(result)
}

func (service *TokenService) PaginateTokenByUserName(stub shim.ChaincodeStubInterface, query []string, pageSize int32,
	bookMark string) peer.Response {
	result, err := service.TokenBalanceRepository.PaginateBalanceByUserName(stub, query, pageSize, bookMark)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(result.ToBytes())
}

func (service *TokenService) PaginateTokenByTokenName(stub shim.ChaincodeStubInterface, query []string, pageSize int32,
	bookMark string) peer.Response {
	result, err := service.TokenBalanceRepository.PaginateBalanceByTokenName(stub, query, pageSize, bookMark)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(result.ToBytes())
}
