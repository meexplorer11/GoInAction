package main

import "fmt"

func main() {

	//----------- Arrays -----------
	var names [5]string
	fmt.Println(names, len(names))

	names = [5]string{"a", "b"}
	fmt.Println(names, len(names))

	nums := [2]int{1, 2}
	fmt.Println(nums, len(nums))

	subjects := [...]string{"Physics"}
	fmt.Println(subjects, len(subjects))

	records := [...]string{1: "ABC", 3: "DEF"}
	fmt.Println(records, len(records))
	records[2] = "K"
	fmt.Println(records, len(records))

	array1 := [2]string{"one", "two"}
	var array2 [2]string

	array2 = array1
	fmt.Println(array2)

	array1[0] = "changed"
	fmt.Println(array1)
	fmt.Println(array2)

	var pointerArray [2]*string
	pointerArray = [2]*string{new(string), new(string)}

	*pointerArray[0] = "somevalue"
	fmt.Println(*pointerArray[0])

	var somePtrArray [2]*string
	somePtrArray = pointerArray
	fmt.Println(*somePtrArray[0])

	*pointerArray[1] = "updated"
	fmt.Println(*pointerArray[0], *pointerArray[1])
	fmt.Println(*somePtrArray[0], *somePtrArray[1])

	a := [...]int{1, 2, 3}
	passByValue(a)
	fmt.Println(a)

	passByReference(&a)
	fmt.Println(a)

	//----------- Slices -----------
	langs := []string{"C", "C++", "Java"}
	fmt.Println(langs)
	langs = append(langs, "GO")
	fmt.Println(langs)

	for _, l := range langs {
		fmt.Println(l)
	}

	shows := make([]string, 3)
	shows = []string{"Breaking Bad", "Dark"}
	shows = append(shows, "Fringe")
	fmt.Println(shows)

	for index, value := range shows {
		fmt.Println(index, value)
	}

	//----------- maps -----------
	m := make(map[string]string)
	fmt.Println(m)
	m["key"] = "value"
	fmt.Println(m)

	var m1 = map[string]int{"one": 1}
	fmt.Println(m1)

	var m2 map[int]int
	fmt.Println(m2)
	m2 = map[int]int{}
	m2[1] = 2
	fmt.Println(m2)

	value, exists := m2[2]
	fmt.Println(value, exists)

	m3 := map[string]string{
		"k1": "v1", "k2": "v2",
	}

	for key, value := range m3 {
		fmt.Println(key, value)
	}
}

func passByValue(a [3]int) {
	a[1] = 4
	fmt.Println(a)
}

func passByReference(a *[3]int) {
	a[1] = 5
	fmt.Println(*a)
}
