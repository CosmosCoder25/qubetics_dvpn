package ante_test

import (
	"time"

	sdkmath "cosmossdk.io/math"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	testutiltx "github.com/qubetics/qubetics-blockchain/v2/testutil/tx"

	//nolint:revive // dot imports are fine for Ginkgo
	. "github.com/onsi/ginkgo/v2"
	//nolint:revive // dot imports are fine for Ginkgo
	. "github.com/onsi/gomega"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qubetics/qubetics-blockchain/v2/crypto/ethsecp256k1"
	"github.com/qubetics/qubetics-blockchain/v2/testutil"
	"github.com/qubetics/qubetics-blockchain/v2/utils"
)

var _ = Describe("when sending a Cosmos transaction", func() {
	var (
		addr sdk.AccAddress
		priv *ethsecp256k1.PrivKey
		msg  sdk.Msg
	)

	Context("and the sender account has enough balance to pay for the transaction cost", Ordered, func() {
		var (
			rewardsAmt = sdkmath.NewInt(1e5)
			balance    = sdkmath.NewInt(1e18)
		)

		BeforeEach(func() {
			addr, priv = testutiltx.NewAccAddressAndKey()

			msg = &banktypes.MsgSend{
				FromAddress: addr.String(),
				ToAddress:   "qubetics16awr0rsqy72rgaydjv0fvuvzfdynshf62kkd8h",
				Amount:      sdk.Coins{sdk.Coin{Amount: sdkmath.NewInt(1e14), Denom: utils.BaseDenom}},
			}

			s.ctx, _ = testutil.PrepareAccountsForDelegationRewards(
				s.T(), s.ctx, s.app, addr, balance, rewardsAmt,
			)

			var err error
			s.ctx, err = testutil.CommitAndCreateNewCtx(s.ctx, s.app, time.Second*0, nil)
			Expect(err).To(BeNil())
		})

		It("should succeed & not withdraw any staking rewards", func() {
			res, err := testutil.DeliverTx(s.ctx, s.app, priv, nil, msg)
			Expect(err).To(BeNil())
			Expect(res.IsOK()).To(BeTrue())

			rewards, err := testutil.GetTotalDelegationRewards(s.ctx, s.app.DistrKeeper, addr)
			Expect(err).To(BeNil())
			Expect(rewards).To(Equal(sdk.NewDecCoins(sdk.NewDecCoin(utils.BaseDenom, rewardsAmt))))
		})
	})

	Context("and the sender account neither has enough balance nor sufficient staking rewards to pay for the transaction cost", func() {
		var (
			rewardsAmt = sdkmath.NewInt(0)
			balance    = sdkmath.NewInt(0)
		)

		BeforeEach(func() {
			addr, priv = testutiltx.NewAccAddressAndKey()

			msg = &banktypes.MsgSend{
				FromAddress: addr.String(),
				ToAddress:   "qubetics16awr0rsqy72rgaydjv0fvuvzfdynshf62kkd8h",
				Amount:      sdk.Coins{sdk.Coin{Amount: sdkmath.NewInt(1e14), Denom: utils.BaseDenom}},
			}

			s.ctx, _ = testutil.PrepareAccountsForDelegationRewards(
				s.T(), s.ctx, s.app, addr, balance, rewardsAmt,
			)

			var err error
			s.ctx, err = testutil.CommitAndCreateNewCtx(s.ctx, s.app, time.Second*0, nil)
			Expect(err).To(BeNil())
		})

		It("should fail", func() {
			res, err := testutil.DeliverTx(s.ctx, s.app, priv, nil, msg)
			Expect(res.IsOK()).To(BeTrue())
			Expect(err).To(HaveOccurred())
		})

		It("should not withdraw any staking rewards", func() {
			rewards, err := testutil.GetTotalDelegationRewards(s.ctx, s.app.DistrKeeper, addr)
			Expect(err).To(BeNil())
			Expect(rewards.Empty()).To(BeTrue())
		})
	})

	Context("and the sender account has not enough balance but sufficient staking rewards to pay for the transaction cost", func() {
		var (
			rewardsAmt = sdkmath.NewInt(1e18)
			balance    = sdkmath.NewInt(0)
		)

		BeforeEach(func() {
			addr, priv = testutiltx.NewAccAddressAndKey()

			msg = &banktypes.MsgSend{
				FromAddress: addr.String(),
				ToAddress:   "qubetics16awr0rsqy72rgaydjv0fvuvzfdynshf62kkd8h",
				Amount:      sdk.Coins{sdk.Coin{Amount: sdkmath.NewInt(1), Denom: utils.BaseDenom}},
			}

			s.ctx, _ = testutil.PrepareAccountsForDelegationRewards(
				s.T(), s.ctx, s.app, addr, balance, rewardsAmt,
			)
			var err error
			s.ctx, err = testutil.CommitAndCreateNewCtx(s.ctx, s.app, time.Second*0, nil)
			Expect(err).To(BeNil())
		})

		It("should withdraw enough staking rewards to cover the transaction cost", func() {
			rewards, err := testutil.GetTotalDelegationRewards(s.ctx, s.app.DistrKeeper, addr)
			Expect(err).To(BeNil())
			Expect(rewards).To(Equal(sdk.NewDecCoins(sdk.NewDecCoin(utils.BaseDenom, rewardsAmt))))

			balance := s.app.BankKeeper.GetBalance(s.ctx, addr, utils.BaseDenom)
			Expect(balance.Amount).To(Equal(sdkmath.NewInt(0)))

			res, err := testutil.DeliverTx(s.ctx, s.app, priv, nil, msg)
			Expect(res.IsOK()).To(BeTrue())
			Expect(err).To(BeNil())
		})
	})
})
