package account

import (
	"errors"

	dp "github.com/booksharing/common/domain_with_plugin"
	"github.com/booksharing/config"
)

type AccountLedger struct {
	Prefix string
}

func (al *AccountLedger) Create(context dp.ContractContext, account Account) error {
	dataBytes, err := context.Plugins.LedgerDataProcessor.Marshal(account)
	if err != nil {
		return err
	}
	return context.DataContext.Stub.PutState(al.genStateKey(account.Name), dataBytes)
}

func (al *AccountLedger) Update(context dp.ContractContext, account Account) error {
	dataBytes, err := context.Plugins.LedgerDataProcessor.Marshal(account)
	if err != nil {
		return err
	}
	context.DataContext.Stub.PutState(al.genStateKey(account.Name), dataBytes)
	return nil
}

func (al *AccountLedger) Delete(context dp.ContractContext, account Account) error {
	return context.DataContext.Stub.DelState(al.genStateKey(account.Name))
}

func (al *AccountLedger) CheckExist(context dp.ContractContext, account Account) (bool, error) {
	dataBytes, err := context.DataContext.Stub.GetState(al.genStateKey(account.Name))
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}
	return dataBytes != nil, nil
}

func (al *AccountLedger) Get(context dp.ContractContext, account Account) (Account, error) {
	dataBytes, err := context.DataContext.Stub.GetState(al.genStateKey(account.Name))
	if err != nil {
		return Account{}, err
	}
	if dataBytes == nil {
		return Account{}, errors.New("account not exist")
	}
	finalAccount := Account{}
	context.Plugins.LedgerDataProcessor.UnMarshal(dataBytes, &finalAccount)
	if err != nil {
		return Account{}, err
	}
	return finalAccount, nil
}

func (al *AccountLedger) genStateKey(name string) string {
	return config.PREFIX_ACCOUNT + name
}
