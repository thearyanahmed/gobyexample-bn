// In Go, _variables_ are explicitly declared and used by
// the compiler to e.g. check type-correctness of function
// calls.

// GoLang এ compiler _variables_ (দের) স্পষ্ট (explicite) ভাবে declare আর use করে ।
// একটি উদাহরণ হলো function এর type-correctness check করার জন্য ।

package main

import "fmt"

func main() {

	// `var` দিয়ে ১ বা একাধিক variable declare করা হয়।
	var a = "initial"
	fmt.Println(a)

	// একবারে একাধিক variable এভাবে declare করা হয়।
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go প্রাথমিক ভেরিয়েবলের ধরন অনুমান করবে ।
	var d = true
	fmt.Println(d)

	// যেসব variable এর declaration সময় কোনো value assign করা হবে না,
	// সেগুলো _zero-valued_. যেমন, `int` এর zero value হল `0` ।
	var e int
	fmt.Println(e)

	// `:=` syntax হলো shorthand, একই সাথে variable declare আর initialize করার জন্য ।
	// যেমন এক্ষেত্রে `var f string = "apple"`
	f := "apple"
	fmt.Println(f)
}
