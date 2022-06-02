package yaml

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
	"strings"
)

// YamlParse 读取yaml文件到对应类
func YamlParse(in *[]byte, out interface{}) {
	err := yaml.Unmarshal(*in, out)
	if err != nil {
		panic(err)
	}
	typeOf := reflect.TypeOf(out)
	valueOf := reflect.ValueOf(out)
	configInit(typeOf, valueOf)
}

// 根据系统变量再次重置值
func configInit(t reflect.Type, v reflect.Value) {
	if t.Kind() == reflect.Ptr {
		structRe := t.Elem().Elem()
		value := v.Elem().Elem()
		filed := structRe.NumField()
		for i := 0; i < filed; i++ {
			fieldType := structRe.Field(i).Type
			fieldValue := value.Field(i)
			configInit(fieldType, fieldValue)
		}
	} else if t.Kind() == reflect.Struct {
		filed := t.NumField()
		for i := 0; i < filed; i++ {
			fieldType := t.Field(i).Type
			fieldValue := v.Field(i)
			configInit(fieldType, fieldValue)
		}
	} else if t.Kind() == reflect.Int64 {

	} else if t.Kind() == reflect.Bool {

	} else if t.Kind() == reflect.String {
		s := v.String()
		flagStartIndex := strings.LastIndex(s, "${")
		flagEndIndex := strings.LastIndex(s, "}")
		// 不存在则返回原值
		if flagStartIndex == -1 || flagEndIndex == -1 {
			return
		}
		start := flagStartIndex + 2
		val := s[start:flagEndIndex]
		split := strings.Split(val, ":")
		//查询环境变量
		env := os.Getenv(split[0])
		if len(split) == 1 {
			v.SetString(env)
		} else {
			if env == "" {
				v.SetString(split[1])
			} else {
				v.SetString(env)
			}
		}
	} else if t.Kind() == reflect.Int {

	} else if t.Kind() == reflect.Uint64 {

	}else {
		fmt.Println("解析配置文件，未知类型参数:",t.Kind().String())
	}
}
