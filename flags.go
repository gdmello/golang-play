package main
import (
	"flag"
	"fmt"
)

func main() {

	namePtr := flag.String("name", "Your name", "What is your name?")
	agePtr := flag.Int("age",0,"What is your age?")

	flag.Parse()

	fmt.Println("name: ", *namePtr)
	fmt.Println("age: ", *agePtr)



}
