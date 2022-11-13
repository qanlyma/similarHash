## similarHash

该程序 main.go 可以读取图片、文本文件并生成哈希值（uint64），并可以与已存在的哈希值比较得出是否存在相似文件。

其中图片支持.png .jpg；文本支持 .txt

* dict 文件存放生成文本哈希时使用的词典

* distance 文件存放计算哈希值汉明距离的函数

* readwrite 文件用于读写存放哈希值的文件

* testPic 存放测试用的图片

* dictionary.txt 是测试用的文本

* judge.go 用于判断文本和图片是否存在相似哈希值，若不存在则插入新文件的哈希值到集合里

* phash.txt 存放图片生成哈希的集合

* simHash.go 用于将图片和文本转换为哈希值

* whash.txt 存放文本生成哈希的集合
