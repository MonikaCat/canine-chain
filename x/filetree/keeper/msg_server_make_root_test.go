package keeper_test

import (
	"github.com/MonikaCat/canine-chain/v2/testutil"
	"github.com/MonikaCat/canine-chain/v2/x/filetree/types"
)

func (suite *KeeperTestSuite) TestMsgMakeRoot() {
	suite.SetupSuite()
	msgSrvr, _, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 1)
	suite.Require().NoError(err)

	alice := testAddresses[0]

	msg, err := types.CreateMsgMakeRoot(alice)
	suite.Require().NoError(err)

	cases := []struct {
		preRun    func() *types.MsgMakeRoot
		expErr    bool
		expErrMsg string
		name      string
	}{
		{
			preRun: func() *types.MsgMakeRoot {
				return msg
			},
			expErr: false,
			name:   "make root success",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.MakeRoot(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgMakeRootResponse{}, *res)

			}
		})
	}
}
