package main

type User struct {
	Name string
	age  int
}

type UserHeap []User

func (uh UserHeap) Len() int {
	return len(uh)
}

func (uh UserHeap) Less(i, j int) bool {
	return uh[i].age < uh[j].age
}

func (uh UserHeap) Swap(i, j int) {
	uh[i], uh[j] = uh[j], uh[i]
}

func (uh *UserHeap) Push(u interface{}) {
	u_a := u.(User)
	*uh = append(*uh, u_a)
}

func (uh *UserHeap) Pop() interface{} {
	old := *uh
	n := len(old)
	r_u := old[n-1]

	*uh = old[0 : n-1]

	return r_u
}

// func main() {
// 	h := &UserHeap{}
// 	log.Println(h)
// 	heap.Init(h)

// 	u1 := User{
// 		Name: "minsoo",
// 		age:  12,
// 	}
// 	u2 := User{
// 		Name: "minsoo",
// 		age:  0,
// 	}

// 	u3 := User{
// 		Name: "minsoo",
// 		age:  11,
// 	}

// 	u4 := User{
// 		Name: "minsoo",
// 		age:  -3,
// 	}

// 	heap.Push(h, u1)
// 	heap.Push(h, u2)

// 	heap.Push(h, u3)

// 	heap.Push(h, u4)

// 	fmt.Println(h)

// }
