package yaml

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
	"strings"
)

// YamlParse 读取yaml文件到对应类
func YamlParse(in *[]byte, out interface{}) error {
	err := yaml.Unmarshal(*in, out)
	if err != nil {
		return err
	}
	typeOf := reflect.TypeOf(out)
	valueOf := reflect.ValueOf(out)
	err = configInit(typeOf, valueOf)
	if err != nil {
		return err
	}
	return nil
}

// 根据系统变量再次重置值
func configInit(t reflect.Type, v reflect.Value) error {
	var err error = nil
	if t.Kind() == reflect.Ptr {
		structRe := t.Elem().Elem()
		value := v.Elem().Elem()
		filed := structRe.NumField()
		for i := 0; i < filed; i++ {
			fieldType := structRe.Field(i).Type
			fieldValue := value.Field(i)
			err = configInit(fieldType, fieldValue)
			if err != nil {
				return err
			}
		}
	} else if t.Kind() == reflect.Struct {
		filed := t.NumField()
		for i := 0; i < filed; i++ {
			fieldType := t.Field(i).Type
			fieldValue := v.Field(i)
			err = configInit(fieldType, fieldValue)
			if err != nil {
				return err
			}
		}
	} else if t.Kind() == reflect.Int64 {

	} else if t.Kind() == reflect.Bool {

	} else if t.Kind() == reflect.String {
		s := v.String()
		flagStartIndex := strings.LastIndex(s, "${")
		flagEndIndex := strings.LastIndex(s, "}")
		if flagStartIndex == -1 || flagEndIndex == -1 {
			return err
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

	} else {
		fmt.Println("解析配置文件，未知类型参数")
	}
	return err
}
