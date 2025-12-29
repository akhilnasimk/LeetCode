func isEvenOddTree(root *TreeNode) bool {
    dep := Depth(root)
    arr := make([][]int, dep)
    return Traversal(root, 0, &arr)
}

func Traversal(node *TreeNode, index int, arr *[][]int) bool {
    if node == nil {
        return true
    }

    level := (*arr)[index]

    if index%2 == 0 {
        if node.Val%2 == 0 {
            return false
        }
        if len(level) > 0 && node.Val <= level[len(level)-1] {
            return false
        }
    } else {
        if node.Val%2 == 1 {
            return false
        }
        if len(level) > 0 && node.Val >= level[len(level)-1] {
            return false
        }
    }

    (*arr)[index] = append((*arr)[index], node.Val)

    return Traversal(node.Left, index+1, arr) &&
           Traversal(node.Right, index+1, arr)
}

func Depth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := Depth(root.Left)
    right := Depth(root.Right)
    if left > right {
        return left + 1
    }
    return right + 1
}
