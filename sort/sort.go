package sort

import (
	"time"
	"container/vector"
)

type Algorithm struct {
	Name string
	Time int
}

//function to time events
func TimeEvent(f func(d []int), data []int, name string, done chan Algorithm) {
	curTime := time.Nanoseconds()
	f(data) //call the sort function with unsorted data
	now := time.Nanoseconds() - curTime
	done <- Algorithm{name, int(now / 1e3)} //return time in microseconds
}

//Insertion sort algorithm
func InsertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		for j := i; j > 0 && data[j] < data[j-1]; j-- { //
			data[j], data[j-1] = data[j-1], data[j] //swap data[j] & data[j-1]
		}
	}
}

func QuickSort(data []int) {
	partition := func(left, right, pivotIndex int) int {
		pivot := data[pivotIndex]
		data[right], data[pivotIndex] = data[pivotIndex], data[right] //move pivot to end
		storeIndex := left
		for i := left; i < right; i++ {
			if data[i] <= pivot {
				data[i], data[storeIndex] = data[storeIndex], data[i] //swap them
				storeIndex++
			}
		}
		data[right], data[storeIndex] = data[storeIndex], data[right] //move pivot to its position
		return storeIndex
	}
	var sort func(int, int) //declared to use it recursively
	sort = func(left, right int) {
		if left >= right {
			return
		}
		pivotIndex := (right-left)/2 + left
		newIndex := partition(left, right, pivotIndex)
		sort(left, newIndex-1)
		sort(newIndex+1, right)
	}
	sort(0, len(data)-1)
}

func TreeSort(data []int) {
	if len(data) <= 1 {
		return
	}
	root := NewTree(nil, data[0]) //root node
	for i := 1; i < len(data); i++ {
		root.NewNode(data[i]) //filling up the tree
	}
	var sorted vector.IntVector
	var sortIntoArray func(*Tree) //declared to use recursively
	sortIntoArray = func(node *Tree) {
		if node == nil {
			return
		}
		sortIntoArray(node.left)
		sorted.Push(node.val)
		sortIntoArray(node.right)
	}
	sortIntoArray(root)
	copy(data, sorted.Data()) //copy the sorted to the source
}

//Tree data structure and its methods
type Tree struct {
	left, right, parent *Tree
	val                 int
}

func NewTree(parent *Tree, val int) *Tree {
	tree := new(Tree)
	tree.val = val
	tree.parent = parent
	return tree
}

func (tree *Tree) NewNode(val int) {
	if tree.val > val {
		if tree.left == nil {
			tree.left = NewTree(tree, val)
		} else {
			tree.left.NewNode(val)
		}
	} else {
		if tree.right == nil {
			tree.right = NewTree(tree, val)
		} else {
			tree.right.NewNode(val)
		}
	}
}
