// Copyright 2019 Andyfoo
// [http://andyfoo.com][http://pslib.com]
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package xdate

/*
时间格式参数仿php
https://www.php.net/manual/en/function.date.php

Examples:

> 当前时间:
	xdate.Now().Time

> 指定日期:
	xdate.Date(日期参数).Time

> 指定日期和时间:
	xdate.DateTime(日期和时间参数).Time

> go 默认格式:
	xdate.Now().Format("2006-01-02 15:04:05") // "2018-10-11 23:22:21"

-> php 格式:
	xdate.Now().PFormat("Y-m-d H:i:s") // "2018-10-11 23:22:21"
*/
import (
	"bytes"
	"time"
)

type WeekType uint

const (
	WeekType_def WeekType = iota
	WeekType_cn
	WeekType_cnShort
	WeekType_en
	WeekType_enShort
)

//字段类型
type DateField uint

const (
	DateField_YEAR DateField = iota
	DateField_MONTH
	DateField_DAY
	DateField_HOUR
	DateField_MINUTE
	DateField_SECOND
)

func (p DateField) String() string {
	switch p {
	case DateField_YEAR:
		return "YEAR"
	case DateField_MONTH:
		return "MONTH"
	case DateField_DAY:
		return "DAY"
	case DateField_HOUR:
		return "HOUR"
	case DateField_MINUTE:
		return "MINUTE"
	case DateField_SECOND:
		return "SECOND"
	default:
		return "UNKNOWN"
	}
}

