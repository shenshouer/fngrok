package client

import (
	"fmt"
	"fngrok/conf"

	log "github.com/Sirupsen/logrus"
)

type (
	// Client client interface
	Client interface {
		PrintConfig()
		Init() error
		Start() error
		HTTP() error
	}

	client struct {
		config    *conf.CConfig
		authority *authority
	}
)

// New create client instance
func New(serverAddr string, serverPort uint64, privilegeToken, authToken string) (c Client) {
	c = &client{
		config: conf.NewDefaultCConfig(),
	}

	return
}

func (c *client) Init() (err error) {
	return
}

func (c *client) PrintConfig() {
	fmt.Println("ServerAddr:", c.config.ServerAddr)
	fmt.Println("ServerPort:", c.config.ServerPort)
	fmt.Println("PrivilegeToken:", c.config.PrivilegeToken)
	fmt.Println("AuthToken:", c.config.AuthToken)
	fmt.Println("HeartBeatInterval:", c.config.HeartBeatInterval)
	fmt.Println("HeartBeatTimeout:", c.config.HeartBeatTimeout)
	fmt.Println("LogFile:", c.config.LogFile)
	fmt.Println("LogLevel:", c.config.LogLevel)
	fmt.Println("LogMaxDays:", c.config.LogMaxDays)
}

func (c *client) Start() (err error) {
	if c.authority.IsCertExpire() {
		err = fmt.Errorf("certificates have expired")
		return
	}

	return
}

// HTTP start http proxy
func (c *client) HTTP() (err error) {
	log.Infoln("client.Http")
	return
}
