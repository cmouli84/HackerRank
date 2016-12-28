package main
import "fmt"

func main() {
    var q int
    fmt.Scanf("%d", &q)
    
    nArr := make([]int, q)
    mArr := make([]int, q)
    croadArr := make([]int, q)
    clibArr := make([]int, q)
    uArr := make([][]int, 0)
    vArr := make([][]int, 0)
    for i:=0; i<q; i++ {
        fmt.Scanf("%d", &nArr[i])
        fmt.Scanf("%d", &mArr[i])
        fmt.Scanf("%d", &clibArr[i])
        fmt.Scanf("%d", &croadArr[i])
        
        u := make([]int, mArr[i])
        v := make([]int, mArr[i])
        for j:=0; j<mArr[i]; j++ {
            fmt.Scanf("%d", &u[j])
            fmt.Scanf("%d", &v[j])
        }
        
        uArr = append(uArr, u)
        vArr = append(vArr, v)
    }
    
    for i:=0; i<q; i++ {
        if clibArr[i] <= croadArr[i] {
            fmt.Println(nArr[i] * clibArr[i])
            continue
        } 
        
        uniqueNodes := make(map[int]*Node)
        for j:=0; j<mArr[i]; j++ {
            var firstNode, secondNode *Node
            var found bool
            if firstNode, found = uniqueNodes[uArr[i][j]]; !found {
                firstNode = &Node{Val: uArr[i][j]}
                uniqueNodes[uArr[i][j]] = firstNode
            }
            if secondNode, found = uniqueNodes[vArr[i][j]]; !found {
                secondNode = &Node{Val: vArr[i][j]}
                uniqueNodes[vArr[i][j]] = secondNode
            }
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
        
        cost := 0
        totalLen := 0
        for j:=0; j<len(tree); j++ {
            totalLen += len(tree[j])
            cost += clibArr[i] + ((len(tree[j])-1)*croadArr[i])
        }
        cost += (nArr[i]-totalLen)*clibArr[i]

        fmt.Println(cost)
        
    }

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

type Node struct {
    Val int
    Siblings []*Node
    Visited bool
}