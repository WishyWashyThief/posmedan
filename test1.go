package main

import (
	"fmt"
	"sort"
	"strings"
)

// test number 1
func letterCount(str string) {
	charMap := make(map[rune]int)

	strLower := strings.ToLower(str)
	for _, char := range strLower {
		if char == ' ' {
			continue
		} else {
			charMap[char]++
		}
	}

	for i, j := range charMap {
		fmt.Printf("%c = %d, ", i, j)
	}

	fmt.Println()
}

// test number 2
func charOrganizer(strs []string) {
	//make map
	str := strings.Join(strs, "")
	charMap := make(map[rune]int)
	for _, char := range str {
		charMap[char]++
	}

	//count  freqcount
	var keys []rune
	values := make(map[int]int)
	for key, value := range charMap {
		keys = append(keys, key)
		values[value]++
	}
	//sort char by frequency
	sort.SliceStable(keys, func(i, j int) bool {
		return charMap[keys[i]] > charMap[keys[j]]
	})

	//print by freq order then ascii order
	var asciiSortStr []string
	flag := 1
	for _, char := range keys {
		flag = values[charMap[char]]
		if flag == 1 {
			if len(asciiSortStr) != 0 {
				asciiSortStr = append(asciiSortStr, string(char))
				sort.SliceStable(asciiSortStr, func(i, j int) bool {
					return asciiSortStr[i] < asciiSortStr[j]
				})

				newStr := strings.Join(asciiSortStr, "")

				fmt.Print(newStr)
				asciiSortStr = nil
			} else {
				fmt.Print(string(char))
			}
		} else {
			asciiSortStr = append(asciiSortStr, string(char))
			values[charMap[char]]--
		}
	}
	fmt.Println()
}

func main() {
	var str string
	var strs []string

	str = "We Always Mekar"
	letterCount(str)

	str = "coding is fun"
	letterCount(str)

	strs = []string{"Abc", "bCd"}
	charOrganizer(strs)

	strs = []string{"One", "Oke"}
	charOrganizer(strs)

	strs = []string{"Pendanaan", "Terproteksi", "Untuk", "Dampak", "Berarti"}
	charOrganizer(strs)
}
