package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// 去重相邻重复行（像 uniq）
// ignoreCase 时先转小写再比较
func unique(lines []string, ignoreCase bool) []string {
	var out []string
	prev := ""
	for _, l := range lines {
		if ignoreCase {
			if strings.EqualFold(l, prev) {
				continue
			}
		} else if l == prev {
			continue
		}
		prev = l
		out = append(out, l)
	}
	return out
}

func main() {
	global := flag.Bool("g", false, "全局去重，而非仅合并相邻重复行")
	ignoreCase := flag.Bool("i", false, "比较时忽略大小写")
	count := flag.Bool("c", false, "在每行前显示它出现的次数")
	flag.Parse()

	var lines []string
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	// -g 全局去重和 -c 计数都依赖全集，这里统一走基于全集的去重
	if *global || *count {
		for _, l := range uniqueGlobal(lines, *ignoreCase, *count) {
			fmt.Println(l)
		}
		return
	}
	for _, l := range unique(lines, *ignoreCase) {
		fmt.Println(l)
	}
}

// uniqueGlobal 走全集去重：只保留首次出现，且支持计数
func uniqueGlobal(lines []string, ignoreCase, count bool) []string {
	key := func(s string) string {
		if ignoreCase {
			return strings.ToLower(s)
		}
		return s
	}
	counts := map[string]int{}
	for _, l := range lines {
		counts[key(l)]++
	}
	seen := map[string]bool{}
	var out []string
	for _, l := range lines {
		k := key(l)
		if seen[k] {
			continue
		}
		seen[k] = true
		if count {
			out = append(out, fmt.Sprintf("%4d  %s", counts[k], l))
		} else {
			out = append(out, l)
		}
	}
	return out
}
