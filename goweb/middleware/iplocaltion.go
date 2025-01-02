package middleware

import (
	"fmt"
	"strings"

	"path"          // 路径
	"path/filepath" // 文件路径
	"runtime"       // 运行时

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var ip2region *xdb.Searcher
 // 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// IP归属地查询
func init() {
	var cBuff []byte
	dbPath := filepath.Join(getCurrentAbPathByCaller(), "ip2region.xdb")
	fmt.Println(dbPath)
    var err error
    cBuff, err = xdb.LoadContentFromFile(dbPath)
    if err != nil {
        fmt.Printf("failed to load content from `%s`: %s\n", dbPath, err)
        return
    }
	ip2region,err = xdb.NewWithBuffer(cBuff)
	if err != nil {
		fmt.Printf("failed to initialize ip2region: %s\n", err)
		return
	}
}

// IP归属地查询
func IpLocation(ip string) string {
	if ip2region == nil {
		return "未知"
	}
	result, err := ip2region.SearchByStr(ip)
	if err != nil {
		return "未知"
	}
	tmp := strings.Split(result, "|")
	result = tmp[0]
	if tmp[2] != "0" {
		result = result + "-" + tmp[2]
	}
	if tmp[3] != "0" {
		result = result + "-" + tmp[3]
	}
	return result
}

// 关闭IP归属地查询
func CloseIpLocation() {
	ip2region.Close()
}

//手机号码运营商查询
func MobileLocation(mobile string) string {
	if len(mobile) != 11 {
		return "未知"
	}
	switch mobile[0:3] {
	case "130", "131", "132", "155", "156", "185", "186":
		return "中国联通"
	case "133", "153", "180", "181", "189":
		return "中国电信"
	case "134", "135", "136", "137", "138", "139", "147", "150", "151", "152", "157", "158", "159", "182", "183", "187", "188":
		return "中国移动"
	default:
		return "未知"
	}
}