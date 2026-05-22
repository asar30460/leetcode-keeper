package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	set map[int]bool // 值及位置
}

func Constructor() RandomizedSet {
	var set = make(map[int]bool)
	return RandomizedSet{
		set: set,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if !this.set[val] {
		this.set[val] = true
		return true
	}

	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if this.set[val] {
		delete(this.set, val)
		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.set) == 0 {
		return -1
	}

	rand_val := rand.Intn(len(this.set))
	start := 0;

	for i, _ := range this.set {
		if start == rand_val {
			return i
		}
		start++
	}

	return -1
}

func main() {
	obj := Constructor()
	fmt.Print("Test Case 1: [")
	fmt.Print(obj.Insert(5), ", ")
	fmt.Print(obj.Remove(5), ", ")
	fmt.Print(obj.GetRandom(), ", ")
	fmt.Print(obj.Insert(6), ", ")
	fmt.Print(obj.Insert(7), ", ")
	fmt.Print(obj.GetRandom(), "]\n")

	obj2 := Constructor()
	fmt.Print("Test Case 2: [")
	fmt.Print(obj2.Insert(1), ", ")
	fmt.Print(obj2.Remove(2), ", ")
	fmt.Print(obj2.Insert(2), ", ")
	fmt.Print(obj2.GetRandom(), ", ")
	fmt.Print(obj2.Remove(1), ", ")
	fmt.Print(obj2.Insert(2), ", ")
	fmt.Print(obj2.GetRandom(), "]\n")
}
