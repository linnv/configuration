package demo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/linnv/logx"

	"golang.org/x/crypto/bcrypt"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	save2FileDemo()
	println("-----------------------------JustDemo end>>>")
	return
}

func save2FileDemo() {
	println("//<<-------------------------save2FileDemo start-----------")
	const endpoint = "http://www.baidu.com"
	httpclient := &http.Client{
		Timeout: time.Second,
	}
	req, err := http.NewRequest("GET", endpoint, strings.NewReader(""))
	logx.CheckErr(err)

	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")

	resp, err := httpclient.Do(req)
	logx.CheckErr(err)

	respBody, err := ioutil.ReadAll(resp.Body)
	logx.CheckErr(err)
	resp.Body.Close()

	if len(respBody) == 0 {
		respBody = []byte("{}")
	}

	logx.Debugf("string(respBody): %+v\n", string(respBody))
	println("//---------------------------save2FileDemo end----------->>")
}

func ExampleGet() {
	// res, err := http.Get("http://www.google.com/robots.txt")
	res, err := http.Get("http://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  res.ContentLength: %+v bytes\n", res.ContentLength)
	// _, err = ioutil.ReadAll(res.Body)
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	fmt.Printf("  len(rebots): %+v\n", len(robots))
	fmt.Printf("  cap(rebots): %+v\n", cap(robots))
	// fmt.Printf("%s", robots)
}

func ExampleGetFixLength() {
	// res, err := http.Get("http://www.google.com/robots.txt")
	res, err := http.Get("http://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  res.ContentLength: %+v bytes\n", res.ContentLength)
	// _, err = ioutil.ReadAll(res.Body)
	// robots := make([]byte, 0, res.ContentLength)
	robots := make([]byte, res.ContentLength)
	n, err := io.ReadFull(res.Body, robots)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  readfull n bytes: %+v\n", n)
	fmt.Printf("  len(rebots): %+v\n", len(robots))
	fmt.Printf("  cap(rebots): %+v\n", cap(robots))
	// fmt.Printf("%s", robots)
}

func ExampleGetFixLengthUsingBuffer() {
	// res, err := http.Get("http://www.google.com/robots.txt")
	res, err := http.Get("http://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  res.ContentLength: %+v bytes\n", res.ContentLength)
	// _, err = ioutil.ReadAll(res.Body)
	// robots := make([]byte, 0, res.ContentLength)
	// buffer := bytes.NewBuffer(make([]byte, 0, 65536))
	// buffer := bytes.NewBuffer(make([]byte, 0, 19639))
	buffer := bytes.NewBuffer(make([]byte, 0, 19640))
	io.Copy(buffer, res.Body)
	robots := buffer.Bytes()
	fmt.Printf("  len(rebots): %+v\n", len(robots))
	fmt.Printf("  cap(rebots): %+v\n", cap(robots))
	length := len(robots)
	var body []byte
	//are we wasting more than 10% space?
	if cap(robots) > (length + length/10) {
		body = make([]byte, length)
		copy(body, robots)
	} else {
		body = robots
	}
	fmt.Printf("  len(body): %+v\n", len(body))
	fmt.Printf("  cap(body): %+v\n", cap(body))
}

func TemplateDemo() {
	println("//<<-------------------------TemplateDemo start-----------")
	start := time.Now()

	hc := `
	hello {{.}}
	`
	var t = template.Must(template.New("name").Parse(hc))
	err := t.Execute(os.Stdout, "jialin")
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("TemplateDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------TemplateDemo end----------->>")
}

func toUnicode(s string) (r string) {
	p := []rune(s)
	for i := 0; i < len(p); i++ {
		r += "%u" + fmt.Sprintf("%U", p[i])[2:]
	}
	return r
}

func UrlDemo() {
	println("//<<-------------------------UrlDemo start-----------")
	start := time.Now()
	// a := make(url.Values)
	// v := []rune("安庆")
	v := "安庆"
	r := toUnicode("安庆")
	// r := "%u" + fmt.Sprintf("%U", v[0])[2:]
	// r := strconv.QuoteRuneToASCII(v[0])
	// r = url.QueryEscape(r)
	log.Printf("v: %+v\n", v)
	log.Printf("r: %s\n", r)
	fmt.Printf("UrlDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------UrlDemo end----------->>")
}

func SecureHeadersDemo() {
	println("//<<-------------------------SecureHeadersDemo start-----------")
	start := time.Now()

	const (
		stsHeader           = "Strict-Transport-Security"
		stsSubdomainString  = "; includeSubdomains"
		stsPreloadString    = "; preload"
		frameOptionsHeader  = "X-Frame-Options"
		frameOptionsValue   = "DENY"
		contentTypeHeader   = "X-Content-Type-Options"
		contentTypeValue    = "nosniff"
		xssProtectionHeader = "X-XSS-Protection"
		xssProtectionValue  = "1; mode=block"
		cspHeader           = "Content-Security-Policy"
		hpkpHeader          = "Public-Key-Pins"
	)
	println(xssProtectionHeader, ":", xssProtectionValue)
	println(stsHeader, ":", stsSubdomainString, stsPreloadString) //@Todo
	println(frameOptionsHeader, ":", frameOptionsValue)
	println(contentTypeHeader, ":", contentTypeValue)

	fmt.Printf("SecureHeadersDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------SecureHeadersDemo end----------->>")
}

