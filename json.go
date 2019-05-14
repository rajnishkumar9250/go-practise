package main
import "encoding/json"
import "fmt"
//import "reflect"
import "strconv"

func main() {
  byt := []byte(`{"num":"123", "strs":[{"key":"Pressure","value":"12.0928272"}, {"key":"asset_id", "value":"hmmm"}]}`)
  var dat map[string]interface{}
  json.Unmarshal(byt, &dat)
  //fmt.Println(dat)
  num := dat["num"].(string)
  //fmt.Println(num)
  _, err := strconv.ParseInt(num, 10, 64)
  fmt.Println(err)
  if err != nil {
	fmt.Printf("float64")
	fmt.Print(err)
  } 
  /*var strs1  []interface{}
  strs, _ := json.Marshal(dat["readings"])
  //fmt.Println(string(strs))
  json.Unmarshal(strs, &strs1)
  var result = make(map[string]interface{})
  result["name"] = dat["device"]
  for _, v := range strs1 {
    //fmt.Println(k, v)
    //fmt.Println(reflect.TypeOf(v).Kind())
    //var rec map[string]interface{}
    //json.Unmarshal(v, &rec)
    key := v.(map[string]interface{})["name"]
    //fmt.Println(v.(map[string]interface{})["key"])
    //fmt.Println(key)
    value := v.(map[string]interface{})["value"]
    //fmt.Println(value)
    result[key.(string)] = value
  }
  fmt.Println(result)
  */
}
