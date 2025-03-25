package grdp

import (
	"fmt"
	"github.com/hi-unc1e/grdp/glog"
	"testing"
)

func testrdp(target string) {
	domain := ""
	username := ""
	password := ""
	var err error
	g := NewClient(target, glog.NONE)
	/* 支持默认情况及开启NLA时的认证 */
	//SSL协议登录测试
	err = g.LoginForSSL(domain, username, password)
	if err == nil {
		fmt.Println("SSL Login Success---")
		return
	} else {
		fmt.Println("SSL Login Error:", err)
		fmt.Println("Try RDP Login...")
		//RDP协议登录测试
		err = g.LoginForRDP(domain, username, password)
		if err == nil {
			fmt.Println("RDP Login Success....")
			return
		} else {
			fmt.Println("RDP Login Error:", err)
			return
		}
	}
}

func TestName(t *testing.T) {
	targetArr := []string{
		"127.0.0.1:3389",
	}
	for _, target := range targetArr {
		fmt.Println(target)
		testrdp(target)
	}
}
