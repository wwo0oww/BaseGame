package main

func main() {
	var i []int
	i = append(i, 1)
	i = append(i, 2)
	i = append(i, 3)
	i = i[3:]
	for _, item := range i {
		println(item)
	}
	println(len(i))
}
