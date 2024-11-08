package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"log"
	"math/rand/v2"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
)

type vpn struct {
	protol         string
	password       string
	server_address string
	server_port    string
	sni            string
}

var (
	stopFlag    = make(chan bool)
	reStartFlag = make(chan bool)

	option_server = []vpn{
		{protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.xgzl.fogvip-zz.uk", server_port: "50069", sni: "ld.xgzl02.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.xgzl.fogvip-zz.uk", server_port: "50069", sni: "ld.xgzl02.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.xgzl.fogvip-zz.uk", server_port: "50069", sni: "ld.xgzl02.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.xgzl.fogvip-zz.uk", server_port: "50069", sni: "ld.xgzl02.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.xg01.fogvip-zz.uk", server_port: "40001", sni: "ld.xgg01.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.xgzl.fogvip-zz.uk", server_port: "50069", sni: "ld.xgzl02.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.xgb.fogvip-zz.uk", server_port: "40024", sni: "ld.xghk03.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.xgb.fogvip-zz.uk", server_port: "40025", sni: "ld.xghk03.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.xgys.fogvip-zz.uk", server_port: "40052", sni: "ld.xgys01.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.hg01.fogvip-zz.uk", server_port: "40021", sni: "ld.hg01.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.jp01.fogvip-zz.uk", server_port: "40014", sni: "ld.jp01.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.tw01.fogvip-zz.uk", server_port: "40036", sni: "ld.tw01.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.xinjp01.fogvip-zz.uk", server_port: "40013", sni: "ld.xinjp01.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.mgxyt.fogvip-zz.uk", server_port: "40005", sni: "ld.mgxyt.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.mgxyt.fogvip-zz.uk", server_port: "40006", sni: "ld.mgxyt.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.usa02.fogvip-zz.uk", server_port: "40044", sni: "ld.meiguo02.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.hgkneko.viptap-tcb-zz.cc", server_port: "50111", sni: "ld.hkgneko.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.hgkneko.viptap-tcb-zz.cc", server_port: "50112", sni: "ld.hkgneko.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.2hkmix.fogvip-zz.uk", server_port: "50306", sni: "ld.2heix.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.2pkshk.fogvip-zz.uk", server_port: "50307", sni: "ld.2heix.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.xgbig.viptap-tcb-zz.cc", server_port: "50070", sni: "ld.xgbig.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.hg01.fogvip-zz.uk", server_port: "40021", sni: "ld.hg01.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.tw01.fogvip-zz.uk", server_port: "40037", sni: "ld.tw01.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.usa01.fogvip-zz.uk", server_port: "50072", sni: "ld.usa01.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.usa02.fogvip-zz.uk", server_port: "40045", sni: "ld.meiguo02.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.mglsj.fogvip-zz.uk", server_port: "50108", sni: "ld.mglsj.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.ysls.fogvip-zz.uk", server_port: "50308", sni: "ld.vegas.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.tgcloud.fogvip-zz.uk", server_port: "50055", sni: "ld.tgcloud.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.mlxy-1.fogvip-zz.uk", server_port: "50117", sni: "ld.mlxy-cn.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.jp01.fogvip-zz.uk", server_port: "40015", sni: "ld.jp01.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.jp03.fogvip-zz.uk", server_port: "50056", sni: "ld.rbv2.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.xinjp02.fogvip-zz.uk", server_port: "40017", sni: "ld.xinjp01.yuiuingkig.icu"}, {protol: "trojan", password: "fa737811-14ea-40b2-b07e-c9b0427b05a9", server_address: "zz.sghkg.fogvip-zz.uk", server_port: "50311", sni: "ld.xjpsghkg.yuiuingkig.icu"},
	}
)

type configStruct struct {
	Policy struct {
		System struct {
			StatsOutboundUplink   bool `json:"statsOutboundUplink"`
			StatsOutboundDownlink bool `json:"statsOutboundDownlink"`
		} `json:"system"`
	} `json:"policy"`
	Log struct {
		Access   string `json:"access"`
		Error    string `json:"error"`
		Loglevel string `json:"loglevel"`
	} `json:"log"`
	Inbounds []struct {
		Tag      string `json:"tag"`
		Port     int    `json:"port"`
		Listen   string `json:"listen"`
		Protocol string `json:"protocol"`
		Sniffing struct {
			Enabled      bool     `json:"enabled"`
			DestOverride []string `json:"destOverride"`
		} `json:"sniffing,omitempty"`
		Settings struct {
			Auth             string `json:"auth"`
			UDP              bool   `json:"udp"`
			AllowTransparent bool   `json:"allowTransparent"`
		} `json:"settings,omitempty"`
		Settings0 struct {
			UDP              bool `json:"udp"`
			AllowTransparent bool `json:"allowTransparent"`
		} `json:"settings,omitempty"`
		Settings1 struct {
			UDP              bool   `json:"udp"`
			Address          string `json:"address"`
			AllowTransparent bool   `json:"allowTransparent"`
		} `json:"settings,omitempty"`
	} `json:"inbounds"`
	Outbounds []struct {
		Tag      string `json:"tag"`
		Protocol string `json:"protocol"`
		Settings struct {
			Servers []struct {
				Address  string `json:"address"`
				Method   string `json:"method"`
				Ota      bool   `json:"ota"`
				Password string `json:password`
				Port     int    `json:"port"`
				Level    int    `json:"level"`
			} `json:"servers"`
		} `json:"settings,omitempty"`
		StreamSettings struct {
			Network     string `json:"network"`
			Security    string `json:"security"`
			TLSSettings struct {
				AllowInsecure bool   `json:"allowInsecure"`
				ServerName    string `json:"serverName"`
			} `json:"tlsSettings"`
		} `json:"streamSettings,omitempty"`
		Mux struct {
			Enabled     bool `json:"enabled"`
			Concurrency int  `json:"concurrency"`
		} `json:"mux,omitempty"`
		Settings0 struct {
		} `json:"settings,omitempty"`
		Settings1 struct {
			Response struct {
				Type string `json:"type"`
			} `json:"response"`
		} `json:"settings,omitempty"`
	} `json:"outbounds"`
	Stats struct {
	} `json:"stats"`
	API struct {
		Tag      string   `json:"tag"`
		Listen   string   `json:"listen"`
		Services []string `json:"services"`
	} `json:"api"`
	Routing struct {
		DomainStrategy string `json:"domainStrategy"`
		Rules          []struct {
			Type        string   `json:"type"`
			InboundTag  []string `json:"inboundTag"`
			OutboundTag string   `json:"outboundTag"`
			Port        string   `json:"port,omitempty"`
			Domain      []string `json:"domain,omitempty"`
			IP          []string `json:"ip,omitempty"`
		} `json:"rules"`
	} `json:"routing"`
}

func interfaceServer() {

	router := gin.Default()
	router.GET("/flag", func(c *gin.Context) {

		stopFlag <- true

		c.JSON(200, gin.H{"succ": true})
	})

	router.Run("0.0.0.0:8081")
}

func run() {
	server, err := startXray()
	if err != nil {
		fmt.Println("Failed to start:", err)
		// Configuration error. Exit with a special value to prevent systemd from restarting.
		os.Exit(23)
	}

	if *test {
		fmt.Println("Configuration OK.")
		os.Exit(0)
	}

	if err := server.Start(); err != nil {
		fmt.Println("Failed to start:", err)
		os.Exit(-1)
	}
	defer func() {
		server.Close()

		reStartFlag <- true
	}()

	runtime.GC()
	debug.FreeOSMemory()

	<-stopFlag
}

func handleFlag() {

	for {
		select {
		case <-reStartFlag:
			if err := modifyConfigFile(); err == nil {
				go run()
			}
		default:
		}
	}
}

func modifyConfigFile() error {
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println("当前目录配置文件：", filepath.Join(workingDir, "config.json"))
	file, err := os.ReadFile(filepath.Join(workingDir, "config.json"))
	if err != nil {
		log.Fatal(err)
	}

	string_file := string(file)

	protol := gjson.Get(string_file, "outbounds.0.protocol").String()
	address := gjson.Get(string_file, "outbounds.0.settings.servers.0.address").String()
	password := gjson.Get(string_file, "outbounds.0.settings.servers.0.password").String()
	port := gjson.Get(string_file, "outbounds.0.settings.servers.0.port").String()
	sni := gjson.Get(string_file, "outbounds.0.streamSettings.tlsSettings.serverName").String()

	// 选取随机值
	index := rand.IntN(len(option_server))
	randomValue := option_server[index]
	fmt.Println("当前选中服务：", randomValue)

	string_file = strings.Replace(string_file, protol, randomValue.protol, 1)
	string_file = strings.Replace(string_file, address, randomValue.server_address, 1)
	string_file = strings.Replace(string_file, password, randomValue.password, 1)
	string_file = strings.Replace(string_file, port, randomValue.server_port, 1)
	string_file = strings.Replace(string_file, sni, randomValue.sni, 1)

	err = os.WriteFile("config.json", []byte(string_file), 0666)

	return err
}

func timerRun() {

	go func() {
		handleFlag()
	}()

	go func() {
		run()
	}()

	interfaceServer()
}
