package main

func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)   //返回值
	queue := make([]*Node, 0) //队列
	if root == nil {
		return res
	}
	queue = append(queue, root)

	depthMap := map[*Node]int{} //记录节点深度
	depthMap[root] = 0

	for len(queue) > 0 { //队列不空
		cur := queue[0]
		queue = queue[1:]
		curDepth := depthMap[cur]
		if len(res) <= curDepth { //新的一层
			res = append(res, make([]int, 0))
		}
		res[curDepth] = append(res[curDepth], cur.Val)
		for _, child := range cur.Children {
			depthMap[child] = curDepth + 1
			queue = append(queue, child)
		}
	}
	return res
}
