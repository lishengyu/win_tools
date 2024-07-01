package proc_time

import (
	"fmt"
	"strconv"
	"time"
	"tools/global"
)

func init() {
	global.Register(Int2Str, "时间int转字符")
	global.Register(Str2Int, "时间字符转int")
}

func Int2Str(in string) string {
	t, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return fmt.Sprintf("Parse Int err: %v\n", err)
	}

	var buf string
	buf += fmt.Sprintf("%s\n", time.Unix(t, 0).Format("2006-01-02 15:04:05"))
	buf += fmt.Sprintf("%s\n", time.Unix(t, 0).Format("20060102150405"))
	return buf
}

func Str2Int(in string) string {
	t, err := time.Parse("2006-01-02 15:04:05", in)
	if err != nil {
		return fmt.Sprintf("Parse timestr err: %v\n", err)
	}
	return string(t.Unix())
}
