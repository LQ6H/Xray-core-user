```
2024/11/12 15:18:00 Using default config:  E:\GolangProject\src\Xray-core-user\config.json
2024/11/12 15:18:00 [Info] infra/conf/serial: Reading config: &{Name:E:\GolangProject\src\Xray-core-user\config.json Format:json}
2024/11/12 15:18:00 [Debug] app/log: Logger started
2024/11/12 15:18:00 [Info] app/dns: DNS: created UDP client initialized for 223.5.5.5:53
2024/11/12 15:18:00 [Info] app/dns: DNS: created UDP client initialized for 1.1.1.1:53
2024/11/12 15:18:00 [Info] app/dns: DNS: created UDP client initialized for 8.8.8.8:53
2024/11/12 15:18:00 [Info] app/dns: DNS: created Remote DOH client for https://dns.google/dns-query
2024/11/12 15:18:00 [Info] transport/internet/udp: listening UDP on 0.0.0.0:10085
2024/11/12 15:18:00 [Info] transport/internet/tcp: listening TCP on 0.0.0.0:10086
2024/11/12 15:18:00 [Warning] core: Xray 24.10.16 started

2024/11/12 15:18:07 [Info] [2419787188] proxy/socks: TCP Connect request to tcp:142.250.196.100:443
2024/11/12 15:18:07 [Info] [2419787188] app/dispatcher: sniffed domain: www.google.com
2024/11/12 15:18:07 [Info] [2419787188] app/dispatcher: taking detour [proxy12] for [tcp:www.google.com:443]
2024/11/12 15:18:07 [Info] [2419787188] transport/internet/tcp: dialing TCP to tcp:zz.xinjp01.fogvip-zz.uk:40013
2024/11/12 15:18:07 [Debug] [2419787188] transport/internet: dialing to tcp:zz.xinjp01.fogvip-zz.uk:40013
2024/11/12 15:18:07 from tcp:192.168.31.22:65258 accepted tcp:142.250.196.100:443 [socks -> proxy12]
2024/11/12 15:18:07 [Info] [2419787188] proxy/trojan: tunneling request to tcp:www.google.com:443 via zz.xinjp01.fogvip-zz.uk:40013
2024/11/12 15:18:08 [Info] [2419787188] app/proxyman/inbound: connection ends > proxy/socks: connection ends > context canceled


程序启动流程：

    1. main/run.go/getConfigFilePath   
    2. infra/conf/serial/builder.go/mergeConfigs
    3. app/log/log.go/New
    4. app/dns/nameserver_udp.go/NewClassicNameServer
    5. app/dns/nameserver_doh.go/NewDoHNameServer
    6. transport/internet/udp/hub.go/ListenUDP
    7. transport/internet/tcp/hub.go/ListenTCP
    8. core/xray.go/Start


单个`socks5` 请求处理流程：

    1. proxy/socks/server.go/Process
    2. proxy/socks/server.go/processTCP
    3. app/dispatcher/default.go/Dispatch
    4. app/dispatcher/default.go/routedDispatch
    5. transport/internet/tcp/dialer.go/Dial
    6. transport/internet/system_dialer.go/Dial
    7. ******
    8. ******
    9. app/proxyman/inbound/worker.go/callback
```





整个配置内容如下：

1. `api`

```go
type APIConfig struct {
	Tag      string   `json:"tag"`
	Listen   string   `json:"listen"`
	Services []string `json:"services"`
}
```

```
reflectionservice, handlerservice, loggerservice, statsservice,
observatoryservice, routingservice
```



2. `log` 

```go
type LogConfig struct {
	AccessLog   string `json:"access"`
	ErrorLog    string `json:"error"`
	LogLevel    string `json:"loglevel"`
	DNSLog      bool   `json:"dnsLog"`
	MaskAddress string `json:"maskAddress"`
}
```



3. `routing` 

```go
type RouterConfig struct {
    RuleList       []json.RawMessage `json:"rules"`
    DomainStrategy *string           `json:"domainStrategy"`
    Balancers      []*BalancingRule  `json:"balancers"`

    DomainMatcher string `json:"domainMatcher"`
}
```



4. `dns` 

```go
type DNSConfig struct {
    Servers                []*NameServerConfig `json:"servers"`
    Hosts                  *HostsWrapper       `json:"hosts"`
    ClientIP               *Address            `json:"clientIp"`
    Tag                    string              `json:"tag"`
    QueryStrategy          string              `json:"queryStrategy"`
    DisableCache           bool                `json:"disableCache"`
    DisableFallback        bool                `json:"disableFallback"`
    DisableFallbackIfMatch bool                `json:"disableFallbackIfMatch"`
}
```



5. `inbounds` 

```go
type InboundDetourConfig struct {
    Protocol       string                         `json:"protocol"`
    PortList       *PortList                      `json:"port"`
    ListenOn       *Address                       `json:"listen"`
    Settings       *json.RawMessage               `json:"settings"`
    Tag            string                         `json:"tag"`
    Allocation     *InboundDetourAllocationConfig `json:"allocate"`
    StreamSetting  *StreamConfig                  `json:"streamSettings"`
    SniffingConfig *SniffingConfig                `json:"sniffing"`
}
```



6. `outbounds` 

```go
type OutboundDetourConfig struct {
    Protocol      string           `json:"protocol"`
    SendThrough   *string          `json:"sendThrough"`
    Tag           string           `json:"tag"`
    Settings      *json.RawMessage `json:"settings"`
    StreamSetting *StreamConfig    `json:"streamSettings"`
    ProxySettings *ProxyConfig     `json:"proxySettings"`
    MuxSettings   *MuxConfig       `json:"mux"`
}
```



一个`socks5` 请求进来的完整流程：

```
1. core/xray.go/Start 中的 s.features中包含配置内容，然后调用其Start方法
2. app/proxyman/inbound/inbound.go/Start 获取taggedHandlers, 调用Start方法
3. app/proxyman/inbound/always.go/Start 获取workers, 然后调用其Start方法
4. app/proxyman/inbound/worker.go/Start 协程调用
5. app/proxyman/inbound/worker.go/callback 
6. proxy/socks/server.go/Process
```

