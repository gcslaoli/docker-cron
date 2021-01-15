set -e
# # GOOS：目标平台的操作系统（darwin、freebsd、linux、windows）
# # GOARCH：目标平台的体系架构（386、amd64、arm）
# # 交叉编译不支持 CGO 所以要禁用它，上面的命令编译 64 位可执行程序，同理使用 386 编译 32 位可执行程序
# # 通过docker编译：docker run --rm -it -v "$PWD":/usr/src/myapp -w /usr/src/myapp gcslaoli/golang-cn sh build.sh
# # 关闭CGO
CGO_ENABLED=0

# 编译windows 64位
GOOS=windows
GOARCH=amd64
go build -o ./dist/docker-cron-${GOOS}-${GOARCH}.exe

# 编译mac 64位
GOOS=darwin
GOARCH=amd64
go build -o ./dist/docker-cron-${GOOS}-${GOARCH}

# 编译linux 64位
GOOS=linux
GOARCH=amd64
go build -o ./dist/docker-cron-${GOOS}-${GOARCH}




