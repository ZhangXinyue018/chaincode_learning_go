package service

import (
	"github.com/chaincode_learning_go/chaincode/token/repo"
	"github.com/chaincode_learning_go/chaincode/token/service/impl"
	"github.com/chaincode_learning_go/chaincode/token/service/iterf"
)

var TokenService iterf.ITokenService

func init() {
	TokenService = &impl.TokenService{
		TokenCreationRepository: repo.TokenCreationRepository,
		TokenBalanceRepository:  repo.TokenBalanceRepository,
		TokenLogRepository:      repo.TokenLogRepository,
	}
}
