package main

import "fmt"

type PageNode struct {
	url  string
	prev *PageNode
	next *PageNode
}

type BrowserHistory struct {
	current *PageNode
}

func New(homepage string) *BrowserHistory {

	root := &PageNode{
		url:  homepage,
		prev: nil,
		next: nil,
	}

	return &BrowserHistory{
		current: root,
	}
}

func (b *BrowserHistory) Visit(url string) {
	newPage := PageNode{
		url:  url,
		prev: b.current,
		next: nil,
	}

	b.current.next = &newPage

	b.current = &newPage
}

func (b *BrowserHistory) Forward(steps int) string {

	for steps > 0 && b.current.next != nil {
		b.current = b.current.next
		steps--
	}
	return b.current.url
}

func (b *BrowserHistory) Back(steps int) string {

	for steps > 0 && b.current.prev != nil {
		b.current = b.current.prev
		steps--
	}

	return b.current.url
}

func main() {

	bh := New("leetcode.com")
	bh.Visit("google.com")
	bh.Visit("facebook.com")
	bh.Visit("youtube.com")

	fmt.Println(bh.Back(1))
	fmt.Println(bh.Back(1))
	fmt.Println(bh.Forward(1))

	bh.Visit("linkdin.com")

	fmt.Println(bh.Forward(2))
	fmt.Println(bh.Back(2))
	fmt.Println(bh.Back(7))
}
