package main

import "fmt"
import "bufio"
import "io"
import "os"
import "strings"
import "strconv"

type passport struct {
	fields map[string]string
}

func isValidYear(input string, min int, max int) bool {
	if len(input) != 4 {
		return false
	}

	inputNumber, err := strconv.Atoi(input)
	if err != nil {
		return false
	} else if inputNumber < min || inputNumber > max {
		return false
	}
	return true
}

func isValidHairColor(input string) bool {
	if len(input) != 7 {
		return false
	} else if input[0] != '#' {
		return false
	}

	for _, c := range input[1:7] {
		_, err := strconv.ParseInt(string(c), 16, 16)
		if err != nil {
			return false
		}
	}
	return true
}

func isValidHeight(input string) bool {
	if len(input) < 3 {
		return false
	}
	unit := input[len(input)-2:]
	if unit == "in" {
		value, err := strconv.Atoi(input[0 : len(input)-2])
		return err == nil && value >= 59 && value <= 76
	} else if unit == "cm" {
		value, err := strconv.Atoi(input[0 : len(input)-2])
		return err == nil && value >= 150 && value <= 193
	}

	return false
}

func isValidEyeColor(input string) bool {
	return input == "amb" ||
		input == "blu" ||
		input == "brn" ||
		input == "gry" ||
		input == "grn" ||
		input == "hzl" ||
		input == "oth"
}

func isValidPassportID(input string) bool {
	if len(input) != 9 {
		return false
	}
	_, err := strconv.Atoi(input)
	return err == nil
}

func (p *passport) isValid() bool {
	return isValidYear(p.fields["byr"], 1920, 2002) &&
		isValidYear(p.fields["iyr"], 2010, 2020) &&
		isValidYear(p.fields["eyr"], 2020, 2030) &&
		isValidHeight(p.fields["hgt"]) &&
		isValidHairColor(p.fields["hcl"]) &&
		isValidEyeColor(p.fields["ecl"]) &&
		isValidPassportID(p.fields["pid"])
}

func parsePassports(bufReader *bufio.Reader) ([]passport, error) {
	var passports []passport
	for {
		passport, _ := parsePassport(bufReader)

		if passport != nil {
			passports = append(passports, *passport)
		} else {
			break
		}

	}
	return passports, nil
}

func parsePassport(bufReader *bufio.Reader) (*passport, error) {
	fields := make(map[string]string)
	for {
		line, err := bufReader.ReadString('\n')

		if len(line) == 1 && line[0] == '\n' {
			return &passport{fields}, nil
		}

		if len(line) > 0 {
			pieces := strings.Split(strings.TrimSpace(line), " ")
			for _, piece := range pieces {
				piecePieces := strings.Split(piece, ":")
				if len(piecePieces) != 2 {
					return nil, fmt.Errorf("Bad piece count for '%s': %d", piece, len(piecePieces))
				}
				fields[piecePieces[0]] = piecePieces[1]
			}
		}
		if len(line) == 0 || err == io.EOF {
			if len(fields) == 0 {
				return nil, nil
			}
			return &passport{fields}, nil

		} else if err != nil {
			return nil, err

		}
	}
}

func main() {
	passports, err := parsePassports(bufio.NewReader(os.Stdin))
	if err != nil {
		panic(err)
	}

	count := 0
	for _, passport := range passports {
		if passport.isValid() {
			count++
		}
	}
	fmt.Printf("%d\n", count)
}
