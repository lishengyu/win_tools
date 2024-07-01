# MagicBox

图形库
    go get github.com/lxn/walk

golang的图形exe需要依赖于manifest 清单文件才能正常运行。要把manifest嵌入exe文件中
需要用到rsrc，manifest文件内容固定，只需运行一次
只有修改test.mainfest或者图标文件才需要重新生成一次
    go get github.com/akavel/rsrc
    rsrc -manifest test.manifest -o rsrc.syso

打包
    go build -ldflags="-H windowsgui"

编译系统
    go env -w GOOS=windows
