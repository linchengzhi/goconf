package goconf

import (
	"testing"
)


func TestGet(t *testing.T){
	conf :=New("conf/config.conf")
	str:=conf.Get("user","account")
	if str!="root"{
		t.Errorf("测试Get函数失败", str)
	}
}

func TestRead(t *testing.T){
	conf :=New("conf/config.conf")
	str:=conf.Read()
	if str!=""{
		t.Errorf("测试Read函数失败", str)
	}
}

func Benchmark_Get(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		conf :=New("conf/config.conf")
		conf.Read()
	}
}