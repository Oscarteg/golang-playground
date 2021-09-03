package src_test

import (
	"github.com/stretchr/testify/assert"
	"golang-playground/src"
	"golang.org/x/tour/tree"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	ch := make(chan int)
	go src.Walk(tree.New(1), ch)
	assert.False(t,  src.Same(tree.New(1), tree.New(2)))
	assert.True(t, src.Same(tree.New(1), tree.New(1)))
	assert.False(t, src.Same(tree.New(2), tree.New(1)))

}
