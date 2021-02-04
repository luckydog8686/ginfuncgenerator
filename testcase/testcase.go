
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/luckydog8686/logs"
	"net/http"
	"reflect"
)

type SS struct {
	Name string
}

func main(){
	r := GenRunc(TestGen)
	router := gin.Default()
	router.POST("/ping",r)
	router.Run("127.0.0.1:80")
}



func GenRunc(fun interface{}) gin.HandlerFunc {
	t :=reflect.TypeOf(fun)
	logs.Info(t)
	inV := t.In(0)
	in_Kind := inV.Kind()
	fmt.Printf("Kind: %v   Name: %v -----------",in_Kind,inV.Name())

	tx := func(c *gin.Context){
		s := reflect.New(inV.Elem()).Interface()
		logs.Info(s)
		logs.Info(reflect.TypeOf(s))
		err2 := c.Bind(s)
		logs.Info(reflect.TypeOf(s))
		if err2 != nil {
			logs.Error(err2)
		}
		logs.Info("S:",s)
		res:=Call(fun,s)
		logs.Info(len(res))
		data := res[0].Interface()
		err := res[1].Interface()
		logs.Info(data)
		logs.Info(res[1])

		//data,err := TestGen(s)
		c.JSON(http.StatusOK,gin.H{"error":err,"data":data})
	}
	return tx
}

func TestGen(ss *SS) (*SS,error)  {
	logs.Info(ss)
	return ss,nil
}

func  Call(fun interface{},params ...interface{}) []reflect.Value {
	f := reflect.ValueOf(fun)
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	return f.Call(in)
}