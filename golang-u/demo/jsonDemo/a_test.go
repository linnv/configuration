package demo

import "testing"

// . "github.com/smartystreets/goconvey/convey"

// func TestMarshalDemo(t *testing.T) {
// 	Convey("value inherit without name", t, func() {
// 		MarshalDemo()
// 		MarshalDemoP()
// 		t.Log("value inherit without name :good")
// 		// So(1, ShouldNotEqual, 1)
// 	})
// }
//
// func TestMarshalDemoA(t *testing.T) {
// 	Convey("pinter inherit withouth name", t, func() {
// 		MarshalDemoA()
// 		MarshalDemoAP()
// 		t.Log("pinter inherit withouth name fault")
// 	})
// }
//
// func TestUnMarshalDemo(t *testing.T) {
// 	Convey("unMarshal ", t, func() {
// 		bs := []byte(`{"name":"a","bName":"bn"}`)
// 		err := UnMarshalDemo(bs)
// 		So(err, ShouldEqual, nil)
// 	})
// }
//
// func TestUnMarshalDemoP(t *testing.T) {
// 	Convey("unMarshal P", t, func() {
// 		bs := []byte(`{"name":"a","bName":"bn"}`)
// 		err := UnMarshalDemoP(bs)
// 		So(err, ShouldEqual, nil)
// 	})
// }
//
// func TestUnMarshalDemoA(t *testing.T) {
// 	Convey("unMarshal A", t, func() {
// 		// bs := []byte(`{"name":"a","bName":"bn"}`)
// 		bs := []byte(`{"name":"a","B":{"bName":"bn"}}`)
// 		// unMarshal B with nil while B is anonymous
// 		err := UnMarshalDemoA(bs)
// 		So(err, ShouldEqual, nil)
// 	})
// }
// func TestUnMarshalDemoAP(t *testing.T) {
// 	Convey("unMarshal AP", t, func() {
// 		bs := []byte(`{"name":"a","bName":"bn"}`)
//
// 		// unMarshal B with nil while B is anonymous
// 		// bs := []byte(`{"name":"a","B":{"bName":"bn"}}`)
// 		err := UnMarshalDemoAP(bs)
// 		So(err, ShouldEqual, nil)
// 	})
// }

// func TestMarshalDemo(t *testing.T) {
// 	Convey("value inherit withouth name", t, func() {
// 		MarshalDemo()
// 		MarshalDemoP()
// 		t.Log("good")
// 		// So()
// 	})
// }

// func TestMarshalDemoA(t *testing.T) {
// 	Convey("pointer inherit withouth name", t, func() {
// 		MarshalDemoA()
// 		MarshalDemoAP()
// 		t.Log("good")
// 		// So()
// 	})
// }

func TestJustDemo(t *testing.T) {
	IssurDemo()
	// s := secure.New(
	// 	secure.Options{
	// 		AllowedHosts:            []string{},
	// 		SSLRedirect:             false,
	// 		SSLTemporaryRedirect:    false,
	// 		SSLHost:                 "",
	// 		SSLProxyHeaders:         map[string]string{},
	// 		STSSeconds:              0,
	// 		STSIncludeSubdomains:    false,
	// 		STSPreload:              false,
	// 		ForceSTSHeader:          false,
	// 		FrameDeny:               false,
	// 		CustomFrameOptionsValue: "",
	// 		ContentTypeNosniff:      false,
	// 		BrowserXssFilter:        false,
	// 		ContentSecurityPolicy:   "",
	// 		PublicKey:               "",
	// 		IsDevelopment:           false,
	// 	})
	//
	// err := GeneratTemplateJson(&s)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// DoIt()
	// JustDemo()
	// TimeDemo()
	// BsonDemo()
	// WireUnMarshalDemoAP()
	// td := &T{}
	// td.Times = time.Now()
	// td.C = 1111
	// td.N = "eeee"
	//
	// tbs, err := bson.Marshal(td)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(string(tbs))
	//
	// type Response1 struct {
	// 	Page   int
	// 	Fruits []string
	// }
	// res1D := &Response1{
	// 	Page:   1,
	// 	Fruits: []string{"apple", "peach", "pear"}}
	// res1B, _ := json.Marshal(res1D)
	// fmt.Println(string(res1B))
	// fmt.Printf("  string(tbs): %s\n", string(tbs))
	// postData := adslot.PostBuyRule{}

	// postData := &apiSlot{ViewType: &apiViewType{}, Site: &Site{}}
	// err := GeneratTemplateJson(postData)
	// if err != nil {
	// 	panic(err.Error())
	// 	return
	// }

	// tt := time.Unix(1362984425, 0)
	// nt := tt.Format("2006-01-02 15:04:05")
	// fmt.Println(nt)
}

// func TestLoadJsonFile(t *testing.T) {
//
// 	path := "/Users/Jialin/golang/src/demo/jsonDemo/file.json"
//
// 	fs := []ImageRatio{}
// 	// fs := FS{}
// 	err := LoadJsonFile(path, &fs)
// 	if err != nil {
// 		panic(err.Error())
// 		return
// 	}
// 	for k, v := range fs {
// 		fmt.Printf("%+v: %+v\n", k, v)
// 		fmt.Printf("  v.H/v.W: %+v\n", v.Min.H/v.Max.W)
// 	}
// }
