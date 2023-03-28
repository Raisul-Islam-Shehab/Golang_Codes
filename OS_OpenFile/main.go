// This program explains "os.OpenFile"
package main

import (
	"log"
	"os"
)

func check(err error) {
	if err == nil {
		return
	}
	log.Fatal(err)
}

func main() {
	// **Reads the file

	// file, err := os.OpenFile("file.txt", os.O_RDONLY, os.FileMode(0600))
	// check(err)
	// defer file.Close()
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
	// check(scanner.Err())

	// **Writes the file. Overrides it

	// file, err := os.OpenFile("file.txt", os.O_WRONLY, os.FileMode(0600))
	// check(err)
	// _, err = file.Write([]byte("Golang is amazing\n"))
	// check(err)
	// err = file.Close()
	// check(err)

	// **Writes at the end of the file

	// options := os.O_WRONLY | os.O_APPEND
	// file, err := os.OpenFile("file.txt", options, os.FileMode(0600))
	// check(err)
	// _, err = file.Write([]byte("It's so simple.\n"))
	// check(err)
	// err = file.Close()
	// check(err)

	// **Writes at the end of the file and creates the file if doesn't exit
	// "os.FileMode" declares the permission of the file like UNIX command 'chmod'

	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("file.txt", options, os.FileMode(0600))
	check(err)
	_, err = file.Write([]byte("It's so simple.\n"))
	check(err)
	err = file.Close()
	check(err)
}
