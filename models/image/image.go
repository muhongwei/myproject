package image

import (
	"encoding/json"
	"fmt"

	"github.com/golang/glog"

	curl "github.com/mikemintang/go-curl"
)

//Image Image
type Image struct {
	Name string `json:"name"`
	Tag  string `json:"tag,omitempty"`
}

//ListImages ListImages
func ListImages() *[]Image {
	var data []Image
	url := "http://192.168.34.158:5000/v2/_catalog"
	req := curl.NewRequest()
	resp, err := req.SetUrl(url).Get()
	if err != nil {
		//fmt.Println(err)
		glog.Errorln(err)
	} else {
		if resp.IsOk() {
			var im map[string][]string
			if err := json.Unmarshal([]byte(resp.Body), &im); err != nil {
				glog.Errorln(err)
				panic(err)
			}
			for _, v := range im["repositories"] {
				images := LoadImage(v)
				for _, v := range images {
					data = append(data, v)
				}
			}
		} else {
			fmt.Println(resp.Raw)
		}
	}
	//glog.Infoln(data)
	return &data
}

//DelImage DelImage
func DelImage(imagename string, tag string) error {
	glog.Infoln(imagename)
	glog.Infoln(tag)
	digest := GetDigest(imagename, tag)
	url := "http://192.168.34.158:5000/v2/" + imagename + "/manifests/" + digest
	fmt.Println(url)
	req := curl.NewRequest()
	resp, err := req.Send(url, "DELETE")
	if err != nil {
		//fmt.Println(err)
		glog.Errorln(err)
	} else {
		if resp.IsOk() {
			return nil
		} else {
			//fmt.Println(resp.Raw)
			glog.Infoln(resp.Raw)
		}
	}

	return err

}

//LoadImage LoadImage
func LoadImage(name string) []Image {
	var images []Image
	url := "http://192.168.34.158:5000/v2/" + name + "/tags/list"
	req := curl.NewRequest()
	resp, err := req.SetUrl(url).Get()
	if err != nil {
		fmt.Println(err)
	} else {
		if resp.IsOk() {
			var im map[string]interface{}
			if err := json.Unmarshal([]byte(resp.Body), &im); err != nil {
				panic(err)
			}
			if im["tags"] != nil {
				for _, v := range interface{}(im["tags"]).([]interface{}) {
					var imagedata Image
					imagedata.Name = name
					imagedata.Tag = interface{}(v).(string)
					images = append(images, imagedata)
				}
			}

		} else {
			fmt.Println(resp.Raw)
		}
	}
	//glog.Infoln(data)
	return images
}

//GetDigest GetDigest
func GetDigest(imagename string, tag string) string {
	var hr map[string]string
	var header map[string]string
	header = make(map[string]string)
	header["Accept"] = "application/vnd.docker.distribution.manifest.v2+json"
	url := "http://192.168.34.158:5000/v2/" + imagename + "/manifests/" + tag
	fmt.Println(url)
	req := curl.NewRequest()
	resp, err := req.SetUrl(url).SetHeaders(header).Get()
	if err != nil {
		//fmt.Println(err)
		glog.Errorln(err)
	} else {
		if resp.IsOk() {
			hr = resp.Headers
		} else {
			fmt.Println(resp.Raw)
		}
	}
	fmt.Println(hr["Docker-Content-Digest"])
	return hr["Docker-Content-Digest"]
}
