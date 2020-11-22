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

func HashList(arr []IHashable) {
	for _, v := range arr {
		fmt.Printf("Value: %v / Hash: %d\n", v, v.Hash())
	}
}

// ----  HashMap  ----------------------------------------------------

type HashMap struct {
	content map[uint32]IHashable
}

func NewHashMap() *HashMap {
	hashMap := new(HashMap)
	hashMap.content = make(map[uint32]IHashable)
	return hashMap
}

func (this HashMap) Add(x IHashable) {
	hash := x.Hash()
	this.content[hash] = x
}

// --------------------------------------------------------------

const myConst = 42
const foooo = "aaa"

func main() {
	mixedBag := []IHashable{MyInt(42), MyString("foobar")}
	HashList(mixedBag)

	hashMap := NewHashMap() //new(HashMap)
	hashMap.Add(MyInt(42))
	hashMap.Add(MyString("hi"))
	fmt.Printf("Hashmap contents: %v\n", hashMap.content)

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
