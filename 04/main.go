package main

import "fmt"
import "bufio"
import "io"
import "os"
import "strings"

type passport struct {
	fields map[string]string
}

func (p *passport) isValid() bool {
	for _, field := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		if _, ok := p.fields[field]; !ok {
			return false
		}
	}
	return true
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
			pieces := strings.Split(line, " ")
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
