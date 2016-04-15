# goconf
一个简单的配置文件读取程序

#安装

	go get github.com/linchengzhi/goconf

#使用

>配置文件

	#[config] 如果没有前缀，则默认为config

	[user]
	account = root
	password = mima

>初始化

	conf := goconf.New("conf/config.conf") //New(path) path为配置文件路径

>获取单个配置

	val := goconf.Get("user","account") //Get(prefix,tag) prefix为配置前缀，tag为标识符；如果没有前缀，则默认为config

>获取所有配置

	list := goconf.Read() //Read()


---
