package memories

import "time"

/*
@Time : 2021/11/27 23:35
@Author : onns
@File : /type.go
*/

type Anniversary struct {
	Type    AnniversaryType `json:"type"`
	Name    string          `json:"name"`
	Date    time.Time       `json:"date_format"`
	DateRaw string          `json:"date"`
}

type AnniversaryType int8

const (
	Birthday      AnniversaryType = 1
	LunarBirthday AnniversaryType = 2
	SpecialDay    AnniversaryType = 3
)
