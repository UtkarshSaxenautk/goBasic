package main

/*import "fmt"

type language interface {
	speed()
	code()
}

type lang struct {
	langname  string
	langcode  int
	langspeed string
}

func (l lang) Info() {
	fmt.Println(" Langauge is : ", l.langname, " with languagecode : ", l.langcode, " and has speed : ",
		l.langspeed)
}
func (l *lang) speed() {
	fmt.Println("Speed: ", l.langspeed)
}

func (l *lang) code() {
	fmt.Println("Code: ", l.langcode)
}
func main() {
	utk := true
	us := &utk
	//us = nil
	if us == nil {
		fmt.Println("Value is nil")
	} else if *us {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	numbers := []string{"c++", "java", "python"}
	for i := 0; i < 3; i++ {

		if numbers[i] == "c++" {
			fmt.Println("Value is 5")
		} else if numbers[i] == "java" {
			fmt.Println("value is 6")
		} else {
			fmt.Println("Vakue is 7")
		}
	}

	mymap := make(map[string]interface{})

	mymap["Name"] = "Utk"
	mymap["Lname"] = "Saxena"
	mymap["Sno"] = 90
	for k, v := range mymap {
		fmt.Printf("key: %s and value %v", k, v)
	}

	l1 := lang{
		langname:  "c++",
		langcode:  89,
		langspeed: "very fast",
	}
	l2 := lang{
		langname:  "java",
		langcode:  90,
		langspeed: "slow",
	}
	l3 := lang{
		langname:  "python",
		langcode:  91,
		langspeed: "very slow",
	}

	l1.Info()
	l2.Info()
	l3.Info()
	l1.code()
	l1.speed()
	l2.code()
	l2.speed()
	l3.code()
	l3.speed()

	arr := [][]int{{2, 4, 5}, {4, 6, 7}, {6, 8, 10}}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			fmt.Println(" arr[", i, "][", j, "] = ", arr[i][j])
		}
	}



}*/
