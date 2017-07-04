package demo

import (
	"encoding/json"
	"log"
)

type ResponseBody struct {
	Showapi_res_body ResBody `json:"showapi_res_body"`
}

type ResBody struct {
	F1       DayRecord `json:"f1"` //today,next day, f2...f7
	Ret_code int       `json:"ret_code"`
	Time     string    `json:"time"` //day precission
	Now      NowRecord `json:"now"`
}

type NowRecord struct {
	AqiDetail        AqiDetail `json:"aqiDetail"`
	Temperature_time string    `json:"temperature_time"`
	Wind_power       string    `json:"wind_power"`
	Aqi              int       `json:"aqi"`
	Sd               string    `json:"sd"`
	Temperature      string    `json:"temperature"`
}

type AqiDetail struct {
	Co    float32 `json:"co"`
	So2   float32 `json:"so2"`
	O3    int     `json:"o3"`
	No2   int     `json:"no2"`
	Pm10  int     `json:"pm10"`
	Pm2_5 int     `json:"pm2_5"`
	O3_8h int     `json:"o3_8h"`
}

type DayRecord struct {
	DayTemperature   string `json:"day_air_temperature"`   //hightest temp
	NightTemperature string `json:"night_air_temperature"` //lowest temp
	AirPress         string `json:"air_press"`
	Date             string `json:"day"`
}

func DoIt() {
	var one ResponseBody
	err := json.Unmarshal([]byte(aliyunWeather), &one)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("one: %+v\n", one)
}

