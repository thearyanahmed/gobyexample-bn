// Go language এ _array_ হল নির্দিষ্ট length এর ক্রম অনুযায়ি elements এর sequence ।
package main

import "fmt"

func main() {

	// এখানে আমরা একটি array `a` তৈরি করছি যেটি exactly 5 টি `int` ডাটা বা element রাখবে ।
	// Element এর type ( data type ) আর length, উভয় ই array'র অংশ । Default ভাবে
	// array zero-valued, `int` এর ক্ষেত্রে `0`.
	var a [5]int
	fmt.Println("emp:", a)

	// আমরা array'র একটি index এ value set করতে পারি
	// `array[index] = value` syntax দিয়ে।
	// আর (ঐ / কোনো) index এর value পেতে পারি
	// `array[index]` syntax এর মাধ্যমে।
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// Builtin `len` (function) array'র length return করে ।
	fmt.Println("len:", len(a))

	// এই syntax এর মাধ্যমে একই লাইনে (array) declare and initialize করা যায় ।
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Array one-dimensional (এক-মাত্রিক), তবে আপনি types compose করে
	// multi-dimensional (বহু-মাত্রিক) data structure তৈরি করতে পারেন।
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
