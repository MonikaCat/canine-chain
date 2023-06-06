package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgSignContract_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSignContract
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSignContract{
				Creator: "invalid_address",
				Fid:     "jklf129c5wfsdex0zcpfnzarjhf7ks9nrrspjgqlmv7e6lmcz4z8ecacsrwz7hu",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid cid",
			msg: MsgSignContract{
				Creator: "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Fid:     "invalid_cid",
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "valid address",
			msg: MsgSignContract{
				Creator: "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Fid:     "jklf129c5wfsdex0zcpfnzarjhf7ks9nrrspjgqlmv7e6lmcz4z8ecacsrwz7hu",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
