package models

type EasyuiTree struct {
	Id       string       `json:"id"`
	Text     string       `json:"text"`
	IconCls  string       `json:"iconCls"`
	Children []EasyuiTree `json:"children"`
	Checked  bool         `json:"checked"`
	Status   string       `json:"state"`
}
