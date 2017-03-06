/**
 * Created by WillkYang on 2017/3/6.
 */
package yeego

import (
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
)

type Json struct {
	Data interface{}
}

func Init(data string) *Json {
	j := new(Json)
	var jsonData interface{}
	if err := json.Unmarshal([]byte(data), &jsonData); err != nil {
		return j
	}
	j.Data = jsonData
	return j
}

//return json data by key
func (j *Json) Get(key string) *Json {
	mapData := j.GetData()
	if value, ok := mapData[key]; ok {
		j.Data = value
		return j
	}
	j.Data = nil
	return j
}

//return JsonData in map type
func (j *Json) GetData() map[string]interface{} {
	if mapData, ok := (j.Data).(map[string]interface{}); ok {
		return mapData
	}
	return nil
}

//return Json of index n; warning that index is no unique
func (j *Json) GetIndex(index int) *Json {
	num := index - 1
	if mapData, ok := (j.Data).([]interface{}); ok {
		j.Data = mapData[index]
		return j
	}

	if mapData, ok := (j.Data).(map[string]interface{}); ok {
		n := 0
		data := make(map[string]interface{})

		for k, v := range mapData {
			if n == num {
				switch vv := v.(type) {
				case float64:
					data[k] = strconv.FormatFloat(vv, 'f', -1, 64)
					j.Data = data
					return j
				case string:
					data[k] = vv
					j.Data = data
					return j
				case []interface{}:
					j.Data = vv
					return j
				}
			}
			n++
		}
	}
	j.Data = nil
	return j
}

//return json of arr and key;  json data type must be [](map[string]interface{})
func (j *Json) GetKey(key string, index int) *Json {
	num := index - 1
	arrData, ok := (j.Data).([]interface{})
	if !ok {
		j.Data = errors.New("invalid json data type").Error()
		return j
	} else if index > len(arrData) {
		j.Data = errors.New("index out of range list").Error()
		return j
	}

	if v, ok := arrData[num].(map[string]interface{}); ok {
		if vv, ok := v[key]; ok {
			j.Data = vv
			return j
		}
	}
	j.Data = nil
	return j
}

//return json by multi key
func (j *Json) GetPath(args ...string) *Json {
	d := j
	for _, v := range args {
		mapData := d.GetData()
		if vv, ok := mapData[v]; ok {
			d.Data = vv
		} else {
			d.Data = nil
			return d
		}
	}
	return d
}

//return string index of json; json data type must be []interface{}
func (j *Json) ArrayIndex(index int) string {
	num := index - 1
	arrData, ok := (j.Data).([]interface{})

	if !ok {
		return errors.New("invalid json data type").Error()
	} else if index > len(arrData) {
		return errors.New("index out of range list").Error()
	}

	v := arrData[num]
	switch vv := v.(type) {
	case float64:
		return strconv.FormatFloat(vv, 'f', -1, 64)
	case string:
		return vv
	default:
		return ""
	}
}

func (j *Json) ToData() interface{} {
	return j.Data
}

func (j *Json) ToSlice() []interface{} {
	if arrData, ok := (j.Data).([]interface{}); ok {
		return arrData
	}
	return nil
}

func (j *Json) ToInt() int {
	if data, ok := (j.Data).(int); ok {
		return data
	}
	if data, ok := (j.Data).(float64); ok {
		return int(data)
	}
	if data, ok := (j.Data).(string); ok {
		if data, err := strconv.ParseInt(data, 10, 64); err != nil {
			return int(data)
		}
	}
	return 0
}

func (j *Json) ToFloat() float64 {
	if data, ok := (j.Data).(float64); ok {
		return data
	}
	if data, ok := (j.Data).(string); ok {
		if data, err := strconv.ParseFloat(data, 64); err != nil {
			return data
		}
	}
	return 0
}

func (j *Json) ToString() string {
	if data, ok := (j.Data).(string); ok {
		return data
	}
	if data, ok := (j.Data).(float64); ok {
		if data := strconv.FormatFloat(data, 'f', -1, 64); ok {
			return data
		}
	}
	return ""
}

func (j *Json) ToArray() (k, v []string) {
	if data, ok := (j.Data).([]interface{}); ok {
		for _, arrValue := range data {
			for key, value := range arrValue.(map[string]interface{}) {
				switch vv := value.(type) {
				case float64:
					k = append(k, key)
					v = append(v, strconv.FormatFloat(vv, 'f', -1, 64))
				case string:
					k = append(k, key)
					v = append(v, vv)
				}
			}
		}
		return
	}

	if data, ok := (j.Data).(map[string]interface{}); ok {
		for key, value := range data {
			switch vv := value.(type) {
			case float64:
				k = append(k, key)
				v = append(v, strconv.FormatFloat(vv, 'f', -1, 64))
			case string:
				k = append(k, key)
				v = append(v, vv)
			}
		}
		return
	}
	return
}

func (j *Json) StringToArray() (data []string) {
	for _, value := range (j.Data).([]interface{}) {
		switch vv := value.(type) {
		case string:
			data = append(data, vv)
		case float64:
			data = append(data, strconv.FormatFloat(vv, 'f', -1, 64))
		}
	}
	return
}
