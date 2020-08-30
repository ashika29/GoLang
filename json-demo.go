package main

import(
	"fmt"
	"os"
	"bufio"
	"encoding/json"
)

func main(){
	type Person struct{
		Name string
		ADDRESS string
	}
	sc:=bufio.NewScanner(os.Stdin)
	//s :=sc.Text()
	fmt.Println("Enter name: ")
	sc.Scan()
	nam := sc.Text()
	fmt.Println("Enter address: ")
	sc.Scan()
	adr := sc.Text()
	p1 := Person{
		Name: nam,
		ADDRESS: adr,
	}
	barr, _ := json.Marshal(p1)
	os.Stdout.Write(barr)
}