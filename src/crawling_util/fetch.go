package crawling_util

import (
	"net/http"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"io"
	"fmt"
)

func Fetch_ex(url string){
	transport := getProxy("http://192.168.2.101:8080")

	client := &http.Client{}
	client.Transport = &transport

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))}

func getProxy(proxyAddr string) http.Transport {
	url_i := url.URL{}
	url_proxy, _ := url_i.Parse(proxyAddr)

	transport := http.Transport{}
	transport.Proxy = http.ProxyURL(url_proxy)// set proxy
	//transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return transport
}

func ImageDownload(url string, filename string){
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	fmt.Println("Success!")
}

func Fetch(url string){
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

