package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type CountW struct {
	word  string
	count int
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")
	new_, _ := reader.ReadString('\n')
	s := []string{new_}
	//s := []string{"I am aditya working in think future technology. I currently in Final year to"}
	//s := "Adityaa kjsfd"
	words := strings.Split(strings.Join(s, ""), " ")
	m := make(map[string]int)
	for _, word := range words {
		if _, ok := m[word]; ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}
	WC := make([]CountW, 0, len(m))
	for key, val := range m {
		WC = append(WC, CountW{word: key, count: val})
	}

	sort.Slice(WC, func(i, j int) bool {
		return WC[i].count > WC[j].count
	})
	for i := 0; i < len(WC) && i < 11; i++ {
		//a := WC[i].word
		//b := WC[i].count
		//	fmt.Println(WC[i].word, ":", WC[i].count)
		//	fmt.Println(WC[i].word, ":", WC[i].count)

		fmt.Println(WC[i].word, ":", WC[i].count)

	}

}
