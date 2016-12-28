package main
import "fmt"

func main() {
    var n int64
    fmt.Scanf("%d", &n)
    
    var m int
    fmt.Scanf("%d", &m)
    
    s := make([]int64, m)
    validS := make([]int64, 0)
    for i:=0; i<m; i++ {
        fmt.Scanf("%d", &s[i])
        
        if n%s[i] == 0 {
            found := false
            for j:=0; j<len(validS); j++ {
                if validS[j] == s[i] {
                    found = true
                    break
                }
            }
            if !found {
                validS = append(validS, s[i])
            }
        }
    }
    
    node := &Node{}
    temp := make([]int64, 0)
    findCount(n, &validS, node, temp)
    
    setValue(node, 1, 1)    

    aWins := false
   // q := Queue{}
   // q.Push(node)
    for _, n := range node.childs {
        //n := q.Pop()
        if n.aWins {
            aWins = true
            break
        }
        //fmt.Println(n.cumulativeValue)
       // for _, child := range n.childs {
       //     q.Push(child)
       // }
    }

    if aWins {
        fmt.Println("First")
    } else {
        fmt.Println("Second")
    }
}

func setValue(n *Node, val int64, prod int64) {
    if len(n.childs) == 0 {
        if n.cumulativeValue%2 == 0 {
            n.bWins = true
        } else {
            n.aWins = true
        }
        return
    }
    
    aWins := false
    bWins := false
    for _, child := range n.childs {
        child.cumulativeValue = n.cumulativeValue + (prod*val)
        setValue(child, child.value, (prod*val))
        
        if child.aWins {
            aWins = true
        } else {
            bWins = true
        }
    }
    if (aWins && !bWins) || (aWins && n.cumulativeValue%2==0) {
        n.aWins = true
    } else {
        n.bWins = true
    }
}

func findCount(n int64, sArr *[]int64, node *Node, tmp []int64) {
    if n!=1 {
        for _, s := range *sArr {
            if n%s==0 {
                child := node.AddChild(s)
                tmp1 := append(tmp, s)
                //fmt.Println("n/s", n/s, "s", s, "n", n, "tmp", tmp1)
                findCount(n/s, sArr, child, tmp1)
            }
        }
    }
    return
}

type Node struct {
    parent *Node
    value int64
    cumulativeValue int64
    childs []*Node
    aWins bool
    bWins bool
}

func (n *Node) AddChild(childNodeValue int64) *Node  {
    childNode := new(Node)
    childNode.parent = n
    childNode.childs = make([]*Node, 0)
    childNode.value = childNodeValue
    n.childs = append(n.childs, childNode)
    
    return childNode
}

type Queue struct {
    nodes []*Node
    isEmpty bool
}

func (q *Queue) Push(n *Node) {
    q.nodes = append(q.nodes, n)
    q.isEmpty = false
}

func (q *Queue) Pop() *Node {
    x := q.nodes[0]
    q.nodes = q.nodes[1:]

    if len(q.nodes) == 0 {
        q.isEmpty = true
    }
    return x
}