# easy-func

## 项目介绍
使用 golang 来翻译 php 函数。可以看做php函数在golang的映射字典。

## 项目要求

- 单元测试
    - 单测用例需参考 [https://github.com/php/php-src/tree/master/ext/standard/tests/strings](https://github.com/php/php-src/tree/master/ext/standard/tests/strings)
    - 单测覆盖率：go-carpet 或者 go-carpet --summary 保障单元测试覆盖率
    - 单测规范：goland -> general -> test for function 来生成单元测试用例代码
    - **单测覆盖率要求：100%**
- 圈复杂度
    - 在根目录使用命令 gocyclo -over 10 ./ 检查圈复杂度
    - **保证所有复杂度不超过 10**

## 圈复杂度：

- 安装参考：https://github.com/fzipp/gocyclo
- 示例命令：

    `gocyclo -over 10 ./` #查看圈复杂度超过 10 的 function 
    
## 单元测试：

生成测试用例：goland 在方法名处右键，Generate, Tests for package. 根据生成的代码完善用例
示例命令：(在文件所在的文件夹下)：

 - 方法一：`go test` 
 - 方法二：`go test --cover` 
 - 方法三：在相应的 _test.go 文件下，用例前面有一个三角形，直接点击运行

## 单元测试覆盖率
https://github.com/msoap/go-carpet

