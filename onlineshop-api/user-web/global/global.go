package global

import (
	ut "github.com/go-playground/universal-translator"

	"onlineshop-api/user-web/config"
	"onlineshop-api/user-web/proto"
)

var (
	ServerConfig  *config.ServerConfig = &config.ServerConfig{}
	Trans         ut.Translator
	UserSrvClient proto.UserClient
)
