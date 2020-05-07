package factory

type Logger struct {
	DebugLevel   string `yaml:"debugLevel"`
	ReportCaller bool   `yaml:"ReportCaller"`
}
