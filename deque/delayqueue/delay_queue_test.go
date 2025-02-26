package delayqueue

import (
	"log"
	"testing"
	"time"
)

func TestDelayQueue(t *testing.T) {
	q := NewDelayQueue[int]()
	q.Push(1, time.Now().Add(10*time.Second))
	q.Push(2, time.Now().Add(20*time.Second))
	q.Push(3, time.Now().Add(30*time.Second))
	q.Start()
	log.Printf("NowTime")
	for {
		v, _ := q.Pop()
		log.Printf("Get the value: %d", v)
	}
}
