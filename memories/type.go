package memories

import "time"

/*
@Time : 2021/11/27 23:35
@Author : onns
@File : /type.go
*/

type Anniversary struct {
	Type     AnniversaryType `json:"type"`
	Name     string          `json:"name"`
	Date     time.Time       `json:"date_format"`
	DateRaw  string          `json:"date"`
	Start    time.Time       `json:"start_format"`
	StartRaw string          `json:"start"`
	End      time.Time       `json:"end_format"`
	EndRaw   string          `json:"end"`
	AllDay   bool            `json:"all_day"`
}

type AnniversaryType int8

const (
	Birthday      AnniversaryType = 1
	LunarBirthday AnniversaryType = 2
	SpecialDay    AnniversaryType = 3
	OneDay        AnniversaryType = 4
)
