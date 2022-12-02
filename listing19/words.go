package main

import (
	"fmt"
	"strings"
	"sort"
)

func countS(s string) map[string]int {
	loweredS := strings.ToLower(s)
	fieldedS := strings.Fields(loweredS)
	trimmedS := make([]string, 0, len(fieldedS))
	for _, s := range fieldedS {
		trimmedS = append(trimmedS, strings.Trim(s, ".;()—/,"))
	}

	sort.Strings(trimmedS)
	fmt.Println(trimmedS)
	result := make(map[string]int)

	for _, s := range trimmedS {
		result[s]++
	}

	return result
}

func main() {
	str := `
	As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man.

—C.S. Lewis, Out of the Silent Planet,
(see mng.bz/V7nO)
	`
	result := countS(str)
	fmt.Println(result)

	for s, c := range result {
		if c > 1 {
			fmt.Println(s, c)
		}
	}
}


