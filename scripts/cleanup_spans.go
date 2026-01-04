package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	path := "main.go"
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	s := string(b)

	re := regexp.MustCompile(`\[span_[0-9]+\]\((?:start|end)_span\)`)
	s = re.ReplaceAllString(s, "")

	// Normalize multiple blank lines introduced
	re2 := regexp.MustCompile(`\r?\n{3,}`)
	s = re2.ReplaceAllString(s, "\r\n\r\n")

	if err := ioutil.WriteFile(path, []byte(s), os.FileMode(0644)); err != nil {
		log.Fatal(err)
	}
	log.Println("cleaned main.go")
}
