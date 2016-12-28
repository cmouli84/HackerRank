package main
import "fmt"

func main() {
    var n, q int
    
    fmt.Scanf("%d", &n)
    fmt.Scanf("%d", &q)
    
    nodes := make([]Node, n)
    val := 0
    for i:=0; i<n; i++ {
        fmt.Scanf("%d", &val)
        nodes[i] = Node{Index: i+1, Value: val}
    }
    
    var node1, node2 int 
    for i:=0; i<n-1; i++ {
        fmt.Scanf("%d", &node1)
        fmt.Scanf("%d", &node2)
        
        nodes[node1-1].Siblings = append(nodes[node1-1].Siblings, &(nodes[node2-1]))
        nodes[node2-1].Siblings = append(nodes[node2-1].Siblings, &(nodes[node1-1]))
    }

    uArr := make([]int, q)
    vArr := make([]int, q)
    for i:=0; i<q; i++ {
        fmt.Scanf("%d", &uArr[i])
        fmt.Scanf("%d", &vArr[i])
    }
    
    for i:=0; i<q; i++ {
        path := make([]int, 0) 
        findPath(&(nodes[uArr[i]-1]), &(nodes[vArr[i]-1]), &path)
        
        //fmt.Println(path)

        count := 0
        for j:=0; j<len(path)-1; j++ {
            for k:=j+1; k<len(path); k++ {
                if gcd(path[j], path[k]) == 1 {
                    count++
                }
            }
        }
        
        fmt.Println(count)
        
        for j:=0; j<n; j++ {
            nodes[j].Visited = false
        }
    }
}

func findPath(current, destination *Node, path *[]int) bool {
    if current.Index == destination.Index {
        *path = append(*path, destination.Value)
        return true
    }
    
    current.Visited = true
    for index := range current.Siblings {
        if current.Siblings[index].Visited {
            continue
        }
        
        if findPath(current.Siblings[index], destination, path) {
            *path = append(*path, current.Value)
            return true
        }
    }
    
    return false
}

func gcd(a, b int) int {
    if a%b == 0 {
        return b
    }
    return gcd(b, a%b)
}

type Node struct {
    Index int
    Value int
    Visited bool
    Siblings []*Node
}