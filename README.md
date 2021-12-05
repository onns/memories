# Hopes Memories Last Forever

记录一些特殊的时刻，方便订阅。

当然，也方便移除。

## 使用方式

```bash
$ ./mem -h                                                                          [2021/12/05 16:30:09] 
Usage of ./mem:
  -a	absolute file path (default true)
  -c string
    	config file name (default "config.json")
  -o string
    	output file name with (.ics) (default "special-day")
```
1. 修改`config.json`，运行（保证config.json和mem在相同文件夹）。
2. 制定配置文件和输出名：
```bash
./mem -c other.json -o "taylor"
./mem -o "taylor" -c /Users/onns/Downloads/github/memories/other.json -a
```

## TODO

- [x] 生日
- [x] 农历生日
- [x] 纪念日