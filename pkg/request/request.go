package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func HttpPostForm(requestUrl string, params map[string]string) string {
	form := url.Values{}
	for k, v := range params {
		form.Set(k, v)
	}
	fmt.Println(form)
	resp, err := http.PostForm(requestUrl, form)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}
