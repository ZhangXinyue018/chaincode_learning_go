package repo

import (
	"github.com/chaincode_learning_go/chaincode/token/repo/impl"
	"github.com/chaincode_learning_go/chaincode/token/repo/iterf"
	"sync"
)

var TokenBalanceRepository iterf.ITokenBalanceRepository

var TokenCreationRepository iterf.ITokenCreationRepository

var TokenLogRepository iterf.ITokenLogRepository

var once sync.Once

func init() {
	once.Do(func() {
		TokenBalanceRepository = impl.GenTokenBalanceRepo()
		TokenCreationRepository = impl.GenTokenCreationRepo()
		TokenLogRepository = impl.GenTokenLogRepo()
	})
}
