/**
 * Created by WillkYang on 2017/3/6.
 */
package yeego

import (
	"fmt"
	"testing"
)

func mockJsonData() *Json {
	json := InitJson(`{"ret":0,"msg":"ok","projects":[{"id":302,"name":"测试0109","startTime":1483957800000,"endTime":1485426600000,"tag":"做市","type":"企"},{"id":245,"name":"测试融资项目513-1","startTime":1463108700000,"endTime":1464710100000,"tag":"做市","type":"项"},{"id":244,"name":"测试0512-05","startTime":1463036700000,"endTime":1463286000000,"tag":"做市","type":"项"},{"id":222,"name":"做市企业509-1","startTime":1462755600000,"endTime":1462859400000,"tag":"做市","type":"企"},{"id":201,"name":"测试项目505-1","startTime":1462518000000,"endTime":1462955700000,"tag":"定增","type":"项"},{"id":41,"name":"蘑菇财富X项目123","startTime":1452842100000,"endTime":1453533300000,"tag":"做市","type":"企"}],"stocks":[{"id":"NQ839815","zqdm":"839815","zqjc":"双承生物","ywjc":"","pyjc":"SCSW","scdm":"ngpgsgp","gpsj":20161121},{"id":"NQ839817","zqdm":"839817","zqjc":"力德气体","ywjc":"","pyjc":"LDQT:LDQB","scdm":"ngpgsgp","gpsj":20161116},{"id":"NQ839821","zqdm":"839821","zqjc":"铭博股份","ywjc":"","pyjc":"MBGF","scdm":"ngpgsgp","gpsj":20161109},{"id":"NQ839831","zqdm":"839831","zqjc":"明及电气","ywjc":"","pyjc":"MJDQ","scdm":"ngpgsgp","gpsj":20161118},{"id":"NQ839841","zqdm":"839841","zqjc":"瑞尔泰","ywjc":"","pyjc":"RET","scdm":"ngpgsgp","gpsj":20161116},{"id":"NQ839851","zqdm":"839851","zqjc":"珠峰电气","ywjc":"","pyjc":"ZFDQ","scdm":"ngpgsgp","gpsj":20161114},{"id":"NQ839661","zqdm":"839661","zqjc":"医微讯","ywjc":"","pyjc":"YWX","scdm":"ngpgsgp","gpsj":20161111},{"id":"NQ839671","zqdm":"839671","zqjc":"德利福","ywjc":"","pyjc":"DLF","scdm":"ngpgsgp","gpsj":20161123},{"id":"NQ839681","zqdm":"839681","zqjc":"宝涞精工","ywjc":"","pyjc":"BLJG","scdm":"ngpgsgp","gpsj":20161122},{"id":"NQ839701","zqdm":"839701","zqjc":"桑格尔","ywjc":"","pyjc":"SGE","scdm":"ngpgsgp","gpsj":20161124}]}`)
	return json
}

func mockJsonArrData() *Json {
	json := InitJson(`[{"id":302,"name":"测试0109","startTime":1483957800000,"endTime":1485426600000,"tag":"做市","type":"企"},{"id":245,"name":"测试融资项目513-1","startTime":1463108700000,"endTime":1464710100000,"tag":"做市","type":"项"},{"id":244,"name":"测试0512-05","startTime":1463036700000,"endTime":1463286000000,"tag":"做市","type":"项"},{"id":222,"name":"做市企业509-1","startTime":1462755600000,"endTime":1462859400000,"tag":"做市","type":"企"},{"id":201,"name":"测试项目505-1","startTime":1462518000000,"endTime":1462955700000,"tag":"定增","type":"项"},{"id":41,"name":"蘑菇财富X项目123","startTime":1452842100000,"endTime":1453533300000,"tag":"做市","type":"企"}]`)
	return json
}

