package testutil

import (
	"strconv"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/qubetics/qubetics-blockchain/v2/app"
	"github.com/qubetics/qubetics-blockchain/v2/crypto/ethsecp256k1"
)

// SubmitProposal delivers a submit proposal tx for a given gov content.
// Depending on the content type, the eventNum needs to specify submit_proposal
// event.
func SubmitProposal(
	ctx sdk.Context,
	appQubetics *app.Qubetics,
	pk *ethsecp256k1.PrivKey,
	content govv1beta1.Content,
	eventNum int,
) (id uint64, err error) {
	accountAddress := sdk.AccAddress(pk.PubKey().Address().Bytes())
	stakeDenom := stakingtypes.DefaultParams().BondDenom

	deposit := sdk.NewCoins(sdk.NewCoin(stakeDenom, math.NewInt(100000000)))
	msg, err := govv1beta1.NewMsgSubmitProposal(content, deposit, accountAddress)
	if err != nil {
		return id, err
	}
	res, err := DeliverTx(ctx, appQubetics, pk, nil, msg)
	if err != nil {
		return id, err
	}

	submitEvent := res.GetEvents()[eventNum]
	if submitEvent.Type != "submit_proposal" || submitEvent.Attributes[0].Key != "proposal_id" {
		return id, errorsmod.Wrapf(errorsmod.Error{}, "eventNumber %d in SubmitProposal calls %s instead of submit_proposal", eventNum, submitEvent.Type)
	}

	return strconv.ParseUint(submitEvent.Attributes[0].Value, 10, 64)
}

// Delegate delivers a delegate tx
func Delegate(
	ctx sdk.Context,
	appQubetics *app.Qubetics,
	priv *ethsecp256k1.PrivKey,
	delegateAmount sdk.Coin,
	validator stakingtypes.Validator,
) (abci.ResponseDeliverTx, error) {
	accountAddress := sdk.AccAddress(priv.PubKey().Address().Bytes())

	val, err := sdk.ValAddressFromBech32(validator.OperatorAddress)
	if err != nil {
		return abci.ResponseDeliverTx{}, err
	}

	delegateMsg := stakingtypes.NewMsgDelegate(accountAddress, val, delegateAmount)
	return DeliverTx(ctx, appQubetics, priv, nil, delegateMsg)
}

// Vote delivers a vote tx with the VoteOption "yes"
func Vote(
	ctx sdk.Context,
	appQubetics *app.Qubetics,
	priv *ethsecp256k1.PrivKey,
	proposalID uint64,
	voteOption govv1beta1.VoteOption,
) (abci.ResponseDeliverTx, error) {
	accountAddress := sdk.AccAddress(priv.PubKey().Address().Bytes())

	voteMsg := govv1beta1.NewMsgVote(accountAddress, proposalID, voteOption)
	return DeliverTx(ctx, appQubetics, priv, nil, voteMsg)
}
