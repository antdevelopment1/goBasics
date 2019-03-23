// To run code inside a go file we say - go run and the file name
// Go run can run a handleful of files
// Go build complies it doesn't execute the code but we will see a new file
// We can run that file by typing ./filename
// At the very top of a file we have to declare what package a file belongs to. In this case it belongs to main
// Types of packages are executables and reusable
package main

// Import statement from standard go library
import "fmt"

func main() {
	fmt.Println("Hello World!")
}
