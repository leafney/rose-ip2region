/**
 * @Author:      leafney
 * @Date:        2022-11-06 15:18
 * @Project:     rose-ip2region
 * @HomePage:    https://github.com/leafney
 * @Description: [VORE-API - 提供免费接口服务](https://api.vore.top/)
 */

package ipvore

import (
	"errors"
	"fmt"
	"github.com/leafney/rose/reqx"
	"github.com/tidwall/gjson"
	"log"
	"time"
)

type IpVore struct {
	host    string
	timeout int64
	debug   bool
}

// 注意：该网站请求超市时间一般需要大于2s以上
func NewIpVore() *IpVore {
	return &IpVore{
		host:    "https://api.vore.top/api/IPdata",
		debug:   false,
		timeout: 3000,
	}
}

func (c *IpVore) Parse(ip string) (province string, ok bool, err error) {
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
			province = gjson.Get(body, "ipdata").Get("info1").String()
		} else {
			msg := gjson.Get(body, "msg").String()
			err = errors.New(msg)
		}
	} else {
		err = fmt.Errorf("error of statusCode [%v] body [%v]", resp.StatusCode(), resp.String())
	}

	return province, ok, err
}

// SetTimeout timeout of Millisecond
func (c *IpVore) SetTimeout(t int64) *IpVore {
	if t > 0 {
		c.timeout = t
	}
	return c
}

func (c *IpVore) SetDebug(d bool) *IpVore {
	c.debug = d
	return c
}
