package main

var li = []int{}

func main() {
	li = append(li, 1)

	// staticcheck can't find this, can sonarqube?
	println("%d", li[len(li)])
}
