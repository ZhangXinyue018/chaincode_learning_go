package plugin

import (
	"github.com/booksharing/common/domain"
	"github.com/booksharing/config"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var logger = shim.NewLogger("identityVerifier")

// IdentityVerifier is to verify account identity
type IdentityVerifier struct{}

// VerifyWithLedgerKey verifies input and input sig using account pubkey
func (iv *IdentityVerifier) VerifyWithLedgerKey(context domain.Context) bool {
	stub := context.Stub
	sig := context.Sig
	account := context.AccountName
	input := context.Input
	pubkey, err := stub.GetState(config.PREFIX_ACCOUNT + account)
	if err != nil {
		logger.Errorf("Cannot find pubkey for account: [%v]", account)
		logger.Error(err.Error())
		return false
	}
	// todo: check sig
	logger.Debugf("input: [%v], sig: [%v], pubkey: [%v]", input, sig, pubkey)
	return true
}

// VerifyWithPubKey verifies input and input sig using defined pubkey, this method is for create account
func (iv *IdentityVerifier) VerifyWithPubKey(context domain.Context, pubkey string) bool {
	sig := context.Sig
	input := context.Input
	// todo: check sig
	logger.Debugf("input: [%v], sig: [%v], pubkey: [%v]", input, sig, pubkey)
	return true
}
