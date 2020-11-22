package main

import (
	"fmt"
	"strconv"
)

// ----  IAny  ----------------------------------------------------

type IAny interface{}

func doStuff(thing IAny) {
	theInt, isInt := thing.(int32)
	theString, isString := thing.(string)
	if isInt {
		fmt.Printf("An int: %d\n", theInt)
	} else if isString {
		fmt.Printf("A string: %s\n", theString)
	} else {
		fmt.Printf("Something else: %v\n", thing)
	}
}

// ----  IHashable  ----------------------------------------------------

type IHashable interface {
	Hash() uint32
}

type MyString string
type MyInt int32

func (this MyString) Hash() uint32 {
	return uint32(len(this))
}
func (this MyInt) Hash() uint32 {
	return uint32(this)
}

func HashStuff(arr []IHashable) {
	for _, v := range arr {
		fmt.Printf("Value: %v / Hash: %d\n", v, v.Hash())
	}
}

// ----  HashSet  ----------------------------------------------------

type HashSet struct {
	content map[uint32]IHashable
}

func NewHashSet() *HashSet {
	HashSet := new(HashSet)
	HashSet.content = make(map[uint32]IHashable)
	return HashSet
}

func (this HashSet) Add(x IHashable) {
	this.content[x.Hash()] = x
}

func (this HashSet) Remove(x IHashable) {
	delete(this.content, x.Hash())
}

func (this HashSet) Contains(x IHashable) bool {
	_, ok := this.content[x.Hash()]
	return ok
}

// --------------------------------------------------------------

const myConst = 42
const foooo = "aaa"

func main() {
	mixedBag := []IHashable{MyInt(42), MyString("foobar")}
	HashStuff(mixedBag)

	hashSet := NewHashSet() //new(HashSet)
	hashSet.Add(MyInt(42))
	hashSet.Add(MyString("hi"))
	fmt.Printf("hashSet contains 42: %v\n", hashSet.Contains(MyInt(42)))
	fmt.Printf("hashSet contains 43: %v\n", hashSet.Contains(MyInt(43)))
	fmt.Printf("hashSet contains 'hi': %v\n", hashSet.Contains(MyString("hi")))
	fmt.Printf("hashSet contains 'hiu': %v\n", hashSet.Contains(MyString("hiu")))
	hashSet.Remove(MyInt(42))
	hashSet.Remove(MyInt(42))
	hashSet.Remove(MyInt(43))
	fmt.Printf("hashSet contains 42: %v\n", hashSet.Contains(MyInt(42)))
	fmt.Printf("hashSet contains 43: %v\n", hashSet.Contains(MyInt(43)))

	var i32 int32 = 42
	doStuff(i32)
	doStuff(42)
	doStuff(42.2)
	doStuff("asdf")
	doStuff([]int32{1, 2, 3})

	var i int = 42
	k := 42.1
	k = 11.1
	iStr := strconv.Itoa(i)
	fmt.Printf("my int   : %d %T\n", i, i)
	fmt.Printf("my str   : %s %T\n", iStr, iStr)
	fmt.Printf("my double: %f %T\n", k, k)
}
