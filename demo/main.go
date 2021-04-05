package main

import (
	"fmt"

	"github.com/Andyfoo/go-xdate"
)

func main() {
	fmt.Println(xdate.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(xdate.Unix(1588009973, 0).Format("2006-01-02 15:04:05"))
	fmt.Println(xdate.PFormatConv("Y-m-d H:i:s"))
	fmt.Println(xdate.Now().PFormat("Y-m-d H:i:s"))
	fmt.Println(xdate.Now().Format(xdate.YMD))
	fmt.Println(xdate.Now().UTC().Format(xdate.TT))
	fmt.Println(xdate.Now().UTC().PFormat("Y-m-d H:i:s"))
	fmt.Println(xdate.Now().Unix())
	fmt.Println(xdate.Now().UnixMilli())
	fmt.Println(xdate.Now().UnixNano())
	fmt.Println(xdate.Now().Weekday())
	fmt.Println(xdate.Now().WeekdayStr(xdate.WeekType_cn))
	fmt.Println(xdate.Now().WeekdayStr(xdate.WeekType_cnShort))
	fmt.Println(xdate.Now().WeekdayStr(xdate.WeekType_en))
	fmt.Println(xdate.Now().WeekdayStr(xdate.WeekType_enShort))
	fmt.Println(xdate.Now().UTC().Local().Format(xdate.TT))
	fmt.Println(xdate.Str2Time("2018-04-23 23:11:23", "Y-m-d H:i:s").PFormat("Y-m-d H:i:s"))
}
