package goconf

import (
	"os"
	"bufio"
	"strings"
	"io"
)

type Conf struct {
	path string
	list map[string]map[string]string
}

func New(path string)(*Conf){
	conf :=new(Conf)
	if path==""{
		path = "conf/config.conf"
	}
	conf.path=path
	return conf
}

func (c *Conf)Get(prefix,tag string)(string){
	str :=c.Read()
	if str!=""{
		return "值为空"
	}
	for k, v := range c.list {
		if k==prefix{
			for key, val := range v {
				if key == tag {
					return val
				}
			}
		}
	}
	return "值为空"
}

func (c *Conf)Read()(string){

	var prefix = "config"
	var data = make(map[string]map[string]string)
	data[prefix] = make(map[string]string)

	file,err := os.Open(c.path)
	if err!=nil{
		return "打开配置文件失败"
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	for{
		row,_,err :=buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return "读取文件出错"
			}else{
				break
			}
		}
		line := strings.TrimSpace(string(row))
		switch {
		case len(line) == 0:
		case line[0] == '[' && line[len(line)-1] == ']':
			prefix = strings.TrimSpace(line[1 : len(line)-1])
			data[prefix] = make(map[string]string)
		default:
			arr :=strings.Split(line,"=")
			if len(arr)==2{
				key :=strings.TrimSpace(arr[0])
				value :=strings.TrimSpace(arr[1])
				data[prefix][key] = value
			}
		}
	}
	if len(data["config"])==0{
		delete(data,"config")
	}
	c.list=data
	return ""
}

