/*
Actual solution of the problem will be here
*/

package main

import (
	co "ex3/pkg/config" // This adds coloring
	so "ex3/pkg/work" // Actual Solution
	wo "ex3/pkg/work" // This is the solution
)

func main() {
	co.ColoBanner("Actual excercise with proper tree")
	wo.Tree()
	wo.OpenFile()
	so.T1()

}
