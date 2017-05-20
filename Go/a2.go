package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type tokens [][]byte

func htmlSpecialCharacter(c byte) ([]byte, bool) {
	switch c {
	case '&':
		return []byte("&amp;"), true
	case '\'':
		return []byte("&apos;"), true
	case '<':
		return []byte("&lt;"), true
	case '>':
		return []byte("&gt;"), true
	case '"':
		return []byte("&quot;"), true
	default:
		return []byte{c}, false
	}
}

func (t tokens) convertSpecialCharactersToHTML() {
	for i := range t {
		if t[i][0] == '"' {
			var temp []byte
			for j := range t[i] {
				if seq, specialChar := htmlSpecialCharacter(t[i][j]); specialChar {
					temp = append(temp, seq...)
				} else {
					temp = append(temp, t[i][j])
				}
			}
			t[i] = temp
		}
	}
}

func readJSONFromFile(filename string) ([]byte, bool) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error (function 'getJSONFileText') %v", err)
	}
	return buf, len(buf) == 0
}

func writeHTMLToFile(s string, outputFile string) {
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error (function 'writeHTMLToFile') %v", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = fmt.Fprintf(w, "%v\n", s)
	if err != nil {
		log.Fatalf("Error (function 'writeHTMLToFile') %v", err)
	}
	w.Flush()
}

func escapeCharacter(c byte) bool {
	return c == '\\'
}

func validNum(c byte) bool {
	switch c {
	case '-', '+', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'E', 'e':
		return true
	default:
		return false
	}
}

func stringSequence(b []byte) ([]byte, int) {
	result := []byte{b[0]}
	offset := 1
	for i := 1; i < len(b); i += offset {
		offset = 1
		if escapeCharacter(b[i]) {
			result = append(result, b[i])
			result = append(result, b[i+1])
			offset = 2
		} else if b[i] != '"' {
			result = append(result, b[i])
		} else {
			result = append(result, b[i])
			break
		}
	}
	return result, len(result)
}

func numSequence(b []byte) ([]byte, int) {
	var result []byte
	for i := range b {
		if validNum(b[i]) {
			result = append(result, b[i])
		} else {
			break
		}
	}
	return result, len(result)
}

func parseJSON(b []byte) (t tokens) {
	var seq []byte
	openArrays := []int{0}
	openArraysIndex := 0
	offset := 1
	for i := 0; i < len(b); i += offset {
		offset = 1
		switch b[i] {
		case ':':
			t = append(t, []byte{b[i]})
		case '"':
			seq, offset = stringSequence(b[i:])
			t = append(t, seq)
		case '-', '+', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			seq, offset = numSequence(b[i:])
			t = append(t, seq)
		case 'f':
			t = append(t, b[i:i+5])
			offset = 5
		case 't', 'n':
			t = append(t, b[i:i+4])
			offset = 4
		case '[':
			openArrays[openArraysIndex]++
			t = append(t, []byte{b[i]})
		case ']':
			openArrays[openArraysIndex]--
			t = append(t, []byte{b[i]})
		case ',':
			if openArrays[openArraysIndex] > 0 {
				t = append(t, []byte{b[i], 'a'})
			} else {
				t = append(t, []byte{b[i], 'o'})
			}
		case '{':
			t = append(t, []byte{b[i]})
			if openArrays[openArraysIndex] > 0 {
				openArrays = append(openArrays, 0)
				openArraysIndex++
			}
		case '}':
			t = append(t, []byte{b[i]})
			if openArraysIndex > 0 {
				openArrays = openArrays[:openArraysIndex]
				openArraysIndex--
			}
		}
	}
	return t
}

func openPTag(indent int) string {
	pTag := []string{"<p style=\"text-indent:"}
	pTag = append(pTag, strconv.Itoa(indent))
	pTag = append(pTag, "px\">")
	return strings.Join(pTag, "")
}

func htmlSpecialCharacterString(b []byte) (string, int) {
	var result []byte
	for i := range b {
		result = append(result, b[i])
		if b[i] == ';' {
			break
		}
	}
	return string(result), len(result)
}

func spanTag(b []byte, color string, isString bool) string {
	sTag := []string{"<span style=\"color:"}
	sTag = append(sTag, color)
	sTag = append(sTag, "\">")
	if isString {
		offset := 1
		for i := 0; i < len(b); i += offset {
			offset = 1
			if escapeCharacter(b[i]) {
				sTag = append(sTag, "<span style=\"color:")
				sTag = append(sTag, "#05138F")
				sTag = append(sTag, "\">")
				if b[i+1] == '&' {
					var str string
					str, offset = htmlSpecialCharacterString(b[i:])
					sTag = append(sTag, str)
				} else {
					sTag = append(sTag, string(b[i:i+2]))
					offset = 2
				}
				sTag = append(sTag, "</span>")
			} else {
				sTag = append(sTag, string(b[i]))
			}
		}
	} else {
		str := string(b)
		switch str {
		case ":":
			sTag = append(sTag, " ")
			sTag = append(sTag, str)
			sTag = append(sTag, " ")
		case ",":
			sTag = append(sTag, str)
			sTag = append(sTag, " ")
		default:
			sTag = append(sTag, str)
		}
	}
	sTag = append(sTag, "</span>")
	return strings.Join(sTag, "")
}

func convertToHTML(t tokens) string {
	t.convertSpecialCharactersToHTML()
	html := []string{"<!DOCTYPE html><html><body>"}
	var color string
	var isString bool
	indent := 0
	for i := range t {
		isString = false
		switch t[i][0] {
		case '{':
			color = "#0E0B16"
			if i == 0 {
				html = append(html, openPTag(indent))
				html = append(html, spanTag(t[i], color, isString))
			} else {
				html = append(html, spanTag(t[i], color, isString))
			}
			html = append(html, "</p>")
			indent += 40
			html = append(html, openPTag(indent))
		case '}':
			color = "#0E0B16"
			indent -= 40
			html = append(html, "</p>")
			html = append(html, openPTag(indent))
			html = append(html, spanTag(t[i], color, isString))
			if i == len(t)-1 || (t[i+1][0] != ',' && t[i-1][0] != ']') {
				html = append(html, "</p>")
			}
		case '&':
			color = "#007849"
			isString = true
			html = append(html, spanTag(t[i], color, isString))
		case '-', '+', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			color = "#4717F6"
			html = append(html, spanTag(t[i], color, isString))
		case ',':
			color = "#666633"
			html = append(html, spanTag([]byte{t[i][0]}, color, isString))
			if t[i][1] != 'a' {
				html = append(html, "</p>")
				html = append(html, openPTag(indent))
			}
		case ':':
			color = "orange"
			html = append(html, spanTag(t[i], color, isString))
		case '[', ']':
			color = "red"
			html = append(html, spanTag(t[i], color, isString))
		case 't', 'f', 'n':
			color = "#A239CA"
			html = append(html, spanTag(t[i], color, isString))
		}
	}
	html = append(html, "</body></html>")
	return strings.Join(html, "")
}

func main() {
	filename := os.Args[1]
	if json, empty := readJSONFromFile(filename); empty {
		fmt.Println("File is empty")
	} else {
		jsonTokens := parseJSON(json)
		html := convertToHTML(jsonTokens)
		outputFile := "output.html"
		writeHTMLToFile(html, outputFile)
	}
}
