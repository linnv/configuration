package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

func HttpTimeConsumptionDemo() {
	println("//<<-------------------------HttpTimeConsumptionDemo start-----------")
	start := time.Now()

	var a = struct {
		A      string   `json:"a"`
		Aa     []string `json:"a a"`
		Aaint  []int    `json:"b"`
		Aauint []uint   `json:"c"`
	}{A: "xxx",
		Aa:     []string{"dd", "fefe"},
		Aaint:  []int{2, 5},
		Aauint: []uint{23, 445},
	}
	form, err := Json2HttpForm(a)
	if err != nil {
		panic(err.Error())
	}

	data := form.Encode()
	fmt.Printf("data for preparation HttpTimeConsumptionDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	for {
		start = time.Now()
		req, _ := http.NewRequest("POST", "http://cent.local", bytes.NewBuffer([]byte(data)))
		//this header must be set when posting data in body
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		// fmt.Printf("data for sending HttpTimeConsumptionDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err.Error())
		}

		ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		fmt.Printf("post costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
		fmt.Printf("start.String(): %+v========\n", start.String())
		println("sleeping")
		time.Sleep(2)
		println("wake up")
	}
	println("//---------------------------HttpTimeConsumptionDemo end----------->>")
}

func Json2HttpForm(s interface{}) (form *url.Values, err error) {
	if reflect.TypeOf(s).Kind() != reflect.Struct {
		return nil, errors.New(fmt.Sprintf("usopport convert type %s", reflect.TypeOf(s).Kind()))
	}

	types := reflect.TypeOf(s)
	values := reflect.ValueOf(s)
	form = &url.Values{}
	for i := 0; i < types.NumField(); i++ {
		a := types.Field(i)
		av := values.Field(i)
		switch av.Kind() {
		case reflect.String:
			form.Add(a.Tag.Get("json"), av.String())
		case reflect.Float32, reflect.Float64:
			form.Add(a.Tag.Get("json"), fmt.Sprintf("%f", av.Float()))
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			form.Add(a.Tag.Get("json"), fmt.Sprintf("%d", av.Int()))
		case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			form.Add(a.Tag.Get("json"), fmt.Sprintf("%d", av.Uint()))
		case reflect.Slice:
			vs := reflect.ValueOf(av.Interface())
			vsLen := vs.Len()
			ss := make([]string, vsLen)
			switch t := vs.Interface().(type) {
			case []int, []int16, []int32, []int64:
				for j := 0; j < vsLen; j++ {
					ss[j] = fmt.Sprintf("%d", vs.Index(j).Int())
				}
			case []uint, []uint16, []uint32, []uint64:
				for j := 0; j < vsLen; j++ {
					ss[j] = fmt.Sprintf("%d", vs.Index(j).Uint())
				}
			case []float32, []float64:
				for j := 0; j < vsLen; j++ {
					ss[j] = fmt.Sprintf("%f", vs.Index(j).Float())
				}
			case []string:
				for j := 0; j < vsLen; j++ {
					ss[j] = fmt.Sprintf("%s", vs.Index(j))
				}
			default:
				return nil, errors.New(fmt.Sprintf("%s, element type is unsupported", t))

			}
			form.Add(a.Tag.Get("json"), strings.Join(ss, ","))

		default:
			return nil, errors.New(fmt.Sprintf("%s, element type is unsupported", av.Kind()))

		}
	}

	return
}

