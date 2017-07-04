// Package main provides ...
package demo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const doc = `
  /**
        * 手机号码
        * 移动：134[0-8],135,136,137,138,139,150,151,157,158,159,182,187,188
        * 联通：130,131,132,152,155,156,185,186
        * 电信：133,1349,153,180,189
        */
       NSString * MOBILE = @"^1(3[0-9]|5[0-35-9]|8[025-9])\\d{8}$";
       /**
        10         * 中国移动：China Mobile
        11         * 134[0-8],135,136,137,138,139,150,151,157,158,159,182,187,188
        12         */
       NSString * CM = @"^1(34[0-8]|(3[5-9]|5[017-9]|8[278])\\d)\\d{7}$";
       /**
        15         * 中国联通：China Unicom
        16         * 130,131,132,152,155,156,185,186
        17         */
       NSString * CU = @"^1(3[0-2]|5[256]|8[56])\\d{8}$";
       /**
        20         * 中国电信：China Telecom
        21         * 133,1349,153,180,189
        22         */
       NSString * CT = @"^1((33|53|8[09])[0-9]|349)\\d{7}$";
       /**
        25         * 大陆地区固话及小灵通
        26         * 区号：010,020,021,022,023,024,025,027,028,029
        27         * 号码：七位或八位
        28         */
      // NSString * PHS = @"^0(10|2[0-5789]|\\d{3})\\d{7,8}$";
`

func PhoneDemo(number string) bool {
	println("\\<<-------------------------PhoneDemo start-----------")
	start := time.Now()
	//@TODO testcase
	//mobile,unicom,Telecom
	regMobile := `^1(3[0-9]|5[0-35-9]|8[025-9])\d{8}$`
	// rgx := regexp.MustCompile(regMobile)
	if !regexp.MustCompile(regMobile).MatchString(number) {
		houseNumberReg := `^(0[0-9]{2,3}\-)?([2-9][0-9]{6,7})+(\-[0-9]{1,4})?$`
		return regexp.MustCompile(houseNumberReg).MatchString(number)
	}
	// s := []string{"18505921256", "13489594009", "12759029321", "18290015121"}
	// for _, v := range s {
	// 	fmt.Println(rgx.MatchString(v))
	// }

	// s = append(s, []string{"8585179"}...)
	// log.Println("house number: works")
	// rgx = regexp.MustCompile(houseNumberReg)
	// for _, v := range s {
	// 	fmt.Println(rgx.MatchString(v))
	// }
	// output:
	//    true
	//    true
	//    false
	fmt.Printf("PhoneDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("\\---------------------------PhoneDemo end----------->>")
	return true
}

func PersonalIDDemo(id string) bool {
	println("//<<-------------------------PersonalIDDemo start-----------")
	start := time.Now()
	id = strings.ToUpper(id)
	// id = "120111199004182802"
	var aCity = map[int]string{11: "北京", 12: "天津", 13: "河北", 14: "山西", 15: "内蒙古", 21: "辽宁", 22: "吉林", 23: "黑龙江 ", 31: "上海", 32: "江苏", 33: "浙江", 34: "安徽", 35: "福建", 36: "江西", 37: "山东", 41: "河南", 42: "湖北 ", 43: "湖南", 44: "广东", 45: "广西", 46: "海南", 50: "重庆", 51: "四川", 52: "贵州", 53: "云南", 54: "西藏 ", 61: "陕西", 62: "甘肃", 63: "青海", 64: "宁夏", 65: "新Unbsp;疆", 71: "台湾", 81: "香港", 82: "澳门", 91: "国外 "}
	reg := `^\d{6}(18|19|20)?\d{2}(0[1-9]|1[12])(0[1-9]|[12]\d|3[01])\d{3}(\d|X)$`
	log.Printf("`aCity`: %+v\n", `aCity`)
	if !regexp.MustCompile(reg).MatchString(id) {
		log.Println("length illegal: works")
		return false
	}
	log.Println(": works", string(id[:2]))
	num, _ := strconv.Atoi(string(id[:2]))
	if _, ok := aCity[num]; !ok {
		log.Println("illegal address code: works")
		return false
	}
	if len(id) == 18 {
		var factor = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
		//校验位
		var parity = []int{1, 0, 10, 9, 8, 7, 6, 5, 4, 3, 2}
		var sum = 0
		var ai = 0
		var wi = 0
		for i := 0; i < 17; i++ {
			if id[i] == 'X' {
				ai = 10
			} else {
				num, _ := strconv.Atoi(string(id[i]))
				ai = num
			}
			wi = factor[i]
			sum += ai * wi
			log.Printf("ai: %+v\n", ai)
			log.Printf("wi: %+v\n", wi)
			println()
		}
		// log.Println("int('X'): works", int('X'))
		// log.Printf("last: %+v\n", last)
		var last = parity[sum%11]
		var end int
		if id[17] == 'X' {
			end = 10
		} else {
			//@TODO
			// end, _ = strconv.Atoi(string(id[17]))
			end = int(id[17]) - int('0')
		}
		if last != end {
			log.Printf("last: %+v\n", last)
			log.Printf("end: %+v\n", end)
			log.Println("校验位错误: works")
			return false
		}
	}
	fmt.Printf("PersonalIDDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------PersonalIDDemo end----------->>")
	return true
}

func CarNOValidationDemo(cn string) {
	println("//<<-------------------------CarNOValidationDemo start-----------")
	start := time.Now()
	cn = "6228430120000000000"
	// https: //ccdcapi.alipay.com/validateAndCacheCardInfo.json?_input_charset=utf-8&cardNo=6222021309000088888&cardBinCheck=true
	// https://ccdcapi.alipay.com/validateAndCacheCardInfo.json?_input_charset=utf-8&cardNo=6228430120000000000&cardBinCheck=true
	basic := "https://ccdcapi.alipay.com/validateAndCacheCardInfo.json?_input_charset=utf-8&cardNo=%s&cardBinCheck=true"
	url := fmt.Sprintf(basic, cn)
	const TimeOut = 20
	c := &http.Client{
		Timeout: time.Second * TimeOut,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: TimeOut * time.Second,
			}).Dial,
			TLSHandshakeTimeout: TimeOut * time.Second,
		},
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	response, err := c.Do(req)
	if err != nil {
		return
	}
	// {"bank":"ABC","validated":true,"cardType":"DC","key":"6228430120000000000","messages":[],"stat":"ok"}
	// {"validated":false,"key":"622843012000x000000","stat":"ok","messages":[{"errorCodes":"PARAM_ILLEGAL","name":"cardNo"}]}
	//http://www.voidcn.com/blog/onceing/article/p-4415076.html
	var ret map[string]interface{}
	bs, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(bs, &ret)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("rep: %+v\n", ret)
	fmt.Printf("CarNOValidationDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------CarNOValidationDemo end----------->>")
}

