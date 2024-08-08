package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name       string
	YearBirth  int
	Occupation string
}

func (p *Person) GetAge() int {
	now := time.Now()
	currentYear := now.Year()
	return currentYear - p.YearBirth
}

func (p *Person) CheckGoodOccupation() bool {
	return (p.YearBirth%len(p.Name) == 0)
}

func main() {
	p := Person{
		Name:       "Jake",
		YearBirth:  2000,
		Occupation: "Software Engineer",
	}
	fmt.Printf("%v is %v years old, %v is a %v.\n", p.Name, p.GetAge(), p.Name, p.Occupation)
	if p.CheckGoodOccupation() {
		fmt.Printf("%v has good occupation!\n", p.Name)
		fmt.Printf("Because %v was born in %v and %v's first name has %v letters.\n", p.Name, p.YearBirth, p.Name, len(p.Name))
	} else {
		fmt.Printf("%v does not have good occupation!\n", p.Name)
		fmt.Printf("Because %v was born in %v and %v's first name has %v letters.\n", p.Name, p.YearBirth, p.Name, len(p.Name))
	}

}
