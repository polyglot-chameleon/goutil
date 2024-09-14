package util

import (
	"bufio"
	"os"
	"strings"
)

func LoadDotEnv(fpath string) {
	file, err := os.Open(fpath)
	Check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res := strings.Split(scanner.Text(), "=")
		os.Setenv(res[0], res[1])
	}

}
