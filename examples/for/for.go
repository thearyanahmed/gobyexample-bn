// `for` GoLang এর একমাত্র loop । এখানে কিছু  উদাহরণ দেয়া হলো ।
package main

import "fmt"

func main() {

	// সবথেকে basic, একটি single condition দিয়ে ।
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// একটি classic initial/condition/after `for` loop ।
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// Condition ছাড়া `for` চলতে থাকবে, যতক্ষন না loop থেকে `break` করা হয় বা
	// enclosing function থেকে `return` করা হয়।
	for {
		fmt.Println("loop")
		break
	}

	// Loop এর পরবর্তি iteration এ যাওয়ার জন্য `continue` (ব্যবহার) করা যায়।
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
