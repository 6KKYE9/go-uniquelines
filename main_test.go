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

func TestRemoveEmpty(t *testing.T) {
	in := []string{"a", "", "  ", "b", ""}
	got := removeEmpty(in)
	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Fatalf("删空行不符: %#v", got)
	}
}

func TestDupesAdjacent(t *testing.T) {
	in := []string{"a", "a", "b", "b", "a"}
	got := dupes(in, false, false)
	// 相邻重复：第2个a、第2个b 算重复，第3个a 和前一个a不相邻所以不算
	want := []string{"a", "b"}
	if len(got) != len(want) {
		t.Fatalf("相邻重复不符: %#v", got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("相邻重复结果不符: %#v", got)
		}
	}
}

// 以前 prev 初值是空串，首行正好是空行时会被当成重复直接丢掉
func TestUniqueKeepsLeadingBlankLine(t *testing.T) {
	in := []string{"", "a", "a", ""}
	got := unique(in, false)
	want := []string{"", "a", ""}
	if len(got) != len(want) {
		t.Fatalf("首行空行被吞了: %#v", got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("结果不符: %#v", got)
		}
	}
}

// 同样的坑：首行空行不该被算作「和上一行重复」
func TestDupesIgnoresLeadingBlankLine(t *testing.T) {
	got := dupes([]string{"", "a"}, false, false)
	if len(got) != 0 {
		t.Fatalf("首行不该被算成重复行: %#v", got)
	}
}

func TestDupesGlobal(t *testing.T) {
	in := []string{"a", "a", "b", "b", "a"}
	got := dupes(in, false, true)
	// 全局重复：除第一次出现的 a、b 外其余都是重复
	want := []string{"a", "b", "a"}
	if len(got) != len(want) {
		t.Fatalf("全局重复不符: %#v", got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("全局重复结果不符: %#v", got)
		}
	}
}
