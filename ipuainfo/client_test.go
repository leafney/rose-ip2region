/**
 * @Author:      leafney
 * @Date:        2022-11-06 15:12
 * @Project:     rose-ip2region
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package ipuainfo

import "testing"

func TestNewIpUAInfo(t *testing.T) {
	province, ok, err := NewIpUAInfo().
		SetTimeout(200).
		SetDebug(true).
		//Parse("61.132.188.153")
		Parse("124.171.74.153")
	if ok {
		t.Logf("province: [%v]", province)
	} else {
		t.Errorf("请求失败 [%v]", err)
	}
}