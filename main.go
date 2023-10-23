package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		args := os.Args[1]
		if args != "" {
			file, err := os.Open(args)
			if err != nil {
				log.Fatal(err.Error())
			}
			defer func() {
				if err = file.Close(); err != nil {
					log.Fatal(err)
				}
			}()
			var myArray = ReadInput(file)
			Solve(myArray)
			PrintSol()
		}
	}

}
