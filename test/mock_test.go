package test

import (
	"CM-SSO/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var timeout = 10 * time.Second
func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
func Racer(a,b string) (winner string,err error){
	select {
	case <-ping(a):
		return a,nil
	case <-ping(b):
		return b,nil
	case <-time.After(timeout):
		return "",fmt.Errorf("timed out waiting for %s and %s", a, b)

	}
}

//mock http server
func TestRacer(t *testing.T){
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	defer slowServer.Close()
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastURL := fastServer.URL

	_,err := Racer(slowUrl,fastURL)

	if err != nil{
		t.Errorf("expected an error but didn't get one")
	}
}

var testUrl = "http://127.0.0.1:8090/cmsso/"

func TestSendEmail(t *testing.T) {
	postData := model.SignUpForm{Identifier: "asxxw741@qq.com",Credential:"as1523318199"}
	reqBody, _ := json.Marshal(postData)
	fmt.Println("input:", string(reqBody))
	req, err := http.NewRequest("PUT",testUrl+"mail",bytes.NewReader(reqBody))
	if err != nil {
		t.Error("请求失败",err)
	}
	req.Header.Add("Content-Type", "application/json")
	res,err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("请求失败",err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if err != nil {
		t.Error("请求失败",err)
	}
}
func TestEmailLogin(t *testing.T) {
	postData := model.SignUpForm{Identifier: "asxxw741@qq.com",Credential:"as1523318199"}
	reqBody, _ := json.Marshal(postData)
	fmt.Println("input:", string(reqBody))
	req, err := http.NewRequest("POST",testUrl+"mail",bytes.NewReader(reqBody))
	if err != nil {
		t.Error("请求失败",err)
	}
	req.Header.Add("Content-Type", "application/json")
	res,err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("请求失败",err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if err != nil {
		t.Error("请求失败",err)
	}
}

func TestSSOLogin(t *testing.T) {
	req, err := http.NewRequest("GET",testUrl,nil)
	if err != nil {
		t.Error("请求失败",err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBdXRoSWQiOjMsIlVzZXJJZCI6NSwiZXhwIjoxNTYzMzU5ODg1LCJpYXQiOjE1NjMzNTYyODV9.RnSRpqZwtrNH4O6qOWN0bOXMAnFdQqgXz5kEAnn_EuM")
	res,err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("请求失败",err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if err != nil {
		t.Error("请求失败",err)
	}
}
