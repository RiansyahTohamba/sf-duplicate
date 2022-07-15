package main

import "fmt"

func getStartEnd(page, itemPerPage int) (start, end int) {
	if page == 1 {
		start = 0
	} else {
		start = (page - 1) * itemPerPage
	}

	end = (start - 1) + itemPerPage
	return
}

func printStartEnd(page, itemPerPage int) {
	var start int
	fmt.Printf("page %d \n", page)
	if page == 1 {
		start = 0
	} else {
		start = (page - 1) * itemPerPage
	}

	end := (start - 1) + itemPerPage

	fmt.Printf("start = %d \n", start)
	fmt.Printf("end = %d \n", end)
	fmt.Println(" -------")

	// return start, end
}

// select from 0, 10 (limit, offset)
func main() {
	itemPerPage := 8
	printStartEnd(1, itemPerPage)
	printStartEnd(2, itemPerPage)
	printStartEnd(3, itemPerPage)
	printStartEnd(4, itemPerPage)
	// startpage1 := 0
	// endpage1 := startpage1 + (itemPerPage - 1)

	//  0 7 8 15 16 23 24
	//
	// fmt.Println("page 1")

	// startpage1 := 0
	// endpage1 := startpage1 + (itemPerPage - 1)

	// fmt.Printf("start = %d \n", startpage1)
	// fmt.Printf("end = %d \n", endpage1)

	// fmt.Println("\npage 2")

	// startpage2 := endpage1 + 1

	// // startpage_ke := page_ke * itemPerPage
	// //

	// endpage2 := startpage2 + (itemPerPage - 1)

	// fmt.Printf("start = %d \n", startpage2)
	// fmt.Printf("end = %d \n", endpage2)

	// fmt.Println("\npage 3")

	// startpage3 := endpage2 + 1
	// endpage3 := startpage3 + (itemPerPage - 1)

	// fmt.Printf("start = %d \n", startpage3)
	// fmt.Printf("end = %d \n", endpage3)

}
