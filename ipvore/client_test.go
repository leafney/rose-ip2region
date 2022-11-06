/**
 * @Author:      leafney
 * @Date:        2022-11-06 15:18
 * @Project:     rose-ip2region
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package ipvore

import "testing"

func TestNewIpVore(t *testing.T) {
	province, ok, err := NewIpVore().
		SetTimeout(2000).
		SetDebug(true).
		Parse("61.132.188.153")
	if ok {
		t.Logf("province: [%v]", province)
	} else {
		t.Errorf("请求失败 [%v]", err)
	}
}