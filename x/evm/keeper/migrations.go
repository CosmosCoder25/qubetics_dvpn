package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v7 "github.com/qubetics/qubetics-blockchain/v2/x/evm/migrations/v7"
	"github.com/qubetics/qubetics-blockchain/v2/x/evm/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper         Keeper
	legacySubspace types.Subspace
}

// NewMigrator returns a new Migrator instance.
func NewMigrator(keeper Keeper, legacySubspace types.Subspace) Migrator {
	return Migrator{
		keeper:         keeper,
		legacySubspace: legacySubspace,
	}
}

// Migrate6to7 migrates the store from consensus version 6 to 7.
func (m Migrator) Migrate6to7(ctx sdk.Context) error {
	return v7.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}
