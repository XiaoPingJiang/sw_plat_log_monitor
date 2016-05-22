package biz

import (
	"bufio"

	"database/sql"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/axgle/mahonia"
	_ "github.com/mattn/go-adodb"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"pro_monitor/biz/dtos"
	"pro_monitor/helpers"
	"strconv"
	"strings"
	"time"
)

type AppLogBiz struct {
}

func (c *AppLogBiz) GetPageList(website, logtype string, page, size int) ([]dtos.AppLogDto, int) {
	db, err := sql.Open("adodb", beego.AppConfig.String("connstr"))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	data := make([]dtos.AppLogDto, 0, 15)
	sqlQuery := "select Id,[Time],App,Ip,Msg,WebSite,Level from AppLog where 1=1 "
	sqlCount := "select count(id) from AppLog where 1=1 "
	params := make([]interface{}, 0, 10)
	if website != "" {
		sqlQuery += " and website=? "
		sqlCount += " and website=? "
		params = append(params, website)
	}
	if logtype != "" {
		sqlQuery += " and level=? "
		sqlCount += " and level=? "
		params = append(params, logtype)
	}
	rows, err := db.Query(sqlQuery+" order by id desc offset "+strconv.Itoa((page-1)*size)+" rows fetch next "+strconv.Itoa(size)+" rows only", params...)
	rowCount := db.QueryRow(sqlCount, params...)
	count := 0
	rowCount.Scan(&count)
	helpers.HandleError(err)
	for rows.Next() {
		tmp := dtos.AppLogDto{}
		rows.Scan(&tmp.IP, &tmp.Time, &tmp.App, &tmp.IP, &tmp.Msg, &tmp.WebSite, &tmp.Level)
		data = append(data, tmp)
		log.Println(tmp)
	}
	helpers.HandleError(err)
	return data, count
}

func (c *AppLogBiz) ReadFromTodayFile() {
	logs := make([]dtos.AppLogDto, 0, 10)
	// 取当前日期并格式化: 20160519
	date := time.Now().Format("20060102")
	websites := strings.Split(beego.AppConfig.String("websites"), "|")
	logtypes := strings.Split(beego.AppConfig.String("logtypes"), "|")
	for _, site := range websites {
		for _, logtype := range logtypes {
			logtype:= strings.Split(logtype, ":")[0]
			path := filepath.Join(beego.AppConfig.String("logpath"), site, logtype)
			beego.Debug(path)
			// 读取目录下所有文件
			files, err := ioutil.ReadDir(path)
			helpers.HandleError(err)
			env := mahonia.NewDecoder("GBK")
			for _, v := range files {
				// 如果是文件,并且文件名包含当前日期
				if !v.IsDir() && strings.Contains(v.Name(), date) {
					file, err := os.OpenFile(path+"/"+v.Name(), 0, 0777)
					helpers.HandleError(err)
					rd := bufio.NewReader(file)
					for {
						line, _, err := rd.ReadLine()
						if err != nil || io.EOF == err {
							break
						}
						strLog := env.ConvertString(string(line))
						if strLog != "" {
							tmpLog := dtos.AppLogDto{}
							json.Unmarshal([]byte(strLog), &tmpLog)
							tmpLog.WebSite = site
							logs = append(logs, tmpLog)
						}
					}
					file.Close()
				}
			}
		}
	}
	save(logs)
}

func save(logs []dtos.AppLogDto) {
	db, err := sql.Open("adodb", beego.AppConfig.String("connstr"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for _, v := range logs {
		count:=0
		db.QueryRow("select count(id) from applog where time=? and website=? and level=?", time.Time(v.Time), v.WebSite, v.Level).Scan(&count)
		if count == 0 {
			_, err := db.Exec("INSERT INTO [dbo].[AppLog]([Time],[App],[Host],[IP],[Level],[Class],[Method],[Msg],[WebSite]) VALUES(?,?,?,?,?,?,?,?,?)", time.Time(v.Time), v.App, v.Host, v.IP, v.Level, v.Class, v.Method, v.Msg, v.WebSite)
			helpers.HandleError(err)
		}
	}
}
