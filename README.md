# 

一堆文本要处理，开编辑器点来点去太慢，命令行一把梭最省事。

﻿去掉重复行。默认只合并相邻的重复行（像 uniq），加 -g 就全局去重只留首次出现的。

用法：
  cat 日志.txt | go-uniquelines
  cat 日志.txt | go-uniquelines -g
  cat 日志.txt | go-uniquelines -i           # 忽略大小写
  cat 日志.txt | go-uniquelines -c           # 每行前显示出现次数
  cat 日志.txt | go-uniquelines -d           # 只打印被合并掉的重复行
  cat 日志.txt | go-uniquelines -empty       # 先删掉空行再处理

参数：
  -g        全局去重，只保留每个值第一次出现的位置
  -i        比较时忽略大小写
  -c        在每行前显示它出现的次数
  -d        只打印重复出现的行（第二次及以后出现的）
  -empty    处理前先删掉空白行

测试：
  go test
