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
	latestDay := &Anniversary{
		Date: time.Time{},
	}
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
				lunarNew, err := lunar.AddDate(i, 0, 0)
				if err != nil {
					// 目前阴历只能算到2050年
					break
				}
				days = append(days, &Anniversary{
					Type:   a.Type,
					Name:   fmt.Sprintf("%s的第%d个农历生日", a.Name, i),
					Date:   lunarNew.ToSolar(),
					Start:  a.Start,
					End:    a.End,
					AllDay: a.AllDay,
				})
			}
		case SpecialDay:
			countdown := formatCountdown(a.Countdown, DefaultMax)
			for i := 1; i <= countdown/365; i++ {
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
			sep := formatSep(a.Sep, DefaultSep)
			for i := 1; i*sep <= countdown; i++ {
				days = append(days, &Anniversary{
					Type:   a.Type,
					Name:   fmt.Sprintf("%s的第%d天", a.Name, i*sep),
					Date:   a.Date.AddDate(0, 0, i*sep-1),
					Start:  a.Start,
					End:    a.End,
					AllDay: a.AllDay,
				})
			}
			if countdown%sep != 0 {
				days = append(days, &Anniversary{
					Type:   a.Type,
					Name:   fmt.Sprintf("%s的第%d天", a.Name, countdown),
					Date:   a.Date.AddDate(0, 0, countdown-1),
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
		case Countdown:
			days = append(days, &Anniversary{
				Type:   a.Type,
				Name:   fmt.Sprintf("%s", a.Name),
				Date:   a.Date,
				Start:  a.Start,
				End:    a.End,
				AllDay: a.AllDay,
			})
			sep := formatSep(a.Sep, DefaultSep)
			countdown := formatCountdown(a.Countdown, DefaultCountDown)
			for i := 1; i < sep; i++ {
				days = append(days, &Anniversary{
					Type:   a.Type,
					Name:   fmt.Sprintf("距离%s还有%d天", a.Name, i),
					Date:   a.Date.AddDate(0, 0, -i),
					Start:  a.Start,
					End:    a.End,
					AllDay: a.AllDay,
				})
			}
			for i := 1; i*sep <= countdown; i++ {
				days = append(days, &Anniversary{
					Type:   a.Type,
					Name:   fmt.Sprintf("距离%s还有%d天", a.Name, i*sep),
					Date:   a.Date.AddDate(0, 0, -i*sep),
					Start:  a.Start,
					End:    a.End,
					AllDay: a.AllDay,
				})
			}
			if countdown%sep != 0 {
				days = append(days, &Anniversary{
					Type:   a.Type,
					Name:   fmt.Sprintf("距离%s还有%d天", a.Name, countdown),
					Date:   a.Date.AddDate(0, 0, -countdown),
					Start:  a.Start,
					End:    a.End,
					AllDay: a.AllDay,
				})
			}
		case RepeatedDay:
			sep := formatSep(a.Sep, DefaultRepeat)
			countdown := formatCountdown(a.Countdown, DefaultCountDown)
			for i := 0; i*sep <= countdown; i++ {
				days = append(days, &Anniversary{
					Type:   a.Type,
					Name:   a.Name,
					Date:   a.Date.AddDate(0, 0, i*sep),
					Start:  a.Start,
					End:    a.End,
					AllDay: a.AllDay,
				})
			}
		case NucleicAcidTestDay:
			if a.Date.After(latestDay.Date) {
				latestDay = a
			}
			days = append(days, &Anniversary{
				Type:   a.Type,
				Name:   fmt.Sprintf("%s", a.Name),
				Date:   a.Date,
				Start:  a.Start,
				End:    a.End,
				AllDay: a.AllDay,
			})

		}
	}
	if !latestDay.Date.IsZero() {
		days = append(days, &Anniversary{
			Type:   latestDay.Type,
			Name:   "核酸检测失效",
			Date:   latestDay.Date.AddDate(0, 0, 3),
			Start:  latestDay.Start,
			End:    latestDay.End,
			AllDay: latestDay.AllDay,
		})
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
		alarm := event.AddAlarm()
		if day.AllDay {
			alarm.SetTrigger("P0DT9H0M0S")
		} else {
			alarm.SetTrigger("-P0DT2H0M0S")
		}
		alarm.SetAction(ics.ActionDisplay)
		// BEGIN:VALARM
		// ACTION:DISPLAY
		// DESCRIPTION:This is an event reminder
		// TRIGGER:P0DT9H0M0S
		// END:VALARM

		// alarm.Properties
	}
	res = cal.Serialize()
	return
}

func mergeDate(a, b time.Time) (res time.Time) {
	res = time.Date(a.Year(), a.Month(), a.Day(), b.Hour(), b.Minute(), b.Second(), b.Nanosecond(), b.Location())
	return
}

func generateUid(a *Anniversary) string {
	h := md5.New()
	h.Write([]byte(a.Name))
	return fmt.Sprintf("%s%d%s", hex.EncodeToString(h.Sum(nil)), a.Type, a.Date.Format("20060102"))
}

func formatSep(sep, defaultSep int) (res int) {
	if sep <= 0 {
		res = defaultSep
		return
	}
	res = sep
	return
}

func formatCountdown(countdown, defaultCountdown int) (res int) {
	if countdown <= 0 {
		res = defaultCountdown
		return
	}
	res = countdown
	return
}
