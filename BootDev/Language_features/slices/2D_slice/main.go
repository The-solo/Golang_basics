package main

func createMatrix(rows, cols int) [][]int {
	dashboard := make([][]int, rows) // pre-allocating for the rows.
	for i := range dashboard {
    	dashboard[i] = make([]int, cols) // pre-allocating the cols for this row.
    	for j := range dashboard[i] {
        	dashboard[i][j] = i*j
    	}
	}
	return dashboard
}

// Because we cannot access or write to the memory that has not been granted.
// There are no such things as colums only the sub-slices of rows. 



/*
func createMatrix(rows, cols int) [][]int{
	dashboard := make([][]int, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dashboard[i] = append(dashboard[i], i*j) // Inserting into the i'th row as a subslice
		}
	}
	return dashboard
}
*/

// The append() func changes the underlying array of it's parameter AND returns a new slice.
// This means using the append() on other than itself is usually a bad idea. 
