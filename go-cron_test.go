package cron

import (
	"testing"
	"fmt"
	"time"
)

type testRemoveJob struct {
	id string
}

func (tj *testRemoveJob) ID() string {
	return tj.id
}

func (tj *testRemoveJob) Run() {
	fmt.Printf("demo: %s\n", tj.id)
}

func NewTestRemoveJob(id string) Job {
	return &testRemoveJob{
		id: id,
	}
}

func TestRemoveJob(t *testing.T) {
	var stopChan chan struct{}
	c := New()
	c.AddJob(" */10 * * * *", NewTestRemoveJob("1"))
	c.Start()

	c.AddJob(" */10 * * * *", NewTestRemoveJob("2"))

	<-time.After(time.Second * 10)
	c.RemoveJob("1")
	<-time.After(time.Second * 10)
	c.AddJob(" */5 * * * *", NewTestRemoveJob("1"))
	<-stopChan
}
