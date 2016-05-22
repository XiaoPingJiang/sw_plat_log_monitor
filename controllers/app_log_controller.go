package controllers

import (
	"github.com/astaxie/beego"
	"pro_monitor/biz"
	"pro_monitor/models"
	"strings"
)

type AppLogController struct {
	beego.Controller
}

func (c *AppLogController) List() {
	c.TplName = "app_log_list.html"
}

func (c *AppLogController) GetListJson() {
	website := c.Input().Get("website")
	logtype := c.Input().Get("logtype")
	page, _ := c.GetInt("page", 0)
	size, _ := c.GetInt("rows", 0)
	data, count := new(biz.AppLogBiz).GetPageList(website, logtype, page, size)
	c.Data["json"] = models.EasyuiDataGrid{Total: count, Rows: data}
	c.ServeJSON()
}

func (c *AppLogController) GetWebSiteComboboxJson() {
	data := make([]models.EasyuiCombo, 0, 10)
	arr := strings.Split(beego.AppConfig.String("websites"),"|")
	for _, v := range arr {
		data = append(data, models.EasyuiCombo{Id: v, Text: v})
	}
	c.Data["json"] = data
	c.ServeJSON()
}

func (c *AppLogController) GetLogTypeComboboxJson() {
	data := make([]models.EasyuiCombo, 0, 10)
	arr := strings.Split(beego.AppConfig.String("logtypes"), "|")
	for _, a := range arr {
		brr := strings.Split(a, ":")
		data = append(data, models.EasyuiCombo{Id: brr[1], Text: brr[0]})
	}
	c.Data["json"] = data
	c.ServeJSON()
}

func (c *AppLogController) ReadFromTodayFile() {
	new(biz.AppLogBiz).ReadFromTodayFile()
	c.Data["json"] = models.ResponseJsonVm{Success: true, Message: "操作成功"}
	c.ServeJSON()
}
