package monster
import(
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Monster struct {
	Name string
	Age int
	Skill string
} 

func (this * Monster) Store() bool{
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err =", err)
		return false
	} 

	filePath := "e:/monster.ser"
	err = ioutil.WriteFile(filePath, data,066)
	if err != nil {
		fmt.Println("write file err =", err)
		return false
	}
	return true
}

func (this *Monster) Restore() bool {
	filePath := "e:/monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFile err =", err)
		return false
	}

	err = json.Unmarshal(data,this)
	if err != nil {
		fmt.Println("Unmarshal err =", err)
		return false
	}
	return true
}
