package models

import (
	"fmt"
	"strings"
	"time"
)

type JsonTime time.Time

type AppLogVm struct {
	Time   JsonTime `json:"time"`
	App    string   `json:"app"`
	Host   string   `json:"host"`
	IP     string   `json:"ip"`
	TID    string   `json:"tid"`
	Level  string   `json:"level"`
	Class  string   `json:"class"`
	Method string   `json:"method"`
	Msg    string   `json:"msg"`
}

func (t *JsonTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	tm, err := time.Parse("2006/1/02 15:04:05", strings.Replace(str, "\"", "", -1))
	*t = JsonTime(tm)
	return err
}

func (t JsonTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
