package helpers

import (
	"github.com/astaxie/beego"
)

func HandleError(err error) bool {
	if err != nil {
		beego.Error(err)
		return true
	}
	return false
}
