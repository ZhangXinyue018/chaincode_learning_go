package service

import (
	"github.com/chaincode_learning_go/chaincode/token/repo"
	"github.com/chaincode_learning_go/chaincode/token/service/impl"
	"github.com/chaincode_learning_go/chaincode/token/service/iterf"
	"sync"
)

var TokenService iterf.ITokenService
var TokenLogService iterf.ITokenLogService
var once sync.Once

func init() {
	once.Do(func() {
		TokenService = &impl.TokenService{
			TokenCreationRepository: repo.TokenCreationRepository,
			TokenBalanceRepository:  repo.TokenBalanceRepository,
			TokenLogRepository:      repo.TokenLogRepository,
		}

		TokenLogService = &impl.TokenLogService{
			TokenLogRepository: repo.TokenLogRepository,
		}
	})
}
