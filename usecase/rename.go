package usecase

import (
	"regexp"
	"strconv"
)

func Fibonacci(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]

}

func Rename(name string, num int) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		return "", err
	}
	nameOnly := reg.ReplaceAllString(name, "")

	num = Fibonacci(num)
	nickname := nameOnly + "-" + strconv.Itoa(num)

	return nickname, nil
}
