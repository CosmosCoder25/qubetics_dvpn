package ibc

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	ibctesting "github.com/cosmos/ibc-go/v7/testing"
	teststypes "github.com/qubetics/qubetics-blockchain/v2/types/tests"
)

func init() {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("qubetics", "qubeticspub")
}

func TestGetTransferSenderRecipient(t *testing.T) {
	testCases := []struct {
		name         string
		data         transfertypes.FungibleTokenPacketData
		expSender    string
		expRecipient string
		expError     bool
	}{
		{
			name:         "empty FungibleTokenPacketData",
			data:         transfertypes.FungibleTokenPacketData{},
			expSender:    "",
			expRecipient: "",
			expError:     true,
		},
		{
			name: "invalid sender",
			data: transfertypes.FungibleTokenPacketData{
				Sender:   "cosmos1",
				Receiver: "qubetics16awr0rsqy72rgaydjv0fvuvzfdynshf62kkd8h",
				Amount:   "123456",
			},
			expSender:    "",
			expRecipient: "",
			expError:     true,
		},
		{
			name: "invalid recipient",
			data: transfertypes.FungibleTokenPacketData{
				Sender:   "cosmos1qql8ag4cluz6r4dz28p3w00dnc9w8ueulg2gmc",
				Receiver: "qubetics1",
				Amount:   "123456",
			},
			expSender:    "",
			expRecipient: "",
			expError:     true,
		},
		{
			name: "valid - cosmos sender, qubetics recipient",
			data: transfertypes.FungibleTokenPacketData{
				Sender:   "cosmos1qql8ag4cluz6r4dz28p3w00dnc9w8ueulg2gmc",
				Receiver: "qubetics16awr0rsqy72rgaydjv0fvuvzfdynshf62kkd8h",
				Amount:   "123456",
			},
			expSender:    "qubetics1pvpljz6tzyyg42c33vhtnvynsnfl6q38w6pyue",
			expRecipient: "qubetics16awr0rsqy72rgaydjv0fvuvzfdynshf62kkd8h",
			expError:     false,
		},
		{
			name: "valid - qubetics sender, cosmos recipient",
			data: transfertypes.FungibleTokenPacketData{
				Sender:   "qubetics16awr0rsqy72rgaydjv0fvuvzfdynshf62kkd8h",
				Receiver: "cosmos1qql8ag4cluz6r4dz28p3w00dnc9w8ueulg2gmc",
				Amount:   "123456",
			},
			expSender:    "qubetics16awr0rsqy72rgaydjv0fvuvzfdynshf62kkd8h",
			expRecipient: "qubetics1pvpljz6tzyyg42c33vhtnvynsnfl6q38w6pyue",
			expError:     false,
		},
		{
			name: "valid - osmosis sender, qubetics recipient",
			data: transfertypes.FungibleTokenPacketData{
				Sender:   "osmo1qql8ag4cluz6r4dz28p3w00dnc9w8ueuhnecd2",
				Receiver: "qubetics16awr0rsqy72rgaydjv0fvuvzfdynshf62kkd8h",
				Amount:   "123456",
			},
			expSender:    "qubetics1pvpljz6tzyyg42c33vhtnvynsnfl6q38w6pyue",
			expRecipient: "qubetics16awr0rsqy72rgaydjv0fvuvzfdynshf62kkd8h",
			expError:     false,
		},
	}

	for _, tc := range testCases {
		sender, recipient, _, _, err := GetTransferSenderRecipient(tc.data)
		if tc.expError {
			require.Error(t, err, tc.name)
		} else {
			require.NoError(t, err, tc.name)
			require.Equal(t, tc.expSender, sender.String())
			require.Equal(t, tc.expRecipient, recipient.String())
		}
	}
}

func TestGetTransferAmount(t *testing.T) {
	testCases := []struct {
		name      string
		packet    channeltypes.Packet
		expAmount string
		expError  bool
	}{
		{
			name:      "empty packet",
			packet:    channeltypes.Packet{},
			expAmount: "",
			expError:  true,
		},
		{
			name:      "invalid packet data",
			packet:    channeltypes.Packet{Data: ibctesting.MockFailPacketData},
			expAmount: "",
			expError:  true,
		},
		{
			name: "invalid amount - empty",
			packet: channeltypes.Packet{
				Data: transfertypes.ModuleCdc.MustMarshalJSON(
					&transfertypes.FungibleTokenPacketData{
						Sender:   "cosmos1qql8ag4cluz6r4dz28p3w00dnc9w8ueulg2gmc",
						Receiver: "qubetics16awr0rsqy72rgaydjv0fvuvzfdynshf62kkd8h",
						Amount:   "",
					},
				),
			},
			expAmount: "",
			expError:  true,
		},
		{
			name: "invalid amount - non-int",
			packet: channeltypes.Packet{
				Data: transfertypes.ModuleCdc.MustMarshalJSON(
					&transfertypes.FungibleTokenPacketData{
						Sender:   "cosmos1qql8ag4cluz6r4dz28p3w00dnc9w8ueulg2gmc",
						Receiver: "qubetics16awr0rsqy72rgaydjv0fvuvzfdynshf62kkd8h",
						Amount:   "test",
					},
				),
			},
			expAmount: "test",
			expError:  true,
		},
		{
			name: "valid",
			packet: channeltypes.Packet{
				Data: transfertypes.ModuleCdc.MustMarshalJSON(
					&transfertypes.FungibleTokenPacketData{
						Sender:   "cosmos1qql8ag4cluz6r4dz28p3w00dnc9w8ueulg2gmc",
						Receiver: "qubetics16awr0rsqy72rgaydjv0fvuvzfdynshf62kkd8h",
						Amount:   "10000",
					},
				),
			},
			expAmount: "10000",
			expError:  false,
		},
	}

	for _, tc := range testCases {
		amt, err := GetTransferAmount(tc.packet)
		if tc.expError {
			require.Error(t, err, tc.name)
		} else {
			require.NoError(t, err, tc.name)
			require.Equal(t, tc.expAmount, amt)
		}
	}
}

