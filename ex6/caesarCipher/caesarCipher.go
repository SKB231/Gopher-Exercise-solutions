package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func caesarCipher(s string, k int32) string {
	// Write your code here
	finalString := ""
	for _, char := range s {
		if unicode.IsUpper(char) {
			// Upper case character
			lowerBound := rune('A')
			char = ((char-lowerBound)+k)%26 + lowerBound
			fmt.Println(string(char))
			finalString += string(char)
		} else if unicode.IsLower(char) {
			// Lower case character
			lowerBound := rune('a')
			char = ((char-lowerBound)+k)%26 + lowerBound
			finalString += string(char)
			fmt.Println(string(char))

		} else {
			finalString += string(char)
			fmt.Println(string(char))

		}
	}

	return finalString
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	_, err = strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	// n := int32(nTemp)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
