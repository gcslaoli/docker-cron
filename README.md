一个给 docker 镜像提供计划任务的小程序

编译命令

```
docker run --rm -it -v "$PWD":/usr/src/myapp -w /usr/src/myapp gcslaoli/golang-cn sh build.sh
```

程序运行后将从程序所在目录下查找配置目录

```
minutely  目录内的sh文件每分钟被执行
hourly  目录内的文件每小时被执行
daily 目录内的文件每天被执行
weekly 还没写
monthly 还没写
quarterly 还没写
yearly 还没写
```
上述目录中以`.sh`结尾的文件将按周期以 `sh xxx.sh`的方式执行

下载地址:
https://github.com/GCSLaoLi/docker-cron/releases/