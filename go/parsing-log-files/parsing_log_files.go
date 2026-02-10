package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	r := regexp.MustCompile(`^(\[TRC]|\[DBG]|\[INF]|\[WRN]|\[ERR]|\[FTL])(.)*`)

	return r.Match([]byte(text))
	// [TRC] [DBG] [INF] [WRN] [ERR] [FTL]
}

func SplitLogLine(text string) []string {
	r := regexp.MustCompile(`<[^a-zA-Z]*>`)
	return r.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	ans := 0
	for _, line := range lines {
		if re.MatchString(line) {
			ans++
		}
	}
	return ans
}

func RemoveEndOfLineText(text string) string {
	r := regexp.MustCompile(`end-of-line(\d*)`)
	for {
		position := r.FindStringIndex(text)
		if position == nil {
			return text
		}
		text = text[:position[0]] + text[position[1]:]
	}

}

func TagWithUserName(lines []string) []string {
	var newLines []string
	r := regexp.MustCompile(`(User ( )*(\w)+)`)
	for _, line := range lines {
		find := r.FindString(line)
		if len(find) == 0 {
			newLines = append(newLines, line)
			continue
		}
		user := strings.TrimSpace(r.FindString(line)[4:])
		nLine := "[USR] " + user + " " + line
		newLines = append(newLines, nLine)
	}
	return newLines
}
