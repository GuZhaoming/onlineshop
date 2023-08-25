package global

import (
	ut "github.com/go-playground/universal-translator"

	"onlineshop-api/goods-web/config"
	"onlineshop-api/goods-web/proto"
)

var (
	ServerConfig   *config.ServerConfig = &config.ServerConfig{}
	Trans          ut.Translator
	GoodsSrvClient proto.GoodsClient
)
