package config

type Config struct {
	NodeVersion    string
	BootstrapPeers []string
}

func NodeConfig() *Config {
	return &Config{
		NodeVersion:    "v4.0.1",
		BootstrapPeers: []string{"152.32.140.46:5005"},
	}
}

func GetBootstrapPeers() []string {

	//peer1 :=
	return NodeConfig().BootstrapPeers
}