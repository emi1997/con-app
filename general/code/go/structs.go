package main

import (
	"fmt"
	"time"
)

func main() {
	type Pupil struct {
		FirstName string
		LastName  string
		Age       int
	}
	type School struct {
		Location string
		Type     string
	}
	type Subject struct {
		id    int
		Name  string
		Topic string
	}

	var (
		date = time.Now()
		maja = Pupil{"Maja", "Henderson", 6}
		elementary bool
		majaSchool = School{"Vagen", ""}
		majaSubject = Subject{1, "", ""}
	)
	
	if maja.Age >= 6 && maja.Age <= 10{
		elementary = true
	}
	if elementary{
		majaSchool.Type = "Elementary School"
		majaSubject.Name = "German"
		majaSubject.Topic = "Alphabet"
	}
	// var elementary = School{"Vagen", "Grundschule"}
	// var subject1 = Subject{1, "Deutsch", "Begleiter"}

	// fmt.Println(elementary.Name)
	// fmt.Println(subject1.Name)
	// fmt.Println(subject1.Topic)
	fmt.Println(maja.FirstName)
	fmt.Println(maja.LastName)
	//%d == digit 
	fmt.Printf("%d Jahre\n", maja.Age)
	fmt.Println(majaSchool.Type)
	fmt.Println(majaSubject.Name)
	fmt.Println(majaSubject.Topic)
	fmt.Println(date)

}
