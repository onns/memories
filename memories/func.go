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
					Type:   a.Type,
					Name:   fmt.Sprintf("%s的第%d个生日", a.Name, i),
					Date:   a.Date.AddDate(i, 0, 0),
					Start:  a.Start,
					End:    a.End,
					AllDay: a.AllDay,
				})
			}
		case LunarBirthday:
			lunar := lunar.Parse(a.Date)
			for i := 0; i <= 100; i++ {
				days = append(days, &Anniversary{
					Type:   a.Type,
					Name:   fmt.Sprintf("%s的第%d个农历生日", a.Name, i),
					Date:   lunar.AddDate(i, 0, 0).ToSolar(),
					Start:  a.Start,
					End:    a.End,
					AllDay: a.AllDay,
				})
			}
		case SpecialDay:
			for i := 1; i <= 100; i++ {
				days = append(days, &Anniversary{
					Type:   a.Type,
					Name:   fmt.Sprintf("%s的第%d年", a.Name, i),
					Date:   a.Date.AddDate(i, 0, 0),
					Start:  a.Start,
					End:    a.End,
					AllDay: a.AllDay,
				})
			}
			days = append(days, &Anniversary{
				Type:   a.Type,
				Name:   fmt.Sprintf("%s的第%d天", a.Name, 1),
				Date:   a.Date,
				Start:  a.Start,
				End:    a.End,
				AllDay: a.AllDay,
			})
			for i := 1; i*100 <= 365*100; i++ {
				days = append(days, &Anniversary{
					Type:   a.Type,
					Name:   fmt.Sprintf("%s的第%d天", a.Name, i*100),
					Date:   a.Date.AddDate(0, 0, i*100-1),
					Start:  a.Start,
					End:    a.End,
					AllDay: a.AllDay,
				})
			}
		case OneDay:
			days = append(days, &Anniversary{
				Type:   a.Type,
				Name:   fmt.Sprintf("%s", a.Name),
				Start:  a.Start,
				End:    a.End,
				Date:   a.Date,
				AllDay: a.AllDay,
			})
		}
	}
	return
}

func GenerateIcs(name string, days []*Anniversary) (res string) {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodPublish)
	cal.SetXWRCalName(name)
	cal.SetXWRTimezone("Asia/Shanghai")
	for _, day := range days {
		event := cal.AddEvent(fmt.Sprintf("%s@onns.xyz", generateUid(day)))
		event.SetCreatedTime(time.Now())
		event.SetCreatedTime(time.Now())
		event.SetDtStampTime(time.Now())
		event.SetModifiedAt(time.Now())
		if day.AllDay {
			event.SetAllDayStartAt(day.Date)
			event.SetAllDayEndAt(day.Date)
		} else {
			event.SetStartAt(mergeDate(day.Date, day.Start))
			event.SetEndAt(mergeDate(day.Date, day.End))
		}
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

func mergeDate(a, b time.Time) (res time.Time) {
	res = time.Date(a.Year(), a.Month(), a.Day(), b.Hour(), b.Minute(), b.Second(), b.Nanosecond(), a.Location())
	return
}

func generateUid(a *Anniversary) string {
	h := md5.New()
	h.Write([]byte(a.Name))
	return fmt.Sprintf("%s%d%s", hex.EncodeToString(h.Sum(nil)), a.Type, a.Date.Format("20060102"))
}
