package keeper_test

import (
	gocontext "context"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	amm "github.com/jackalLabs/canine-chain/x/amm"
	"github.com/jackalLabs/canine-chain/x/amm/keeper"
	ammtestutil "github.com/jackalLabs/canine-chain/x/amm/testutil"
	"github.com/jackalLabs/canine-chain/x/amm/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc         codec.Codec
	ctx         sdk.Context
	ammKeeper   *keeper.Keeper
	bankKeeper  *ammtestutil.MockBankKeeper
	queryClient types.QueryClient
	msgSrvr     types.MsgServer
}

func (suite *KeeperTestSuite) SetupSuite() {
	suite.reset()
}

func (suite *KeeperTestSuite) reset() {
	ammKeeper, bankKeeper, encCfg, ctx := setupAMMKeeper(suite.T())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, ammKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	coins := sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000000000)))
	err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, coins)
	suite.NoError(err)

	suite.ctx = ctx
	suite.ammKeeper = ammKeeper
	suite.bankKeeper = bankKeeper
	suite.cdc = encCfg.Codec
	suite.queryClient = queryClient
	suite.msgSrvr = keeper.NewMsgServerImpl(*suite.ammKeeper)
}

func setupMsgServer(suite *KeeperTestSuite) (types.MsgServer, keeper.Keeper, gocontext.Context) {
	k := suite.ammKeeper
	amm.InitGenesis(suite.ctx, *k, *types.DefaultGenesis())
	ctx := sdk.WrapSDKContext(suite.ctx)
	return keeper.NewMsgServerImpl(*k), *k, ctx
}
