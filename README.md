### LeetCode 题目下载助手
爬取所有的题目，保存在每一个文件夹下的README.md中，并且保存对应的代码模板

<img src="demo.gif" width=400>

运行之后会在对应的目录下写入题目描述和对应的代码模板
```
leetcode/
├── LCP 06       
│   ├── README.md
│   └── main.go  
├── LCP 07       
│   ├── README.md
│   └── main.go  
├── LCP 08       
│   ├── README.md
│   └── main.go  
├── LCP 09       
│   ├── README.md
│   └── main.go  
.......
```

使用方式
```
Usage:
  -d string
        directory to store all problems (default "leetcode")
  -lang string
        description language, support [zh, en] (default "zh")
  -r int
        get question detail retry times (default 5)
  -s uint
        sleep time to avoid http code 429 (default 5)
  -t string
        which program language to solve, [C, C++, Python, Python3, Java, etc] (default "Go")    
  -w string
        write log to file
  -a 
		if question directory exists, not download again
```
