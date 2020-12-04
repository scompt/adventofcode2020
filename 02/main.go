package main

import "bufio"
import "strconv"
import "os"
import "fmt"

type passwordRule struct {
	minLength int
	maxLength int
	letter    byte
}

func countValidPasswords(bufReader *bufio.Reader) int {
	count := 0
	for {
		rule, password, err := parseLine(bufReader)

		if rule != nil && rule.isValid(password) {
			count++
		}
		if err != nil {
			break
		}
	}

	return count
}

func (rule *passwordRule) isValid(password string) bool {
	count := 0

	for _, c := range password {
		if byte(c) == rule.letter {
			count++
		}
	}

	return count >= rule.minLength && count <= rule.maxLength
}

func parseLine(bufReader *bufio.Reader) (*passwordRule, string, error) {
	rule, err := parsePasswordRule(bufReader)
	if err != nil {
		return nil, "", err
	}
	_, err = bufReader.ReadByte()
	password, err := bufReader.ReadString('\n')
	if err != nil {
		return rule, "", err
	}
	return rule, password[:len(password)-1], nil
}

func parsePasswordRule(bufReader *bufio.Reader) (*passwordRule, error) {
	str, err := bufReader.ReadString('-')
	if err != nil {
		return nil, err
	}
	minLength, _ := strconv.Atoi(str[:len(str)-1])
	if err != nil {
		return nil, err
	}
	str, _ = bufReader.ReadString(' ')
	if err != nil {
		return nil, err
	}
	maxLength, _ := strconv.Atoi(str[:len(str)-1])
	if err != nil {
		return nil, err
	}
	str, _ = bufReader.ReadString(':')
	if err != nil {
		return nil, err
	}
	letter := byte(str[0])
	return &passwordRule{minLength, maxLength, letter}, nil
}

func main() {
	valid := countValidPasswords(bufio.NewReader(os.Stdin))
	fmt.Printf("%d\n", valid)
}
