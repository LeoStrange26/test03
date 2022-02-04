package main

import (
	listtest "DataStructureTest/listTest"
	treetest "DataStructureTest/treeTest"
	"fmt"
)
func main() {
	var num int = 5
	// fmt.Scanf("%d", &num)
	var head *listtest.Node = listtest.CreateList(num)
	head.ShowList()
	fmt.Println(treetest.Flag)
}