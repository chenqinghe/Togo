# Togo

一些使用Go语言实现的php函数，让php程序猿能够无缝切换到Go，方便学习。我们的口号是： let phper easy to Go!
写这个的目的有两个：
* 封装一些常用函数，不需要再自己实现，也不需要import一大堆的包
* 方便php程序员学习Go，在不熟悉或不了解Go的api的情况下可以用熟悉的php函数进行开发。

# 安装
 只需一条简单的命令：
> go get github.com/chenqinghe/Togo

# 简单示例
```Go

 package main
 
 import "fmt"
 import tg "github.com/chenqinghe/Togo"
 
 func main(){
 
   fmt.Println(tg.Md5("helloworld"))//fc5e038d38a57032085441e7fe7010b0
   
 }
 
``` 

# NOTICE
由于Go的语言规定或者为满足Go的编程风格，以下几点是与php不同的地方，使用时请注意：
### 命名
本库中所有函数的命名均遵循大写开头的驼峰命名方式，下划线会被删除

### 数组
数组是php中最方便的东西，但是由于Go是静态类型语言，不能做到像php那样方便。
php的关联数组对应Go中的map[string]interface{}类型
php的索引数组对应Go中的[]interface类型

# 实现函数列表
### 字符串处理
* [explode](https://github.com/chenqinghe/Togo/blob/master/strings.go#L14) — 使用一个字符串分割另一个字符串
* hex2bin — 转换十六进制字符串为二进制字符串
* htmlspecialchars_decode — 将特殊的 HTML 实体转换回普通字符
* htmlspecialchars — 将特殊字符转换为HTML实体
* ltrim — 删除字符串开头的空白字符（或其他字符）
* rtrim — 删除字符串末端的空白字符（或者其他字符）
* str_repeat — 重复一个字符串 
* str_replace — 子字符串替换 
* str_split — 将字符串转换为数组 
* strpos — 查找字符串首次出现的位置 
* strlen — 获取字符串长度 
* strrev — 反转字符串 
* substr — 返回字符串的子串 
* trim — 去除字符串首尾处的空白字符（或者其他字符） 
* ucfirst — 将字符串的首字母转换为大写 
* lcfirst — 使一个字符串的第一个字符小写

### 数组操作
* array_diff- 计算两个数组的差集

### 加密相关操作
* md5_file — 计算指定文件的 MD5 散列值
* md5 — 计算字符串的 MD5 散列值
* sha1_file — 计算文件的 sha1 散列值 
* sha1 — 计算字符串的 sha1 散列值 
* ord — 返回字符的 ASCII 码值

### 时间日期相关
* date
* time
* sleep


### 数据类型转换
* bin2hex
* hex2bin
* boolval
* intval
* strval