// http://shockerli.leanote.com/post/javascript-regula-idcard
func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}

func CarNOValidationLocalDemo(cn string) bool {
	var regStr string
	switch len(cn) {
	case 16:
		regStr = "^[0-9]{16}$"
	case 19:
		regStr = "^[0-9]{19}$"
	}
	b, _ := regexp.MatchString(regStr, cn)
	return b
	// var regStr, regSpaceStr string
	// regSpaceStr = `^[0-9]{4}((?:\s[0-9]{4}){3})|((?:\s[0-9]{4}){3}\s[0-9]{3})$`
	// if !regexp.MustCompile(regStr).MatchString(cn) {
	// 	println("false")
	// 	return false
	// }
	// return regexp.MustCompile(regSpaceStr).MatchString(cn)
	// return true
}
func RDemo(cn string) bool {
	println("//<<-------------------------RDemo start-----------")
	start := time.Now()
	reg := `^\d{2}$`
	// reg := `^\d{2}$`
	if regexp.MustCompile(reg).MatchString(cn) {
		return true
	}
	fmt.Printf("RDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------RDemo end----------->>")
	return false
}

func NumberRegDemo(s string) bool {
	println("//<<-------------------------NumberRegDemo start-----------")
	start := time.Now()
	reg := `\d*`
	r := regexp.MustCompile(reg).MatchString(s)
	fmt.Printf("NumberRegDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------NumberRegDemo end----------->>")
	return r
}
