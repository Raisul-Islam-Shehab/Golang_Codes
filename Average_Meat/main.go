package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const daysInWeek = 7

func openFile(fileName string) (*os.File, error) {
	fmt.Println("Opening file", fileName)
	return os.Open(fileName)
}

func closeFile(file *os.File) error {
	fmt.Println("Closing the file")
	return file.Close()
}

func getFloats(fileName string) ([]float64, error) {
	file, err := openFile(fileName)

	if err != nil {
		return nil, err
	}
	defer closeFile(file)

	var numbers []float64
	newScanner := bufio.NewScanner(file)

	for newScanner.Scan() {
		number, err := strconv.ParseFloat(newScanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	if newScanner.Err() != nil {
		return nil, newScanner.Err()
	}
	return numbers, nil
}

func main() {
	fileName := "data.txt"

	numbers, err := getFloats(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var sum float64
	for _, value := range numbers {
		sum += value
	}

	fmt.Printf("Average meat needed in a week is %0.2f kg\n", sum/float64(daysInWeek))
}
