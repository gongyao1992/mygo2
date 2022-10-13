package main

import "testing"

//func TestAdd(t *testing.T) {
//	sum := Add(1,2)
//	if sum == 3 {
//		t.Log("the result is ok")
//	} else {
//		t.Fatal("the result is wrong")
//	}
//}

/*
const N  =  1000

func initSlice() []string{
	s:=make([]string,N)
	for i:=0;i<N;i++{
		s[i]="www.flysnow.org"
	}
	return s;
}

func BenchmarkForSlice1(b *testing.B) {
	s:=initSlice()

	b.ResetTimer()
	for i:=0; i<b.N;i++  {
		ForSlice(s)
	}
}

func BenchmarkRangeForSlice(b *testing.B) {
	s:=initSlice()

	b.ResetTimer()
	for i:=0; i<b.N;i++  {
		RangeForSlice(s)
	}
}
*/

func BenchmarkStringPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringPlus()
	}
}

func BenchmarkStringPlusFtm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringPlusFtm()
	}
}

func BenchmarkStringPlusJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringPlusJoin()
	}
}

func BenchmarkStringPlusBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringPlusBuffer()
	}
}
