package account

type Account struct {
	Name      string
	PublicKey string
}

type GetAccountReq struct {
	Name string
}

type CreateAccountReq Account
