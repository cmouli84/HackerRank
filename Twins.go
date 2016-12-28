package main

import (
	"fmt"
	"math"
)

func main() {
    var m, n int
    fmt.Scanf("%d", &m)
    fmt.Scanf("%d", &n)

//	fmt.Println(time.Now())
	arr := make([]int, n+1)
	arr[1] = 0
    if (n > 1) {
	   arr[2] = 2
    }
	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		arr[i] = i
	}

	startVal := sqrt
	if m > startVal {
		startVal = m
	}
	if (startVal % 2) == 0 {
		startVal++
	}
	for i := startVal; i <= n; i += 2 {
		arr[i] = i
	}

//	fmt.Println(time.Now())
	findPrime(&arr, sqrt)

//	fmt.Println(time.Now())
	for i := 3; i <= sqrt; i += 2 {
		if arr[i] == 0 {
			continue
		}

		startVal := int(math.Ceil(float64(m)/float64(arr[i])) * float64(arr[i]))
        if startVal == i {
            startVal += arr[i]
        }
		if startVal%2 == 0 {
			startVal += arr[i]
		}
		//fmt.Println(startVal, i)
		for j := startVal; j <= n; j += 2 * arr[i] {
			arr[j] = 0
		}
	}

//	fmt.Println(time.Now())
	count := 0
	startVal = m
	if m%2 == 0 {
		startVal = m + 1
	}
	for i := startVal; i <= n-2; i += 2 {
		if arr[i] != 0 && arr[i+2] != 0 {
			//			fmt.Print(arr[i], " ")
			count++
		}
	}
	//fmt.Println(time.Now())
	fmt.Println(count)
}

func findPrime(arr *[]int, val int) {
	sqrt := int(math.Sqrt(float64(val)))
	if val >= 4 {
		findPrime(arr, sqrt)
	}
	for i := 3; i <= sqrt; i += 2 {
        curr := (*arr)[i]
		if curr == 0 {
			continue
		}

		for j := 3 * curr; j <= val; j += 2 * curr {
			(*arr)[j] = 0
		}
	}
}