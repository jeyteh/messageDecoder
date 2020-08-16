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
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	// wait for input
	for len(text) > 0 {
		keyString, _ := reader.ReadString('\n')

		keyString = strings.TrimSpace(keyString)

		m := make(map[string]string)

		var defCodeList = generateDefaultBinary(len(text))

		for i := 0; i < len(defCodeList); i++ {
			m[defCodeList[i]] = string(text[i])
		}

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
	}
}

func getCode(code string, subLen int) string {
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
}