var (
	months = map[int]time.Month{
		1:  time.January,
		2:  time.February,
		3:  time.March,
		4:  time.April,
		5:  time.May,
		6:  time.June,
		7:  time.July,
		8:  time.August,
		9:  time.September,
		10: time.October,
		11: time.November,
		12: time.December,
	}
	formats = map[string]string{
		"D": "Mon",
		"d": "02",
		"j": "2",
		"l": "Monday", //（"L"的小写字母）
		"z": "__2",

		"F": "January",
		"m": "01",
		"M": "Jan",
		"n": "1",

		"Y": "2006",
		"y": "06",

		"a": "pm",
		"A": "PM",
		"g": "3",
		"h": "03",
		"H": "15",
		"i": "04",
		"s": "05",
		"u": ".000000",

		"O": "-0700",
		"P": "-07:00",
		"T": "MST",

		"c": "2006-01-02T15:04:05Z07:00",
		"r": "Mon, 02 Jan 2006 15:04:05 -0700",
	}

	weeks = map[WeekType][]string{
		WeekType_cn:      {"星期天", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"},
		WeekType_cnShort: {"周日", "周一", "周二", "周三", "周四", "周五", "周六"},
		WeekType_en:      {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		WeekType_enShort: {"sun", "mon", "tue", "wed", "thu", "fri", "sat"},
	}

	TT  = "2006-01-02 15:04:05"
	YMD = "2006-01-02"
	HMS = "15:04:05"
)

//XDate
type XDate struct {
	time.Time
}

//Unix returns the local Time corresponding to the given Unix time,
func Unix(sec int64, nsec int64) XDate {
	return XDate{time.Unix(sec, nsec)}
}

//Now returns the current local time
func Now() XDate {
	return XDate{time.Now()}
}

//NowDateStr return date string: 2006-01-02
func NowDateStr() string {
	return time.Now().Format(YMD)
}

//NowUtcDateStr return utc date string: 2006-01-02 15:04:05
func NowUtcDateStr() string {
	return time.Now().UTC().Format(YMD)
}

//NowTimeStr return time string: 15:04:05
func NowTimeStr() string {
	return time.Now().Format(HMS)
}

//NowUtcTimeStr return utc time string: 15:04:05
func NowUtcTimeStr() string {
	return time.Now().UTC().Format(HMS)
}

//NowDateTimeStr return date&time string: 2006-01-02 15:04:05
func NowDateTimeStr() string {
	return time.Now().Format(TT)
}

//NowUtcDateTimeStr return utc date&time string: 2006-01-02 15:04:05
func NowUtcDateTimeStr() string {
	return time.Now().UTC().Format(TT)
}

//PFormat return date string, pformat=like php date style
func (t XDate) PFormat(pformat string) string {
	return t.Time.Format(PFormatConv(pformat))
}

//DateStr 2006-01-02
func (t XDate) DateStr() string {
	return t.Format(YMD)
	//return t.PFormat("Y-m-d")
}

//TimeStr 15:04:05
func (t XDate) TimeStr() string {
	return t.Format(HMS)
	//return t.PFormat("H:i:s")
}

//DateTimeStr 2006-01-02 15:04:05
func (t XDate) DateTimeStr() string {
	return t.Format(TT)
	//return t.PFormat("Y-m-d H:i:s")
}

//DayBeginDateTimeStr 2006-01-02 00:00:00
func (t XDate) DayBeginDateTimeStr() string {
	return t.Format(YMD) + " 00:00:00"
}

//DayEndDateTimeStr 2006-01-02 23:59:59
func (t XDate) DayEndDateTimeStr() string {
	return t.Format(YMD) + " 23:59:59"
}

//UnixMilli millisecond
func (t XDate) UnixMilli() int64 {
	return t.Time.UnixNano() / 1e6
}

//WeekdayStr week string
//WeekType_def
//WeekType_cn = { "星期天", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六" };
//WeekType_cnShort = { "周日", "周一", "周二", "周三", "周四", "周五", "周六" };
//WeekType_en = { "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday" };
//WeekType_enShort = { "sun", "mon", "tue", "wed", "thu", "fri", "sat" };
func (t XDate) WeekdayStr(wt WeekType) string {
	wtArr, exist := weeks[wt]
	if exist {
		return wtArr[int(t.Time.Weekday())]
	}
	return t.Time.Weekday().String()
}

//Weekday week day num
//(Sunday 星期日 = 0, ...).
func (t XDate) Weekday() int {
	return int(t.Time.Weekday())
}

// UTC returns t with the location set to UTC.
func (t XDate) UTC() XDate {
	return XDate{t.Time.UTC()}
}

// Local returns t with the location set to local time.
func (t XDate) Local() XDate {
	return XDate{t.Time.Local()}
}

//Offset date operation
func (t XDate) Offset(field DateField, offset int) XDate {
	switch field {
	case DateField_YEAR:
		t.Time = t.Time.AddDate(offset, 0, 0)
	case DateField_MONTH:
		t.Time = t.Time.AddDate(0, offset, 0)
	case DateField_DAY:
		t.Time = t.Time.AddDate(0, 0, offset)
	case DateField_HOUR:
		t.Time = t.Time.Add(time.Hour * time.Duration(offset))
	case DateField_MINUTE:
		t.Time = t.Time.Add(time.Minute * time.Duration(offset))
	case DateField_SECOND:
		t.Time = t.Time.Add(time.Second * time.Duration(offset))
	default:
	}

	return t
}

//PFormatConv php date format string to go date format string
func PFormatConv(pformat string) string {
	var format bytes.Buffer
	len := len(pformat)
	for i := 0; i < len; i++ {
		val, ok := formats[pformat[i:i+1]]
		if ok {
			format.WriteString(val)
		} else {
			format.WriteString(pformat[i : i+1])
		}

	}
	return format.String()
}

//Time time to XDate
func Time(time time.Time) XDate {
	return XDate{
		Time: time,
	}
}

//Date date to XDate
func Date(day int, month int, year int) XDate {
	return XDate{
		Time: time.Date(year, months[month], day, 0, 0, 0, 0, time.UTC),
	}
}

//DateTime date&time to XDate
func DateTime(day int, month int, year int, hour int, min int, sec int) XDate {
	return XDate{
		Time: time.Date(year, months[month], day, hour, min, sec, 0, time.UTC),
	}
}

//Str2Time string date to XDate
func Str2Time(str string, _format ...string) XDate {
	var format = ""
	if len(_format) > 0 {
		format = _format[0]
	} else if len(_format) == 0 && len(str) == 19 {
		format = "Y-m-d H:i:s"
	} else if len(_format) == 0 && len(str) == 10 {
		format = "Y-m-d"
	} else {
		format = "Y-m-d"
	}
	format = PFormatConv(format)
	time, err := time.Parse(format, str)
	if err != nil {
		return XDate{}
	}
	return XDate{time}
}
