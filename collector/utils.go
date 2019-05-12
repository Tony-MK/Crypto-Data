package collector
import (
	"io/ioutil"
	"errors"
	"github.com/kr/pretty"
)
func toLower(s string)(string,error){
	var l = ""
	for i := 0; i < len(s); i++ {
		if(s[i] > 64 && s[i] < 91){
			l += string(s[i]+32)
			continue
		}
		if(s[i] < 96 && s[i] > 123){
			return string(l), errors.New("Input has a non-alphasetic character "+string(s[i]))
		}
		l += string(s[i])
	}
	return string(l), nil
}

func CheckErr(err error)bool{
	if(err != nil){
		pretty.Println(err)
		return true
	}
	return false
}

func getCurrencies(stream_name string)([]string,error){
	if(stream_name[0] == '/'){
		stream_name = string(stream_name)
	}
	dirs ,err := ioutil.ReadDir("./Data/"+stream_name)
	if(err != nil){return nil, err}

	var Currenies = []string{}
	for i := 0;i<len(dirs);i++{
		for n:= 0; i < len(dirs[i].Name()); i++{
			if(dirs[i].Name()[n] < 65 &&  dirs[i].Name()[n] > 90){
				print(string(dirs[i].Name()[n])+" "+dirs[i].Name());
			}
		}
		Currenies = append(Currenies,dirs[i].Name());
	}
	return Currenies,nil
}