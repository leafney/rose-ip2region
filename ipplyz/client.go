/**
 * @Author:      leafney
 * @Date:        2022-11-06 11:31
 * @Project:     rose-ip2region
 * @HomePage:    https://github.com/leafney
 * @Description: [免费ip地址查询 - 免费ip归属地查询接口](http://ip.plyz.net/)
 */

package ipplyz

import (
	"fmt"
	"github.com/leafney/rose"
	"github.com/leafney/rose/reqx"
	"log"
	"strings"
	"time"
)

type IpPlyz struct {
	host    string
	timeout int64
	debug   bool
}

func NewIpPlyz() *IpPlyz {
	return &IpPlyz{
		host:  "http://ip.plyz.net/ip.ashx",
		debug: false,
	}
}

func (c *IpPlyz) Parse(ip string) (province string, ok bool, err error) {
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

		addr := strings.Split(body, "|")
		if len(addr) == 2 {
			addrName := addr[1]
			if !rose.StrIsEmpty(addrName) {
				addrNameS := strings.Split(addrName, " ")
				addrNameL := len(addrNameS)
				if addrNameL == 4 {
					ok = true
					province = addrNameS[1]
				} else if addrNameL == 3 {
					ok = true
					province = addrNameS[0]
				}
			}
		}
	} else {
		err = fmt.Errorf("error of statusCode [%v] body [%v]", resp.StatusCode(), resp.String())
	}

	return province, ok, err
}

// SetTimeout timeout of Millisecond
func (c *IpPlyz) SetTimeout(t int64) *IpPlyz {
	c.timeout = t
	return c
}

func (c *IpPlyz) SetDebug(d bool) *IpPlyz {
	c.debug = d
	return c
}