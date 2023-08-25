package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	
	"onlineshop-api/goods-web/global"
)

func InitTrans(locale string) (err error) {
	//修改gin框架中的validator引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//参数一:备用语言环境，后面是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator failed : %s", locale)
		}

		switch locale {
		case "en":
			enTranslations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			zhTranslations.RegisterDefaultTranslations(v, global.Trans)
		default:
			enTranslations.RegisterDefaultTranslations(v, global.Trans)
		}
		return
	}
	return
}
