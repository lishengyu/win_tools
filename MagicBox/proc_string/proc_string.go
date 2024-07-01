package proc_string

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"tools/global"
)

func init() {

}

func init() {
	global.Register(Upper, "字符串转大写")
	global.Register(Lower, "字符串转小写")
	global.Register(LogSplit, "字符串拆分")
	global.Register(Encry, "字符串base64编码")
	global.Register(Decry, "字符串base64解码")
	global.Register(CalMd5, "字符串计算MD5")
	global.Register(StringLen, "计算字符长度")
	global.Register(JsonMarshal, "json格式美化")
}

func Upper(in string) string {
	return strings.ToUpper(in)
}

func Lower(in string) string {
	return strings.ToLower(in)
}

func LogSplit(in string) string {
	var buf string
	fs := strings.Split(in, "|")
	for i, v := range fs {
		buf += fmt.Sprintf("%d\t%s\r\n", i, v)
	}
	return buf
}

func Encry(in string) string {
	return base64.StdEncoding.EncodeToString([]byte(in))
}

func Decry(in string) string {
	data, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return fmt.Sprintf("decoding str err:%v\n", err)
	}
	return string(data)
}

func CalMd5(in string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(in)))
}

func StringLen(in string) string {
	return fmt.Sprintf("%d", len(in))
}

func JsonMarshal(in string) string {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(in), "", "    ")
	if err != nil {
		return fmt.Sprintf("%v\n", err)
	}

	return string(prettyJSON.Bytes())
}
