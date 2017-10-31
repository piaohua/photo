package tests

import (
	"io/ioutil"
	"testing"

	beetest "github.com/astaxie/beego/testing"
)

func TestUrl(t *testing.T) {
	request := beetest.Post("/domainurl")
	request.Param("url", "http://wx.qlogo.cn/mmopen/...")
	response, _ := request.Response()
	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	t.Logf("contents %v", contents)
}

//http://wx.qlogo.cn/mmopen/vi_32/EnctdgHtll59eeJIRuh56lA0MZfmMtzwfQUPyJbpX5neicxVEZNq2nj3ibp9dIMUhXIu06N3EnKLbPibmcEiabZycg/0
//http://nnyl.iy00.cn/mmopen/?photo=vi_32/EnctdgHtll59eeJIRuh56lA0MZfmMtzwfQUPyJbpX5neicxVEZNq2nj3ibp9dIMUhXIu06N3EnKLbPibmcEiabZycg/0
func TestPhoto(t *testing.T) {
	request := beetest.Get("/mmopen")
	request.Param("photo", "vi_32/EnctdgHtll59eeJIRuh56lA0MZfmMtzwfQUPyJbpX5neicxVEZNq2nj3ibp9dIMUhXIu06N3EnKLbPibmcEiabZycg/0")
	response, _ := request.Response()
	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	t.Logf("contents %v", contents)
}
