package scheduler

import (
	"regexp"
	"strings"
)
//匹配IP
var regexpForIP = regexp.MustCompile(`((?:(?:25[0-5]|2[0-4]\d|[01]?\d?\d)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d?\d))`)

var regexpForDomains = []*regexp.Regexp{
	// *.xx or *.xxx.xx
	regexp.MustCompile(`\.(com|com\.\w{2})$`),
	regexp.MustCompile(`\.(gov|gov\.\w{2})$`),
	regexp.MustCompile(`\.(net|net\.\w{2})$`),
	regexp.MustCompile(`\.(org|org\.\w{2})$`),
	// *.xx
	regexp.MustCompile(`\.me$`),
	regexp.MustCompile(`\.biz$`),
	regexp.MustCompile(`\.info$`),
	regexp.MustCompile(`\.name$`),
	regexp.MustCompile(`\.mobi$`),
	regexp.MustCompile(`\.so$`),
	regexp.MustCompile(`\.asia$`),
	regexp.MustCompile(`\.tel$`),
	regexp.MustCompile(`\.tv$`),
	regexp.MustCompile(`\.cc$`),
	regexp.MustCompile(`\.co$`),
	regexp.MustCompile(`\.\w{2}$`),
}

// getPrimaryDomain 用于获取给定主机名的主域名。
//IP直接返回IP
//www.zhihu.com返回zhihu.com
func getPrimaryDomain(host string) (string, error) {
	host = strings.TrimSpace(host)
	if host == "" {
		return "", genError("empty host")
	}
	//是IPv4的话
	if regexpForIP.MatchString(host) {
		return host, nil
	}
	var suffixIndex int
	for _, re := range regexpForDomains {
		pos := re.FindStringIndex(host)
		if pos != nil {
			//返回类似".com"的开始位置
			suffixIndex = pos[0]
			//fmt.Println(suffixIndex)
			break
		}
	}
	if suffixIndex > 0 {
		var pdIndex int
		firstPart := host[:suffixIndex]//www.zhihu
		index := strings.LastIndex(firstPart, ".")
		if index < 0 {
			pdIndex = 0
		} else {
			pdIndex = index + 1
		}
		return host[pdIndex:], nil//zhihu.com
	}
	return "", genError("unrecognized host")
}