const aliyunWeather = `
{
  "showapi_res_code": 0,
  "showapi_res_error": "",
  "showapi_res_body": {
    "f6": {
      "day_weather": "多云",
      "night_weather": "多云",
      "night_weather_code": "01",
      "index": {
        "yh": {
          "title": "不适宜",
          "desc": "风力较大，请您挑选避风的地点。"
        },
        "ls": {
          "title": "不适宜",
          "desc": "天气阴沉，不太适宜晾晒。"
        },
        "clothes": {
          "title": "较冷",
          "desc": "建议着厚外套加毛衣等服装。"
        },
        "dy": {
          "title": "不适宜钓鱼",
          "desc": "天气不好，不适合垂钓"
        },
        "sports": {
          "title": "较适宜",
          "desc": "温较低，推荐您进行室内运动。"
        },
        "travel": {
          "title": "较适宜",
          "desc": "较弱降水和微风将伴您共赴旅程。"
        },
        "beauty": {
          "title": "保湿",
          "desc": "请选用中性保湿型霜类化妆品。"
        },
        "xq": {
          "title": "较好",
          "desc": "温度适宜，心情会不错。"
        },
        "hc": {
          "title": "较适宜",
          "desc": "风大，需注意及时添衣。"
        },
        "zs": {
          "title": "不容易中暑",
          "desc": "气温不高，中暑几率极低。"
        },
        "cold": {
          "title": "少发",
          "desc": "感冒机率较低，避免长期处于空调屋中。"
        },
        "gj": {
          "title": "较适宜",
          "desc": "风稍大，出门逛街前记得给秀发定型。"
        },
        "uv": {
          "title": "弱",
          "desc": "辐射较弱，涂擦SPF12-15、PA+护肤品。"
        },
        "cl": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "glass": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "aqi": {
          "title": "良",
          "desc": "气象条件有利于空气污染物扩散。"
        },
        "ac": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "wash_car": {
          "title": "较不适宜",
          "desc": "风力较大，洗车后会蒙上灰尘。"
        },
        "mf": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "ag": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "pj": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "nl": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "pk": {
          "title": "暂缺",
          "desc": "暂缺"
        }
      },
      "air_press": "1004 hPa",
      "jiangshui": "15%",
      "night_wind_power": "4-5级 8.0~10.7m/s",
      "day_wind_power": "4-5级 8.0~10.7m/s",
      "day_weather_code": "01",
      "3hourForcast": [
        {
          "wind_direction": "北风",
          "temperature_max": "15",
          "wind_power": "3-4级",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/day/01.png",
          "temperature_min": "12",
          "hour": "8时-14时",
          "temperature": "14"
        },
        {
          "wind_direction": "北风",
          "temperature_max": "15",
          "wind_power": "4-5级",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/day/01.png",
          "temperature_min": "14",
          "hour": "14时-20时",
          "temperature": "15"
        },
        {
          "wind_direction": "北风",
          "temperature_max": "15",
          "wind_power": "3-4级",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/night/01.png",
          "temperature_min": "15",
          "hour": "20时-2时",
          "temperature": "15"
        },
        {
          "wind_direction": "北风",
          "temperature_max": "15",
          "wind_power": "3-4级",
          "weather": "晴",
          "weather_pic": "http://app1.showapi.com/weather/icon/night/00.png",
          "temperature_min": "15",
          "hour": "2时-8时",
          "temperature": "15"
        }
      ],
      "sun_begin_end": "07:20|18:11",
      "ziwaixian": "弱",
      "day_weather_pic": "http://app1.showapi.com/weather/icon/day/01.png",
      "weekday": 3,
      "night_air_temperature": "15",
      "day_air_temperature": "16",
      "day_wind_direction": "北风",
      "day": "20161228",
      "night_weather_pic": "http://app1.showapi.com/weather/icon/night/01.png",
      "night_wind_direction": "北风"
    },
    "f7": {
      "night_weather_code": "01",
      "day_weather": "多云",
      "night_weather": "多云",
      "night_wind_power": "3-4级",
      "day_wind_power": "4-5级",
      "day_weather_code": "01",
      "3hourForcast": [
        {
          "wind_direction": "北风",
          "temperature_max": "15",
          "wind_power": "3-4级",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/day/01.png",
          "temperature_min": "12",
          "hour": "8时-14时",
          "temperature": "14"
        },
        {
          "wind_direction": "北风",
          "temperature_max": "15",
          "wind_power": "4-5级",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/day/01.png",
          "temperature_min": "14",
          "hour": "14时-20时",
          "temperature": "15"
        },
        {
          "wind_direction": "北风",
          "temperature_max": "15",
          "wind_power": "3-4级",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/night/01.png",
          "temperature_min": "15",
          "hour": "20时-2时",
          "temperature": "15"
        },
        {
          "wind_direction": "北风",
          "temperature_max": "15",
          "wind_power": "3-4级",
          "weather": "晴",
          "weather_pic": "http://app1.showapi.com/weather/icon/night/00.png",
          "temperature_min": "15",
          "hour": "2时-8时",
          "temperature": "15"
        }
      ],
      "sun_begin_end": "07:20|18:11",
      "day_weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
      "weekday": 4,
      "night_air_temperature": "12",
      "day_air_temperature": "18",
      "day_wind_direction": "北风",
      "day": "20161229",
      "night_weather_pic": "http://appimg.showapi.com/images/weather/icon/night/01.png",
      "night_wind_direction": "北风"
    },
    "ret_code": 0,
    "time": "20161223000000",
    "now": {
      "aqiDetail": {
        "co": 1.65,
        "so2": 17,
        "area": "北海",
        "o3": 183,
        "no2": 20,
        "area_code": "beihai",
        "quality": "良",
        "aqi": 80,
        "pm10": 67,
        "pm2_5": 42,
        "o3_8h": 121,
        "primary_pollutant": "臭氧1小时"
      },
      "weather_code": "01",
      "wind_direction": "东南风",
      "temperature_time": "17:32",
      "wind_power": "2级",
      "aqi": 80,
      "sd": "57%",
      "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
      "weather": "多云",
      "temperature": "23"
    },
    "cityInfo": {
      "c6": "guangxi",
      "c5": "北海",
      "c4": "beihai",
      "c3": "北海",
      "c9": "中国",
      "c8": "china",
      "c7": "广西",
      "c17": "+8",
      "c16": "AZ9771",
      "c1": "101301301",
      "c2": "beihai",
      "c11": "0779",
      "longitude": 109.105,
      "c10": "2",
      "latitude": 21.485,
      "c12": "536000",
      "c15": "13"
    },
    "alarmList": [],
    "hourDataList": [
      {
        "aqiDetail": {
          "co": 1.55,
          "so2": 12,
          "area": "北海",
          "o3": 103,
          "no2": 19,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 63,
          "pm10": 44,
          "pm2_5": 35,
          "o3_8h": 114,
          "primary_pollutant": "臭氧8小时"
        },
        "weather_code": "01",
        "wind_direction": "东北风",
        "temperature_time": "00:02",
        "wind_power": "2级",
        "aqi": 63,
        "sd": "69%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/night/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.55,
          "so2": 12,
          "area": "北海",
          "o3": 103,
          "no2": 19,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 63,
          "pm10": 44,
          "pm2_5": 35,
          "o3_8h": 114,
          "primary_pollutant": "臭氧8小时"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "00:31",
        "wind_power": "2级",
        "aqi": 63,
        "sd": "68%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/night/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.525,
          "so2": 10,
          "area": "北海",
          "o3": 104,
          "no2": 17,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 61,
          "pm10": 42,
          "pm2_5": 35,
          "o3_8h": 112,
          "primary_pollutant": "臭氧8小时"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "01:03",
        "wind_power": "2级",
        "aqi": 61,
        "sd": "68%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/night/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.525,
          "so2": 10,
          "area": "北海",
          "o3": 104,
          "no2": 17,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 61,
          "pm10": 42,
          "pm2_5": 35,
          "o3_8h": 112,
          "primary_pollutant": "臭氧8小时"
        },
        "weather_code": "01",
        "wind_direction": "东北风",
        "temperature_time": "01:30",
        "wind_power": "2级",
        "aqi": 61,
        "sd": "68%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/night/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.55,
          "so2": 10,
          "area": "北海",
          "o3": 100,
          "no2": 19,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 53,
          "pm10": 47,
          "pm2_5": 36,
          "o3_8h": 100,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "02:00",
        "wind_power": "2级",
        "aqi": 53,
        "sd": "69%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/night/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.55,
          "so2": 10,
          "area": "北海",
          "o3": 100,
          "no2": 19,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 53,
          "pm10": 47,
          "pm2_5": 36,
          "o3_8h": 100,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "东北风",
        "temperature_time": "02:31",
        "wind_power": "2级",
        "aqi": 53,
        "sd": "67%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/night/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.525,
          "so2": 15,
          "area": "北海",
          "o3": 93,
          "no2": 20,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 58,
          "pm10": 51,
          "pm2_5": 40,
          "o3_8h": 96,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "03:00",
        "wind_power": "2级",
        "aqi": 58,
        "sd": "68%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/night/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.525,
          "so2": 15,
          "area": "北海",
          "o3": 93,
          "no2": 20,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 58,
          "pm10": 51,
          "pm2_5": 40,
          "o3_8h": 96,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "03:31",
        "wind_power": "1级",
        "aqi": 58,
        "sd": "68%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/night/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.525,
          "so2": 14,
          "area": "北海",
          "o3": 90,
          "no2": 19,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 54,
          "pm10": 46,
          "pm2_5": 37,
          "o3_8h": 94,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "04:01",
        "wind_power": "1级",
        "aqi": 54,
        "sd": "68%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/night/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.525,
          "so2": 14,
          "area": "北海",
          "o3": 90,
          "no2": 19,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 54,
          "pm10": 46,
          "pm2_5": 37,
          "o3_8h": 94,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "04:31",
        "wind_power": "2级",
        "aqi": 54,
        "sd": "66%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/night/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.55,
          "so2": 15,
          "area": "北海",
          "o3": 86,
          "no2": 21,
          "area_code": "beihai",
          "quality": "优",
          "aqi": 50,
          "pm10": 45,
          "pm2_5": 35,
          "o3_8h": 92,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "05:00",
        "wind_power": "2级",
        "aqi": 50,
        "sd": "67%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/night/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.55,
          "so2": 15,
          "area": "北海",
          "o3": 86,
          "no2": 21,
          "area_code": "beihai",
          "quality": "优",
          "aqi": 50,
          "pm10": 45,
          "pm2_5": 35,
          "o3_8h": 92,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "05:30",
        "wind_power": "2级",
        "aqi": 50,
        "sd": "67%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/night/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.55,
          "so2": 15,
          "area": "北海",
          "o3": 81,
          "no2": 20,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 53,
          "pm10": 43,
          "pm2_5": 36,
          "o3_8h": 90,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "东北风",
        "temperature_time": "06:00",
        "wind_power": "2级",
        "aqi": 53,
        "sd": "67%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/night/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.55,
          "so2": 15,
          "area": "北海",
          "o3": 81,
          "no2": 20,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 53,
          "pm10": 43,
          "pm2_5": 36,
          "o3_8h": 90,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "06:31",
        "wind_power": "2级",
        "aqi": 53,
        "sd": "67%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/night/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.525,
          "so2": 12,
          "area": "北海",
          "o3": 74,
          "no2": 23,
          "area_code": "beihai",
          "quality": "优",
          "aqi": 50,
          "pm10": 42,
          "pm2_5": 35,
          "o3_8h": 87,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "07:01",
        "wind_power": "2级",
        "aqi": 50,
        "sd": "67%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/night/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.525,
          "so2": 12,
          "area": "北海",
          "o3": 74,
          "no2": 23,
          "area_code": "beihai",
          "quality": "优",
          "aqi": 50,
          "pm10": 42,
          "pm2_5": 35,
          "o3_8h": 87,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "东北风",
        "temperature_time": "07:32",
        "wind_power": "2级",
        "aqi": 50,
        "sd": "68%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/night/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.55,
          "so2": 11,
          "area": "北海",
          "o3": 71,
          "no2": 23,
          "area_code": "beihai",
          "quality": "优",
          "aqi": 50,
          "pm10": 45,
          "pm2_5": 35,
          "o3_8h": 85,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "东北风",
        "temperature_time": "08:01",
        "wind_power": "1级",
        "aqi": 50,
        "sd": "67%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.55,
          "so2": 11,
          "area": "北海",
          "o3": 71,
          "no2": 23,
          "area_code": "beihai",
          "quality": "优",
          "aqi": 50,
          "pm10": 45,
          "pm2_5": 35,
          "o3_8h": 85,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "东北风",
        "temperature_time": "08:31",
        "wind_power": "2级",
        "aqi": 50,
        "sd": "66%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.575,
          "so2": 12,
          "area": "北海",
          "o3": 70,
          "no2": 23,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 51,
          "pm10": 50,
          "pm2_5": 34,
          "o3_8h": 83,
          "primary_pollutant": "颗粒物(PM10)"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "09:02",
        "wind_power": "1级",
        "aqi": 51,
        "sd": "65%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "17"
      },
      {
        "aqiDetail": {
          "co": 1.575,
          "so2": 12,
          "area": "北海",
          "o3": 70,
          "no2": 23,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 51,
          "pm10": 50,
          "pm2_5": 34,
          "o3_8h": 83,
          "primary_pollutant": "颗粒物(PM10)"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "09:32",
        "wind_power": "2级",
        "aqi": 51,
        "sd": "65%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "18"
      },
      {
        "aqiDetail": {
          "co": 1.575,
          "so2": 11,
          "area": "北海",
          "o3": 70,
          "no2": 23,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 54,
          "pm10": 56,
          "pm2_5": 36,
          "o3_8h": 80,
          "primary_pollutant": "颗粒物(PM10)"
        },
        "weather_code": "01",
        "wind_direction": "东北风",
        "temperature_time": "10:01",
        "wind_power": "2级",
        "aqi": 54,
        "sd": "64%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "18"
      },
      {
        "aqiDetail": {
          "co": 1.575,
          "so2": 11,
          "area": "北海",
          "o3": 70,
          "no2": 23,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 54,
          "pm10": 56,
          "pm2_5": 36,
          "o3_8h": 80,
          "primary_pollutant": "颗粒物(PM10)"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "10:34",
        "wind_power": "2级",
        "aqi": 54,
        "sd": "61%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "19"
      },
      {
        "aqiDetail": {
          "co": 1.575,
          "so2": 12,
          "area": "北海",
          "o3": 76,
          "no2": 22,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 56,
          "pm10": 51,
          "pm2_5": 39,
          "o3_8h": 77,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "11:01",
        "wind_power": "2级",
        "aqi": 56,
        "sd": "61%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "19"
      },
      {
        "aqiDetail": {
          "co": 1.575,
          "so2": 12,
          "area": "北海",
          "o3": 76,
          "no2": 22,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 56,
          "pm10": 51,
          "pm2_5": 39,
          "o3_8h": 77,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "11:32",
        "wind_power": "2级",
        "aqi": 56,
        "sd": "59%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "20"
      },
      {
        "aqiDetail": {
          "co": 1.575,
          "so2": 12,
          "area": "北海",
          "o3": 87,
          "no2": 20,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 59,
          "pm10": 57,
          "pm2_5": 41,
          "o3_8h": 77,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "12:01",
        "wind_power": "2级",
        "aqi": 59,
        "sd": "58%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "20"
      },
      {
        "aqiDetail": {
          "co": 1.575,
          "so2": 12,
          "area": "北海",
          "o3": 87,
          "no2": 20,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 59,
          "pm10": 57,
          "pm2_5": 41,
          "o3_8h": 77,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "12:34",
        "wind_power": "2级",
        "aqi": 59,
        "sd": "57%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "21"
      },
      {
        "aqiDetail": {
          "co": 1.575,
          "so2": 13,
          "area": "北海",
          "o3": 109,
          "no2": 20,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 64,
          "pm10": 70,
          "pm2_5": 45,
          "o3_8h": 80,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "北风",
        "temperature_time": "13:01",
        "wind_power": "3级",
        "aqi": 64,
        "sd": "56%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "21"
      },
      {
        "aqiDetail": {
          "co": 1.575,
          "so2": 13,
          "area": "北海",
          "o3": 109,
          "no2": 20,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 64,
          "pm10": 70,
          "pm2_5": 45,
          "o3_8h": 80,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "西北风",
        "temperature_time": "13:32",
        "wind_power": "1级",
        "aqi": 64,
        "sd": "56%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "22"
      },
      {
        "aqiDetail": {
          "co": 1.575,
          "so2": 13,
          "area": "北海",
          "o3": 109,
          "no2": 20,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 64,
          "pm10": 70,
          "pm2_5": 45,
          "o3_8h": 80,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "东北风",
        "temperature_time": "14:01",
        "wind_power": "2级",
        "aqi": 64,
        "sd": "56%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "22"
      },
      {
        "aqiDetail": {
          "co": 1.525,
          "so2": 12,
          "area": "北海",
          "o3": 134,
          "no2": 16,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 63,
          "pm10": 72,
          "pm2_5": 44,
          "o3_8h": 87,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "东南风",
        "temperature_time": "14:32",
        "wind_power": "2级",
        "aqi": 63,
        "sd": "53%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "24"
      },
      {
        "aqiDetail": {
          "co": 2.133,
          "so2": 12,
          "area": "北海",
          "o3": 151,
          "no2": 15,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 63,
          "pm10": 63,
          "pm2_5": 44,
          "o3_8h": 97,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "东南风",
        "temperature_time": "15:02",
        "wind_power": "2级",
        "aqi": 63,
        "sd": "55%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "24"
      },
      {
        "aqiDetail": {
          "co": 2.133,
          "so2": 12,
          "area": "北海",
          "o3": 151,
          "no2": 15,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 63,
          "pm10": 63,
          "pm2_5": 44,
          "o3_8h": 97,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "南风",
        "temperature_time": "15:31",
        "wind_power": "2级",
        "aqi": 63,
        "sd": "53%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "24"
      },
      {
        "aqiDetail": {
          "co": 2.133,
          "so2": 12,
          "area": "北海",
          "o3": 151,
          "no2": 15,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 63,
          "pm10": 63,
          "pm2_5": 44,
          "o3_8h": 97,
          "primary_pollutant": "颗粒物(PM2.5)"
        },
        "weather_code": "01",
        "wind_direction": "南风",
        "temperature_time": "16:02",
        "wind_power": "1级",
        "aqi": 63,
        "sd": "53%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "24"
      },
      {
        "aqiDetail": {
          "co": 1.6,
          "so2": 13,
          "area": "北海",
          "o3": 167,
          "no2": 16,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 60,
          "pm10": 64,
          "pm2_5": 42,
          "o3_8h": 106,
          "primary_pollutant": "颗粒物(PM2.5)臭氧1小时"
        },
        "weather_code": "01",
        "wind_direction": "东风",
        "temperature_time": "16:31",
        "wind_power": "2级",
        "aqi": 60,
        "sd": "55%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "24"
      },
      {
        "aqiDetail": {
          "co": 1.65,
          "so2": 17,
          "area": "北海",
          "o3": 183,
          "no2": 20,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 80,
          "pm10": 67,
          "pm2_5": 42,
          "o3_8h": 121,
          "primary_pollutant": "臭氧1小时"
        },
        "weather_code": "01",
        "wind_direction": "东南风",
        "temperature_time": "17:02",
        "wind_power": "2级",
        "aqi": 80,
        "sd": "54%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "24"
      },
      {
        "aqiDetail": {
          "co": 1.65,
          "so2": 17,
          "area": "北海",
          "o3": 183,
          "no2": 20,
          "area_code": "beihai",
          "quality": "良",
          "aqi": 80,
          "pm10": 67,
          "pm2_5": 42,
          "o3_8h": 121,
          "primary_pollutant": "臭氧1小时"
        },
        "weather_code": "01",
        "wind_direction": "东南风",
        "temperature_time": "17:32",
        "wind_power": "2级",
        "aqi": 80,
        "sd": "57%",
        "weather_pic": "http://appimg.showapi.com/images/weather/icon/day/01.png",
        "weather": "多云",
        "temperature": "23"
      }
    ],
    "f1": {
      "day_weather": "多云",
      "night_weather": "阴",
      "night_weather_code": "02",
      "index": {
        "yh": {
          "title": "较适宜",
          "desc": "不用担心天气来调皮捣乱而影响了兴致。 "
        },
        "ls": {
          "title": "不适宜",
          "desc": "天气阴沉，不太适宜晾晒。"
        },
        "clothes": {
          "title": "较舒适",
          "desc": "建议穿薄外套或牛仔裤等服装。"
        },
        "dy": {
          "title": "不适宜钓鱼",
          "desc": "天气不好，不适合垂钓"
        },
        "sports": {
          "title": "较适宜",
          "desc": "推荐您进行室内运动。"
        },
        "travel": {
          "title": "较适宜",
          "desc": "天热注意防晒，可选择水上娱乐项目。"
        },
        "beauty": {
          "title": "保湿",
          "desc": "请选用中性保湿型霜类化妆品。"
        },
        "xq": {
          "title": "较好",
          "desc": "好天气会带来一天的好心情。"
        },
        "hc": {
          "title": "较适宜",
          "desc": "风大，需注意及时添衣。"
        },
        "zs": {
          "title": "不容易中暑",
          "desc": "气温不高，中暑几率极低。"
        },
        "cold": {
          "title": "少发",
          "desc": "感冒机率较低，避免长期处于空调屋中。"
        },
        "gj": {
          "title": "较适宜",
          "desc": "风稍大，出门逛街前记得给秀发定型。"
        },
        "uv": {
          "title": "弱",
          "desc": "辐射较弱，涂擦SPF12-15、PA+护肤品。"
        },
        "cl": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "glass": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "aqi": {
          "title": "良",
          "desc": "气象条件有利于空气污染物扩散。"
        },
        "ac": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "wash_car": {
          "title": "适宜",
          "desc": "无雨且风力较小，易保持清洁度。"
        },
        "mf": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "ag": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "pj": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "nl": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "pk": {
          "title": "暂缺",
          "desc": "暂缺"
        }
      },
      "air_press": "1004 hPa",
      "jiangshui": "14%",
      "night_wind_power": "3-4级 5.5~7.9m/s",
      "day_wind_power": "3-4级 5.5~7.9m/s",
      "day_weather_code": "01",
      "3hourForcast": [
        {
          "wind_direction": "北风",
          "temperature_max": "18",
          "wind_power": "微风",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/day/01.png",
          "temperature_min": "16",
          "hour": "8时-11时",
          "temperature": "16"
        },
        {
          "wind_direction": "东北风",
          "temperature_max": "22",
          "wind_power": "3-4级",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/day/01.png",
          "temperature_min": "16",
          "hour": "11时-14时",
          "temperature": "18"
        },
        {
          "wind_direction": "东北风",
          "temperature_max": "22",
          "wind_power": "微风",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/day/01.png",
          "temperature_min": "18",
          "hour": "14时-17时",
          "temperature": "22"
        },
        {
          "wind_direction": "东北风",
          "temperature_max": "22",
          "wind_power": "微风",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/day/01.png",
          "temperature_min": "20",
          "hour": "17时-20时",
          "temperature": "22"
        },
        {
          "wind_direction": "东北风",
          "temperature_max": "22",
          "wind_power": "微风",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/night/01.png",
          "temperature_min": "18",
          "hour": "20时-23时",
          "temperature": "20"
        },
        {
          "wind_direction": "东风",
          "temperature_max": "20",
          "wind_power": "微风",
          "weather": "阴",
          "weather_pic": "http://app1.showapi.com/weather/icon/night/02.png",
          "temperature_min": "17",
          "hour": "23时-2时",
          "temperature": "18"
        },
        {
          "wind_direction": "东风",
          "temperature_max": "18",
          "wind_power": "微风",
          "weather": "阴",
          "weather_pic": "http://app1.showapi.com/weather/icon/night/02.png",
          "temperature_min": "17",
          "hour": "2时-5时",
          "temperature": "17"
        },
        {
          "wind_direction": "东风",
          "temperature_max": "17",
          "wind_power": "微风",
          "weather": "阴",
          "weather_pic": "http://app1.showapi.com/weather/icon/night/02.png",
          "temperature_min": "17",
          "hour": "5时-8时",
          "temperature": "17"
        }
      ],
      "sun_begin_end": "07:18|18:08",
      "ziwaixian": "弱",
      "day_weather_pic": "http://app1.showapi.com/weather/icon/day/01.png",
      "weekday": 5,
      "night_air_temperature": "17",
      "day_air_temperature": "23",
      "day_wind_direction": "东北风",
      "day": "20161223",
      "night_weather_pic": "http://app1.showapi.com/weather/icon/night/02.png",
      "night_wind_direction": "东风"
    },
    "f3": {
      "day_weather": "多云",
      "night_weather": "多云",
      "night_weather_code": "01",
      "index": {
        "yh": {
          "title": "较适宜",
          "desc": "天气较好，适宜约会"
        },
        "ls": {
          "title": "不适宜",
          "desc": "天气阴沉，不太适宜晾晒。"
        },
        "clothes": {
          "title": "舒适",
          "desc": "建议穿长袖衬衫单裤等服装。"
        },
        "dy": {
          "title": "不适宜钓鱼",
          "desc": "天气不好，不适合垂钓"
        },
        "sports": {
          "title": "较适宜",
          "desc": "推荐您进行室内运动。"
        },
        "travel": {
          "title": "较不宜",
          "desc": "天气很热，如外出可选择水上娱乐项目。"
        },
        "beauty": {
          "title": "保湿",
          "desc": "请选用中性保湿型霜类化妆品。"
        },
        "xq": {
          "title": "较好",
          "desc": "晴朗天气，将为您的身心带来温暖。"
        },
        "hc": {
          "title": "较适宜",
          "desc": "风大，需注意及时添衣。"
        },
        "zs": {
          "title": "不容易中暑",
          "desc": "气温不高，中暑几率极低。"
        },
        "cold": {
          "title": "少发",
          "desc": "感冒机率较低，避免长期处于空调屋中。"
        },
        "gj": {
          "title": "较适宜",
          "desc": "风稍大，出门逛街前记得给秀发定型。"
        },
        "uv": {
          "title": "弱",
          "desc": "辐射较弱，涂擦SPF12-15、PA+护肤品。"
        },
        "cl": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "glass": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "aqi": {
          "title": "良",
          "desc": "气象条件有利于空气污染物扩散。"
        },
        "ac": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "wash_car": {
          "title": "适宜",
          "desc": "无雨且风力较小，易保持清洁度。"
        },
        "mf": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "ag": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "pj": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "nl": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "pk": {
          "title": "暂缺",
          "desc": "暂缺"
        }
      },
      "air_press": "1004 hPa",
      "jiangshui": "16%",
      "night_wind_power": "3-4级 5.5~7.9m/s",
      "day_wind_power": "3-4级 5.5~7.9m/s",
      "day_weather_code": "01",
      "3hourForcast": [
        {
          "wind_direction": "东风",
          "temperature_max": "23",
          "wind_power": "3-4级",
          "weather": "小雨",
          "weather_pic": "http://app1.showapi.com/weather/icon/day/07.png",
          "temperature_min": "18",
          "hour": "8时-14时",
          "temperature": "19"
        },
        {
          "wind_direction": "东风",
          "temperature_max": "23",
          "wind_power": "3-4级",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/day/01.png",
          "temperature_min": "19",
          "hour": "14时-20时",
          "temperature": "23"
        },
        {
          "wind_direction": "东风",
          "temperature_max": "23",
          "wind_power": "微风",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/night/01.png",
          "temperature_min": "20",
          "hour": "20时-2时",
          "temperature": "21"
        },
        {
          "wind_direction": "东风",
          "temperature_max": "21",
          "wind_power": "3-4级",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/night/01.png",
          "temperature_min": "20",
          "hour": "2时-8时",
          "temperature": "20"
        }
      ],
      "sun_begin_end": "07:19|18:09",
      "ziwaixian": "弱",
      "day_weather_pic": "http://app1.showapi.com/weather/icon/day/01.png",
      "weekday": 7,
      "night_air_temperature": "20",
      "day_air_temperature": "24",
      "day_wind_direction": "东风",
      "day": "20161225",
      "night_weather_pic": "http://app1.showapi.com/weather/icon/night/01.png",
      "night_wind_direction": "东风"
    },
    "f2": {
      "day_weather": "小雨",
      "night_weather": "小雨",
      "night_weather_code": "07",
      "index": {
        "yh": {
          "title": "较适宜",
          "desc": "不用担心天气来调皮捣乱而影响了兴致。 "
        },
        "ls": {
          "title": "不适宜",
          "desc": "有降水会淋湿衣物，不适宜晾晒。"
        },
        "clothes": {
          "title": "较舒适",
          "desc": "建议穿薄外套或牛仔裤等服装。"
        },
        "dy": {
          "title": "不适宜钓鱼",
          "desc": "天气不好，不适合垂钓"
        },
        "sports": {
          "title": "较不宜",
          "desc": "有降水,推荐您在室内进行休闲运动。"
        },
        "travel": {
          "title": "较不宜",
          "desc": "天气很热，如外出可选择水上娱乐项目。"
        },
        "beauty": {
          "title": "保湿",
          "desc": "请选用保湿型霜类化妆品。"
        },
        "xq": {
          "title": "较差",
          "desc": "天气阴沉，会感觉压抑，情绪低落。"
        },
        "hc": {
          "title": "较适宜",
          "desc": "温度适宜，注意着凉。"
        },
        "zs": {
          "title": "不容易中暑",
          "desc": "气温不高，中暑几率极低。"
        },
        "cold": {
          "title": "易发",
          "desc": "天冷湿度大，注意增加衣服。"
        },
        "gj": {
          "title": "较不宜",
          "desc": "有较强降水，坚持出门需带雨具。"
        },
        "uv": {
          "title": "弱",
          "desc": "辐射较弱，涂擦SPF12-15、PA+护肤品。"
        },
        "cl": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "glass": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "aqi": {
          "title": "良",
          "desc": "气象条件有利于空气污染物扩散。"
        },
        "ac": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "wash_car": {
          "title": "不适宜",
          "desc": "有雨，雨水和泥水会弄脏爱车。"
        },
        "mf": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "ag": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "pj": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "nl": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "pk": {
          "title": "暂缺",
          "desc": "暂缺"
        }
      },
      "air_press": "1004 hPa",
      "jiangshui": "82%",
      "night_wind_power": "3-4级 5.5~7.9m/s",
      "day_wind_power": "3-4级 5.5~7.9m/s",
      "day_weather_code": "07",
      "3hourForcast": [
        {
          "wind_direction": "东风",
          "temperature_max": "18",
          "wind_power": "3-4级",
          "weather": "阴",
          "weather_pic": "http://app1.showapi.com/weather/icon/day/02.png",
          "temperature_min": "17",
          "hour": "8时-11时",
          "temperature": "17"
        },
        {
          "wind_direction": "东风",
          "temperature_max": "22",
          "wind_power": "微风",
          "weather": "小雨",
          "weather_pic": "http://app1.showapi.com/weather/icon/day/07.png",
          "temperature_min": "17",
          "hour": "11时-14时",
          "temperature": "18"
        },
        {
          "wind_direction": "东风",
          "temperature_max": "22",
          "wind_power": "微风",
          "weather": "小雨",
          "weather_pic": "http://app1.showapi.com/weather/icon/day/07.png",
          "temperature_min": "18",
          "hour": "14时-17时",
          "temperature": "22"
        },
        {
          "wind_direction": "东风",
          "temperature_max": "22",
          "wind_power": "微风",
          "weather": "小雨",
          "weather_pic": "http://app1.showapi.com/weather/icon/day/07.png",
          "temperature_min": "18",
          "hour": "17时-20时",
          "temperature": "21"
        },
        {
          "wind_direction": "东风",
          "temperature_max": "21",
          "wind_power": "3-4级",
          "weather": "小雨",
          "weather_pic": "http://app1.showapi.com/weather/icon/night/07.png",
          "temperature_min": "18",
          "hour": "20时-23时",
          "temperature": "18"
        },
        {
          "wind_direction": "东风",
          "temperature_max": "18",
          "wind_power": "微风",
          "weather": "小雨",
          "weather_pic": "http://app1.showapi.com/weather/icon/night/07.png",
          "temperature_min": "17",
          "hour": "23时-2时",
          "temperature": "18"
        },
        {
          "wind_direction": "东风",
          "temperature_max": "18",
          "wind_power": "微风",
          "weather": "小雨",
          "weather_pic": "http://app1.showapi.com/weather/icon/night/07.png",
          "temperature_min": "17",
          "hour": "2时-5时",
          "temperature": "17"
        },
        {
          "wind_direction": "东风",
          "temperature_max": "19",
          "wind_power": "微风",
          "weather": "小雨",
          "weather_pic": "http://app1.showapi.com/weather/icon/night/07.png",
          "temperature_min": "17",
          "hour": "5时-8时",
          "temperature": "18"
        }
      ],
      "sun_begin_end": "07:19|18:08",
      "ziwaixian": "弱",
      "day_weather_pic": "http://app1.showapi.com/weather/icon/day/07.png",
      "weekday": 6,
      "night_air_temperature": "17",
      "day_air_temperature": "23",
      "day_wind_direction": "东风",
      "day": "20161224",
      "night_weather_pic": "http://app1.showapi.com/weather/icon/night/07.png",
      "night_wind_direction": "东风"
    },
    "f5": {
      "day_weather": "多云",
      "night_weather": "多云",
      "night_weather_code": "01",
      "index": {
        "yh": {
          "title": "较不适宜",
          "desc": "风力较大，建议尽量不要去室外约会。"
        },
        "ls": {
          "title": "不适宜",
          "desc": "天气阴沉，不太适宜晾晒。"
        },
        "clothes": {
          "title": "较舒适",
          "desc": "建议穿薄外套或牛仔裤等服装。"
        },
        "dy": {
          "title": "不适宜钓鱼",
          "desc": "天气不好，不适合垂钓"
        },
        "sports": {
          "title": "较适宜",
          "desc": "推荐您进行室内运动。"
        },
        "travel": {
          "title": "较不宜",
          "desc": "天气很热，如外出可选择水上娱乐项目。"
        },
        "beauty": {
          "title": "保湿",
          "desc": "请选用中性保湿型霜类化妆品。"
        },
        "xq": {
          "title": "较好",
          "desc": "温度适宜，心情会不错。"
        },
        "hc": {
          "title": "较适宜",
          "desc": "风大，需注意及时添衣。"
        },
        "zs": {
          "title": "不容易中暑",
          "desc": "气温不高，中暑几率极低。"
        },
        "cold": {
          "title": "少发",
          "desc": "无明显降温，感冒机率较低。"
        },
        "gj": {
          "title": "较适宜",
          "desc": "风稍大，出门逛街前记得给秀发定型。"
        },
        "uv": {
          "title": "弱",
          "desc": "辐射较弱，涂擦SPF12-15、PA+护肤品。"
        },
        "cl": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "glass": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "aqi": {
          "title": "良",
          "desc": "气象条件有利于空气污染物扩散。"
        },
        "ac": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "wash_car": {
          "title": "较不适宜",
          "desc": "风力较大，洗车后会蒙上灰尘。"
        },
        "mf": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "ag": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "pj": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "nl": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "pk": {
          "title": "暂缺",
          "desc": "暂缺"
        }
      },
      "air_press": "1004 hPa",
      "jiangshui": "12%",
      "night_wind_power": "4-5级 8.0~10.7m/s",
      "day_wind_power": "4-5级 8.0~10.7m/s",
      "day_weather_code": "01",
      "3hourForcast": [
        {
          "wind_direction": "北风",
          "temperature_max": "20",
          "wind_power": "4-5级",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/day/01.png",
          "temperature_min": "14",
          "hour": "8时-14时",
          "temperature": "14"
        },
        {
          "wind_direction": "北风",
          "temperature_max": "20",
          "wind_power": "4-5级",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/day/01.png",
          "temperature_min": "14",
          "hour": "14时-20时",
          "temperature": "20"
        },
        {
          "wind_direction": "北风",
          "temperature_max": "20",
          "wind_power": "3-4级",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/night/01.png",
          "temperature_min": "12",
          "hour": "20时-2时",
          "temperature": "16"
        },
        {
          "wind_direction": "北风",
          "temperature_max": "16",
          "wind_power": "4-5级",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/night/01.png",
          "temperature_min": "12",
          "hour": "2时-8时",
          "temperature": "12"
        }
      ],
      "sun_begin_end": "07:20|18:10",
      "ziwaixian": "弱",
      "day_weather_pic": "http://app1.showapi.com/weather/icon/day/01.png",
      "weekday": 2,
      "night_air_temperature": "11",
      "day_air_temperature": "21",
      "day_wind_direction": "北风",
      "day": "20161227",
      "night_weather_pic": "http://app1.showapi.com/weather/icon/night/01.png",
      "night_wind_direction": "北风"
    },
    "f4": {
      "day_weather": "多云",
      "night_weather": "多云",
      "night_weather_code": "01",
      "index": {
        "yh": {
          "title": "较适宜",
          "desc": "天气较好，适宜约会"
        },
        "ls": {
          "title": "不适宜",
          "desc": "天气阴沉，不太适宜晾晒。"
        },
        "clothes": {
          "title": "舒适",
          "desc": "建议穿长袖衬衫单裤等服装。"
        },
        "dy": {
          "title": "不适宜钓鱼",
          "desc": "天气不好，不适合垂钓"
        },
        "sports": {
          "title": "较适宜",
          "desc": "推荐您进行室内运动。"
        },
        "travel": {
          "title": "较适宜",
          "desc": "天热注意防晒，可选择水上娱乐项目。"
        },
        "beauty": {
          "title": "去油",
          "desc": "请选用露质面霜打底，水质无油粉底霜。"
        },
        "xq": {
          "title": "好",
          "desc": "温度适宜，心情会不错。"
        },
        "hc": {
          "title": "较适宜",
          "desc": "风稍大会对划船产生一定影响。"
        },
        "zs": {
          "title": "不容易中暑",
          "desc": "气温不高，中暑几率极低。"
        },
        "cold": {
          "title": "少发",
          "desc": "感冒机率较低，避免长期处于空调屋中。"
        },
        "gj": {
          "title": "较适宜",
          "desc": "风稍大，出门逛街前记得给秀发定型。"
        },
        "uv": {
          "title": "弱",
          "desc": "辐射较弱，涂擦SPF12-15、PA+护肤品。"
        },
        "cl": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "glass": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "aqi": {
          "title": "良",
          "desc": "气象条件有利于空气污染物扩散。"
        },
        "ac": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "wash_car": {
          "title": "适宜",
          "desc": "无雨且风力较小，易保持清洁度。"
        },
        "mf": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "ag": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "pj": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "nl": {
          "title": "暂缺",
          "desc": "暂缺"
        },
        "pk": {
          "title": "暂缺",
          "desc": "暂缺"
        }
      },
      "air_press": "1004 hPa",
      "jiangshui": "16%",
      "night_wind_power": "4-5级 8.0~10.7m/s",
      "day_wind_power": "3-4级 5.5~7.9m/s",
      "day_weather_code": "01",
      "3hourForcast": [
        {
          "wind_direction": "东风",
          "temperature_max": "24",
          "wind_power": "微风",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/day/01.png",
          "temperature_min": "20",
          "hour": "8时-14时",
          "temperature": "20"
        },
        {
          "wind_direction": "东风",
          "temperature_max": "24",
          "wind_power": "微风",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/day/01.png",
          "temperature_min": "20",
          "hour": "14时-20时",
          "temperature": "24"
        },
        {
          "wind_direction": "东风",
          "temperature_max": "24",
          "wind_power": "3-4级",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/night/01.png",
          "temperature_min": "15",
          "hour": "20时-2时",
          "temperature": "22"
        },
        {
          "wind_direction": "北风",
          "temperature_max": "22",
          "wind_power": "3-4级",
          "weather": "多云",
          "weather_pic": "http://app1.showapi.com/weather/icon/night/01.png",
          "temperature_min": "14",
          "hour": "2时-8时",
          "temperature": "15"
        }
      ],
      "sun_begin_end": "07:19|18:09",
      "ziwaixian": "弱",
      "day_weather_pic": "http://app1.showapi.com/weather/icon/day/01.png",
      "weekday": 1,
      "night_air_temperature": "14",
      "day_air_temperature": "25",
      "day_wind_direction": "东风",
      "day": "20161226",
      "night_weather_pic": "http://app1.showapi.com/weather/icon/night/01.png",
      "night_wind_direction": "北风"
    }
  }
}
`
