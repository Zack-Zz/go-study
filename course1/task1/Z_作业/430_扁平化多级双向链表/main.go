package main

func main() {

}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {

	dfs(root)

	return root
}

func dfs(node *Node) (tail *Node) {
	cur := node
	for cur != nil {
		next := cur.Next

		if cur.Child != nil {
			cTail := dfs(cur.Child)

			next = cur.Next

			// 连接当前节点和子节点
			cur.Next = cur.Child
			cur.Child.Prev = cur

			// 连接尾结点和下一个节点
			if next != nil {
				cTail.Next = next
				next.Prev = cTail
			}

			cur.Child = nil
			tail = cTail
		} else {
			tail = cur
		}
		cur = next
	}
	return
}
