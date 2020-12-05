package main

import "testing"
import "github.com/stretchr/testify/assert"
import "bufio"
import "strings"
import "fmt"

func TestParseBadPassport(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("e3cm"))

	passport, err := parsePassport(reader)
	assert.Nil(t, passport)
	assert.Error(t, err)
}

func TestParseGoodPassport(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm"))

	passport, err := parsePassport(reader)
	assert.NotNil(t, passport)
	assert.NoError(t, err)

	assert.Equal(t, 8, len(passport.fields))
}

func TestParseGoodMultiplePassport(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\n\nbyr:1937 iyr:2017 cid:147 hgt:183cm"))

	passport, err := parsePassport(reader)
	assert.NotNil(t, passport)
	assert.NoError(t, err)

	assert.Equal(t, 4, len(passport.fields))
}

func TestParseMultiplePassports(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\n\nbyr:1937 iyr:2017 cid:147 hgt:183cm"))

	passports, err := parsePassports(reader)
	assert.NotNil(t, passports)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(passports))

	assert.Equal(t, 4, len(passports[0].fields))
	assert.Equal(t, 4, len(passports[1].fields))
}

func TestParseMultiplePassports2(t *testing.T) {
	input := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`
	reader := bufio.NewReader(strings.NewReader(input))

	passports, err := parsePassports(reader)
	assert.NotNil(t, passports)
	assert.NoError(t, err)

	assert.Equal(t, 4, len(passports))

	assert.Equal(t, 8, len(passports[0].fields))
	assert.Equal(t, 7, len(passports[1].fields))
	assert.Equal(t, 7, len(passports[2].fields))
	assert.Equal(t, 6, len(passports[3].fields))
}

func TestIsValid(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm"))

	passport, err := parsePassport(reader)
	assert.NotNil(t, passport)
	assert.NoError(t, err)
	assert.Equal(t, 8, len(passport.fields))
	assert.True(t, passport.isValid())

	input := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`
	reader = bufio.NewReader(strings.NewReader(input))
	passports, err := parsePassports(reader)
	assert.True(t, passports[0].isValid())
	assert.False(t, passports[1].isValid())
	assert.True(t, passports[2].isValid())
	assert.False(t, passports[3].isValid())
}

func TestIsValidHeight(t *testing.T) {
	tests := []struct {
		input string
		valid bool
	}{
		{"", false},
		{"123", false},
		{"20cm", false},
		{"20in", false},
		{"200cm", false},
		{"100in", false},
		{"160cm", true},
		{"150cm", true},
		{"50in", false},
		{"60in", true},
		{"50incm", false},
	}

	for _, tt := range tests {
		if tt.valid {
			assert.True(t, isValidHeight(tt.input), fmt.Sprintf("%s == %t", tt.input, tt.valid))
		} else {
			assert.False(t, isValidHeight(tt.input), fmt.Sprintf("%s == %t", tt.input, tt.valid))
		}
	}
}

func TestIsValidHairColor(t *testing.T) {
	tests := []struct {
		input string
		valid bool
	}{
		{"", false},
		{"123", false},
		{"#123456", true},
		{"123456", false},
		{"#123abc", true},
		{"#abcdef", true},
		{"#qwerty", false},
		{"#123456asdf", false},
	}

	for _, tt := range tests {
		if tt.valid {
			assert.True(t, isValidHairColor(tt.input), fmt.Sprintf("%s == %t", tt.input, tt.valid))
		} else {
			assert.False(t, isValidHairColor(tt.input), fmt.Sprintf("%s == %t", tt.input, tt.valid))
		}
	}
}
