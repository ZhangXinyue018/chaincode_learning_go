package feature

import (
	dp "github.com/booksharing/common/domain_with_plugin"
	"github.com/booksharing/feature/account"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("services")

type ServicePack struct {
	AccountService account.AccountService
	// BookService    book.BookService
}

func (sp *ServicePack) Invoke(funcName string, context dp.ContractContext) peer.Response {
	switch funcName {
	case "CreateAccount":
		return sp.AccountService.CreateAccount(context)
	case "GetAccount":
		context.Plugins.IdentityVerifier.VerifyWithLedgerKey(*context.DataContext)
		return sp.AccountService.GetAccount(context)
	case "PingSigVerify":
		context.Plugins.IdentityVerifier.VerifyWithLedgerKey(*context.DataContext)
		return shim.Success([]byte("Success"))
	}
	return shim.Error("error")
}
