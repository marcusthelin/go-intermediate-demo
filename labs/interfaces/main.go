package main

import "fmt"

type I interface {
  val() int
}

type S struct {
  num int
}

func (s S) val() int {
  return s.num
}

type B struct {
  num int
}

func (b B) val() int {
  return 0
}

func main() {
  var v I = S{10}
  fmt.Println(v.val())

  v = B{}

  fmt.Println(v.val())
}
