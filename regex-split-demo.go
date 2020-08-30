package main

import (
	f "fmt"
	re "regexp"
)

// Simple example code to understand
// 1. Function of Split function
// 2. Parameters of split function
// regex-object.Split(string: , n: )

func main() {
	// Sample string that will be used in this
	// example "GeeksforGeeks loves bananas"
	str := "GeeksforGeeks loves bananas"
	f.Println(str)

	f.Println("\nPart-1: Excluding all vowels from given string")
	// a regexp object (geek) for storing all vowels
	geek := re.MustCompile("[aeiou]")
	f.Print("Printing all substring lists = ")
	// Checking split for n = -1
	f.Println(geek.Split(str, -1))
	f.Print("For n = 0 substring list = ")
	// Checking split for n = 0
	f.Println(geek.Split(str, 0))
	f.Print("For n = 1 substring list = ")
	// Checking split for n = 1
	f.Println(geek.Split(str, 1))
	f.Print("For n = 10 substring list = ")
	// Checking split for n = 10
	f.Println(geek.Split(str, 10))
	f.Print("For n = 100 substring list = ")
	// Checking split for n = 100
	f.Println(geek.Split(str, 100))

	f.Println("\n\nPart-2: Extracting all vowels from given string")
	// a regexp object (geek) for storing all consonants
	geek = re.MustCompile("[^aeiou]")
	f.Print("Printing all substring lists = ")
	// Checking split for n = -1
	f.Println(geek.Split(str, -1))
	f.Print("For n = 0 substring list = ")
	// Checking split for n = 0
	f.Println(geek.Split(str, 0))
	f.Print("For n = 1 substring list = ")
	// Checking split for n = 1
	f.Println(geek.Split(str, 1))
	f.Print("For n = 10 substring list = ")
	// Checking split for n = 10
	f.Println(geek.Split(str, 10))
	f.Print("For n = 100 substring list = ")
	// Checking split for n = 100
	f.Println(geek.Split(str, 100))

	// Did you notice that split function
	// does not modify the original regex
	// matching object?
}