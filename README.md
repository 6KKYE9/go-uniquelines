去掉重复行。默认只合并相邻的重复行（像 uniq），加 -g 就全局去重只留首次出现的。

用法：
  cat 日志.txt | go-uniquelines
  cat 日志.txt | go-uniquelines -g