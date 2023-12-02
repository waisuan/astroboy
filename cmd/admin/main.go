package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	res := 0

	for _, g := range strings.Split(input(), "\n") {
		fmt.Println(g)

		tmp := strings.Split(g, ":")
		//tmp2 := strings.Split(tmp[0], " ")
		//gameId, _ := strconv.Atoi(tmp2[1])

		tmp3 := strings.Split(tmp[1], ";")

		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		for _, s := range tmp3 {
			s := strings.TrimSpace(s)
			s = strings.Replace(s, ", ", "|", -1)
			tmp4 := strings.Split(s, "|")

			for _, s2 := range tmp4 {
				tmp5 := strings.Split(s2, " ")
				if tmp5[1] == "red" {
					count, _ := strconv.Atoi(tmp5[0])
					if count > maxRed {
						maxRed = count
					}
				}
				if tmp5[1] == "green" {
					count, _ := strconv.Atoi(tmp5[0])
					if count > maxGreen {
						maxGreen = count
					}
				}
				if tmp5[1] == "blue" {
					count, _ := strconv.Atoi(tmp5[0])
					if count > maxBlue {
						maxBlue = count
					}
				}
			}
		}

		res += maxRed * maxGreen * maxBlue
	}

	fmt.Println(res)
}

func sample() string {
	return "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
}

