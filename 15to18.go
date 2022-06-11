package sf15to18

import (
	"os"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func HandleArgs(source string, args []string) []string {
	if source != "" {
		dat, err := os.ReadFile(source)
		check(err)
		args = []string{string(dat)}
	}
	return args
}

func GetValidIds(in []string) []string {
	r, _ := regexp.Compile("([a-zA-Z0-9]{15})")
	return r.FindAllString(strings.Join(in, " "), -1)
}

func Convert(input string) string {
	if len(input) == 18 {
		return input
	}

	sum := ""
	for i := 0; i < 3; i++ {
		loop := 0
		for j := 0; j < 5; j++ {
			current := rune(input[i*5+j])
			if current >= 'A' && current <= 'Z' {
				loop = loop + (1 << j)
			}
		}
		sum += string([]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ012345")[loop])
	}
	return input + sum
}
