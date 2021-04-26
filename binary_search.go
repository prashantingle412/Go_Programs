package main

import "fmt"

func main() {
	searchField := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	fmt.Println("half slice is", searchField[5+1:])
	result := binarySearch(searchField, 23)
	fmt.Printf("Your number was found in position %d steps.\n\n", result)
}

func binarySearch(a []int, search int) (result int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1 // not found
	case a[mid] > search:
		result = binarySearch(a[:mid], search)
	case a[mid] < search:
		result = binarySearch(a[mid+1:], search)
		result += mid + 1
	default: // a[mid] == search
		result = mid // found
	}
	return result
}
