package main

import (
	"container/list"
)

type Subject interface {
	Register(Observer)
	Release(Observer)
	notify()
}

type Observer interface {
	Update(Subject)
}

type ConcreteSubject struct {
	observers *list.List
	value     int
}

func NewConcreteSubject() *ConcreteSubject {
	s := &ConcreteSubject{}
	s.observers = list.New()
	return s
}

func (s *ConcreteSubject) Register(o Observer) {
	s.observers.PushBack(o)
}

func (s *ConcreteSubject) Release(o Observer) {
	for ob := s.observers.Front(); ob != nil; ob = ob.Next() {
		if ob.Value == o {
			s.observers.Remove(ob)
			break
		}
	}
}

func (s *ConcreteSubject) notify() {
	for ob := s.observers.Front(); ob != nil; ob = ob.Next() {
		ob.Value.(Observer).Update(s)
	}
}

func (s *ConcreteSubject) setValue(value int) {
	s.value = value
	s.notify()
}

func (s *ConcreteSubject) getValue() int {
	return s.value
}

type ConcreteObserver1 struct{}
type ConcreteObserver2 struct{}
type ConcreteObserver3 struct{}

func (c *ConcreteObserver1) Update(subject Subject) {
	println("ConcreteObserver1 value is ", subject.(*ConcreteSubject).getValue())
}

func (c *ConcreteObserver2) Update(subject Subject) {
	println("ConcreteObserver2 value is ", subject.(*ConcreteSubject).getValue())
}

func (c *ConcreteObserver3) Update(subject Subject) {
	println("ConcreteObserver3 value is ", subject.(*ConcreteSubject).getValue())
}

func main() {
	subject := NewConcreteSubject()
	observer1 := &ConcreteObserver1{}
	observer2 := &ConcreteObserver2{}
	observer3 := &ConcreteObserver3{}
	subject.Register(observer1)
	subject.Register(observer2)
	subject.Register(observer3)
	subject.setValue(5)

	subject.Release(observer3)
	subject.setValue(3)
}