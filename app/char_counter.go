package app
func CharCounter(s string) map[string]int{
	count := map[string]int{}

	for _, v := range s {
		if _, ok := count[string(v)]; ok {
			count[string(v)] += 1
		}else{
			count[string(v)] =1
		}
	}
	return count
}