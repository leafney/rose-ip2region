/**
 * @Author:      leafney
 * @Date:        2022-11-06 15:12
 * @Project:     rose-ip2region
 * @HomePage:    https://github.com/leafney
 * @Description: [我的IP信息 - useragent info](https://ip.useragentinfo.com/)
 */

package ipuainfo

import (
	"fmt"
	"github.com/leafney/rose/reqx"
	"github.com/tidwall/gjson"
	"log"
	"strings"
	"time"
)

type IpUAInfo struct {
	host    string
	timeout int64
	debug   bool
}

func NewIpUAInfo() *IpUAInfo {
	return &IpUAInfo{
		host:    "https://ip.useragentinfo.com/json",
		debug:   false,
		timeout: 3000,
	}
}

func (c *IpUAInfo) Parse(ip string) (province string, ok bool, err error) {
	ok = false
	resp, err := reqx.
		Get(c.host).
		SetDebug(c.debug).
		SetQueryParam("ip", ip).
		SetTimeout(time.Duration(c.timeout) * time.Millisecond).
		Do()

	if err != nil {
		log.Printf("request error [%v]", err)
		return
	}
	if resp.IsSuccess() {
		body := resp.String()
		code := gjson.Get(body, "code").Int()
		if code == 200 {
			ok = true
			if strings.EqualFold(gjson.Get(body, "short_name").String(), "CN") {
				//	国内 province
				province = gjson.Get(body, "province").String()
			} else {
				//	国外 country
				province = gjson.Get(body, "country").String()
			}
		}
	} else {
		err = fmt.Errorf("error of statusCode [%v] body [%v]", resp.StatusCode(), resp.String())
	}

	return province, ok, err
}

// SetTimeout timeout of Millisecond
func (c *IpUAInfo) SetTimeout(t int64) *IpUAInfo {
	if t > 0 {
		c.timeout = t
	}
	return c
}

func (c *IpUAInfo) SetDebug(d bool) *IpUAInfo {
	c.debug = d
	return c
}
