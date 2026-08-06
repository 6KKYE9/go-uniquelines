package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// 去重相邻重复行（像 uniq），或全局去重只保留首次出现的行
func unique(lines []string, global bool) []string {
	seen := map[string]bool{}
	var out []string
	prev := ""
	for _, l := range lines {
		if global {
			if seen[l] {
				continue
			}
			seen[l] = true
			out = append(out, l)
			continue
		}
		if l == prev {
			continue
		}
		prev = l
		out = append(out, l)
	}
	return out
}

func main() {
	global := flag.Bool("g", false, "全局去重，而非仅合并相邻重复行")
	flag.Parse()

	var lines []string
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	for _, l := range unique(lines, *global) {
		fmt.Println(l)
	}
}
