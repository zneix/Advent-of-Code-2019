package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	allPassports := strings.Split(string(data), "\n\n")

	validEyeColours := map[string]struct{}{
		"amb": {},
		"blu": {},
		"brn": {},
		"gry": {},
		"grn": {},
		"hzl": {},
		"oth": {},
	}
	validPassportCount := 0
	for _, passport := range allPassports {
		goodFields := 0
		//bad := false
		fields := strings.Fields(passport)
		if len(fields) < 7 {
			continue // passport won't be valid without at least 7 good fields
		}
		for _, field := range fields {
			key, val := field[:3], field[4:]
			switch key {
			case "byr": // birth year
				year, _ := strconv.Atoi(val)
				if year >= 1920 && year <= 2002 {
					goodFields++
				}
			case "iyr": // issue year
				year, _ := strconv.Atoi(val)
				if year >= 2010 && year <= 2020 {
					goodFields++
				}
			case "eyr": // expiration year
				year, _ := strconv.Atoi(val)
				if year >= 2020 && year <= 2030 {
					goodFields++
				}
			case "hgt": // heightStr
				heightStr, unit := val[:len(val)-2], val[len(val)-2:]
				height, err := strconv.Atoi(heightStr)
				if err != nil {
					//fmt.Println("Failed for format height:", err)
					continue
				}
				switch unit {
				case "cm":
					if height >= 150 && height <= 193 {
						//fmt.Println("good cm", val)
						goodFields++
					}
				case "in":
					if height >= 59 && height <= 76 {
						//fmt.Println("good inch", val)
						goodFields++
					}
				}
			case "hcl": // hair colour
				if val[0] == '#' {
					ok, _ := regexp.MatchString("^#[0-9a-f]{6}$", val)
					if ok {
						goodFields++
					}
				}
			case "ecl": // eye colour
				if _, ok := validEyeColours[val]; ok {
					//fmt.Printf("%q %q\n", key, val)
					goodFields++
				}
			case "pid": // passport ID
				if ok, _ := regexp.MatchString(`^\d{9}$`, val); ok {
					goodFields++
				}
			case "cid": // country ID (optional)
			}
		}
		if goodFields >= 7 {
			validPassportCount++
			//fmt.Printf("%d, %d/%d %#v\n", i, goodFields, len(fields), fields)
		}
	}

	fmt.Println(validPassportCount)
}
