package main
import "fmt"

func main() {
    var t int
    fmt.Scanf("%d", &t)
    
    nArr := make([]int, t)
    townArr := make([][]int64, 0)
    for i:=0; i<t; i++ {
        fmt.Scanf("%d", &nArr[i])
        
        towns := make([]int64, nArr[i]-1)
        for j:=0; j<nArr[i]-1; j++ {
            fmt.Scanf("%d", &towns[j])
        }
        townArr = append(townArr, towns)
    }
    
    for i:=0; i<t; i++ {
        multiple := make([]int64, 1)
        multiple[0] = 1
        
        for j:=0; j<nArr[i]-1; j++ {
            Multiply(townArr[i][j], &multiple)
            //fmt.Println(townArr[i][j], multiple)
        }
        
        fmt.Println(findDivisor(1234567, &multiple))
    }
}

func Multiply(value int64, multiple *[]int64) {
    var carryOver int64 = 0
    for i:=0; i<len(*multiple); i++ {
        (*multiple)[i] *= value
        (*multiple)[i] += carryOver
        
        if (*multiple)[i] >= 100000000000 {
            carryOver = (*multiple)[i]/100000000000
            (*multiple)[i] = (*multiple)[i]%100000000000
        } else {
            carryOver = 0
        }
    }
    
    if carryOver > 0 {
        *multiple = append(*multiple, carryOver)
    }
}

func findDivisor(divisor int64, multiple *[]int64) int64 {
    var carryOver int64 = 0
    for i:=len(*multiple)-1; i>=0; i-- {
        value := (carryOver * 100000000000) + (*multiple)[i]
        carryOver = value%divisor
    }
    return carryOver
}