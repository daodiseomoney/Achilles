package v102

import (
	"context"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

const UpgradeName = "v1.0.2"

func CreateUpgradeHandler(
	mm *module.Manager,
	cfg module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {

		return mm.RunMigrations(ctx, cfg, fromVM)
	}
}
