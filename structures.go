package main

import (
	"sort"
	"runtime"
	"rand"
	"graph"
)

var (
	x = []int{0, 1, 2, 3, 4, 5, 6}
	insertionVals = graph.Values{x, make([]int, 7)}
	quickVals = graph.Values{x, make([]int, 7)}
	treeVals = graph.Values{x, make([]int, 7)}
)

func main() {
	runtime.GOMAXPROCS(12) //max number of processes for concurrency
	for i := 1; i < 7; i++ {
		done := make(chan sort.Algorithm) //channel for communication
		nums := rand.Perm(1000 * i)
		println("\nusing", 1 * graph.Pow(10,i), "elements")
		nums1, nums2 := make([]int, len(nums)), make([]int, len(nums))
		copy(nums1, nums) //create copies of the array for others to sort
		copy(nums2, nums)
		go sort.TimeEvent(sort.InsertionSort, nums, "Insertion Sort", done)
		go sort.TimeEvent(sort.QuickSort, nums1, "Quick Sort", done)
		go sort.TimeEvent(sort.TreeSort, nums2, "Tree Sort", done)
		n := 0
		for alg := range done {//wait for the events to finish execution
			println(alg.Name, "finished in", alg.Time)
			switch alg.Name {
				case "Insertion Sort":
					insertionVals.Y[i] = alg.Time
				case "Quick Sort":
					quickVals.Y[i] = alg.Time
				case "Tree Sort":
					treeVals.Y[i] = alg.Time
			}
			n++
			if n > 2 { close(done) }//close the channel
		}
	}
	values := []graph.Values{treeVals, insertionVals, quickVals}
	graph.DrawToFile("output.svg", graph.NewGraph(100, 100, values))
}
