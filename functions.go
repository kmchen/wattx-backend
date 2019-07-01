package main

// http://rosettacode.org/wiki/Category:String_manipulation

import (
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Find min in int slice
func findMin(v []int) int {
	var m int
	for i, e := range v {
		if i == 0 || e < m {
			m = e
		}
	}
	return m
}

func even(a []int) (r []int) {
	for _, e := range a {
		if e%2 == 0 {
			r = append(r, e)
		}
	}
	return
}

func removeDuplicatesInt(elements []int) []int {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func removeDuplicatesStringSlice(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}

func sortKeysInMap(codes map[string]int) map[string]int {
	keys := []string{}
	newMap := map[string]int{}
	for key, _ := range codes {
		keys = append(keys, key)
	}

	// Sort string keys.
	sort.Strings(keys)

	// Loop over sorted key-value pairs.
	for i := range keys {
		key := keys[i]
		value := codes[key]
		newMap[key] = value
	}
	return newMap
}

func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}

func matchString(first, second string) {
	fmt.Printf("1. %s starts with %s: %t\n",
		first, second, strings.HasPrefix(first, second))
	i := strings.Index(first, second)
	fmt.Printf("2. %s contains %s: %t,\n", first, second, i >= 0)
	if i >= 0 {
		fmt.Printf("2.1. at location %d,\n", i)
		for start := i + 1; ; {
			if i = strings.Index(first[start:], second); i < 0 {
				break
			}
			fmt.Printf("2.2. at location %d,\n", start+i)
			start += i + 1
		}
		fmt.Println("2.2. and that's all")
	}
	fmt.Printf("3. %s ends with %s: %t\n",
		first, second, strings.HasSuffix(first, second))
}

func show(label, str string) {
	fmt.Printf("%s: |%s| %v\n", label, str, []rune(str))
}

func TokenizeString(s string, sep, escape rune) (tokens []string, err error) {
	var runes []rune
	inEscape := false
	for _, r := range s {
		switch {
		case inEscape:
			inEscape = false
			fallthrough
		default:
			runes = append(runes, r)
		case r == escape:
			inEscape = true
		case r == sep:
			tokens = append(tokens, string(runes))
			runes = runes[:0]
		}
	}
	tokens = append(tokens, string(runes))
	if inEscape {
		err = errors.New("invalid terminal escape")
	}
	return tokens, err
}

func reverseBytes(s string) string {
	r := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	return string(r)
}

type pair struct {
	name, value string
}
type csArray []pair

// three methods satisfy sort.Interface
func (a csArray) Less(i, j int) bool { return a[i].name < a[j].name }
func (a csArray) Len() int           { return len(a) }
func (a csArray) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func li(is ...int) *big.Int {
	ps := make(cc, len(is))
	ss := make([]c, len(is))
	ml := 0
	for j, i := range is {
		p := &ss[j]
		ps[j] = p
		p.i = i
		p.s = strconv.Itoa(i)
		if len(p.s) > ml {
			ml = len(p.s)
		}
	}
	for _, p := range ps {
		p.rs = strings.Repeat(p.s, (ml+len(p.s)-1)/len(p.s))
	}
	sort.Sort(ps)
	s := make([]string, len(ps))
	for i, p := range ps {
		s[i] = p.s
	}
	b, _ := new(big.Int).SetString(strings.Join(s, ""), 10)
	return b
}

func main() {
	fmt.Println(findMin([]int{5, 4, 3, 2, 1, -1, 0, -1, -1, 1}))
	fmt.Println(removeDuplicatesInt([]int{5, 4, 3, 2, 1, -1, 0, -1, -1, 1}))
	fmt.Println(removeDuplicatesInt([]int{10, 20, 30, 10, 10, 20, 40}))
	fmt.Println(removeDuplicatesStringSlice([]string{"cat", "dog", "cat", "bird"}))
	stringSlice := []string{"cat", "dog", "cat", "bird"}
	sort.Strings(stringSlice)
	fmt.Println(stringSlice)
	codes := map[string]int{
		"xyz": 1,
		"ghi": 1,
		"abc": 1,
		"def": 1,
	}
	fmt.Println(sortKeysInMap(codes))
	fmt.Println(stripchars("She was a soul stripper. She took my heart!",
		"aei"))
	matchString("abracadabra", "abr")

	var simple = `
	    simple   `
	show("original", simple)
	show("leading ws removed", strings.TrimLeftFunc(simple, unicode.IsSpace))
	show("trailing ws removed", strings.TrimRightFunc(simple, unicode.IsSpace))
	// equivalent to strings.TrimFunc(simple, unicode.IsSpace)
	show("both removed", strings.TrimSpace(simple))

	// Substring
	s := "ABCDEFGH"
	n, m := 2, 3
	// for reference
	fmt.Println("Index: ", "01234567")
	fmt.Println("String:", s)
	// starting from n characters in and of m length
	fmt.Printf("Start %d, length %d:    %s\n", n, m, s[n:n+m])
	// starting from n characters in, up to the end of the string
	fmt.Printf("Start %d, to end:      %s\n", n, s[n:])
	// whole string minus last character
	fmt.Printf("All but last:         %s\n", s[:len(s)-1])
	// starting from a known character within the string and of m length
	dx := strings.IndexByte(s, 'D')
	fmt.Printf("Start 'D', length %d:  %s\n", m, s[dx:dx+m])
	// starting from a known substring within the string and of m length
	sx := strings.Index(s, "DE")
	fmt.Printf(`Start "DE", length %d: %s`+"\n", m, s[sx:sx+m])

	// Tokenize string
	const sample = "one^|uno||three^^^^|four^^^|^cuatro|"
	const separator = '|'
	const escape = '^'

	fmt.Printf("Input:   %q\n", sample)
	tokens, err := TokenizeString(sample, separator, escape)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("Tokens: %q\n", tokens)
	}

	// Reverse a string
	s = "asdf"
	fmt.Println("\noriginal:      ", []byte(s), s)
	fmt.Println(reverseBytes(s))

	// Even number filter
	a := rand.Perm(20)
	fmt.Println("original: ", a)
	fmt.Println("even number filter: ", even(a))

	// Sort object in slice
	var x = csArray{
		pair{"joe", "120"},
		pair{"foo", "31"},
		pair{"bar", "251"},
	}

	sort.Sort(x)
	for _, p := range x {
		fmt.Printf("%5s: %s\n", p.name, p.value)
	}

	// largest int from slice int
	type c struct {
		i     int
		s, rs string
	}
	fmt.Println(li(1, 34, 3, 98, 9, 76, 45, 4))
	fmt.Println(li(54, 546, 548, 60))
}