func AuthoricationDemo(token string) bool {
	println("//<<-------------------------AuthoricationDemo start-----------")
	start := time.Now()
	passwordHashToDB, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	logx.CheckErr(err)
	logx.Debugf("passwordHashToDB: %+v\n", passwordHashToDB)
	// bcrypt.CompareHashAndPassword(passwordHashFromDB, []byte(token))
	fmt.Printf("AuthoricationDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	// http.StatusNonAuthoritativeInfo
	println("//---------------------------AuthoricationDemo end----------->>")
	return false
}

func RawDemo() error {
	println("//<<-------------------------RawDemo start-----------")
	start := time.Now()

	req, err := http.NewRequest("GET", "https://jialinwu.com", strings.NewReader("form-data"))
	logx.CheckErr(err)
	r, err := http.DefaultClient.Do(req)
	logx.CheckErr(err)
	bs, err := ioutil.ReadAll(r.Body)
	logx.CheckErr(err)
	r.Body.Close()
	logx.Debugf("string(bs): %+v\n", string(bs))

	fmt.Printf("RawDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------RawDemo end----------->>")
	return nil
}

type apiRequest struct {
	Auth    apiAuthHeader `json:"auth"`
	Content interface{}   `json:"content"`
}

type apiAuthHeader struct {
	DspId int    `json:"dspId"`
	Token string `json:"token"`
}

func PostOneDemo() error {
	println("//<<-------------------------RawDemo start-----------")
	start := time.Now()

	const data = `
{
  "auth": {
    "dspId": 21,
    "token": "4b300e739645ef7f80146e449bfeee04f858144fe2f67bf6"
  },
  "content":[
    {
      "name": "nn22224"
    }
  ]
}
	`

	a := apiRequest{}
	err := json.Unmarshal([]byte(data), &a)
	logx.CheckErr(err)
	logx.Debugf("a: %+v\n", a)
	bs, err := json.Marshal(a)
	logx.CheckErr(err)
	// req, err := http.NewRequest("POST", "http://sspmbv.local/api/ad/one", strings.NewReader(string(bs)))
	// req, err := http.NewRequest("GET", "http://sspmbv.local/api/ad/one", strings.NewReader(""))
	req, err := http.NewRequest("POST", "http://127.0.0.1:10000/api/ad/upload", strings.NewReader(string(bs)))
	logx.CheckErr(err)
	req.Header.Set("Cookie", "")
	r, err := http.DefaultClient.Do(req)
	logx.CheckErr(err)
	bs, err = ioutil.ReadAll(r.Body)
	logx.CheckErr(err)
	r.Body.Close()
	logx.Debugf("string(bs): %+v\n", string(bs))

	fmt.Printf("RawDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------RawDemo end----------->>")
	return nil
}

var defaultClient *http.Client

func init() {
	defaultClient = &http.Client{
		Timeout: time.Second * 10,
	}
}

func ExactMagnetDemo(urlAddr string) []string {
	println("//<<-------------------------ExactMagnetDemo start-----------")
	start := time.Now()
	resp, err := ioutil.ReadFile("/Users/Jialin/Downloads/1.htm")
	logx.CheckErr(err)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))
	logx.PanicErr(err)
	hd := map[string]string{
		"Accept":     "*/*",
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.153 Safari/537.36",
	}

	var bs []byte
	var bufByte = bytes.NewBuffer(bs)
	count := 0
	pageCount := 1
	const macCount = 15
	const curPath = "/Users/Jialin/myGit/OpenDemo/golang/demos/httpDemo/magnet"

	doc.Find(".item").Each(func(i int, s *goquery.Selection) {
		magnet, ok := s.Find("a").Attr("href")
		if !ok {
			println("not found")
			return
		}
		req, err := http.NewRequest("GET", magnet, nil)
		logx.CheckErr(err)
		for k, v := range hd {
			req.Header.Set(k, v)
		}

		resp, err := defaultClient.Do(req)
		logx.CheckErr(err)
		if resp.StatusCode != http.StatusOK {
			fmt.Errorf("获取远程文件失败[%d]", resp.StatusCode)
		}
		subDoc, err := goquery.NewDocumentFromResponse(resp)
		logx.CheckErr(err)
		subDoc.Find(".detail .magnet").Each(func(i int, s *goquery.Selection) {
			magnet, ok := s.Find("a").Attr("href")
			if !ok {
				return
			}

			count++
			bufByte.WriteString(magnet)
			err = bufByte.WriteByte(byte('\n'))
			logx.PanicErr(err)
			println(magnet)
			if count == macCount {
				err = ioutil.WriteFile(curPath+strconv.Itoa(pageCount), bufByte.Bytes(), os.ModePerm)
				logx.CheckErr(err)
				count = 0
				pageCount++
				bufByte.Reset()
			}
		})
	})
	err = ioutil.WriteFile(curPath+strconv.Itoa(pageCount), bufByte.Bytes(), os.ModePerm)
	logx.CheckErr(err)
	fmt.Printf("ExactMagnetDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------ExactMagnetDemo end----------->>")
	return nil
}
