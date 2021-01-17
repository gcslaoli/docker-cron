## 前言

一个给 docker 镜像提供计划任务的小程序
[GitHub](https://github.com/GCSLaoLi/docker-cron)

因交叉编译稳定性问题，暂时只编译mac和linux两个版本(其实以设计用途来说有linux版本就够了)
## 编译命令

```
sh build.sh v1.0.2
```

## 使用说明

程序运行后将从程序所在目录下查找配置目录

```
minutely  目录内的sh文件每分钟被执行 * * * * *
hourly  目录内的文件每小时被执行 0 * * * *
daily 目录内的文件每天被执行 0 0 * * *
weekly 目录内的文件每周被执行 0 0 * * 0
monthly 目录内的文件每月被执行 0 0 1 * *
quarterly 目录内的文件每季度被执行 0 0 1 1,4,7,10 *
yearly 目录内的文件每年被执行 0 0 1 1 *
```

上述目录中以`.sh`结尾的文件将按周期以 `sh xxx.sh`的方式执行

下载地址:
https://github.com/GCSLaoLi/docker-cron/releases/

推荐一个 crontab 表达式测试地址:

https://www.box3.cn/page/crontab.html
