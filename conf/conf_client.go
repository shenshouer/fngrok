package conf

import (
	"github.com/spf13/viper"
)

// CConfig config for client
type CConfig struct {
	ServerAddr        string
	ServerPort        int64
	PrivilegeToken    string
	AuthToken         string
	HeartBeatInterval int64
	HeartBeatTimeout  int64
	LogFile           string
	LogLevel          string
	LogMaxDays        int
}

// NewDefaultCConfig create default config for client
func NewDefaultCConfig() (cc *CConfig) {
	cc = &CConfig{
		ServerAddr:        "0.0.0.0",
		ServerPort:        7000,
		LogFile:           "console",
		LogLevel:          "info",
		LogMaxDays:        3,
		PrivilegeToken:    "",
		AuthToken:         "",
		HeartBeatInterval: 20,
		HeartBeatTimeout:  90,
	}

	cc.PrivilegeToken = viper.GetString("privilege_token")
	cc.AuthToken = viper.GetString("auth_token")
	cc.LogFile = viper.GetString("log_file")
	cc.LogLevel = viper.GetString("log_level")
	cc.LogMaxDays = viper.GetInt("log_max_days")
	return
}
