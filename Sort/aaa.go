package Sort

type Node struct {
	left *Node
	right *Node
}

func getHeight(p *Node) int {
	height := 0
	if p != nil {
		leftH := getHeight(p.left)
		rightH := getHeight(p.left)
		if leftH>rightH {
			height = leftH+1
		}else {
			height = rightH+1
		}
	}
	return height
}
func getWidth(p *Node) int {
	if p == nil {
		return 0
	}
	nLastLevelWidth :=0
	nCurLevelWidth :=0
	var queue []*Node
	queue = append(queue,p)
	width := 1
	nLastLevelWidth = 1
	var pCur *Node =nil
	for len(queue)>0 {
		for nLastLevelWidth != 0{
			pCur = queue[0]
			queue = queue[1:len(queue)]
			if pCur.left != nil{
				queue = append(queue, pCur.left)
			}
			if pCur.right != nil{
				queue = append(queue, pCur.right)
			}
			nLastLevelWidth--
		}
		nCurLevelWidth = len(queue)-1
		if nCurLevelWidth >width{
			width = nCurLevelWidth
		}
		nLastLevelWidth = nCurLevelWidth
	}
	return width
}

func getArea(p *Node) int {
	return getHeight(p)*getWidth(p)
}