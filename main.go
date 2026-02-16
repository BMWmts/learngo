package main

import "fmt"

func main() {
	name := "Pongsapak"
	age := 24
	fmt.Printf("Name: %s\nAge: %d\n", name, age)
	theDefaultGopher()
}

func theDefaultGopher() {
	//theDefaultGopher (ตัวเล็กนำ = private function ใช้แค่ในไฟล์/package นี้)
	//TheDefaultGopher (ตัวใหญ่นำ = public function ให้คนอื่นเรียกใช้ได้)
	var score int
	var isPassed bool
	var remark string

	var bonus float64 = 5.5

	totalScore := float64(score) + bonus

	fmt.Println("Score: %d \nPassed: %t \nRemark: %q \n", score, isPassed, remark)
	fmt.Println("Total: ", totalScore)
}

