package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rotateArray(array []string, size, rotation int) string {

	var newArray []string

	for i := 0; i < rotation; i++ {
		newArray = array[1:size]

		newArray = append(newArray, array[0])

		array = newArray
	}

	return strings.Join(array, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()

	r := strings.Split(string(line), " ")

	size, _ := strconv.Atoi(r[0])
	rotation, _ := strconv.Atoi(r[1])

	line2, _, _ := reader.ReadLine()

	array := strings.Split(string(line2), " ")

	fmt.Println(rotateArray(array, size, rotation))
}
