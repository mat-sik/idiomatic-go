package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed english_rights.txt
var english string

//go:embed french_rights.txt
var french string

//go:embed spanish_rights.txt
var spanish string

func main()  {
    if len(os.Args) != 2 {
        panic("Incorrect argument used, use either: english, french or spanish.")
    }

    switch os.Args[1] {
    case "english":
        fmt.Println(english)
    case "french":
        fmt.Println(french)
    case "spanish":
        fmt.Println(spanish)
    default:
        panic(fmt.Sprintf("Wrong argument provided: %s.", os.Args[1]))
    }
}
