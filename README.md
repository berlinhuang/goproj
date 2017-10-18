# goproj
 the project goproj shows the file structure


### 项目目录
- src 包含项目的源代码文件
- pkg 包含编译后生成的包/库文件
- bin 包含编译后生成的可执行文件

### command
- go build 编译包，如果是main包则在当前目录生成可执行文件，其他包不会生成.a文件
- go install 编译包，同时复制结果到$GOPATH/bin，$GOPATH/pkg等对应目录下
- go run gofiles... 编译列出的文件，并生成可执行文件然后执行。注意只能用于main包，否则会出现go run: cannot run non-main package的错误
- go run是不需要设置$GOPATH的，但go build和go install必须设置。go run常用来测试一些功能，这些代码一般不包含在最终的项目中
