package main

import (
	"fmt"
	"sort"
)


var(
	unsortedMap map[string]int = map[string]int{"India": 20, "Canada": 70, "Germany": 15}
)

func sortMapKeys(){

	keys := make([]string, 0, len(unsortedMap))

	for k := range unsortedMap{
		keys = append(keys, k)
	}

	sort.Strings(keys)

	fmt.Println(keys)

	for _, k := range keys{
		fmt.Println(k, unsortedMap[k])
	}
}

func sortValues(){

	values := make([]int, 0 , len(unsortedMap))

	for _, v := range unsortedMap{
		values = append(values, v)
	}

	sort.Ints(values)

	for _, v := range values{
		fmt.Println(v)
	}
}

func mergeMaps(){
	first := map[string]int{"a": 1, "b": 2, "c": 3}
	second := map[string]int{"a": 1, "e": 5, "c": 3, "d": 4}
 
	for k, v := range second {
		first[k] = v
	}

	
 
	fmt.Println(first)
}

func main(){
	sortMapKeys()
	sortValues()
	mergeMaps()
}