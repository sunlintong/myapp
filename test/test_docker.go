package test

import (
	"myapp/modles"
	"github.com/golang/glog"
	"fmt"
	"os"
	"io"
)

//测试modles中的GetImages()
func TestGetImages() {
	fmt.Println("-------------测试GetImages(),本机镜像列表如下：------------")
	images, err := modles.GetImages()
	if err != nil {
		glog.Infoln(err)
	}
	for _, image := range images {
		fmt.Println(image.ID, image.RepoTags)
	}
	fmt.Println("----------------------------------------------------------")
}

//测试modles中的GetContainers()
func TestGetRunningContainers() {
	fmt.Println("-------测试GetContainers(),本机正在运行的容器列表如下：---------")
	containers, err := modles.GetRunningContainers()
	if err != nil {
		glog.Infoln(err)
	}
	for _, container := range containers {
		fmt.Println(container.Image, container.ID, container.Names)
	}
	fmt.Println("--------------------------------------------------------------")
}

//测试modles中的PullImage()
func TestPullImage() {
	fmt.Println("-------测试PullImage()----------------------------------------")
	out, err := modles.PullImage("alpine")
	if err != nil {
		glog.Infoln(err)
	}
	// 打印输出流中信息至stdout
	io.Copy(os.Stdout, out)
	out.Close()
	fmt.Println("---------pull完成------------------------------------------------")
}


// 测试modle中的StartContainer
func TestStartContainer() {
	fmt.Println("-------测试StartContainer()----------------------------------------")
	err := modles.StartContainer("6b65ebe15a1e")
	if err != nil {
		glog.Infoln(err)
	}
	fmt.Println("---------start完成----------------------------------------------")
}

// 测试modle中的StopContainer
func TestStopContainer() {
	fmt.Println("-------测试StopContainer()----------------------------------------")
	err := modles.StopContainer("167ffa211dfa")
	if err != nil {
		glog.Infoln(err)
	}
	fmt.Println("----------------------stop完成-----------------------------------")
}

// 测试StopAllContainers
func TestStopAllContainers() {
	fmt.Println("-----------测试StopAllContainers--------------------------------")
	err := modles.StopAllContainers()
	if err != nil {
		glog.Infoln(err)
	}
	fmt.Println("----------------------stop 完成---------------------------------")
}


