package global

import (
	ut "github.com/go-playground/universal-translator"
	"onlineshop-api/oss-web/config"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}
)
