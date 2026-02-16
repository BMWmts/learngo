package main

import "fmt"

func main() {
	name := "Pongsapak"
	age := 24
	fmt.Printf("Name: %s\nAge: %d\n", name, age)
	theDefaultGopher()
	ifelse_stm()
}

func theDefaultGopher() {
	//theDefaultGopher (ตัวเล็กนำ = private function ใช้แค่ในไฟล์/package นี้)
	//TheDefaultGopher (ตัวใหญ่นำ = public function ให้คนอื่นเรียกใช้ได้)
	var score int
	var isPassed bool
	var remark string

	var bonus float64 = 5.5

	totalScore := float64(score) + bonus

	fmt.Printf("Score: %d \nPassed: %t \nRemark: %q \n", score, isPassed, remark)
	fmt.Printf("Total: %.2f\n", totalScore)
}

func ifelse_stm() {

	var msg string
	age := 18

	if age < 18 {
		msg = "Entry Denied: Too young"
	} else if age >= 18 && age < 20 {
		msg = "Entry Allowed: Limited Access (No Alcohol)"
	} else {
		msg = "Entry Allowed: Full Access"
	}

	if age%2 == 0 {
		msg = msg + " (Lucky Day)"
	}

	fmt.Println(msg)
}