func TestInit(t *testing.T) {
	json := InitJson(`{"ret":0,"msg":"ok","projects":[{"id":302,"name":"测试0109","startTime":1483957800000,"endTime":1485426600000,"tag":"做市","type":"企"},{"id":245,"name":"测试融资项目513-1","startTime":1463108700000,"endTime":1464710100000,"tag":"做市","type":"项"},{"id":244,"name":"测试0512-05","startTime":1463036700000,"endTime":1463286000000,"tag":"做市","type":"项"},{"id":222,"name":"做市企业509-1","startTime":1462755600000,"endTime":1462859400000,"tag":"做市","type":"企"},{"id":201,"name":"测试项目505-1","startTime":1462518000000,"endTime":1462955700000,"tag":"定增","type":"项"},{"id":41,"name":"蘑菇财富X项目123","startTime":1452842100000,"endTime":1453533300000,"tag":"做市","type":"企"}],"stocks":[{"id":"NQ839815","zqdm":"839815","zqjc":"双承生物","ywjc":"","pyjc":"SCSW","scdm":"ngpgsgp","gpsj":20161121},{"id":"NQ839817","zqdm":"839817","zqjc":"力德气体","ywjc":"","pyjc":"LDQT:LDQB","scdm":"ngpgsgp","gpsj":20161116},{"id":"NQ839821","zqdm":"839821","zqjc":"铭博股份","ywjc":"","pyjc":"MBGF","scdm":"ngpgsgp","gpsj":20161109},{"id":"NQ839831","zqdm":"839831","zqjc":"明及电气","ywjc":"","pyjc":"MJDQ","scdm":"ngpgsgp","gpsj":20161118},{"id":"NQ839841","zqdm":"839841","zqjc":"瑞尔泰","ywjc":"","pyjc":"RET","scdm":"ngpgsgp","gpsj":20161116},{"id":"NQ839851","zqdm":"839851","zqjc":"珠峰电气","ywjc":"","pyjc":"ZFDQ","scdm":"ngpgsgp","gpsj":20161114},{"id":"NQ839661","zqdm":"839661","zqjc":"医微讯","ywjc":"","pyjc":"YWX","scdm":"ngpgsgp","gpsj":20161111},{"id":"NQ839671","zqdm":"839671","zqjc":"德利福","ywjc":"","pyjc":"DLF","scdm":"ngpgsgp","gpsj":20161123},{"id":"NQ839681","zqdm":"839681","zqjc":"宝涞精工","ywjc":"","pyjc":"BLJG","scdm":"ngpgsgp","gpsj":20161122},{"id":"NQ839701","zqdm":"839701","zqjc":"桑格尔","ywjc":"","pyjc":"SGE","scdm":"ngpgsgp","gpsj":20161124}]}`)
	NotEqual(json.Data, nil)
}

func TestJson_Get(t *testing.T) {
	json := mockJsonData()
	fmt.Print(json.Get("projects"))
	NotEqual(json.Get("projects"), nil)
}

func TestJson_GetData(t *testing.T) {
	json := mockJsonData()
	mapData := json.GetData()
	fmt.Print(mapData)
	Equal(mapData["ret"], 0.)
	Equal(mapData["msg"], "ok")
	NotEqual(json.Get("projects"), nil)
}

func TestJson_GetIndex(t *testing.T) {
	json := mockJsonData()
	json1 := json.GetIndex(1)
	NotEqual(json1.Data, nil)
}

func TestJson_GetKey(t *testing.T) {
	json := mockJsonArrData()
	Equal(json.GetKey("id", 1).Data, 302.)
}

func TestJson_GetPath(t *testing.T) {
	json := InitJson(`{"projects":{"id":302,"name":"测试0109","startTime":1483957800000,"endTime":1485426600000,"tag":"做市","type":"企"}}`)
	NotEqual(json.GetPath("projects", "id").Data, nil)
}

func TestJson_ArrayIndex(t *testing.T) {
	json := InitJson(`["data1","data2"]`)
	Equal(json.ArrayIndex(1), "data1")
}

func TestJson_ToData(t *testing.T) {
	json := mockJsonData()
	NotEqual(json, nil)
}

func TestJson_ToSlice(t *testing.T) {
	json := mockJsonArrData()
	data := json.ToSlice()
	NotEqual(data, nil)
	Equal(len(data), 6)
}

func TestJson_ToInt(t *testing.T) {
	json := InitJson(`{"value":1}`)
	Equal(json.Get("value").ToInt(), 1)
}

func TestJson_ToFloat(t *testing.T) {
	json := InitJson(`{"value":1}`)
	Equal(json.Get("value").ToFloat(), 1.)
}

func TestJson_ToString(t *testing.T) {
	json := InitJson(`{"key":"value"}`)
	Equal(json.Get("key").ToString(), "value")
}

func TestJson_ToArray(t *testing.T) {
	json := mockJsonData()
	k, v := json.ToArray()
	fmt.Println(k, v)
	NotEqual(k, nil)
	NotEqual(v, nil)

	json1 := mockJsonArrData()
	k1, v1 := json1.ToArray()
	NotEqual(k1, nil)
	NotEqual(v1, nil)
}

func TestJson_StringToArray(t *testing.T) {
	json := InitJson(`["data1","data2"]`)
	data := json.StringToArray()
	fmt.Println(data)
	NotEqual(data, nil)
}