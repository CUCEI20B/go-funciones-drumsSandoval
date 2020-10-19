package main

import "fmt"

func Burbuja(s []int64)  {
	n := cap(s) -1
	aux := int64(0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[j] > s[j+1] {
				aux = s[j]
				s[j] = s[j+1]
				s[j+1] = aux
			}
		}
	}
}

func main()  {
	s := []int64 {5,4,3,2,1}
	Burbuja(s)
	fmt.Println(s)
}