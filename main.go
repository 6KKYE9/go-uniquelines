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
	// 不能拿 prev=="" 当初始状态，否则输入的第一行如果是空行会被当成重复吞掉
	first := true
	for _, l := range lines {
		if !first {
			if ignoreCase {
				if strings.EqualFold(l, prev) {
					continue
				}
			} else if l == prev {
				continue
			}
		}
		first = false
		prev = l
		out = append(out, l)
	}
	return out
}

func main() {
	global := flag.Bool("g", false, "全局去重，而非仅合并相邻重复行")
	ignoreCase := flag.Bool("i", false, "比较时忽略大小写")
	count := flag.Bool("c", false, "在每行前显示它出现的次数")
	dupe := flag.Bool("d", false, "只打印被合并掉的重复行（和 -c -g 互斥，单独用）")
	dropEmpty := flag.Bool("empty", false, "先删掉空行再处理")
	flag.Parse()

	var lines []string
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	// 单行超过 1MB 时 Scan 会静默停住，不查就等于悄悄少读一半
	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "读取输入失败: %v\n", err)
		os.Exit(1)
	}
	if *dropEmpty {
		lines = removeEmpty(lines)
	}
	// -d 单独用：把重复出现的行挖出来
	if *dupe {
		for _, l := range dupes(lines, *ignoreCase, *global) {
			fmt.Println(l)
		}
		return
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

// removeEmpty 删掉纯空白的行
func removeEmpty(lines []string) []string {
	var out []string
	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}
		out = append(out, l)
	}
	return out
}

// dupes 返回重复出现的行（第二次及以后出现的）。global 为真时按全集判断，
// 否则只算相邻重复
func dupes(lines []string, ignoreCase, global bool) []string {
	key := func(s string) string {
		if ignoreCase {
			return strings.ToLower(s)
		}
		return s
	}
	if global {
		seen := map[string]bool{}
		var out []string
		for _, l := range lines {
			k := key(l)
			if seen[k] {
				out = append(out, l)
			} else {
				seen[k] = true
			}
		}
		return out
	}
	var out []string
	prev := ""
	first := true
	for _, l := range lines {
		// 同上：首行不该拿空串去比
		if !first && ((ignoreCase && strings.EqualFold(l, prev)) || (!ignoreCase && l == prev)) {
			out = append(out, l)
		}
		first = false
		prev = l
	}
	return out
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
