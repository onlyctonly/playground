package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)
type Weather struct {
	daily map[string]interface{} `json:"daily"`
}
func main() {
	res,err:=http.Get("https://api.darksky.net/forecast/7643e3b3670839205cd300233d5e24f2/37.9908997,23.7033199?lang=zh&units=si")
	check(err)
	w,err:=ioutil.ReadAll(res.Body)
	check(err)
	//fmt.Println(string(w))
	var ws map[string]interface{}
	json.Unmarshal(w, &ws)
	fmt.Println(`雅典天气: `, ws["daily"].(map[string]interface{})["summary"])
	fmt.Println(`最低气温: `, ws["daily"].(map[string]interface{})["data"].([]interface{})[0].(map[string]interface{})["temperatureMin"])
	fmt.Println(`最高气温: `, ws["daily"].(map[string]interface{})["data"].([]interface{})[0].(map[string]interface{})["temperatureMax"])

}

func check(e error)  {
	if e!=nil{
		fmt.Println(e)
	}
}
//https://api.darksky.net/forecast/7643e3b3670839205cd300233d5e24f2/37.9908997,23.7033199?lang=zh&units=si
//ws["daily"]    map[summary:小雨持续至明天，周四，且周日温度剧增到30°C。 icon:rain data:[map[icon:partly-cloudy-night precipIntensity:0 temperatureMax:28.14 dewPoint:14.38 windBearing:125 time:1.4966964e+09 moonPhase:0.4 apparentTemperatureMinTime:1.4967144e+09 apparentTemperatureMax:28.38 humidity:0.58 pressure:1012.95 precipProbability:0 temperatureMin:18.93 apparentTemperatureMin:18.93 apparentTemperatureMaxTime:1.4967396e+09 cloudCover:0.07 windSpeed:0.58 ozone:324.08 summary:局部多云持续至当晚。 sunriseTime:1.496718258e+09 sunsetTime:1.496771178e+09 precipIntensityMax:0 temperatureMinTime:1.4967144e+09 temperatureMaxTime:1.4967396e+09] map[ozone:329.19 precipIntensityMax:0.2083 temperatureMaxTime:1.4968368e+09 apparentTemperatureMaxTime:1.4968368e+09 humidity:0.67 windBearing:215 cloudCover:0.54 icon:rain precipIntensity:0.0559 precipType:rain temperatureMax:28.16 apparentTemperatureMinTime:1.4968008e+09 apparentTemperatureMin:17.62 apparentTemperatureMax:28.54 dewPoint:15.29 sunriseTime:1.496804645e+09 sunsetTime:1.496857612e+09 precipIntensityMaxTime:1.4968476e+09 precipProbability:0.42 temperatureMin:17.62 pressure:1013.17 time:1.4967828e+09 summary:毛毛雨持续至晚上。 moonPhase:0.43 temperatureMinTime:1.4968008e+09 windSpeed:0.7] map[sunsetTime:1.496944045e+09 precipIntensityMax:0.7112 temperatureMax:29.53 apparentTemperatureMin:18.29 apparentTemperatureMinTime:1.4968872e+09 apparentTemperatureMaxTime:1.4969268e+09 icon:rain sunriseTime:1.496891034e+09 precipType:rain apparentTemperatureMax:30.22 dewPoint:15.61 cloudCover:0.54 pressure:1012.25 precipIntensity:0.2159 precipIntensityMaxTime:1.4968944e+09 temperatureMinTime:1.4968872e+09 humidity:0.63 windSpeed:1.4 windBearing:348 ozone:339.49 summary:小雨持续至早上，晚上。 temperatureMin:18.29 precipProbability:0.65 temperatureMaxTime:1.4969268e+09 time:1.4968692e+09 moonPhase:0.46] map[icon:clear-day sunriseTime:1.496977425e+09 sunsetTime:1.497030477e+09 precipIntensityMaxTime:1.4969736e+09 temperatureMin:19.78 temperatureMax:29.07 apparentTemperatureMaxTime:1.4970132e+09 apparentTemperatureMin:19.78 humidity:0.59 ozone:320.92 windBearing:8 time:1.4969556e+09 moonPhase:0.49 precipIntensityMax:0.0508 precipProbability:0.05 precipType:rain apparentTemperatureMax:29.75 windSpeed:3.4 cloudCover:0.02 pressure:1014.78 summary:晴朗将持续一整天。 precipIntensity:0.033 temperatureMinTime:1.4969772e+09 temperatureMaxTime:1.4970132e+09 apparentTemperatureMinTime:1.4969772e+09 dewPoint:15.42] map[temperatureMin:18.69 apparentTemperatureMin:18.69 apparentTemperatureMinTime:1.49706e+09 ozone:319.34 temperatureMinTime:1.49706e+09 temperatureMaxTime:1.4970996e+09 dewPoint:14.22 time:1.497042e+09 summary:晴朗将持续一整天。 icon:clear-day precipIntensityMaxTime:1.497042e+09 precipProbability:0.01 apparentTemperatureMaxTime:1.4970996e+09 windSpeed:0.83 windBearing:98 sunriseTime:1.497063817e+09 sunsetTime:1.497116907e+09 precipIntensity:0.0051 precipIntensityMax:0.0229 pressure:1016.01 humidity:0.57 cloudCover:0.03 moonPhase:0.53 precipType:rain temperatureMax:29.02 apparentTemperatureMax:29.07] map[precipIntensity:0.0152 precipProbability:0.03 precipIntensityMaxTime:1.4972112e+09 precipType:rain apparentTemperatureMax:29.88 apparentTemperatureMaxTime:1.497186e+09 dewPoint:14.64 windBearing:356 icon:partly-cloudy-night moonPhase:0.56 temperatureMaxTime:1.497186e+09 apparentTemperatureMin:18.9 apparentTemperatureMinTime:1.4971464e+09 windSpeed:1.45 cloudCover:0.04 pressure:1011.89 ozone:324.25 time:1.4971284e+09 summary:多云持续至当晚。 sunriseTime:1.497150211e+09 sunsetTime:1.497203336e+09 precipIntensityMax:0.0381 temperatureMin:18.9 temperatureMinTime:1.4971464e+09 temperatureMax:30.17 humidity:0.57] map[temperatureMin:18.33 apparentTemperatureMaxTime:1.4972724e+09 windBearing:11 precipProbability:0.04 moonPhase:0.59 icon:partly-cloudy-night sunsetTime:1.497289763e+09 precipIntensityMax:0.0457 precipIntensityMaxTime:1.4972904e+09 temperatureMinTime:1.4972364e+09 temperatureMax:26.66 temperatureMaxTime:1.4972724e+09 time:1.4972148e+09 apparentTemperatureMinTime:1.4972364e+09 apparentTemperatureMax:26.66 humidity:0.53 windSpeed:4.86 cloudCover:0.21 pressure:1010.65 ozone:329.59 apparentTemperatureMin:18.33 sunriseTime:1.497236607e+09 precipIntensity:0.0356 precipType:rain dewPoint:12.14 summary:局部多云持续至早上。] map[moonPhase:0.62 temperatureMinTime:1.4973192e+09 dewPoint:13.32 pressure:1011.41 ozone:317.73 time:1.4973012e+09 sunriseTime:1.497323005e+09 temperatureMaxTime:1.4973552e+09 cloudCover:0 sunsetTime:1.497376189e+09 temperatureMax:29.17 temperatureMin:16.82 apparentTemperatureMaxTime:1.4973588e+09 humidity:0.56 windSpeed:0.43 precipIntensity:0.0279 precipIntensityMax:0.0381 precipIntensityMaxTime:1.4973012e+09 precipProbability:0.03 precipType:rain apparentTemperatureMin:16.82 apparentTemperatureMinTime:1.4973192e+09 apparentTemperatureMax:28.91 summary:晴朗将持续一整天。 icon:clear-day windBearing:23]]]