func input() string {
	return "Game 1: 4 blue; 1 green, 2 red; 4 blue, 1 green, 6 red\nGame 2: 7 red, 1 green, 4 blue; 13 red, 11 blue; 6 red, 2 blue; 9 blue, 9 red; 4 blue, 11 red; 15 red, 1 green, 3 blue\nGame 3: 1 blue, 10 green; 4 green, 8 blue; 3 blue, 14 green, 1 red\nGame 4: 1 green, 2 blue; 1 blue, 4 green; 8 red, 3 blue, 3 green; 8 red, 2 green, 1 blue; 7 green, 3 blue, 2 red; 1 red, 1 green, 3 blue\nGame 5: 3 red, 7 blue, 4 green; 12 blue, 16 red, 4 green; 9 red, 2 green; 1 blue, 1 green, 1 red\nGame 6: 15 blue; 15 red, 14 blue, 2 green; 8 red, 2 blue, 2 green; 2 green, 11 blue, 1 red\nGame 7: 6 green, 6 red, 2 blue; 1 blue, 2 red, 7 green; 12 red; 5 green, 3 red, 1 blue; 16 red, 10 green\nGame 8: 2 green, 10 red, 3 blue; 1 blue, 5 green, 11 red; 6 red, 1 blue, 2 green; 11 red; 4 red, 1 blue, 5 green; 5 green, 3 blue\nGame 9: 5 blue, 11 red; 2 blue, 2 green; 11 red, 2 green\nGame 10: 3 red, 5 green; 3 blue, 5 green; 3 red, 2 blue, 14 green\nGame 11: 1 green, 1 blue, 6 red; 2 blue, 7 red, 2 green; 2 green, 2 red, 3 blue; 10 red; 2 red, 2 blue; 11 red, 3 blue\nGame 12: 6 blue, 8 red, 6 green; 15 green, 4 red; 1 red, 10 green, 1 blue; 1 red, 3 blue, 11 green\nGame 13: 2 blue, 6 red; 15 red, 6 blue; 20 blue, 10 red, 3 green; 17 blue, 1 red, 4 green\nGame 14: 3 green, 7 blue, 7 red; 2 blue; 7 blue, 10 red\nGame 15: 3 green, 7 blue; 9 green, 8 blue, 5 red; 6 green, 13 red; 14 red, 2 green, 15 blue; 15 red, 7 green\nGame 16: 13 blue; 6 green, 4 blue, 11 red; 15 red, 2 green, 6 blue; 1 green, 13 red, 8 blue; 8 green, 7 blue, 14 red\nGame 17: 3 red; 17 red, 4 green; 1 blue, 11 red; 3 blue, 20 red, 3 green\nGame 18: 4 red, 2 blue, 3 green; 9 red, 6 green; 11 red, 1 blue\nGame 19: 1 green, 4 blue; 1 green, 2 red; 2 blue, 1 green; 4 red, 2 blue\nGame 20: 15 blue, 6 green, 6 red; 13 green, 1 blue, 1 red; 13 green, 13 blue, 5 red; 7 red, 16 green, 4 blue\nGame 21: 10 blue, 5 green, 8 red; 6 blue, 12 red, 4 green; 2 green, 16 blue, 3 red; 6 red, 12 blue, 3 green; 1 red, 3 green; 7 blue, 6 green, 7 red\nGame 22: 3 blue, 3 green, 17 red; 1 green, 18 red; 3 green, 10 red, 5 blue; 2 green, 2 red, 4 blue; 2 blue, 13 red; 1 green, 3 blue, 9 red\nGame 23: 4 green, 16 red, 2 blue; 10 red, 10 green, 7 blue; 2 green, 6 red, 13 blue; 1 green, 13 blue, 16 red; 7 green, 16 blue, 9 red\nGame 24: 6 blue, 7 red, 4 green; 6 blue, 2 green; 2 green, 6 blue, 2 red; 5 red, 3 green, 11 blue\nGame 25: 4 red, 2 green; 1 green; 4 green, 4 blue, 8 red; 5 red, 1 blue\nGame 26: 9 red, 4 blue, 13 green; 3 blue, 10 red, 7 green; 13 blue, 5 green, 9 red\nGame 27: 1 green, 12 red, 2 blue; 2 blue, 13 red, 2 green; 2 blue, 7 red; 4 green, 9 red, 2 blue; 1 blue, 2 green, 15 red; 3 red, 4 green, 1 blue\nGame 28: 1 red, 9 blue, 17 green; 14 green, 15 blue, 2 red; 4 red, 18 green, 13 blue\nGame 29: 16 green, 5 blue, 1 red; 6 green, 6 red, 16 blue; 4 red, 9 green, 12 blue; 5 green, 14 blue, 1 red\nGame 30: 3 red, 2 blue, 12 green; 13 green, 4 red; 13 green, 2 red, 1 blue; 2 blue, 6 red, 4 green; 3 blue, 13 green, 5 red\nGame 31: 3 red; 6 red, 2 green; 5 red; 3 green, 2 red; 1 green, 2 red, 1 blue; 1 blue, 6 red\nGame 32: 1 red, 7 green; 9 green, 5 blue; 1 green, 2 red; 4 blue, 2 red, 1 green; 4 blue, 1 green, 3 red\nGame 33: 11 green; 12 blue, 2 green; 5 green, 1 blue; 10 green, 3 blue; 4 blue, 1 red, 4 green; 4 green, 5 blue\nGame 34: 4 red, 8 blue, 2 green; 8 green, 4 red, 14 blue; 11 green, 6 red, 8 blue; 16 green, 3 blue, 5 red; 3 blue, 3 red, 13 green\nGame 35: 7 green, 12 red, 1 blue; 1 red; 13 red; 14 red, 2 blue, 9 green\nGame 36: 3 red, 4 green, 1 blue; 3 red, 4 blue; 6 red, 4 blue, 3 green; 3 green, 4 blue, 3 red; 2 blue, 4 green, 7 red\nGame 37: 2 green, 1 blue, 5 red; 1 green; 3 blue; 3 blue, 1 green\nGame 38: 1 red, 12 blue, 17 green; 4 blue, 2 red, 8 green; 7 blue, 20 green; 6 red, 3 blue; 6 green, 7 red, 6 blue; 10 green, 3 red\nGame 39: 3 green, 3 blue, 2 red; 4 blue, 4 red, 4 green; 4 blue, 4 red; 1 blue, 5 green, 2 red; 5 green, 3 blue, 4 red; 4 green, 2 blue\nGame 40: 18 green, 1 red; 17 green, 1 blue; 2 green, 1 blue, 1 red; 9 green, 1 blue; 3 green, 1 red; 1 red, 10 green\nGame 41: 2 red, 4 blue, 3 green; 8 blue, 2 red; 5 blue; 2 green, 2 red, 3 blue; 1 green, 7 blue\nGame 42: 1 green, 2 blue; 9 green, 2 blue, 15 red; 1 green, 4 blue, 9 red\nGame 43: 5 blue, 3 red; 2 blue, 8 red, 7 green; 17 red, 4 blue, 7 green\nGame 44: 13 red, 3 green, 12 blue; 15 green, 10 blue; 8 green, 11 red, 2 blue; 10 blue, 16 red, 2 green; 12 blue, 5 green, 5 red; 14 green, 8 red, 13 blue\nGame 45: 1 red, 3 green; 4 green, 5 blue, 2 red; 6 red, 2 blue, 6 green; 3 blue, 2 green; 5 blue, 3 green, 4 red; 5 red, 5 blue, 6 green\nGame 46: 12 red, 2 blue, 3 green; 15 red, 14 blue, 11 green; 6 red, 11 blue, 6 green; 4 red, 1 green; 7 blue, 14 red; 14 red, 18 blue, 6 green\nGame 47: 3 blue, 5 red, 4 green; 1 blue, 10 red; 6 blue, 5 green, 7 red; 3 red, 4 green; 2 blue, 2 green, 13 red; 4 blue, 13 red, 2 green\nGame 48: 2 green, 3 blue, 7 red; 12 red, 1 green, 2 blue; 5 red, 2 blue; 4 blue, 3 green, 10 red\nGame 49: 8 green, 13 blue, 3 red; 14 blue, 1 green; 14 blue, 2 green\nGame 50: 1 red, 2 green, 3 blue; 2 green, 2 red; 1 green, 5 blue; 4 green\nGame 51: 10 green, 5 red; 10 green, 2 blue, 2 red; 2 blue, 13 red, 1 green; 6 blue, 10 green, 3 red\nGame 52: 8 green, 1 blue, 6 red; 4 green, 5 blue; 4 green, 7 red; 3 blue, 6 green, 3 red; 7 red, 6 blue, 7 green; 4 red, 8 green, 4 blue\nGame 53: 11 blue, 10 green, 1 red; 6 blue, 1 green, 12 red; 6 green, 12 blue, 13 red; 1 blue, 10 green, 10 red; 11 green, 2 blue; 7 green, 7 red, 5 blue\nGame 54: 3 blue, 1 green, 7 red; 18 blue, 3 red, 1 green; 11 blue, 6 red\nGame 55: 9 blue, 1 red; 3 blue, 1 green, 2 red; 1 green, 6 blue, 5 red; 1 green, 5 red, 12 blue; 5 red, 3 green, 12 blue; 12 blue\nGame 56: 3 red, 1 green, 11 blue; 2 red, 20 blue; 12 blue, 4 red; 3 red, 2 blue, 6 green\nGame 57: 1 green, 13 red, 1 blue; 7 green, 2 red, 2 blue; 6 red, 3 blue; 6 blue, 4 red, 3 green; 1 green, 11 red\nGame 58: 3 red, 13 blue, 2 green; 6 green, 6 red, 19 blue; 4 blue, 9 green, 1 red; 1 blue, 6 red\nGame 59: 11 red, 2 blue, 2 green; 1 blue, 13 red; 12 red, 6 blue\nGame 60: 8 blue, 4 red, 11 green; 10 green; 5 blue, 3 red, 8 green; 6 blue, 6 red, 12 green\nGame 61: 1 green, 1 blue, 3 red; 1 blue, 2 green, 5 red; 4 red, 1 green, 1 blue; 5 red, 2 green\nGame 62: 14 blue, 2 green, 11 red; 11 red, 2 green, 8 blue; 5 blue, 14 red, 5 green; 17 red, 2 blue, 3 green; 2 red, 3 green, 5 blue; 11 blue, 10 red, 3 green\nGame 63: 2 blue, 2 green; 9 blue, 3 red; 1 green, 2 red, 12 blue\nGame 64: 14 green, 1 blue, 5 red; 4 red, 14 green, 12 blue; 10 blue, 3 red, 10 green\nGame 65: 1 green, 6 red, 6 blue; 7 red, 7 blue, 3 green; 14 blue, 5 red\nGame 66: 10 blue, 2 red, 7 green; 3 red, 16 blue; 10 green, 7 red, 17 blue; 10 red, 5 green, 5 blue; 13 blue, 10 green, 6 red\nGame 67: 9 blue, 6 green; 1 red, 8 blue, 9 green; 3 blue, 1 green, 1 red; 2 blue, 6 green, 1 red\nGame 68: 4 green, 9 red, 3 blue; 6 blue, 5 green, 2 red; 6 blue, 9 red, 3 green; 4 red, 2 green; 4 red, 9 green\nGame 69: 1 green, 1 blue, 2 red; 2 red, 7 green; 3 red, 1 blue, 5 green; 8 red, 7 green; 2 green, 1 blue; 6 red, 1 blue, 7 green\nGame 70: 13 blue, 3 green, 5 red; 1 red, 1 green, 6 blue; 4 red, 11 blue; 14 blue, 5 red, 1 green; 8 red, 16 blue, 1 green\nGame 71: 1 blue, 1 green; 6 blue, 2 red; 5 green, 1 red, 4 blue; 4 green, 3 red\nGame 72: 4 green, 2 blue, 11 red; 4 red, 7 green, 4 blue; 3 red, 6 green, 14 blue; 4 green, 12 red, 15 blue; 4 blue, 14 red; 6 blue, 13 red, 6 green\nGame 73: 4 green, 6 red, 7 blue; 11 red, 4 blue, 6 green; 8 red, 2 blue, 5 green; 3 red, 1 green, 7 blue\nGame 74: 5 blue, 10 green; 6 green, 5 blue, 10 red; 4 green, 2 red, 1 blue; 3 blue, 11 green\nGame 75: 3 red, 3 green, 15 blue; 6 blue, 3 green, 5 red; 11 blue, 1 red, 3 green; 7 blue, 3 green, 4 red; 9 blue, 1 red, 3 green\nGame 76: 11 red; 7 green, 12 red; 2 red, 1 blue, 2 green; 2 red, 1 blue, 6 green; 5 red, 7 green; 1 blue, 8 green\nGame 77: 2 blue, 15 green, 1 red; 6 blue, 1 red; 1 green, 5 blue, 1 red; 2 blue, 1 red, 1 green; 15 green, 8 blue, 1 red; 19 green, 5 blue\nGame 78: 14 red, 2 green, 7 blue; 2 green, 14 red, 3 blue; 1 blue, 7 red\nGame 79: 15 red, 2 green, 1 blue; 3 red, 1 green; 12 red, 2 blue; 12 red, 1 green; 1 blue, 2 red, 1 green\nGame 80: 2 red, 1 green, 7 blue; 7 red, 6 blue, 5 green; 6 blue, 6 red; 6 green, 2 blue, 3 red; 5 red, 5 blue, 1 green\nGame 81: 10 red, 1 green, 3 blue; 6 green, 13 blue, 3 red; 1 green, 2 red, 10 blue\nGame 82: 4 blue, 1 red, 7 green; 4 red, 14 blue, 8 green; 1 red, 11 blue, 6 green\nGame 83: 10 red, 3 blue, 9 green; 3 green, 3 red, 1 blue; 4 blue, 11 green, 8 red; 2 blue, 8 green, 2 red; 2 green, 2 red\nGame 84: 2 green, 2 blue, 14 red; 7 red, 5 blue, 11 green; 4 red, 6 blue, 5 green; 3 blue, 13 green, 14 red; 6 red, 7 blue, 8 green; 2 blue, 3 red, 18 green\nGame 85: 8 green, 14 blue; 6 green, 9 red, 15 blue; 9 red, 12 green, 15 blue; 12 green, 6 red; 9 green, 10 red, 15 blue; 12 blue, 6 green\nGame 86: 1 blue, 1 green, 4 red; 6 green, 4 red, 6 blue; 1 red, 4 blue, 4 green; 6 green, 2 blue, 1 red\nGame 87: 17 blue, 13 green; 8 blue, 3 red; 16 green, 4 red, 6 blue\nGame 88: 11 red, 16 blue, 6 green; 10 red, 2 blue, 1 green; 5 blue, 2 green, 14 red\nGame 89: 3 blue, 2 green; 2 red; 9 blue, 8 green, 1 red; 2 green, 2 blue, 3 red; 4 red, 3 green\nGame 90: 2 blue, 14 red, 2 green; 6 blue, 2 red, 2 green; 17 red, 1 blue, 6 green; 1 blue, 8 green, 1 red\nGame 91: 6 green, 1 blue, 13 red; 10 red, 4 green, 12 blue; 9 green, 17 red, 3 blue; 12 blue, 5 red, 2 green; 2 green, 9 red, 14 blue\nGame 92: 2 red, 4 green, 6 blue; 9 blue, 3 green, 6 red; 5 blue, 4 green; 3 blue, 2 green, 7 red; 4 red, 4 green, 11 blue\nGame 93: 4 red, 11 blue, 9 green; 10 blue, 3 green, 9 red; 3 green, 11 red, 1 blue\nGame 94: 11 green, 3 red, 1 blue; 3 green, 2 red, 6 blue; 2 red, 6 blue, 5 green; 4 blue, 5 green, 5 red; 17 green, 6 red, 6 blue; 5 green, 6 red, 7 blue\nGame 95: 1 red, 3 blue, 15 green; 5 green, 6 blue; 11 green, 2 red, 11 blue; 15 green, 5 red, 7 blue\nGame 96: 13 red, 3 blue; 3 red, 13 blue; 5 blue, 1 red, 2 green; 7 red, 7 blue; 12 red, 9 blue, 3 green; 8 red, 15 blue, 2 green\nGame 97: 4 blue, 9 green, 2 red; 2 red, 5 green, 13 blue; 9 blue, 2 red, 16 green\nGame 98: 3 red; 1 green, 10 red; 2 blue, 8 red; 1 green, 11 red, 2 blue\nGame 99: 6 red, 14 green; 8 green, 15 red; 1 red, 4 green; 2 blue, 7 green, 13 red; 14 green, 5 red, 1 blue; 1 blue, 5 red, 8 green\nGame 100: 9 blue, 18 green, 4 red; 5 green, 10 blue, 11 red; 1 green, 1 red; 16 green, 5 red, 1 blue"
}
