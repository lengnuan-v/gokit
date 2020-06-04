// +----------------------------------------------------------------------
// | http方法
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2020年04月02日
// +----------------------------------------------------------------------

package gohelp

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

// UA组
var userAgentGroup = []string{
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/22.0.1207.1 Safari/537.1",
	"Mozilla/5.0 (X11; CrOS i686 2268.111.0) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.57 Safari/536.11",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6",
	"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1090.0 Safari/536.6",
	"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/19.77.34.5 Safari/537.1",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.9 Safari/536.5",
	"Mozilla/5.0 (Windows NT 6.0) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.36 Safari/536.5",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1062.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1062.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3",
	"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3",
	"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.24 (KHTML, like Gecko) Chrome/19.0.1055.1 Safari/535.24",
	"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/535.24 (KHTML, like Gecko) Chrome/19.0.1055.1 Safari/535.24",
}

// 阿布云代理服务器api
const proxyServer = "http-dyn.abuyun.com:9010"

type AbuyunProxy struct {
	AppID     string
	AppSecret string
}

func (p AbuyunProxy) ProxyClient() *http.Client {
	proxyUrl, _ := url.Parse("http://" + p.AppID + ":" + p.AppSecret + "@" + proxyServer)
	return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
}

// 随机UA
func UA() string {
	rand.Seed(time.Now().Unix())
	return userAgentGroup[rand.Intn(len(userAgentGroup))]
}

// http proxy GET POST
// proxy 指定阿布云通行证(专业版), 不使用填:nil
// method 指定 GET POST
// url 指定网址
// data 请求参数
// header 请求头
func HttpRequest(proxy map[string]string, method string, url string, data []byte, header map[string]string) ([]byte, error) {
	var err error
	var request *http.Request
	if request, err = http.NewRequest(method, url, bytes.NewReader(data)); err != nil {
		return nil, err
	}
	var client *http.Client
	if IsEmpty(proxy) == false && IsEmpty(proxy["appId"]) == false && IsEmpty(proxy["appSecret"]) == false {
		client = AbuyunProxy{AppID: proxy["appId"], AppSecret: proxy["appSecret"]}.ProxyClient()
		request.Header.Set("Proxy-Switch-Ip", "yes")
	} else {
		client = &http.Client{}
	}
	if IsEmpty(header) == false && header != nil {
		for k, v := range header {
			request.Header.Set(k, v)
		}
	}
	if response, err := client.Do(request); err != nil {
		return nil, err
	} else {
		if body, err := ioutil.ReadAll(response.Body); err != nil {
			return nil, err
		} else {
			defer response.Body.Close()
			return body, nil
		}
	}
}
