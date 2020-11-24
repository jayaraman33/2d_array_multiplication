package main

import "fmt"

func main() {
	var m1 [10][10]int
	var m2 [10][10]int
	var result [10][10]int
	var i, j, k, m, n, p, q, total int
	total = 0
	fmt.Print("Enter the number of rows the 1st matrix: ")
	fmt.Scanln(&m)
	fmt.Print("Enter the number of columns the 1st matrix: ")
	fmt.Scanln(&n)
	fmt.Print("Enter the number of rows the 2nd matrix: ")
	fmt.Scanln(&p)
	fmt.Print("Enter the number of columns the 2nd matrix : ")
	fmt.Scanln(&q)
	if n != p {
		fmt.Println("Error: The matrix cannot be multiplied")
	} else {
		fmt.Println("Enter the 1st matrix elements: ")
		for i = 0; i < m; i++ {
			for j = 0; j < n; j++ {
				fmt.Scan(&m1[i][j])
			}
		}
		fmt.Println("Enter the 2nd matrix elements: ")
		for i = 0; i < p; i++ {
			for j = 0; j < q; j++ {
				fmt.Scan(&m2[i][j])
			}
		}
		for i = 0; i < m; i++ {
			for j = 0; j < q; j++ {
				for k = 0; k < p; k++ {
					total = total + m1[i][k]*m2[k][j]
				}
				result[i][j] = total
				total = 0
			}
		}
		fmt.Println("Results:")
		for i = 0; i < m; i++ {
			for j = 0; j < n; j++ {
				fmt.Print(result[i][j], "\t")
			}
			fmt.Print("\n")
		}
	}
}
