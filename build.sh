# 交叉编译不稳定，暂时改用分开打包，我的开发机是mac,mac包就直接打了，linux包通过容器里的linux构建。
# windows版本暂时放弃。
set -e
set -x
CGO_ENABLED=0
GOOS=darwin
GOARCH=amd64
go build -o ./dist/docker-cron-${GOOS}-${GOARCH}-$1

#构建linux版本

docker run --rm -it -v "$PWD":/usr/src/myapp -w /usr/src/myapp gcslaoli/golang-cn sh shell-in-docker.sh $1