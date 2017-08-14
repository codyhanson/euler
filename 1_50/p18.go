package main

/*
By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.

   3
  7 4
 2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom of the triangle below:

				   75
				  95 64
				17 47 82
			  18 35 87 10
			 20 04 82 47 65
			19 01 23 75 03 34
		   88 02 77 73 07 63 67
		 99 65 04 28 06 16 70 92
		41 41 26 56 83 40 80 70 33
	   41 48 72 33 47 32 37 16 94 29
	  53 71 44 65 25 43 91 52 97 51 14
	 70 11 33 28 77 73 17 78 39 68 17 57
	91 71 52 38 17 14 91 43 58 50 27 29 48
   63 66 04 68 89 53 67 30 73 16 69 87 40 31
  04 62 98 27 23 09 70 98 73 93 38 53 60 04 23

NOTE: As there are only 16384 routes, it is possible to solve this problem by trying every route. However, Problem 67, is the same challenge with a triangle containing one-hundred rows; it cannot be solved by brute force, and requires a clever method! ;o)
*/

import (
	"fmt"
)

type node struct {
	val    int
	lchild *node
	rchild *node
}

func (n *node) String() string {
	lchildVal := -1
	rchildVal := -1
	if n.lchild != nil {
		lchildVal = n.lchild.val
	}
	if n.rchild != nil {
		rchildVal = n.rchild.val
	}
	return fmt.Sprintf("val: %d l:%d, r:%d", n.val, lchildVal, rchildVal)
}

func childOffsetInLevel(parentOffsetInLevel int, isLeft bool) int {
	if isLeft {
		return parentOffsetInLevel
	} else {
		return parentOffsetInLevel + 1
	}
}

func offsetInNodeVals(level, offsetInLevel int) int {
	//1, 0 = 0
	//2, 0 = 1
	//2, 1 = 2
	//3, 0 = 3
	//3, 1 = 4
	//3, 2 = 5
	//4, 0 = 6
	//4, 1 = 7
	//4, 2 = 8
	//4, 3 = 9
	//5, 0 = 10
	//5, 1 = 11
	//5, 2 = 12
	//5, 3 = 13
	//5, 4 = 14
	//6, 0 = 15
	//6, 1 = 15
	//6, 2 = 15
	//6, 3 = 15
	//6, 4 = 15
	offset := 0
	for i := 1; i < level; i++ {
		offset += i
	}
	return offset + offsetInLevel
}

func nodeAtKey(level, offsetInLevel int) string {
	return fmt.Sprintf("%d,%d", level, offsetInLevel)
}

//caches nodes, since we are not building a perfect tree
var nodeAt = make(map[string]*node)

func buildTree(level, offsetInLevel int) *node {
	curNodeOffset := offsetInNodeVals(level, offsetInLevel)
	if curNodeOffset >= len(nodeVals)-1 {
		return nil
	}
	//check the node cache to see if we've created it before
	cachedNode, ok := nodeAt[nodeAtKey(level, offsetInLevel)]
	if ok {
		return cachedNode
	}
	curNode := &node{
		val:    nodeVals[curNodeOffset],
		lchild: buildTree(level+1, childOffsetInLevel(offsetInLevel, true)),
		rchild: buildTree(level+1, childOffsetInLevel(offsetInLevel, false)),
	}
	//cache node
	nodeAt[nodeAtKey(level, offsetInLevel)] = curNode
	return curNode
}

var nodeVals = []int{75, 95, 64, 17, 47, 82, 18, 35, 87, 10, 20, 04, 82, 47, 65, 19, 01, 23, 75, 03, 34, 88, 02, 77, 73, 07, 63, 67, 99, 65, 04, 28, 06, 16, 70, 92, 41, 41, 26, 56, 83, 40, 80, 70, 33, 41, 48, 72, 33, 47, 32, 37, 16, 94, 29, 53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14, 70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57, 91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48, 63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31, 04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23}

var traverseMaxSum = 0

func traversePaths(curNode *node, accum int) {

	if curNode == nil {
		//at the final level.
		if accum > traverseMaxSum {
			fmt.Printf("New Max %d\n", accum)
			traverseMaxSum = accum

		}
		return
	}
	traversePaths(curNode.lchild, accum+curNode.val)
	traversePaths(curNode.rchild, accum+curNode.val)
}

func main() {

	//depth first traverse each path of the tree and compute sums.
	fmt.Println(len(nodeVals))
	root := buildTree(1, 0)
	traversePaths(root, 0)
	fmt.Println(traverseMaxSum)
}
