package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name      string
	Job       string
	BirthYear int
}

func main() {
	people := []Person{}
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		split_line := strings.Split(line, ",")
		if len(split_line) != 3 {
			log.Println("Invalid input")
			continue
		}
		birthYear, err := strconv.Atoi(strings.TrimSpace(split_line[2]))
		if err != nil {
			log.Println(err)
			continue
		} else {
			person := Person{
				Name:      strings.ToUpper(strings.TrimSpace(split_line[0])),
				Job:       strings.ToLower(strings.TrimSpace(split_line[1])),
				BirthYear: birthYear,
			}
			people = append(people, person)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Write slice people to a file for manual checking
	jsonData, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}
	output, err := os.Create("output.json")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()
	_, err = output.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}
