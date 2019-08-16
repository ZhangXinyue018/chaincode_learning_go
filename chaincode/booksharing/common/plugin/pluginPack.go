package plugin

type PluginPack struct {
	IdentityVerifier    IdentityVerifier
	APIDataProcessor    DataProcessor
	LedgerDataProcessor DataProcessor
}
