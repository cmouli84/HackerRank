package main
import "fmt"

func main() {
    var n, p int
    fmt.Scanf("%d", &n)
    fmt.Scanf("%d", &p)
    
    pairs := make([]Pair, p)
    uniqueNodes := make(map[int]*Node)
    for i:=0; i<p; i++ {
        fmt.Scanf("%d", &pairs[i].Node1)
        fmt.Scanf("%d", &pairs[i].Node2)
        uniqueNodes[pairs[i].Node1] = &Node{Val: pairs[i].Node1}
        uniqueNodes[pairs[i].Node2] = &Node{Val: pairs[i].Node2}
    }
    
    for i:=0; i<p; i++ {
        firstNode := uniqueNodes[pairs[i].Node1]
        secondNode := uniqueNodes[pairs[i].Node2]
        firstNode.Siblings = append(firstNode.Siblings, secondNode)
        secondNode.Siblings = append(secondNode.Siblings, firstNode)
    }
    
    tree := make([][]int, 0)
    for _, node := range uniqueNodes {
        branch := make([]int, 0)
        if node.Visited {
            continue
        }
        
        addSiblings(node, &branch)
        tree = append(tree, branch)
    }
    
    //fmt.Println(tree)
    //fmt.Println(len(tree))
    
    count := 0
    totalLen := 0
    for i:=0; i<len(tree); i++ {
        totalLen += len(tree[i])
        for j:=i+1; j<len(tree); j++ {
            count += len(tree[i])*len(tree[j])
        }
    }
    count += (n-totalLen)*totalLen
    count += ((n-totalLen)*(n-totalLen-1))/2
    
    fmt.Println(count)
}

func addSiblings(node *Node, branch *[]int) {
    if node.Visited {
        return
    }
    node.Visited = true
    *branch = append(*branch, node.Val)
    for index := range node.Siblings {
        addSiblings(node.Siblings[index], branch)
    }
}

type Pair struct {
    Node1 int
    Node2 int
}

type Node struct {
    Val int
    Siblings []*Node
    Visited bool
}