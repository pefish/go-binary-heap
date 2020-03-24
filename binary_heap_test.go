package go_binary_heap

import "fmt"

type Data struct {
	priority int
	test string
}

func (d *Data) GetPriority() int {
	return d.priority
}

func (d *Data) GetData() interface{} {
	return d.test
}

func ExampleMaxHeap_Push() {
	heap := NewMaxHeap([]NodeType{
		&Data{
			priority: 56,
			test:     "56",
		},
		&Data{
			priority: 12,
			test:     "12",
		},
		&Data{
			priority: 34,
			test:     "34",
		},
		&Data{
			priority: 2,
			test:     "2",
		},
		&Data{
			priority: 100,
			test:     "100",
		},
	}, HeapType_MaxBinaryHeap)

	for _, a := range heap.Elements {
		fmt.Println(a)
	}

	fmt.Println("")

	for range heap.Elements {
		fmt.Println(heap.Remove(0))
	}

	// Output:
	// &{100 100}
	//&{56 56}
	//&{34 34}
	//&{2 2}
	//&{12 12}
	//
	//&{100 100}
	//&{56 56}
	//&{34 34}
	//&{12 12}
	//&{2 2}
}
