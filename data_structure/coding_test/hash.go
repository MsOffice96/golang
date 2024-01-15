package main

import "fmt"

func main() {

	var clothes [][]string

	sort_map := make(map[string][]string)

	// for _, cloth := range clothes {
	// 	if value, exist := sort_map[cloth[1]]; exist {
	// 		value = append(value, cloth[0])
	// 		sort_map[cloth[1]] = value
	// 	} else {
	// 		sort_map[cloth[1]] = []string{cloth[0]}
	// 	}
	// }

	for _, cloth := range clothes {
		my_cloth_list, exist := sort_map[cloth[1]]
		if exist {
			my_cloth_list = append(my_cloth_list, cloth[0])
			sort_map[cloth[1]] = my_cloth_list
		} else {
			sort_map[cloth[1]] = []string{cloth[0]}
		}
	}

	fmt.Printf("%+v\n", sort_map)

	total := 1
	only := 0

	for _, cloths := range sort_map {
		total *= len(cloths)
	}

	return total + only

}
