/**
 * @Author:      leafney
 * @Date:        2022-11-06 10:35
 * @Project:     rose-ip2region
 * @HomePage:    https://github.com/leafney
 * @Description: [IP-API.com - Geolocation API - Documentation - JSON](https://ip-api.com/docs/api:json)
 */

package ipapi

import (
	"fmt"
	"github.com/leafney/rose"
	"github.com/leafney/rose/reqx"
	"github.com/tidwall/gjson"
	"log"
	"strings"
	"time"
)

type IpApi struct {
	host    string
	timeout int64
	lang    string
	debug   bool
}

func NewIpApi() *IpApi {
	return &IpApi{
		host:  "http://ip-api.com/json/{ip}",
		lang:  "zh-CN",
		debug: false,
	}
}

func (c *IpApi) Parse(ip string) (province string, ok bool, err error) {
	ok = false
	resp, err := reqx.
		Get(c.host).
		SetDebug(c.debug).
		SetPathParam("ip", ip).
		SetQueryParam("lang", c.lang).
		SetTimeout(time.Duration(c.timeout) * time.Millisecond).
		Do()

	if err != nil {
		log.Printf("request error [%v]", err)
		return
	}
	if resp.IsSuccess() {
		body := resp.String()
		status := gjson.Get(body, "status").String()
		if strings.EqualFold(status, "success") {
			ok = true

			// countryCode
			if strings.EqualFold(gjson.Get(body, "countryCode").String(), "CN") {
				// 国内地区 regionName
				province = gjson.Get(body, "regionName").String()
			} else {
				// 国外地区 country
				province = gjson.Get(body, "country").String()
			}
		}
	} else {
		err = fmt.Errorf("error of statusCode [%v] body [%v]", resp.StatusCode(), resp.String())
	}

	return province, ok, err
}

func (c *IpApi) SetLang(lang string) *IpApi {
	if !rose.StrIsEmpty(lang){
		c.lang = lang
	}
	return c
}

// SetTimeout timeout of Millisecond
func (c *IpApi) SetTimeout(t int64) *IpApi {
	c.timeout = t
	return c
}

func (c *IpApi) SetDebug(d bool) *IpApi {
	c.debug = d
	return c
}