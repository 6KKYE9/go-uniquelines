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
	got := uniqueGlobal(in, false, false)
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

func TestUniqueIgnoreCase(t *testing.T) {
	in := []string{"A", "a", "B"}
	got := uniqueGlobal(in, true, false)
	want := []string{"A", "B"}
	if len(got) != len(want) {
		t.Fatalf("忽略大小写去重不符: %#v", got)
	}
}

func TestUniqueCount(t *testing.T) {
	in := []string{"a", "a", "b"}
	got := uniqueGlobal(in, false, true)
	if len(got) != 2 {
		t.Fatalf("计数模式行数不符: %#v", got)
	}
	if got[0] != "   2  a" {
		t.Fatalf("计数格式不符: %q", got[0])
	}
}
