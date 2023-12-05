package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ParseInt(s string) int {
	s = strings.TrimSpace(s)
	k, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return k
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min[T int | uint](a, b T) T {
	if a > b {
		return b
	}
	return a
}

func Abs[T int | uint](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

func Index[T comparable](x T, s []T) (int, bool) {
	for i, t := range s {
		if x == t {
			return i, true
		}
	}
	return -1, false
}

func IndexKey[K comparable, T interface{}](key K, s []T, f func(T) K) (int, bool) {
	for i, t := range s {
		if key == f(t) {
			return i, true
		}
	}
	return -1, false
}

func Delete[T comparable](t T, s []T) ([]T, bool) {
	for i, t2 := range s {
		if t == t2 {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return s, false
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func ListOfNumbers(s, sep string) []int {
	numbers := make([]int, 0)
	for _, n := range strings.Split(strings.TrimSpace(s), sep) {
		n = strings.TrimSpace(n)
		if len(n) > 0 {
			numbers = append(numbers, ParseInt(n))
		}
	}
	return numbers
}

func Equals[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func CheckErrorP(err error) {
	if err != nil {
		panic(err)
	}
}
func ReadFileLinesForDay(day string, test bool) []string {
	path := "./day" + day + "/input.txt"
	if test {
		path = "./day" + day + "/test.txt"
	}
	return ReadFileLines(path)
}

func ReadFileLines(name string) []string {
	readFile, err := os.Open(name)
	defer readFile.Close()
	CheckErrorP(err)
	fileScanner := bufio.NewScanner(readFile)
	lines := make([]string, 0)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines
}
