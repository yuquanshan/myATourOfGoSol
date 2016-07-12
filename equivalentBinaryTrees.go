package main

import "golang.org/x/tour/tree"
import "fmt"
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	if t != nil{
		Walk(t.Left,ch)
		ch <- t.Value
		Walk(t.Right,ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1,ch1)
	go Walk(t2,ch2)
	for i:=0; i<10; i++{
		v1 := <- ch1
		v2 := <- ch2
		if v1 != v2{
			return false
		}
	}
	return true
}

func main() {
	var t *tree.Tree
	t = tree.New(2)
	if t.Left.Left.Left == nil{
		fmt.Println("haha")
	}else{
		fmt.Println("hoho")
	}
	if Same(tree.New(1),tree.New(1)){
		fmt.Println("test1 passed")
	}
	if !Same(tree.New(1),tree.New(2)){
		fmt.Println("test2 passed")   
	}	   
}
