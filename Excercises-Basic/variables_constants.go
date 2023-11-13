package main

import "fmt"

func main() {

	var (
		/* ||---------------------------------------------------||
		     Forma en las que se pueden declarar las variables
		   ||---------------------------------------------------||
		*/
		/*title  string = "The Fellowship of the Ring"
		author string = "J.R.R. Tolkien"
		pages  int    = 420*/
		title, author, pages = "The Fellowship of the Ring",
			"J.R.R. Tolkien",
			420
	)

	fmt.Printf("The book with %d pages %s written by %s", pages, title, author)
}
