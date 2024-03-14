package utils

import (
	"bufio"
	"net/http"
	"os"
	"strings"
)

func LoadCookie(FileName string) []*http.Cookie {
	file, err := os.Open(FileName)
	if err != nil {
		return nil
	}
	defer file.Close()

	// 创建一个切片来保存每一行数据
	Cookie := make([]*http.Cookie, 0)

	// 创建一个 Scanner 来按行读取文件
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 将每一行文本添加到切片中
		data := strings.Split(scanner.Text(), ";")
		for _, v := range data {
			t := strings.Split(v, "=")
			tc := new(http.Cookie)
			tc.Name = t[0]
			tc.Value = t[1]
			Cookie = append(Cookie, tc)
		}
	}
	// 检查扫描时是否发生错误
	if err := scanner.Err(); err != nil {
		return nil
	}
	return Cookie
}
