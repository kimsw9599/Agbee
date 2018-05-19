package main


import(
	"os"
	"crawling_util"
)

func main(){
	for _, url := range os.Args[1:]{
		crawling_util.Fetch(url)
	}

	image_url := "http://cfile21.uf.tistory.com/image/187CAC4A50EC2FF233A87F"
	crawling_util.ImageDownload(image_url, "./eunbin.jpg")

	image_url2 := "http://cfile5.uf.tistory.com/image/2106824D50EC2FF923F720"
	crawling_util.ImageDownload(image_url2, "./eunbin2.jpg")
}
