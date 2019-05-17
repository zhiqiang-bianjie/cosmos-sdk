package genaccounts

import (
	"github.com/cosmos/cosmos-sdk"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

// expected account keeper
type AccountKeeper interface {
	NewAccount(sdk.Context, auth.Account) auth.Account
	SetAccount(sdk.Context, auth.Account)
	IterateAccounts(ctx sdk.Context, process func(auth.Account) (stop bool))
}
