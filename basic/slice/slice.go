package main

import "log"

func main() {
	// Slice Test
	log.Println("Slice Test")

	slice_1 := make([]int, 1, 2)
	log.Println("slice_1", slice_1, "len(slice_1)", len(slice_1), "cap(slice_1)", cap(slice_1), "&slice_1 ", &slice_1[0])

	slice_1 = append(slice_1, 1)
	log.Println("slice_1", slice_1, "len(slice_1)", len(slice_1), "cap(slice_1)", cap(slice_1), "&slice_1 ", &slice_1[0])

	cpy_slice_1 := slice_1

	slice_1 = append(slice_1, 2)
	log.Println("slice_1", slice_1, "len(slice_1)", len(slice_1), "cap(slice_1)", cap(slice_1), "&slice_1 ", &slice_1[0])
	log.Println("cpy_slice_1", cpy_slice_1, "len(cpy_slice_1)", len(cpy_slice_1), "cap(cpy_slice_1)", cap(cpy_slice_1), "&cpy_slice_1 ", &cpy_slice_1[0])

	slice_2 := append(slice_1, 1, 2)
	log.Println("slice_2", slice_2, "len(slice_2)", len(slice_2), "cap(slice_2)", cap(slice_2), "&slice_2 ", &slice_2[0])

	for i, v := range slice_2 {
		log.Println("index", i, "value", v)
	}

}
