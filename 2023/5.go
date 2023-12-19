package main

import (
	"aoc/2023/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func mapValue(input []string, value int) int {
	desMapping := value
	for i := 0; i < len(input) && len(strings.TrimSpace(input[i])) != 0; i++ {
		values := strings.Split(input[i], " ")
		des, _ := strconv.ParseInt(values[0], 10, 64)
		src, _ := strconv.ParseInt(values[1], 10, 64)
		rng, _ := strconv.ParseInt(values[2], 10, 64)
		if int(src) <= value && value <= int(src)+int(rng) {
			desMapping = int(des) + (value - int(src))
		}
	}

	return desMapping
}

func main() {
	input := strings.Split(utils.ReadFile("2023/input/5.1.txt"), "\n")
	answer := int(1e9)
	seeds := strings.Split(input[0][7:], " ")
	for _, v := range seeds {
		parsedSeed, _ := strconv.ParseInt(v, 10, 64)
		location :=
			mapValue(input[slices.Index(input, "humidity-to-location map:")+1:],
				mapValue(input[slices.Index(input, "temperature-to-humidity map:")+1:],
					mapValue(input[slices.Index(input, "light-to-temperature map:")+1:],
						mapValue(input[slices.Index(input, "water-to-light map:")+1:],
							mapValue(input[slices.Index(input, "fertilizer-to-water map:")+1:],
								mapValue(input[slices.Index(input, "soil-to-fertilizer map:")+1:],
									mapValue(input[slices.Index(input, "seed-to-soil map:")+1:], int(parsedSeed))))))))
		answer = int(math.Min(float64(location), float64(answer)))
	}

	fmt.Println(answer)

}
