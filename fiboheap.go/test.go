package main

import (
    "fmt"
    "container/heap"
)

type ClassRecord struct {
    name  string
    grade int
}

type RecordHeap []*ClassRecord

func (p RecordHeap) Len() int { return len(p) }

func (p RecordHeap) Less(i, j int) bool {
    return p[i].grade < p[j].grade
}

func (p *RecordHeap) Swap(i, j int) {
    a := *p
    a[i], a[j] = a[j], a[i]
}

func (p *RecordHeap) Push(x interface{}) {
    a := *p
    n := len(a)
    a = a[0 : n+1]
    r := x.(*ClassRecord)
    a[n] = r
    *p = a
}

func (p *RecordHeap) Pop() interface{} {
    a := *p
    *p = a[0 : len(a)-1]
    r := a[len(a)-1]
    return r
}

func main() {
    a := make([]ClassRecord, 6)
    a[0] = ClassRecord{"John", 80}
    a[1] = ClassRecord{"Dan", 85}
    a[2] = ClassRecord{"Aron", 90}
    a[3] = ClassRecord{"Mark", 65}
    a[4] = ClassRecord{"Rob", 99}
    a[5] = ClassRecord{"Brian", 78}
    h := make(RecordHeap, 0, 100)
    for _, c := range a {
        fmt.Println(c)
		t := c
        heap.Push(&h, &t)
        fmt.Println("Push: heap has", h.Len(), "items")
    }

	fmt.Println()
	for h.Len() > 0 {
		x := heap.Pop(&h).(*ClassRecord)
	    fmt.Println(*x)
	}

}