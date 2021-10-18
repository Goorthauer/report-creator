package models

import (
	"fmt"
	"sync/atomic"
)

var MaxTotalCount = int64(89)
var Step = int64(5)
var MaxPercent = int64(100)

type Progress struct {
	Total    int64
	Current  int64
	Callback CallbackInterface
}

func (pb *Progress) Increment() int {
	incrementValue := int(MaxTotalCount / pb.Total)
	return pb.Add(incrementValue)
}

func (pb *Progress) Finish() {
	pb.setCurrent64(MaxPercent)
}
func (pb *Progress) StandartStep() {
	pb.Add64(Step)
}

func (pb *Progress) Add(add int) int {
	return int(pb.Add64(int64(add)))
}

func (pb *Progress) Add64(add int64) int64 {
	progressString := fmt.Sprint(pb.Current + add)
	pb.Callback.SendStatus(progressString)
	return atomic.AddInt64(&pb.Current, add)
}

//func (pb *Progress) setCurrent(cur int) int {
//	return int(pb.setCurrent64(int64(cur)))
//}

func (pb *Progress) setCurrent64(cur int64) int64 {
	return atomic.AddInt64(&pb.Current, cur)
}
func (pb *Progress) setTotal(add int) int {
	return int(pb.setTotal64(int64(add)))
}
func (pb *Progress) setTotal64(add int64) int64 {
	return atomic.AddInt64(&pb.Total, add)
}

func NewProgress(callbackInterface CallbackInterface) *Progress {
	return &Progress{Callback: callbackInterface}
}
