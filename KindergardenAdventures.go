package main
import "fmt"

func main() {
    var n int
    fmt.Scanf("%d", &n)
    
    arr := make([]int, n)
    for i:=0; i<n; i++ {
        fmt.Scanf("%d", &arr[i])
    }
    
	timeArr := make([]int, n+n+1)
	maxCount := 0
	minIndex := 0
	index := 0
	for i := n - 1; i >= 0; i-- {
		count := 0
		if i == n-1 {
			timeTaken := 0
			for j := 0; j < n; j++ {
				if (arr[(i+j)%n] - timeTaken) <= 0 {
					timeArr[0]++
					count++
				} else {
					timeArr[arr[(i+j)%n]-timeTaken]++
				}
				timeTaken++
			}
			//fmt.Println(time.Now())
		} else {
			timeArr[0] += timeArr[index]
			timeArr[0]--
			//timeArr = append(timeArr[:1], timeArr[2:]...)
			//timeArr = append(timeArr, 0)
            if arr[i] == 0 {
			    timeArr[0]++
            } else {
			    timeArr[arr[i]+index]++
            }
			count = timeArr[0]
		}
		//fmt.Println(timeArr)
		//fmt.Println(count)
		if (maxCount < count) || (maxCount == count && minIndex > i) {
			maxCount = count
			minIndex = i
		}
		index++
	}
	//fmt.Println(time.Now())
	fmt.Println(minIndex + 1)
}