package memories

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/onns/lunar"
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
		case LunarBirthday:
			lunar := lunar.Parse(a.Date)
			for i := 0; i <= 100; i++ {
				days = append(days, &Anniversary{
					Type: a.Type,
					Name: fmt.Sprintf("%s的第%d个农历生日", a.Name, i),
					Date: lunar.AddDate(i, 0, 0).ToSolar(),
				})
			}
		case SpecialDay:
			for i := 1; i <= 100; i++ {
				days = append(days, &Anniversary{
					Type: a.Type,
					Name: fmt.Sprintf("%s的第%d年", a.Name, i),
					Date: a.Date.AddDate(i, 0, 0),
				})
			}
			days = append(days, &Anniversary{
				Type: a.Type,
				Name: fmt.Sprintf("%s的第%d天", a.Name, 1),
				Date: a.Date,
			})
			for i := 1; i*100 <= 365*100; i++ {
				days = append(days, &Anniversary{
					Type: a.Type,
					Name: fmt.Sprintf("%s的第%d天", a.Name, i*100),
					Date: a.Date.AddDate(0, 0, i*100-1),
				})
			}
			// days = append(days, &Anniversary{
			// 	Type: a.Type,
			// 	Name: fmt.Sprintf("%s的第%d天", a.Name, 520),
			// 	Date: a.Date.AddDate(520, 0, 0),
			// })
			// days = append(days, &Anniversary{
			// 	Type: a.Type,
			// 	Name: fmt.Sprintf("%s的第%d天", a.Name, 1314),
			// 	Date: a.Date.AddDate(1314, 0, 0),
			// })
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
		event := cal.AddEvent(fmt.Sprintf("%s@onns.xyz", generateUid(day)))
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

func generateUid(a *Anniversary) string {
	h := md5.New()
	h.Write([]byte(a.Name))
	return fmt.Sprintf("%s%d%s", hex.EncodeToString(h.Sum(nil)), a.Type, a.Date.Format("20060102"))
}
