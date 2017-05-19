package a1

import (
	"fmt"
	"testing"
)

func TestCountPrimes(t *testing.T) {
	testArr := []int{-6, 0, 1, 5, 10000, 25}
	corrValArr := []int{0, 0, 0, 3, 1229, 9}
	for i := range testArr {
		actual := countPrimes(testArr[i])
		if actual == corrValArr[i] {
			fmt.Printf("Success. The number of primes less than or equal to %v is %v.\n", testArr[i], actual)
		} else {
			t.Error("The number of primes calculated for n =", testArr[i], "is incorrect. Expected:", corrValArr[i], "Got:", actual)
		}
	}
	fmt.Printf("\n")
}

func TestCountStrings(t *testing.T) {
	testData := []string{"txt1.txt", "txt2.txt", "txt3.txt"}
	corrData := []map[string]int{
		{"The": 1, "the": 1, "big": 3, "dog": 1, "ate": 1, "apple": 1},
		{},
		{"this": 3, "This": 2, "thIs": 1, "gone": 1, "by": 1, "baby": 1, "|": 1, "it's.": 1},
	}
	for i := range testData {
		fmt.Printf("Checking file '%s'.\n", testData[i])
		testMap := countStrings(testData[i])
		corrMap := corrData[i]
		if len(testMap) == 0 && len(testMap) == len(corrMap) {
			fmt.Printf("Success. File is empty.\n")
		} else if len(testMap) == len(corrMap) {
			for key, value := range testMap {
				if value == corrMap[key] {
					fmt.Printf("Success. Word: %s Instances: %v\n", key, value)
				} else {
					t.Error("The number of instances of word '", key, "' does not match expected result. Expected:", corrMap[key], "Got:", value)
				}
			}
		} else {
			t.Error("The number of keys returned by 'countStrings' does not match expected number of keys. Expected:", len(corrMap), "Got:", len(testMap))
		}
		fmt.Printf("\n")
	}
}

func TestTime24(t *testing.T) {
	testData := [][]Time24{
		{
			{hour: 5, minute: 39, second: 8},
			{hour: 0, minute: 1, second: 54},
			{hour: 20, minute: 60, second: 54},
			{hour: 25, minute: 0, second: 10},
			{hour: 8, minute: 23, second: 100},
			{hour: 31, minute: 23, second: 61},
		},
		{},
		{
			{hour: 5, minute: 39, second: 8},
			{hour: 20, minute: 8, second: 4},
			{hour: 2, minute: 0, second: 10},
			{hour: 8, minute: 23, second: 10},
			{hour: 1, minute: 23, second: 51},
			{hour: 0, minute: 1, second: 54},
		},
		{
			{hour: 5, minute: 39, second: 8},
			{hour: 23, minute: 18, second: 44},
			{hour: 5, minute: 39, second: 7},
			{hour: 5, minute: 39, second: 7},
		},
	}
	corrData := []Time24{
		{hour: 0, minute: 0, second: 0},
		{hour: 0, minute: 0, second: 0},
		{hour: 0, minute: 1, second: 54},
		{hour: 5, minute: 39, second: 7},
	}
	for i := range testData {
		minTimeData, err := minTime24(testData[i])
		minTimeCorr := corrData[i]
		if err != nil {
			t.Error(err)
		} else {
			if equalsTime24(minTimeData, minTimeCorr) {
				fmt.Printf("Success. Mininum time is '%v'\n", minTimeData)
			} else {
				t.Error("The minimum time is incorrect. Expected:", minTimeCorr, "Got:", minTimeData)
			}
		}
	}
	fmt.Printf("\n")
}

type SearchListInt struct {
	x   int
	lst []int
}

type SearchListStr struct {
	x   string
	lst []string
}

func TestLinearSearch(t *testing.T) {
	testDataInt := []SearchListInt{
		{5, []int{4, 2, -1, 5, 0}},
		{3, []int{4, 2, -1, 5, 0}},
	}
	testDataStr := []SearchListStr{
		{"egg", []string{"cat", "nose", "egg"}},
		{"up", []string{"cat", "nose", "egg"}},
	}
	corrData := []int{3, -1, 2, -1}
	for i := range corrData {
		var testVal int
		if i < len(testDataInt) {
			testVal = linearSearch(testDataInt[i].x, testDataInt[i].lst)
		} else {
			testVal = linearSearch(testDataStr[i-len(testDataInt)].x, testDataStr[i-len(testDataInt)].lst)
		}
		corrVal := corrData[i]
		if testVal == corrVal {
			if testVal == -1 {
				fmt.Printf("Success. Search value is not in list\n")
			} else {
				fmt.Printf("Success. Position of search value is '%v'\n", testVal)
			}
		} else {
			t.Error("Position of search value is incorrect. Expected:", corrVal, "Got:", testVal)
		}
	}
	//	testVal := linearSearch(2, []string{"cat", "nose", "egg"})
	//	corrVal := -1
	//	if testVal == corrVal {
	//		fmt.Printf("Success. Search value is not in list\n")
	//	}
	fmt.Printf("\n")
}

func TestAllBitSeqs(t *testing.T) {
	testData := []int{0, 1, 2, 3, 4}
	corrData := [][][]int{
		{},
		{{0}, {1}},
		{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		{{0, 0, 0}, {0, 0, 1}, {0, 1, 0}, {0, 1, 1}, {1, 0, 0}, {1, 0, 1}, {1, 1, 0}, {1, 1, 1}},
		{{0, 0, 0, 0}, {0, 0, 0, 1}, {0, 0, 1, 0}, {0, 0, 1, 1}, {0, 1, 0, 0}, {0, 1, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1},
			{1, 0, 0, 0}, {1, 0, 0, 1}, {1, 0, 1, 0}, {1, 0, 1, 1}, {1, 1, 0, 0}, {1, 1, 0, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}},
	}
	for i := range testData {
		testSlice := allBitSeqs(testData[i])
		corrSlice := corrData[i]
		match, bitSeq := multiSliceMatch(testSlice, corrSlice)
		if match {
			fmt.Printf("Success. 'allBitSeqs' returned the correct bit sequences for n = %v\n", testData[i])
		} else if !match && bitSeq == nil {
			t.Error("'allBitSeqs' did not return the correct number of bit sequences. Expected:", len(corrSlice), "Got:", len(testSlice))
		} else {
			t.Error("Bit sequence", bitSeq, "is not contained in results returned by 'allBitSeqs'")
		}
	}
	fmt.Printf("\n")
}
