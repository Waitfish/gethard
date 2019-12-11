package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/imroc/req"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"strconv"
)

const URL  = "https://oapi.dingtalk.com/robot/send"

func genDingPost(content string) []byte  {
	//把消息内容组成阿里云钉钉文字消息接口允许的json 格式
	type Content struct {
		// golang 的 json 解析只解析大写字母开头的内容,所以需要用 json 反射将 json 内容反射为小写字母
		Content string `json:"content"`
		At map[string]bool `json:"at"`
	}
	type DingPost struct {
		MsgType string `json:"msgtype"`
		Text Content `json:"text"`
	}
	dingPost :=DingPost{
		MsgType: "text",
		Text: Content{
			Content: content,
			At:  map[string]bool{
				"isAtAll":true,
			},
		},
	}
	stu,_ :=json.Marshal(dingPost)
	return stu
}

func doPostDing(url, accessToken string,content []byte) *req.Resp {
	// 把消息推送到钉钉机器人
	resp ,err := req.Post(url,req.Param{"access_token": accessToken},req.BodyJSON(content))
	if err !=nil{
	}
	fmt.Println(resp)
	return resp
}
func pushMsg(accessToken,msg string){
	content := genDingPost(msg)
	doPostDing(URL,accessToken,content)

}
func getMem() (usage float64){
	//打印内存信息,并返回内存利用率
	v, _ := mem.VirtualMemory()
	fmt.Printf("Total: %v M, Free:%v M, UsedPercent:%f%%\n  Available:%v", v.Total/1024/1024, v.Available/1024/1024, v.UsedPercent, v.Available/1024/1024)
	return v.UsedPercent

}
func getDisk(path string) (usage float64) {
	//打印硬盘信息,并返回硬盘利用率
	disk, _ := disk.Usage(path)
	fmt.Printf("        HD        : %v GB  Free: %v GB Usage:%f%%\n", disk.Total/1024/1024/1024, disk.Free/1024/1024/1024, disk.UsedPercent)
	usage = disk.UsedPercent
	return usage

}

func main() {
	var Host = flag.String("h", "myServer", "The host name you want to push")
	var Mode = flag.String("m", "disk", "The mode you want to show")
	var Path = flag.String("p", "/", "The disk mount you want to know")
	var Ding = flag.String("d", "dingcode", "The DingDing token")
	var Warn = flag.Float64("w", 90, "UsedPercent to warn dingding")

	flag.Parse()
	//fmt.Println(flag.NFlag())
	if flag.NFlag() < 4 {
		usage := "使用示例:  gethard -m mem|disk  -p /data -w 90 -d theDingCode"
		fmt.Println(usage)
		return
	}
	host := *Host
	mode := *Mode
	path := *Path
	ding := *Ding
	warn := *Warn


	fmt.Println(host,mode,path,ding,warn)

	switch  {
	case mode=="disk":{
		diskUsage := getDisk(path)
		if diskUsage >warn{
			msg := host +"  "+ path + " 硬盘利用率已经超出 " + strconv.FormatFloat(*Warn,'f', -1, 64)+"%"
			pushMsg(ding,msg)
		}
	}
	case mode=="mem":{
		memUsage:= getMem()
		if memUsage >warn{
			msg := host+" 内存利用率已经超出 " + strconv.FormatFloat(*Warn,'f', -1, 64) +"%"
			pushMsg(ding,msg)
		}
	}



	}

}
