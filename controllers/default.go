package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"pro_monitor/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	log.Println(beego.AppConfig.String("websites"))
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func (c *MainController) GetNavTreeJson() {
	data := make([]models.EasyuiTree, 0, 10)
	data = append(data, models.EasyuiTree{Text: "日志管理", Id: "1000", Status: "open", Children: []models.EasyuiTree{models.EasyuiTree{Text: "应用日志管理", Id: "1001", Status: "open"}, models.EasyuiTree{Text: "系统日志管理", Id: "1002", Status: "open"}}})
	c.Data["json"] = data
	c.ServeJSON()
}
