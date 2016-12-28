package main
import "fmt"

func main() {
    var g int
    fmt.Scanf("%d", &g)
    
    nArr := make([]int, g)
    mArr := make([]int, g)
    kArr := make([]int, g)
    strArr := make([][]string, 0)
    
    for i:=0; i<g; i++ {
        fmt.Scanf("%d", &nArr[i])
        fmt.Scanf("%d", &mArr[i])
        fmt.Scanf("%d", &kArr[i])
        
        gArr := make([]string, nArr[i])
        for j:=0; j<nArr[i]; j++ {
            fmt.Scanf("%s", &gArr[j])
        }
        strArr = append(strArr, gArr)
    }
    
    for i:=0; i<g; i++ {
        alexWins := PlayerWins(strArr[i], nArr[i], mArr[i], kArr[i], "O")
        otherWins := PlayerWins(strArr[i], nArr[i], mArr[i], kArr[i], "X")
        if alexWins && !otherWins {
            fmt.Println("WIN")
        } else if !alexWins && otherWins {
            fmt.Println("LOSE")
        } else {
            fmt.Println("NONE")
        }
    }
}

func PlayerWins(strArr []string, n, m, k int, symbol string) bool {
    for i:=0; i<n; i++ {
        count := 0
        for j:=0; j<m; j++ {
            if string(strArr[i][j]) == symbol {
                count++
                if count == k {
                    return true
                }
            } else {
                count = 0
                if (j+k)>=m {
                    break
                }
            }
        }
    }
    for j:=0; j<m; j++ {
        count := 0
        for i:=0; i<n; i++ {
            if string(strArr[i][j]) == symbol {
                count++
                if count == k {
                    return true
                }
            } else {
                count = 0
                if (i+k)>=n {
                    break
                }
            }
        }
    }
    
    min := n
    if m<n {
        min = m
    }
    if k>min {
        return false
    }

    for i:=0; i<n; i++ {
        count := 0
        for j:=0; j<m && (j+i)<m; j++ {
            if string(strArr[j][j+i]) == symbol {
                count++
                if count == k {
                    return true
                }
            } else {
                count = 0
                if (j+k)>=n {
                    break
                }
            }
        }
        if (i+k)>=m {
            break
        }
    }
    for i:=0; i<n; i++ {
        count := 0
        for j:=0; j<m && (j+i)<n; j++ {
            if string(strArr[j+i][j]) == symbol {
                count++
                if count == k {
                    return true
                }
            } else {
                count = 0
                if (j+k)>=m {
                    break
                }
            }
        }
        if (i+k)>=n {
            break
        }
    }

    for i:=0; i<m; i++ {
        count := 0
        for j:=n-1; j>=0 && (n-1-j)+i<m; j-- {
            if string(strArr[j][(n-1-j)+i]) == symbol {
                count++
                if count == k {
                    return true
                }
            } else {
                count = 0
                if ((n-1-j)+k)>=n {
                    break
                }
            }
        }
        if (i+k)>=m {
            break
        }
    }
    for i:=0; i<n; i++ {
        count := 0
        for j:=n-1; j>=0 && j-i>=0; j-- {
            if string(strArr[j-i][(n-1-j)]) == symbol {
                count++
                if count == k {
                    return true
                }
            } else {
                count = 0
                if ((n-1-j)+k)>=m {
                    break
                }
            }
        }
        if (i+k)>=n {
            break
        }
    }
    return false
}