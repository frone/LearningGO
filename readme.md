# GO语言学习笔记

## GO的优点

- Go代码是可以移植的，特别是类Unix系统之间；
- Go已经支持**面向过程、并发和分布式编程**；
- Go支持**垃圾回收机制**，所以开发者无需关心内存分配与释放的问题；
- Go没有预处理器，所以**编译速度很快**。因此，可以作为脚本语言使用；
- Go可以构建Web应用，提供简单的Web服务器，供测试使用；
- Go提供很多标准的库和包，简化开发者的开发工作。这些标准库和包经过大量的测试，可以安全使用；
- Go默认启用静态链接，二进制文件很容易在同种操作系统间移植。一旦Go代码编译后，无需关心它所依赖的其他库，就可以很方便的在其他地方执行该二进制文件；
- Go提供命令行式编译命令，无需GUI操作，深受Unix开发者的青睐；
- Go**支持Unicode编码**，意味着不需要编写很多代码适配多种语言；
- go的多个特性之间都是正交的，保证语言的稳定性和简单性。

## GO的缺点

- Go不直接支持面向对象编程，这就使得那些已经习惯于面向对象编程的开发者来说非常痛苦。然而，你可以用组合的方式去模拟面向对象的实现方式。

  Is Go an object-oriented language？
  Yes and no. Although Go has types and methods and allows an object oriented style of programming, there is no type hierarchy. The concept of"interface"in Go provides a different approach that we believe is easy to use and in some ways more general Also, the lack of a type hierarchy makes "objects"in Go feel much more lightweight than in languages such as C++ or Java

## 注意点

- 不允许隐式类型转换
- 切片不支持负数索引，不能比较
- 没有内置Set
- String是不可变的byte slice
- unicode是一种字符集，utf-8是Unicode的存储实现

## 函数是一等公民

1. 可以有多个返回值
2. 所有参数都是值传递：slice，map，channel会有传引用的错觉
3. 函数可以作为变量的值
4. 函数可以作为参数和返回值

## GO接口

1. 接口为非侵入性，实现不依赖于接口定义
2. 所以接口的定义可以包含在接口使用者包内

## GO错误机制

1. 没有异常机制
2. error类型实现了error接口
3. 可以通过errors.New来快速创建错误实例

### PANIC

1. panic用于不可以恢复的错误
2. panic退出前会执行defer指定的内容

### OS.Exit

1. 退出时不会调用defer指定的函数
2. 退出时不输出当前调用栈信息

## VS CODE开发配置

默认运行 go test 不会输出 testing.T.Log() 的内容。

要显示这些内容，需要加上开关 -v

```
go test -v -timeout 30s xxx/xxx/package -run ^TestXXXFunction$
```

在 Visual Studio Code IDE 环境中，可以设置 Workspace Settings。打开 .vscode/settings.json，添加：

```
"go.testFlags": ["-v"]
```

这样，在 IDE 编辑器中，点击函数上方的 run test，自动运行 go test，会被加上 -v 标志，在 OUTPUT 窗口就可以看到 t.Logf("xxx%s","xxx") 的输出内容了。

