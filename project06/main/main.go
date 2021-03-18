package main
import(
	"fmt"
	"encoding/json"
)

type Monster struct{
	Name string `json:"monstar_name"`//反射机制
	Age int
	Bir string
}

func testStruct()  {
	monster := Monster{
		Name : "牛魔",
		Age : 40,
		Bir :"2020-1-12",
	}

	data, err :=json.Marshal(&monster)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(string(data))
}

func testMap()  {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "ho"
	a["age"] = 20

	data, err :=json.Marshal(&a)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(string(data))
}

func testSlice()  {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "ja"
	m1["age"] = 10
	slice = append(slice,m1)
	
	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "ja"
	m2["age"] = 10
	m2["add"] = [2]string{"内","经济"}
	slice = append(slice,m2)

	data, err :=json.Marshal(&slice)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(string(data))
}

func unmarshalStruct()  {
	str := "{\"monstar_name\":\"牛魔\",\"Age\":40,\"Bir\":\"2020-1-12\"}"

	var monster Monster

	err := json.Unmarshal([]byte(str),&monster)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(monster)
}


func main()  {

//json序列化
	// testStruct()
	//testMap()
	// testSlice()
	
//json反序列化
	unmarshalStruct()

}
//json的序列号和反序列化