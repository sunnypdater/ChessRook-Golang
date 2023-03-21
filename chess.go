package main

import "fmt"

func main() {
    var n, k int
    fmt.Scan(&n, &k)

    rooks := make([][2]int, k)
    for i := 0; i < k; i++ {
        fmt.Scan(&rooks[i][0], &rooks[i][1])
    }

    count := n*n
    for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			hasRook := false
			for _, r := range rooks {
				if (r[0] == i || r[1] == j) {
					hasRook = true
					break
				}
			}
			if hasRook {
				count--
			}
		}
    }

    fmt.Println(count)
}