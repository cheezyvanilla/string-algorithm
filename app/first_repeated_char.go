package app

import "errors"

func FirstRepeatedChar(a string) (string, error){
	temp := map[string]int{}

	for _, v := range a {
		if _, ok := temp[string(v)];ok{
			temp[string(v)] += 1
		}else{
			temp[string(v)] = 1
		}
	}

	for _, v := range a {
		if temp[string(v)] > 1 {
			return string(v), nil
		}
	}
	return "", errors.New("no repeated char found")
}