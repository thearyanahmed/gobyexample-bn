// GoLang character, string, boolean আর numeric value'র ধ্রবক (_constants_) support করে ।

package main

import (
	"fmt"
	"math"
)

// `const` দিয়ে constant value declare করে ।
const s string = "constant"

func main() {
	fmt.Println(s)

	// `const` statement সেসব যায়গায় ব্যবহার করা যেতে পারে যেসব যায়গায় `var` ব্যবহার করা যায়।
	const n = 500000000

	// ধ্রুবক/constant expression নির্ভুলতার সাথে পাটিগণিত
	// সম্পাদন করে।
	const d = 3e20 / n
	fmt.Println(d)

	// Numeric constant এর কোনো data type থাকে না যতক্ষন না দেয়া হচ্ছে।
	// যেমন explicite conversion এর মাধ্যমে।
	fmt.Println(int64(d))

	// একটি number কে নির্দিষ্ট context এ ব্যবহার করে একটি টাইপ
	// দেওয়া যেতে পারে যার জন্য প্রয়োজন, যেমন একটি
	// অ্যাসাইনমেন্ট বা ফাংশন কল। উদাহরণ স্বরূপ
	// `math.Sin` except করে `float64`.
	fmt.Println(math.Sin(n))
}
