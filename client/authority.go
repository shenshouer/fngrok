package client

import (
	log "github.com/Sirupsen/logrus"
)

const (
	// AUTHHttp http proxy
	AUTHHttp AuthType = "HTTP"
)

type (
	// AuthType authority type
	AuthType string
	// Authority the authority of client
	authority struct{}
	// AuthorityHandler authority handler
	AuthorityHandler map[AuthType]func(*client)
)

// HasHTTP if client has http proxy authority
func (a *authority) HasHTTP() bool {
	// TODO
	log.Infoln("authority.HasHTTP", "是否有HTTP proxy 功能权限")
	return false
}

// IsCertExpire if the certificate expire
func (a *authority) IsCertExpire() bool {
	// TODO
	log.Infoln("authority.HasHTTP", "证书是否有效")
	return false
}
