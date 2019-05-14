package main

import(
	"fmt"
	"encoding/json"
	"strconv"
	"reflect"
)

func missingNo() int{
	rangelimit := []int{1,2,3,4,5,6}
	datalist := []byte(`
	{
		"1":1,
		"6":6,
		"3":3,
		"4":4,
		"5":5
	}`)
	jsonArr := make(map[string]interface{})
	err := json.Unmarshal(datalist, &jsonArr)
	if err != nil{
		return 0
	}
 
	for i :=1; i<= len(rangelimit); i++ {
		fmt.Println("checking")
		var flag bool =   false
		fmt.Println("arr ele is %X", rangelimit[i-1])
		key := strconv.Itoa(i)
		fmt.Println("Key is %X", key)
		fmt.Println(jsonArr[key])
		var val  int
		if jsonArr[key] != nil{
			val = int(jsonArr[key].(float64))
		}
		fmt.Println(reflect.TypeOf(val))
		fmt.Println(rangelimit[i-1])
		fmt.Println(reflect.TypeOf(rangelimit[i-1]))
		if int(rangelimit[i-1]) == val{
			fmt.Println("number %d is match", i)
			flag = true
		}
		if  flag == false{
			return  i
		}
	}
	return 0
}


func main(){
	fmt.Println("search missing no")
	var missedNo int = missingNo()
	fmt.Println("Missing Number is %X", missedNo)
}
