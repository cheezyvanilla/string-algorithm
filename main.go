package main

import (
	"fmt"
)

// import "fmt"

func main(){
	// fmt.Println(isAnagram("kodok", "dokok"))
	repeatedChar, err := FirstRepeatedChar("ads")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(repeatedChar)
}

func isAnagram(a, b string) bool {
	param1 := map[string]int{}
	param2 := map[string]int{}
	
	for _, v := range a{
		if _, ok := param1[string(v)]; ok{
			param1[string(v)] += 1
		}else{
			param1[string(v)] = 1
		}
		
	}

	if len(a) == len(b){
		for _, v := range b{
			if _, ok := param2[string(v)]; ok{
				param2[string(v)] += 1
			}else{
				param2[string(v)] = 1
			}
		}
		
		for _, char := range b {
			if param1[string(char)] != param2[string(char)] {
				return false
			}
		}
		return true
	}else{
		return false
	}
}