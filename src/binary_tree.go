package src

import "golang.org/x/tour/tree"

func Walk(t *tree.Tree, ch chan int) {
	WalkRecursive(t, ch)
	close(ch)
}

func WalkRecursive(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	WalkRecursive(t.Left, ch)
	ch <- t.Value
	WalkRecursive(t.Right, ch)
}

// determines if trees t1 and t2 are same values
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		n1, ok1 := <- ch1
		n2, ok2 := <- ch2
		if ok1 != ok2 || n1 != n2 {
			return false
		}
		if !ok1 {
			break;
		}
	}
	return true
}
