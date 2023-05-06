package route

import (
	"reflect"
	"strings"

	"github.com/androidsr/sgin/annotation"
	"github.com/gin-gonic/gin"
)

var (
	ctrls []interface{}
	docs  map[string]map[string]string
	pre   = "@"
)

type BaseController struct {
}

func New(router *gin.RouterGroup, ctrlDir string, controller ...interface{}) {
	docs = annotation.Scan(ctrlDir, pre)
	ctrls = append(ctrls, controller...)
	autoRegister(router)
}

func autoRegister(router *gin.RouterGroup) {

	for _, ctrl := range ctrls {
		value := reflect.ValueOf(ctrl).Elem()
		doc := docs[value.Type().Name()]
		for i := 0; i < value.NumMethod(); i++ {
			method := value.Type().Method(i)
			comment := doc[method.Name]
			if comment == "" {
				continue
			}
			cms := strings.Split(comment, ":")
			if len(cms) < 2 {
				continue
			}
			httpMethod := strings.ReplaceAll(cms[0], pre, "")
			relativePath := cms[1]
			m := value.MethodByName(method.Name)

			router.Handle(httpMethod, strings.TrimSpace(relativePath), func(c *gin.Context) {
				inputs := []reflect.Value{reflect.ValueOf(c)}
				m.Call(inputs)
			})
		}
	}
}
