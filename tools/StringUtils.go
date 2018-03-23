package tools

import (
	"sort"
	"strconv"
)

func IsContain(arrs []string,el string) bool{
	for _,v := range arrs{
		if el == v{
			return true
		}
	}
	return false
}

func RemoveDuplicatesAndEmpty(a []string) (ret []string){
	sort.Strings(a)
	a_len := len(a)
	for i:=0; i < a_len; i++{
		if (i > 0 && a[i-1] == a[i]) || len(a[i])==0{
			continue;
		}
		ret = append(ret, a[i])
	}
	return
}

func GetString(v interface{})string{
	var result string
	switch v := v.(type) {
	case uint:
		result = strconv.Itoa(int(v))
	case uint8:
		result = strconv.Itoa(int(v))
	case uint32:
		result = strconv.Itoa(int(v))
	case uint64:
		result = strconv.Itoa(int(v))
	case int:
		result = strconv.Itoa(v)
	case int8:
		result = strconv.Itoa(int(v))
	case int32:
		result = strconv.Itoa(int(v))
	case int64:
		result = strconv.Itoa(int(v))
	case string:
		result = v
	case float32:
		result = strconv.FormatFloat(float64(v),'E', -1, 64)
	case float64:
		result = strconv.FormatFloat(v,'E', -1, 64)
	default:
		result = fmt.Sprintf("%v",v)
		if result == "<nil>"{
			result = ""
		}
	}
	return result
}

func DistinctByMap(arr []string ) []string{
	tmp_arr := make([]string, 0)
	if arr == nil || len(arr) == 0{
		return tmp_arr
	}
	tm := make(map[string]int,0)
	for _,e := range arr{
		tm[e] = 0
	}
	for k,_ := range tm{
		tmp_arr = append(tmp_arr, k)
	}
	return tmp_arr
}

func DistinctByCircle(arr []string ) []string{
	if arr == nil{
		return nil
	}
	size := len(arr)
	tmpArr := make([]string,0)
	if size == 0 {
		return tmpArr
	}
	for _,e := range arr{
		if !IsContain(tmpArr,e){
			tmpArr = append(tmpArr, e)
		}
	}
	return tmpArr
}
// 比较两个类型元素是否相等
func Compare(base interface{}, target interface{}) bool{
	// 先比较类型
	baseType := reflect.TypeOf(base)
	targetType := reflect.TypeOf(target)
	if targetType.String() != baseType.String(){
		return false
	}
	// 再比较值
	switch base := base.(type) {
	case []uint8:
		// 比较长度
		targetObj := target.([]uint8)
		if len(base) != len(targetObj){
			return false
		}
		// 遍历匹配元素
		for i,e := range base{
			if e != targetObj[i]{
				return false
			}
		}
		return true
	case string:
		return base == target
	default:
		return base == target
	}
}
