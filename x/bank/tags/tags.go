package tags

import (
	"github.com/cosmos/cosmos-sdk"
)

// Tag keys and values
var (
	ActionUndelegateCoins = "undelegateCoins"
	ActionDelegateCoins   = "delegateCoins"
	TxCategory            = "bank"

	Action    = sdk.TagAction
	Category  = sdk.TagCategory
	Recipient = "recipient"
	Sender    = "sender"
)
