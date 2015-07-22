package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	info    []net.Addr
	ip      string
	geturl  string
	request *http.Request
	resp    *http.Response
	client  *http.Client
	ipx     string
)

func main() {

	info, _ = net.InterfaceAddrs()
	ip = strings.Split(info[0].String(), "/")[0]
	fmt.Println(ip)
	fmt.Println("alarm will begin in 30 second..")
	fmt.Println("please DO NOT close..")
	i := 30
	for i > 0 {
		time.Sleep(time.Second * 1)
		fmt.Println(i)
		i--
	}

	geturl = "http://" + ip + "/alarmbegin"
	request, _ = http.NewRequest("GET", geturl, nil)
	///not must
	client = &http.Client{}
	resp, _ = client.Do(request)
	fmt.Println(resp)
	fmt.Println("done")
	go Iva()
	go Ivs()
	Ref()
}

func IvsOffLine() {

	info, _ = net.InterfaceAddrs()
	ipx = strings.Split(info[0].String(), "/")[0]
	//fmt.Println(ipx)
	if ipx == "0.0.0.0" {
		fmt.Println("network problem....")
		fmt.Println(strings.Split(info[0].String(), "/")[0])
		for strings.Split(info[0].String(), "/")[0] == "0.0.0.0" {
			time.Sleep(time.Second * 5)
			info, _ = net.InterfaceAddrs()
		}
		fmt.Println("network problem fixed.. alarm will restart in 1 min")
		time.Sleep(time.Second * 30)
		geturl = "http://" + strings.Split(info[0].String(), "/")[0] + "/refresh"
		request, _ = http.NewRequest("GET", geturl, nil)
		resp, _ = client.Do(request)
		fmt.Println(resp)
		time.Sleep(time.Second * 2)
		fmt.Println("alarm begin..")
		geturl = "http://" + strings.Split(info[0].String(), "/")[0] + "/alarmbegin"
		request, _ = http.NewRequest("GET", geturl, nil)
		resp, _ = client.Do(request)
		fmt.Println(resp)
		fmt.Println("done..")
	}

}

func IvaOffLine() {

	defer func() { 
		if err := recover(); err != nil {
			fmt.Println(err) //panic
		}
	}()
	geturl = "http://192.168.10.57/milsng/SVSProxy/"
	request, _ = http.NewRequest("GET", geturl, nil)
	resp, _ = client.Do(request)
	if resp == nil {
		fmt.Println("iva connect fail...try to reconnect")
		for resp == nil {
			time.Sleep(time.Second * 5)
			geturl = "http://192.168.10.57/milsng/SVSProxy/"
			request, _ = http.NewRequest("GET", geturl, nil)
			resp, _ = client.Do(request)
			fmt.Println("iva connect fail...try to reconnect")
		}
		fmt.Println("iva connect refreshing... pls wait")
		info, _ = net.InterfaceAddrs()
		if strings.Split(info[0].String(), "/")[0] != "0.0.0.0" {
			time.Sleep(time.Minute * 10) 
			geturl = "http://" + strings.Split(info[0].String(), "/")[0] + "/refresh"
			request, _ = http.NewRequest("GET", geturl, nil)
			resp, _ = client.Do(request)
			fmt.Println(resp)
			time.Sleep(time.Second * 2)
			fmt.Println("alarm begin..")
			geturl = "http://" + strings.Split(info[0].String(), "/")[0] + "/alarmbegin"
			request, _ := http.NewRequest("GET", geturl, nil)
			resp, _ = client.Do(request)
			fmt.Println(resp)

		}
	}
	if resp != nil {
		resp.Body.Close()
	}

	//fmt.Println("iva connect success")
}

func Refresh() {
	geturl = "http://" + strings.Split(info[0].String(), "/")[0] + "/refresh"
	request, _ = http.NewRequest("GET", geturl, nil)
	resp, _ = client.Do(request)
	fmt.Println(resp)
	time.Sleep(time.Second * 2)
	fmt.Println("alarm begin..")
	geturl = "http://" + strings.Split(info[0].String(), "/")[0] + "/alarmbegin"
	request, _ := http.NewRequest("GET", geturl, nil)
	resp, _ = client.Do(request)
	fmt.Println(resp)
	if resp != nil {
		resp.Body.Close()
	}
}

func Ivs() {
	ivs := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ivs.C:
			IvsOffLine()
		}
	}
}

func Iva() {
	iva := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-iva.C:
			IvaOffLine()
		}
	}
}

func Ref() {
	ref := time.NewTicker(1 * time.Hour)
	for {
		select {
		case <-ref.C:
			Refresh()
		}
	}
}
