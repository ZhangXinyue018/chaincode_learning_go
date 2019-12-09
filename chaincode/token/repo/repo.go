package repo

import (
	"github.com/chaincode_learning_go/chaincode/token/repo/impl"
	"github.com/chaincode_learning_go/chaincode/token/repo/iterf"
)

var TokenBalanceRepository iterf.ITokenBalanceRepository

var TokenCreationRepository iterf.ITokenCreationRepository

var TokenLogRepository iterf.ITokenLogRepository

func init() {
	TokenBalanceRepository = impl.GenTokenBalanceRepo()
	TokenCreationRepository = impl.GenTokenCreationRepo()
	TokenLogRepository = impl.GenTokenLogRepo()
}
