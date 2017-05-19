package a1

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
)

type Time24 struct {
	hour, minute, second uint8
}

func countPrimes(n int) int {
	count := 0
	if n <= 1 {
		return count
	}
	boolArr := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		if boolArr[i] {
			continue
		}
		for j := i * i; j < n+1; j += i {
			boolArr[j] = true
		}
		count++
	}
	return count
}

func countStrings(filename string) map[string]int {
	m := make(map[string]int)
	absPath, _ := filepath.Abs(filename)
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatalf("Error (function 'countStrings') %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if _, ok := m[scanner.Text()]; ok {
			m[scanner.Text()]++
		} else {
			m[scanner.Text()] = 1
		}
	}
	return m
}

func equalsTime24(a Time24, b Time24) bool {
	return a.hour == b.hour && a.minute == b.minute && a.second == b.second
}

func lessThanTime24(a Time24, b Time24) bool {
	if a.hour > b.hour {
		return false
	} else if a.hour == b.hour {
		if a.minute > b.minute {
			return false
		} else if a.minute == b.minute {
			if a.second < b.second {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	} else {
		return true
	}
}

func (t Time24) String() string {
	var prefix_hr, prefix_min, prefix_sec string
	if t.hour < 10 {
		prefix_hr = "0"
	}
	if t.minute < 10 {
		prefix_min = "0"
	}
	if t.second < 10 {
		prefix_sec = "0"
	}
	return prefix_hr + strconv.Itoa(int(t.hour)) + ":" + prefix_min + strconv.Itoa(int(t.minute)) + ":" + prefix_sec + strconv.Itoa(int(t.second))
}

func (t Time24) validTime24() bool {
	return t.hour < 24 && t.minute < 60 && t.second < 60
}

func minTime24(times []Time24) (Time24, error) {
	minTime := Time24{hour: 23, minute: 59, second: 59}
	if len(times) == 0 {
		minTime := Time24{hour: 0, minute: 0, second: 0}
		return minTime, errors.New("Error: input must not be empty.")
	}
	for i := 0; i < len(times); i++ {
		if !times[i].validTime24() {
			minTime := Time24{hour: 0, minute: 0, second: 0}
			return minTime, fmt.Errorf("Error: time '%s' at position '%v' is not valid.", times[i].String(), i+1)
		}
	}
	for i := 0; i < len(times); i++ {
		if lessThanTime24(times[i], minTime) {
			minTime = times[i]
		}
	}
	return minTime, nil
}

func linearSearch(x interface{}, lst interface{}) int {
	index := -1
	switch reflect.TypeOf(lst).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(lst)
		for i := 0; i < s.Len(); i++ {
			if reflect.TypeOf(x) != s.Index(i).Type() {
				message := "type '" + reflect.TypeOf(x).String() + "' of search value does not match type '" + s.Index(i).Type().String() + "' of list values"
				panic(message)
			}
			switch x.(type) {
			case int:
				if int64(x.(int)) == s.Index(i).Int() {
					index = i
				}
			case float64:
				if x.(float64) == s.Index(i).Float() {
					index = i
				}
			case string:
				if x.(string) == s.Index(i).String() {
					index = i
				}
			default:
				panic("type is not supported")
			}
		}
	}
	return index
}

func binarySliceSize(n int) int {
	if n <= 0 {
		return 0
	} else {
		size := 2
		for i := 1; i < n; i++ {
			size *= 2
		}
		return size
	}
}

func binarySequence(n int, bStr string) []int {
	b := make([]int, n)
	j := n - 1
	for i := len(bStr) - 1; i >= 0; i-- {
		b[j], _ = strconv.Atoi(string(bStr[i]))
		j--
	}
	return b
}

func allBitSeqs(n int) [][]int {
	b := make([][]int, binarySliceSize(n))
	for i := 0; i < len(b); i++ {
		b[i] = binarySequence(n, strconv.FormatInt(int64(i), 2))
	}
	return b
}

func sliceMatch(a []int, b []int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func sliceContains(a [][]int, b []int) bool {
	for i := range a {
		if sliceMatch(a[i], b) {
			return true
		}
	}
	return false
}

func multiSliceMatch(a [][]int, b [][]int) (bool, []int) {
	if len(a) != len(b) {
		return false, nil
	}
	for i := range b {
		if !sliceContains(a, b[i]) {
			return false, b[i]
		}
	}
	return true, nil
}
