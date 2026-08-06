package main

import "testing"

func TestUniqueAdjacent(t *testing.T) {
	in := []string{"a", "a", "b", "b", "a"}
	got := unique(in, false)
	want := []string{"a", "b", "a"}
	if len(got) != len(want) {
		t.Fatalf("相邻去重长度不符: %#v", got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("相邻去重结果不符: %#v", got)
		}
	}
}

func TestUniqueGlobal(t *testing.T) {
	in := []string{"a", "a", "b", "b", "a"}
	got := unique(in, true)
	want := []string{"a", "b"}
	if len(got) != len(want) {
		t.Fatalf("全局去重长度不符: %#v", got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("全局去重结果不符: %#v", got)
		}
	}
}
