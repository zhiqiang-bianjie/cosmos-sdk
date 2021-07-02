package keeper_test

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/cosmos/cosmos-sdk/x/nft/types"
)

func (suite *KeeperTestSuite) TestIssueType() {
	k := suite.app.NFTKeeper
	err := k.IssueType(suite.ctx, testType, testName,
		testSymbol,
		testDescription,
		testMintRestricted,
		testEditRestricted,
		suite.addrs[0],
	)
	suite.Require().NoError(err)

	expect := types.Metadata{
		Type:           testType,
		Name:           testName,
		Symbol:         testSymbol,
		Description:    testDescription,
		MintRestricted: testMintRestricted,
		EditRestricted: testEditRestricted,
	}
	metadata, has := k.GetMetadata(suite.ctx, testType)
	suite.Require().True(has)
	suite.Require().EqualValues(expect, metadata)

	issuer := k.GetTypeIssuer(suite.ctx, testType)
	suite.Require().EqualValues(suite.addrs[0].String(), issuer)
}

func (suite *KeeperTestSuite) TestMintNFT() {
	suite.TestIssueType()

	k := suite.app.NFTKeeper
	data := &testdata.Cat{
		Moniker: "kitty",
		Lives:   5,
	}
	dataWrap, err := codectypes.NewAnyWithValue(data)
	suite.Require().NoError(err)

	expect := types.NFT{
		Type: testType,
		ID:   testNFTId,
		URI:  "",
		Data: dataWrap,
	}
	err = k.MintNFT(suite.ctx,
		testType,
		testNFTId,
		"",
		dataWrap,
		suite.addrs[0],
	)
	suite.Require().NoError(err)
	nft, has := k.GetNFT(suite.ctx, testType, testNFTId)
	suite.Require().True(has)
	suite.Require().True(dataWrap.Compare(nft.Data) == 0)

	expect.Data = nil
	nft.Data = nil
	suite.Require().EqualValues(expect, nft)
}
