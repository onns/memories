package memories

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	ics "github.com/arran4/golang-ical"
)

/*
@Time : 2021/11/27 23:34
@Author : onns
@File : /func.go
*/

func GenerateDays(as []*Anniversary) (days []*Anniversary) {
	days = make([]*Anniversary, 0)
	for _, a := range as {
		switch a.Type {
		case Birthday:
			for i := 0; i <= 100; i++ {
				days = append(days, &Anniversary{
					Type: a.Type,
					Name: fmt.Sprintf("%s的第%d个生日", a.Name, i),
					Date: a.Date.AddDate(i, 0, 0),
				})
			}
		}
	}
	return
}

func GenerateIcs(days []*Anniversary) (res string) {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodPublish)
	cal.SetXWRCalName("Special Day")
	cal.SetXWRTimezone("Asia/Shanghai")
	for _, day := range days {
		event := cal.AddEvent(fmt.Sprintf("%s_%s@onns.xyz", generateUid(day.Name), day.Date.Format("20060102")))
		event.SetCreatedTime(time.Now())
		event.SetCreatedTime(time.Now())
		event.SetDtStampTime(time.Now())
		event.SetModifiedAt(time.Now())
		event.SetAllDayStartAt(day.Date)
		event.SetAllDayEndAt(day.Date)
		event.SetSummary(day.Name)
		event.SetLocation("")
		event.SetDescription("")
		event.SetSequence(0)
		event.SetStatus(ics.ObjectStatusConfirmed)
		// alarm := event.AddAlarm()
		// alarm.Properties
	}
	res = cal.Serialize()
	return
}

func generateUid(name string) string {
	h := md5.New()
	h.Write([]byte(name))
	return hex.EncodeToString(h.Sum(nil))
}
