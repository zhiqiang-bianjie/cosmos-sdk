package tags

import (
	"github.com/cosmos/cosmos-sdk"
)

// Distribution tx tags
var (
	Rewards    = "rewards"
	Commission = "commission"
	TxCategory = "distribution"

	Validator = sdk.TagSrcValidator
	Category  = sdk.TagCategory
	Sender    = sdk.TagSender
)