func TestGetReceivedCoin(t *testing.T) {
	testCases := []struct {
		name       string
		srcPort    string
		srcChannel string
		dstPort    string
		dstChannel string
		rawDenom   string
		rawAmount  string
		expCoin    sdk.Coin
	}{
		{
			"transfer unwrapped coin to destination which is not its source",
			"transfer",
			"channel-0",
			"transfer",
			"channel-0",
			"uosmo",
			"10",
			sdk.Coin{Denom: teststypes.UosmoIbcdenom, Amount: math.NewInt(10)},
		},
		{
			"transfer ibc wrapped coin to destination which is its source",
			"transfer",
			"channel-0",
			"transfer",
			"channel-0",
			"transfer/channel-0/tics",
			"10",
			sdk.Coin{Denom: "tics", Amount: math.NewInt(10)},
		},
		{
			"transfer 2x ibc wrapped coin to destination which is its source",
			"transfer",
			"channel-0",
			"transfer",
			"channel-2",
			"transfer/channel-0/transfer/channel-1/uatom",
			"10",
			sdk.Coin{Denom: teststypes.UatomIbcdenom, Amount: math.NewInt(10)},
		},
		{
			"transfer ibc wrapped coin to destination which is not its source",
			"transfer",
			"channel-0",
			"transfer",
			"channel-0",
			"transfer/channel-1/uatom",
			"10",
			sdk.Coin{Denom: teststypes.UatomOsmoIbcdenom, Amount: math.NewInt(10)},
		},
	}

	for _, tc := range testCases {
		coin := GetReceivedCoin(tc.srcPort, tc.srcChannel, tc.dstPort, tc.dstChannel, tc.rawDenom, tc.rawAmount)
		require.Equal(t, tc.expCoin, coin)
	}
}

func TestGetSentCoin(t *testing.T) {
	testCases := []struct {
		name      string
		rawDenom  string
		rawAmount string
		expCoin   sdk.Coin
	}{
		{
			"get unwrapped tics coin",
			"tics",
			"10",
			sdk.Coin{Denom: "tics", Amount: math.NewInt(10)},
		},
		{
			"get ibc wrapped tics coin",
			"transfer/channel-0/tics",
			"10",
			sdk.Coin{Denom: teststypes.TicsIbcdenom, Amount: math.NewInt(10)},
		},
		{
			"get ibc wrapped uosmo coin",
			"transfer/channel-0/uosmo",
			"10",
			sdk.Coin{Denom: teststypes.UosmoIbcdenom, Amount: math.NewInt(10)},
		},
		{
			"get ibc wrapped uatom coin",
			"transfer/channel-1/uatom",
			"10",
			sdk.Coin{Denom: teststypes.UatomIbcdenom, Amount: math.NewInt(10)},
		},
		{
			"get 2x ibc wrapped uatom coin",
			"transfer/channel-0/transfer/channel-1/uatom",
			"10",
			sdk.Coin{Denom: teststypes.UatomOsmoIbcdenom, Amount: math.NewInt(10)},
		},
	}

	for _, tc := range testCases {
		coin := GetSentCoin(tc.rawDenom, tc.rawAmount)
		require.Equal(t, tc.expCoin, coin)
	}
}

func TestDeriveDecimalsFromDenom(t *testing.T) {
	testCases := []struct {
		name      string
		baseDenom string
		expDec    uint8
		expFail   bool
		expErrMsg string
	}{
		{
			name:      "fail: empty string",
			baseDenom: "",
			expDec:    0,
			expFail:   true,
			expErrMsg: "Base denom cannot be an empty string",
		},
		{
			name:      "fail: invalid prefix",
			baseDenom: "nqubetics",
			expDec:    0,
			expFail:   true,
			expErrMsg: "Should be either micro ('u[...]') or atto ('a[...]'); got: \"nqubetics\"",
		},
		{
			name:      "success: micro 'u' prefix",
			baseDenom: "uqubetics",
			expDec:    6,
			expFail:   false,
			expErrMsg: "",
		},
	}

	for _, tc := range testCases {
		dec, err := DeriveDecimalsFromDenom(tc.baseDenom)
		if tc.expFail {
			require.Error(t, err, tc.expErrMsg)
			require.Contains(t, err.Error(), tc.expErrMsg)
		} else {
			require.NoError(t, err)
		}
		require.Equal(t, tc.expDec, dec)
	}
}

func TestIsBaseDenomFromSourceChain(t *testing.T) {
	tests := []struct {
		name     string
		denom    string
		expected bool
	}{
		{
			name:     "one hop",
			denom:    "transfer/channel-0/uatom",
			expected: false,
		},
		{
			name:     "no hop with factory prefix",
			denom:    "factory/owner/uatom",
			expected: false,
		},
		{
			name:     "multi hop",
			denom:    "transfer/channel-0/transfer/channel-1/uatom",
			expected: false,
		},
		{
			name:     "no hop",
			denom:    "uatom",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsBaseDenomFromSourceChain(tt.denom)
			require.Equal(t, tt.expected, result)
		})
	}
}
