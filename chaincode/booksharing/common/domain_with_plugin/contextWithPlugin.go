package domain_with_plugin

import (
	"github.com/booksharing/common/domain"
	"github.com/booksharing/common/plugin"
)

type ContractContext struct {
	DataContext *domain.Context
	Plugins     plugin.PluginPack
}
