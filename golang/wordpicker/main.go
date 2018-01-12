package main

import (
	"github.com/herlegs/programQ/golang/hard"
	"os"
	"fmt"
	"bufio"
	"regexp"
	"strings"
	"strconv"
	"sync"
	"github.com/herlegs/programQ/golang/wordpicker/dictionary"
)

const(
	infoMsg = `Please enter characters and word length. Eg. "ccaatt 3"
Or enter q to quit:
`
	errorMsg = `format should be "{characters} {length}"`
	resultMsg = "Possible words:"
	regex = "^([a-zA-Z]+)\\s+(\\d+)$"
)

var(
	replyMap = map[string]string{
		"(?i)(i)?[[:space:]]*love[[:space:]]*you":"I love you too!",
	}
	initOnce = &sync.Once{}
	wordRepo map[string]bool
)

//go:generate go-bindata -pkg=dictionary -o dictionary/dictionary.go dictionary/wordlist

func main() {
	initWordRepo()
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print(infoMsg)
	for {
		input, _ := inputReader.ReadString('\n')
		exit, hit := controlCenter(input)
		if exit {
			return
		}
		if hit {
			continue
		}
		charList, length, err := parseParams(input)
		if err != nil {
			fmt.Println(err)
		}
		allWords := hard.PickWords(charList, length)
		filterdWords := filterWords(allWords)
		fmt.Println(resultMsg)
		fmt.Println(strings.Join(filterdWords,"\n"))
	}
}

//return (exit, hit)
func controlCenter(s string) (bool, bool) {
	s = strings.ToLower(strings.Trim(s, " \n"))
	if s == "q" {
		return true, false
	}
	for reg, reply := range replyMap {
		regex := regexp.MustCompile(reg)
		if regex.MatchString(s) {
			fmt.Println(reply)
			return false, true
		}
	}
	return false, false
}

func parseParams(params string) (string, int, error) {
	params = strings.Trim(params, " \n")
	reg := regexp.MustCompile(regex)
	if !reg.MatchString(params) {
		return "", 0, fmt.Errorf(errorMsg)
	}
	groups := reg.FindStringSubmatch(params)
	length,_ := strconv.ParseInt(groups[2], 10, 64)
	return groups[1], int(length), nil
}

func filterWords(list []string) []string {
	result := []string{}
	for _, item := range list {
		if isWord(item) {
			result = append(result, item)
		}
	}
	return result
}

func isWord(str string) bool {
	return wordRepo[strings.ToLower(str)]
}

func initWordRepo() {
	initOnce.Do(func() {
		wordRepo = make(map[string]bool)
		bytes, err := dictionary.Asset("dictionary/wordlist")
		if err != nil {
			fmt.Printf("Something wrong with loading dictionary: %v\n", err)
		}
		wordList := strings.Split(string(bytes), "\n")
		for _, word := range wordList {
			wordRepo[word] = true
		}
	})
}
