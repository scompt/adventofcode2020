package main

import "strings"
import "testing"
import "bufio"
import "github.com/stretchr/testify/assert"

func TestCountValidPasswords(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("1-3 a: abcde\n2-4 b: bcdef\n"))
	count := countValidPasswords(reader)

	assert.Equal(t, 1, count)
}

func TestPasswordRuleIsValid2(t *testing.T) {
	rule := passwordRule{1, 3, 'a'}
	assert.True(t, rule.isValid2("aa"))
	assert.False(t, rule.isValid2("ba"))
	assert.False(t, rule.isValid2("aaaa"))
	assert.False(t, rule.isValid2("bb"))
}

func TestPasswordRuleIsValid(t *testing.T) {
	rule := passwordRule{1, 3, 'a'}
	assert.True(t, rule.isValid("aa"))
	assert.False(t, rule.isValid("aaaa"))
	assert.False(t, rule.isValid("bb"))
}

func TestParseLine(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("1-3 a: abcde\n2-4 b: bcdef\n"))
	rule, password, err := parseLine(reader)
	assert.NoError(t, err)
	assert.NotNil(t, rule)
	assert.Equal(t, "abcde", password)

	rule, password, err = parseLine(reader)
	assert.NoError(t, err)
	assert.NotNil(t, rule)
	assert.Equal(t, "bcdef", password)
}

func TestParsePasswordRule(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("1-3 a: abcde"))

	rule, err := parsePasswordRule(reader)

	assert.NoError(t, err)
	assert.NotNil(t, rule)
	assert.Equal(t, 1, rule.minLength)
	assert.Equal(t, 3, rule.maxLength)
	assert.Equal(t, byte('a'), rule.letter)

	buf := []byte{1}
	n, err := reader.Read(buf)

	assert.Equal(t, 1, n)
	assert.Equal(t, byte(' '), buf[0])
}
