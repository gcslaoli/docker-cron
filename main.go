package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	cron "github.com/robfig/cron/v3"
)

func main() {
	// 如果带了version参数
	args := os.Args
	var (
		VERSION string = "1.0.1"
	)
	if len(args) >= 2 && args[1] == "version" {
		fmt.Println("docker-cron:一个go实现的任务管理器")
		fmt.Println("Version:" + VERSION)

		os.Exit(1)
	}
	crontab := cron.New()
	// task := func() {
	// 	fmt.Println("hello world")
	// }
	// 添加定时任务, * * * * * 是 crontab,表示每分钟执行一次
	// crontab.AddFunc("* * * * *", task)
	crontab.AddFunc("* * * * *", shellMinutely)
	crontab.AddFunc("0 * * * *", shellHourly)
	crontab.AddFunc("0 0 * * *", shellDaily)

	// 启动定时器
	crontab.Start()
	// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	// 根据实际情况进行控制
	select {}
}
func shell(script string) {
	cmd := exec.Command("sh", script)

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:", err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n", string(output))
}

// 每分钟执行
func shellMinutely() {
	// 按目录增加定时任务
	files, err := ListDir("./minutely", ".sh")
	fmt.Println(files, err)
	for i := 0; i < len(files); i++ {
		cmd := exec.Command("sh", files[i])
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("%s Execute Shell:%s failed with error:%s\n", time.Now().String(), files[i], err.Error())
			return
		}
		fmt.Printf("%s Execute Shell:%s finished with output:%s\n", time.Now().String(), files[i], string(output))
	}
}

// 每小时执行
func shellHourly() {
	// 按目录增加定时任务
	files, err := ListDir("./hourly", ".sh")
	fmt.Println(files, err)
	for i := 0; i < len(files); i++ {
		cmd := exec.Command("sh", files[i])
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("%s Execute Shell:%s failed with error:%s\n", time.Now().String(), files[i], err.Error())
			return
		}
		fmt.Printf("%s Execute Shell:%s finished with output:%s\n", time.Now().String(), files[i], string(output))
	}
}

// 每天执行
func shellDaily() {
	// 按目录增加定时任务
	files, err := ListDir("./daily", ".sh")
	fmt.Println(files, err)
	for i := 0; i < len(files); i++ {
		cmd := exec.Command("sh", files[i])
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("%s Execute Shell:%s failed with error:%s\n", time.Now().String(), files[i], err.Error())
			return
		}
		fmt.Printf("%s Execute Shell:%s finished with output:%s\n", time.Now().String(), files[i], string(output))
	}
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}
