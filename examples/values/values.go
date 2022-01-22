// GoLang এ বিভিন্ন ধরণের types (data types) আছে। তাদের মধ্যে কিছু হলো
// strings, integers (পূর্ণ সংখ্যা) , floats (ভগ্নাংশ), booleans (সত্য-মিথ্যা value) ইত্যাদি। এখানে কিছু basic example দেয়া হলো ।
package main

import "fmt"

func main() {

	// Strings, which can be added together with `+`.
	// ২ বা তার বেশি String একসাথে যুক্ত করতে বা জোড়া লাগাতে (সংযুক্ত নয়) `+` ব্যবহার করা হয়।
	fmt.Println("go" + "lang")

	// পূর্ণ সংখ্যা ও ভগ্নাংশ ।
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleans, with boolean operators যেভাবে আপনি expect করবেন ।
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
