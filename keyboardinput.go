package keyboardinput

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)
//коммент для тестов 2
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	getInputFronKeyb, err := reader.ReadString('\n')
	if err != nil {
		error := errors.New("Что-то не то прилетело в ввод!")
		return 0, error
	}

	getInputFronKeyb = strings.TrimSpace(getInputFronKeyb)
	number, err := strconv.ParseFloat(getInputFronKeyb, 64)
	if err != nil {
		error := errors.New("Что-то не то при конвертации ввода в число!")
		return 0, error
	}

	return number, nil

}
