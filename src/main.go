package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Queue struct {
	item []string
}

func (q *Queue) Enqueue(code string) {
	q.item = append(q.item, code)
}

func (q *Queue) Dequeue() string {
	poll := q.item[0]
	q.item = q.item[1:]
	return poll
}

func (q *Queue) Peek() string {
	return q.item[0]
}

func main() {

	// const text = "TNM AEIOU"
	// var keyString = "0010101100011101000100111011001111000"

	// const text = "$#**\\"
	// var keyString = "0100000101101100011100001000"

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	// var keyString string
	text = strings.TrimSpace(text)

	for len(text) > 0 {
		// fmt.Print("-> ")
		keyString, _ := reader.ReadString('\n')

		keyString = strings.TrimSpace(keyString)

		// if strings.Contains(keyString, "0") || strings.Contains(keyString, "1") {

		m := make(map[string]string)

		var defCodeList = generateDefaultBinary(len(text))

		for i := 0; i < len(defCodeList); i++ {
			m[defCodeList[i]] = string(text[i])
		}

		// fmt.Println(m)
		defaultCode := 3
		code := getCode(keyString, defaultCode) //string to count

		sum := countLength(code)

		keyString = keyString[defaultCode:len(keyString)] //remove first 3 char from keyString

		for {
			if len(keyString) <= 0 {
				break
			}

			codeToDecode := getCode(keyString, sum)

			if !strings.Contains(codeToDecode, "0") {

				// remove the terminated keys
				keyString = keyString[sum:len(keyString)]

				// get the next first 3 key to count
				codeKey := keyString[0:defaultCode]

				//remove the  first 3 key
				keyString = keyString[defaultCode:len(keyString)]

				sum = countLength(codeKey)

				if sum == 0 {
					continue
				}

				codeToDecode = getCode(keyString, sum)
				continue
			}

			fmt.Print(m[codeToDecode])
			keyString = keyString[sum:len(keyString)]
		}
		fmt.Println("")
		// }
		break
	}
}

func getCode(code string, subLen int) string {
	// fmt.Println(code)
	// fmt.Println(subLen)
	cd := code[:subLen]
	return cd
}

func countLength(check string) int {
	sum := 0
	if check[2] == '1' {
		sum = sum + 1
	}

	if check[1] == '1' {
		sum = sum + 2
	}

	if check[0] == '1' {
		sum = sum + 4
	}

	return sum
}

func generateDefaultBinary(len int) []string {
	var list []string
	q := Queue{}
	q.Enqueue("0")
	q.Enqueue("1")

	for i := 1; i <= len; i++ {
		q.Enqueue(q.Peek() + "0")
		q.Enqueue(q.Peek() + "1")

		// pop the front element and print it
		poll := q.Dequeue()

		if !strings.Contains(poll, "0") {
			i--
			continue
		}
		list = append(list, poll)
	}

	return list
	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(defCodeList[i])
	// }
}
