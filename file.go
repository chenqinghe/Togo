package Togo

import (
	"os"
)

func BaseName(path,suffix string)(string,error){
	fi,err:=os.Stat(path)
	if err!=nil{
		return "",err
	}
	return fi.Name(),nil
}
