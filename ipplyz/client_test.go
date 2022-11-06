/**
 * @Author:      leafney
 * @Date:        2022-11-06 11:31
 * @Project:     rose-ip2region
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package ipplyz

import "testing"

func TestNewIpPlyz(t *testing.T) {
	province, ok, err := NewIpPlyz().
		SetTimeout(2000).
		SetDebug(true).
		Parse("61.132.188.153")
	//Parse("124.171.74.153")
	if ok {
		t.Logf("province: [%v]", province)
	} else {
		t.Errorf("请求失败 [%v]", err)
	}
}