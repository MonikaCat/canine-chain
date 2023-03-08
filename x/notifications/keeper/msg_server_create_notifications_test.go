package keeper_test

import (
	"github.com/jackalLabs/canine-chain/testutil"
	"github.com/jackalLabs/canine-chain/x/notifications/types"
)

func (suite *KeeperTestSuite) TestMsgCreateNotifications() {
	suite.SetupSuite()
	msgSrvr, _, context := setupMsgServer(suite)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	alice := testAddresses[0]
	bob := testAddresses[1]

	// set noti counter for bob
	notiCounter := types.NotiCounter{
		bob,
		0,
		"",
	}
	suite.Require().NoError(err)
	suite.notificationsKeeper.SetNotiCounter(suite.ctx, notiCounter)

	cases := []struct {
		preRun    func() *types.MsgCreateNotifications
		expErr    bool
		expErrMsg string
		name      string
	}{
		{ // alice sends a notification to bob
			preRun: func() *types.MsgCreateNotifications {
				return types.NewMsgCreateNotifications(
					alice,
					"hey bob it's alice here",
					bob,
				)
			},
			expErr: false,
			name:   "alice successfully sends a notification to bob",
		},
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			msg := tc.preRun()
			suite.Require().NoError(err)
			res, err := msgSrvr.CreateNotifications(context, msg)
			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {

				suite.Require().NoError(err)
				suite.Require().EqualValues(types.MsgCreateNotificationsResponse{}, *res)

			}
		})
	}
}
