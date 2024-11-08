package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetInteger() (int, error) {
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')

	if err != nil {
		return 0, err
	}
	i = strings.TrimSpace(i)
	n, _ := strconv.Atoi(i)

	if err != nil {
		return 0, err
	}

	return n, nil
}

func GetFloat() (float64, error) {
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')

	if err != nil {
		return 0, err
	}
	i = strings.TrimSpace(i)
	n, _ := strconv.ParseFloat(i, 64)

	if err != nil {
		return 0, err
	}

	return n, nil
}
