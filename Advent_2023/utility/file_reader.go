package utility

import (
	"bufio"
	"fmt"
	"os"
)

func LineIterator(filename string) (func() (string, bool), error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	return func() (string, bool) {
		if scanner.Scan() {
			return scanner.Text(), true
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from file:", err)
		}
		file.Close()

		return "", false
	}, nil
}
func LineIteratorBytes(filename string) (func() ([]byte, bool), error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	return func() ([]byte, bool) {
		if scanner.Scan() {
			return scanner.Bytes(), true
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from file:", err)
		}
		file.Close()

		return nil, false
	}, nil
}
func ReadWholeFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
