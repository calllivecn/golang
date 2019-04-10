package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randint0_9() (n int) {
	n = rand.Intn(10)
	return
}

func main() {

	t1 := new(List)
	fmt.Printf("t1 的类型：%t,t1 的值：%v\n", t1, t1)
	t2 := &List{}
	fmt.Printf("t2 的类型：%t,t2 的值：%v\n", t2, t2)

	//

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ->", randint0_9())
	}

	fmt.Println()

	//
	var l *List
	if l == nil {
		fmt.Printf("l type: %v\n", l)
	}

	var i *int
	if i == nil {
		fmt.Printf("i type: %v\n", l)
	}

    l = Create(rand.Intn(10))
    fmt.Print("l: ")
	l.Show()
	fmt.Println("l lenght : ", l.Len())

    fmt.Println("------------------")
    // 指针
	var list *List
	//var list List

	rand_len := rand.Intn(30)
	fmt.Println("rand len: ", rand_len)

    fmt.Println("list %p: ", list)
	list.Create2(rand_len)
    fmt.Println("list %p: ", list)
	fmt.Println("list lenght: ", list.Len())
	list.Show()
}

type List struct {
	Data int
	Next *List
}

func Create(n int) *List {

	var p, head *List

	for i := 0; i < n; i++ {
		if head == nil {
			head = &List{Data: randint0_9()}
			p = head
		} else {
			tmp := &List{Data: randint0_9()}
			p.Next = tmp
			p = p.Next
		}
	}
	return head
}

func (head *List) Create2(n int) {

    p := new(List)
    //fmt.Printf("p := head, p 的类型：%t, p 的值：%v", p, p)
    head = p

	for i := 0; i < n; i++ {
		fmt.Println("create：", i, "head state: ", *head, "head lenght: ", head.Len())
        tmp := new(List)
        tmp.Data = randint0_9()
		fmt.Println("p state: ", p)
        p.Next = tmp
        p = p.Next
	}
	//return p
}

func (t *List) Show() {
    fmt.Println("t.Show() --> ", t)
	p := t

	for p != nil {
		fmt.Printf("%d -->", p.Data)
		p = p.Next
	}

	fmt.Printf("\n")
}

func (t *List) Len() (n int) {
	n = 0
	p := t
	for p != nil {
		n++
		p = p.Next
	}

	return
}
