package main

import "fmt"

type TreeNodes struct {
    Val   int
    Left  *TreeNodes
    Right *TreeNodes
}

func lowestCommonAncestor(root, p, q *TreeNodes) *TreeNodes {
    if root == nil || root == p || root == q {
        return root
    }

    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)

    if left != nil && right != nil {
        return root
    }
    if left != nil {
        return left
    }
    return right
}

func main() {
    root := &TreeNodes{3, nil, nil}
    root.Left = &TreeNodes{5, nil, nil}
    root.Right = &TreeNodes{1, nil, nil}
    root.Left.Left = &TreeNodes{6, nil, nil}
    root.Left.Right = &TreeNodes{2, nil, nil}
    root.Right.Left = &TreeNodes{0, nil, nil}
    root.Right.Right = &TreeNodes{8, nil, nil}

    p := root.Left.Right       // 5
    q := root.Right.Right      // 1
    ancestor := lowestCommonAncestor(root, p, q)
    fmt.Println("LCA:", ancestor.Val) // LCA: 3
}
