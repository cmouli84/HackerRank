package main
import (
    "fmt"
    "math"
)

func main() {
    var g int
    fmt.Scanf("%d", &g)
    
    maxN := 0
    nArr := make([]int, g)
    for i:=0; i<g; i++ {
        fmt.Scanf("%d", &nArr[i])
        if nArr[i] > maxN {
            maxN = nArr[i]
        }
    }
    
    prime := make([]int, maxN + 1)
    prime[1] = 0
    if maxN >= 2 {
        prime[2] = 2
    }
    for i := 3; i <= maxN; i += 2 {
		prime[i] = i
	}
    
    findPrime(&prime, maxN)
    //fmt.Println(prime)
    
    for i:=0; i<g; i++ {
        count := 0
        for j:=2; j<=nArr[i]; j++ {
            if prime[j] != 0 {
                count++
            }
        }
        if count%2 == 0 {
            fmt.Println("Bob")
        } else {
            fmt.Println("Alice")
        }
    }
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
            //fmt.Println(j)
			(*arr)[j] = 0
		}
	}
}