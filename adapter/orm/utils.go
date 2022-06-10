// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package orm

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type fn func(string) string

var (
	nameStrategyMap = map[string]fn{
		defaultNameStrategy:      snakeString,
		SnakeAcronymNameStrategy: snakeStringWithAcronym,
	}
	defaultNameStrategy      = "snakeString"
	SnakeAcronymNameStrategy = "snakeStringWithAcronym"
	nameStrategy             = defaultNameStrategy
)

// StrTo is the target string
type StrTo orm.StrTo

// Set string
func (f *StrTo) Set(v string) {
	(*orm.StrTo)(f).Set(v)
}

// Clear string
func (f *StrTo) Clear() {
	(*orm.StrTo)(f).Clear()
}

// Exist check string exist
func (f StrTo) Exist() bool {
	return orm.StrTo(f).Exist()
}

// Bool string to bool
func (f StrTo) Bool() (bool, error) {
	return orm.StrTo(f).Bool()
}

// Float32 string to float32
func (f StrTo) Float32() (float32, error) {
	return orm.StrTo(f).Float32()
}

// Float64 string to float64
func (f StrTo) Float64() (float64, error) {
	return orm.StrTo(f).Float64()
}

// Int string to int
func (f StrTo) Int() (int, error) {
	return orm.StrTo(f).Int()
}

// Int8 string to int8
func (f StrTo) Int8() (int8, error) {
	return orm.StrTo(f).Int8()
}

// Int16 string to int16
func (f StrTo) Int16() (int16, error) {
	return orm.StrTo(f).Int16()
}

// Int32 string to int32
func (f StrTo) Int32() (int32, error) {
	return orm.StrTo(f).Int32()
}

// Int64 string to int64
func (f StrTo) Int64() (int64, error) {
	return orm.StrTo(f).Int64()
}

// Uint string to uint
func (f StrTo) Uint() (uint, error) {
	return orm.StrTo(f).Uint()
}

// Uint8 string to uint8
func (f StrTo) Uint8() (uint8, error) {
	return orm.StrTo(f).Uint8()
}

// Uint16 string to uint16
func (f StrTo) Uint16() (uint16, error) {
	return orm.StrTo(f).Uint16()
}

// Uint32 string to uint32
func (f StrTo) Uint32() (uint32, error) {
	return orm.StrTo(f).Uint32()
}

// Uint64 string to uint64
func (f StrTo) Uint64() (uint64, error) {
	return orm.StrTo(f).Uint64()
}

// String string to string
func (f StrTo) String() string {
	return orm.StrTo(f).String()
}

// ToStr interface to string
func ToStr(value interface{}, args ...int) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 32))
	case float64:
		s = strconv.FormatFloat(v, 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 64))
	case int:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int8:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int16:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int32:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int64:
		s = strconv.FormatInt(v, argInt(args).Get(0, 10))
	case uint:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint8:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint16:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint32:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint64:
		s = strconv.FormatUint(v, argInt(args).Get(0, 10))
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

// ToInt64 interface to int64
func ToInt64(value interface{}) (d int64) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		panic(fmt.Errorf("ToInt64 need numeric not `%T`", value))
	}
	return
}

func snakeStringWithAcronym(s string) string {
	data := make([]byte, 0, len(s)*2)
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		before := false
		after := false
		if i > 0 {
			before = s[i-1] >= 'a' && s[i-1] <= 'z'
		}
		if i+1 < num {
			after = s[i+1] >= 'a' && s[i+1] <= 'z'
		}
		if i > 0 && d >= 'A' && d <= 'Z' && (before || after) {
			data = append(data, '_')
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data))
}

// snake string, XxYy to xx_yy , XxYY to xx_y_y
func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data))
}

// SetNameStrategy set different name strategy
func SetNameStrategy(s string) {
	if SnakeAcronymNameStrategy != s {
		nameStrategy = defaultNameStrategy
	}
	nameStrategy = s
}

// camel string, xx_yy to XxYy
func camelString(s string) string {
	data := make([]byte, 0, len(s))
	flag, num := true, len(s)-1
	for i := 0; i <= num; i++ {
		d := s[i]
		if d == '_' {
			flag = true
			continue
		} else if flag {
			if d >= 'a' && d <= 'z' {
				d = d - 32
			}
			flag = false
		}
		data = append(data, d)
	}
	return string(data)
}

type argString []string

// get string by index from string slice
func (a argString) Get(i int, args ...string) (r string) {
	if i >= 0 && i < len(a) {
		r = a[i]
	} else if len(args) > 0 {
		r = args[0]
	}
	return
}

type argInt []int

// get int by index from int slice
func (a argInt) Get(i int, args ...int) (r int) {
	if i >= 0 && i < len(a) {
		r = a[i]
	}
	if len(args) > 0 {
		r = args[0]
	}
	return
}

// parse time to string with location
func timeParse(dateString, format string) (time.Time, error) {
	tp, err := time.ParseInLocation(format, dateString, DefaultTimeLoc)
	return tp, err
}

// get pointer indirect type
func indirectType(v reflect.Type) reflect.Type {
	switch v.Kind() {
	case reflect.Ptr:
		return indirectType(v.Elem())
	default:
		return v
	}
}