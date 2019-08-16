package account

import (
	dp "github.com/booksharing/common/domain_with_plugin"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("AccountService")

type AccountService struct {
	AccountLedger AccountLedger
}

func (service *AccountService) CreateAccount(context dp.ContractContext) peer.Response {
	// step1: check sig
	createAccountReq := CreateAccountReq{}
	context.Plugins.LedgerDataProcessor.UnMarshal(context.DataContext.Input, &createAccountReq)
	context.DataContext.AccountName = createAccountReq.Name
	context.Plugins.IdentityVerifier.VerifyWithPubKey(*context.DataContext, createAccountReq.PublicKey)

	// step2: check account exist
	account := Account(createAccountReq)
	exist, err := service.AccountLedger.CheckExist(context, account)
	if exist {
		return shim.Error("Account Already exist")
	} else if err != nil {
		return shim.Error(err.Error())
	}

	// step3: create account
	err = service.AccountLedger.Create(context, account)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("success"))
}

func (service *AccountService) GetAccount(context dp.ContractContext) peer.Response {
	return shim.Success([]byte("success"))
}