func response(w http.ResponseWriter, obj interface{}) {
	var (
		body []byte
		err  error
	)
	switch obj := obj.(type) {
	case error:
		w.WriteHeader(400)
		body, _ = json.Marshal(map[string]string{
			"error": obj.Error(),
		})
		log.Println(obj)
	case []byte:
		body = obj
	default:
		body, err = json.Marshal(map[string]interface{}{
			"result": obj,
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
	// json
	w.Header().Set("Content-Type", "application/json")
	n, err := w.Write(body)
	if err == nil && n < len(body) {
		err = fmt.Errorf("short written")
	}
	if err != nil {
		log.Println("wrote back to client: ", err)
	}
}

func HasPort(s string) bool { return strings.LastIndex(s, ":") > strings.LastIndex(s, "]") }

func MockTransportDemo() {
	println("//<<-------------------------MockTransportDemo start-----------")
	start := time.Now()

	client := http.DefaultClient
	client.Transport = newMockTransport()

	// resp, err := client.Get("http://ifconfig.co/all.json")
	resp, err := client.Get("jialinwu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("GET http://ifconfig.co/all.json")
	fmt.Println(string(body))

	fmt.Printf("MockTransportDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------MockTransportDemo end----------->>")
}

type mockTransport struct{}

func newMockTransport() http.RoundTripper {
	return &mockTransport{}
}

// Implement http.RoundTripper
func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Create mocked http.Response
	response := &http.Response{
		Header:     make(http.Header),
		Request:    req,
		StatusCode: http.StatusOK,
	}
	response.Header.Set("Content-Type", "application/json")

	responseBody :=
		`{
    "Accept-Encoding": [
        "info generated by mocking"
    ],
    "User-Agent": [
        "mock"
    ],
    "X-Ip-Country": [
        "Japan(Mock)"
    ],
    "X-Real-Ip": [
        "192.168.1.1"
    ]
}`
	response.Body = ioutil.NopCloser(strings.NewReader(responseBody))
	return response, nil
}

// func Json2HttpForm(s interface{}) (form *url.Values, err error) {
// 	if reflect.TypeOf(s).Kind() != reflect.Struct {
// 		return nil, fmt.Errorf("usopport convert type %s", reflect.TypeOf(s).Kind())
// 	}
//
// 	types := reflect.TypeOf(s)
// 	values := reflect.ValueOf(s)
// 	form = &url.Values{}
// 	for i := 0; i < types.NumField(); i++ {
// 		a := types.Field(i)
// 		av := values.Field(i)
// 		switch av.Kind() {
// 		case reflect.String:
// 			form.Add(a.Tag.Get("json"), av.String())
// 		case reflect.Float32, reflect.Float64:
// 			form.Add(a.Tag.Get("json"), fmt.Sprintf("%f", av.Float()))
// 		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
// 			form.Add(a.Tag.Get("json"), fmt.Sprintf("%d", av.Int()))
// 		case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
// 			form.Add(a.Tag.Get("json"), fmt.Sprintf("%d", av.Uint()))
// 		case reflect.Slice:
// 			vs := reflect.ValueOf(av.Interface())
// 			vsLen := vs.Len()
// 			ss := make([]string, vsLen)
// 			switch t := vs.Interface().(type) {
// 			case []int, []int16, []int32, []int64:
// 				for j := 0; j < vsLen; j++ {
// 					ss[j] = fmt.Sprintf("%d", vs.Index(j).Int())
// 				}
// 			case []uint, []uint16, []uint32, []uint64:
// 				for j := 0; j < vsLen; j++ {
// 					ss[j] = fmt.Sprintf("%d", vs.Index(j).Uint())
// 				}
// 			case []float32, []float64:
// 				for j := 0; j < vsLen; j++ {
// 					ss[j] = fmt.Sprintf("%f", vs.Index(j).Float())
// 				}
// 			case []string:
// 				for j := 0; j < vsLen; j++ {
// 					ss[j] = fmt.Sprintf("%s", vs.Index(j))
// 				}
// 			default:
// 				return nil, fmt.Errorf("%s, element type is unsupported", t)
//
// 			}
// 			form.Add(a.Tag.Get("json"), strings.Join(ss, ","))
//
// 		default:
// 			return nil, fmt.Errorf("%s, element type is unsupported", av.Kind())
//
// 		}
// 	}
//
// 	return
// }
