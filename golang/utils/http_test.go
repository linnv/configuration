package utils

import (
	"bytes"
	"net/http"
	"testing"
)

func TestMockTransportDemo(t *testing.T) {
	MockTransportDemo()
	t.Log("test http transport func :good")
}

func TestHttpTimeConsumptionDemo(t *testing.T) {
	HttpTimeConsumptionDemo()
}

func TestJson2HttpForm(t *testing.T) {
	var a = struct {
		A      string   `json:"Aefjfejf"`
		Aa     []string `json:"str arrsy"`
		Aaint  []int    `json:"strint"`
		Aauint []uint   `json:"struint"`
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
	req, _ := http.NewRequest("POST", "http://www.jilainwu.com", bytes.NewBuffer([]byte(data)))
	//this header must be set when posting data in body
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// http.DefaultClient.Do()

	// or set Form directly
	// req.Form = *form
	// r := req.FormValue("strint")

	ra := req.FormValue("Aefjfejf")
	r := req.FormValue("strint")
	rr := req.FormValue("str arrsy")

	if ra != "xxx" {
		t.Log("unexpect value")
		return
	}
	if r != "2,5" {
		t.Log("unexpect value")
		return
	}

	if rr != "dd,fefe" {
		t.Log("unexpect value")
		return
	}

	// b, _ := ioutil.ReadAll(bytes.NewBuffer([]byte(data)))
	// vs, _ := url.ParseQuery(string(b))
	// vs should be same  with form
}
