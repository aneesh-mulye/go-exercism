package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[=~*-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	i := 0
	for _, line := range lines {
		if re.MatchString(line) {
			i++
		}
	}
	return i
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	var ret []string
	for _, line := range lines {
		if re.MatchString(line) {
			ret = append(ret,
				"[USR] "+re.FindStringSubmatch(line)[1]+" "+line)
		} else {
			ret = append(ret, line)
		}
	}
	return ret
}
