package service

import (
	"github.com/chaincode_learning_go/chaincode/token/repo"
	"github.com/chaincode_learning_go/chaincode/token/service/impl"
	"github.com/chaincode_learning_go/chaincode/token/service/iterf"
)

var TokenService iterf.ITokenService
var TokenLogService iterf.ITokenLogService

func init() {
	TokenService = &impl.TokenService{
		TokenCreationRepository: repo.TokenCreationRepository,
		TokenBalanceRepository:  repo.TokenBalanceRepository,
		TokenLogRepository:      repo.TokenLogRepository,
	}

	TokenLogService = &impl.TokenLogService{
		TokenLogRepository: repo.TokenLogRepository,
	}
}
