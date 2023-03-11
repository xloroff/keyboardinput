// keyboardinput Пакет который позволяет читать данные из файла
package keyboardinput

import (
	"bufio"
	"os"
	"strconv"
)

// Readdata функци по чтению данных
func Readdata(fileName string) ([]float64, error) {
	var myaso []float64
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		stroka := scanner.Text()
		tempnum, err := strconv.ParseFloat(stroka, 64)
		if err != nil {
			return nil, err
		}
		myaso = append(myaso, tempnum)
	}

	err = file.Close()
	if err != nil {
		return myaso, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return myaso, nil
}
