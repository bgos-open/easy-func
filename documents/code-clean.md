## 代码整洁
### 圈复杂度：

安装参考：https://github.com/fzipp/gocyclo
示例命令：gocyclo -over 10 ./ #查看圈复杂度超过10的 function 

### 单元测试：
生成测试用例：goland 在方法名处右键，Generate, Tests for package.根据生成的代码完善用例
示例命令：(在文件所在的文件夹下)：
 - 方法一：go test 
 - 方法二：go test --cover 
 - 方法三：在相应的_test.go文件下，用例前面有一个三角形，直接点击运行

### 单元测试覆盖率
https://github.com/msoap/go-carpet
