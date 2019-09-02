package exterror

import (
	"fmt"
	"bytes"
	"runtime"
	"path"
	"strings"
)


type Error map[string]interface{}

func Warp(err interface{}) Error {
	this := Error{}
	switch vType := err.(type) {
	case string:
		{
			this["zerror"] = fmt.Errorf(vType)
		}
	case Error:
		{
			for key, val := range vType {
				this[key] = val
			}
		}
	case *Error:
		{
			if nil != vType {
				for key, val := range *vType {
					this[key] = val
				}
			}
		}
	case error:
		{
			this["zerror"] = vType
		}
	default:
		{
			panic(fmt.Errorf("Not Support Error Type:%#v", err))
		}
	}
	return this
}

func New(err interface{}) Error {
	this := Warp(err)
	if _, ok := this["zcallstack"]; !ok {
		this["zcallstack"] = WithCs()
	}
	return this
}

func Newf(f string, v ...interface{}) Error {
	this := Warp(fmt.Errorf(f, v...))
	if _, ok := this["zcallstack"]; !ok {
		this["zcallstack"] = WithCs()
	}
	return this
}

func (this Error) WithField(key string, val interface{}) Error {
	this[key] = val
	return this
}

func (this Error) WithFields(nFields Error) Error {
	for k, v := range nFields {
		this[k] = v
	}
	return this
}

func (this Error) WithErr(err interface{}) Error {
	for key, val := range Warp(err) {
		this[key] = val
	}

	if _, ok := this["zcallstack"]; !ok {
		this["zcallstack"] = WithCs()
	}
	return this
}

func (this Error) WithCs() Error {
	if _, ok := this["zcallstack"]; !ok {
		this["zcallstack"] = WithCs()
	}
	return this
}

func (this Error) Addr() *Error {
	return &this
}

func (this Error) Cause() error {
	if err, ok := this["zerror"]; ok {
		switch vType := err.(type) {
		case error:
			{
				return vType
			}
		}
	}

	return nil
}

func (this Error) Error() string {
	return this.String()
}

func (this Error) String() string {
	// TODO 使用 string.buffer
	var buffer bytes.Buffer
	buffer.Write([]byte(" "))
	for key, item := range this {
		fmt.Fprintf(&buffer, " %s=%s", key, item)
	}
	buffer.Write([]byte(" "))
	return string(buffer.Bytes())
}

func (this Error) ToMaps() map[string]interface{} {
	return this
}

func ToMaps(err interface{}) map[string]interface{} {
	val := Warp(err)
	return val
}

func WithCs() string {
	pc := make([]uintptr, 7)
	max := runtime.Callers(3, pc) // 参数2表示忽略 Callers() WithCs() 这两层调用

	callStacks := []string{}
	for index, pcItem := range pc {
		if index >= max {
			break
		}
		f := runtime.FuncForPC(pcItem)
		file, line := f.FileLine(pcItem)

		fullname := f.Name()
		funcname := ""
		index := strings.LastIndexAny(fullname, ".")
		if index >= 0 {
			funcname = fullname[index+1:]
		}

		callfunc := fmt.Sprintf("%s@%s:%d", funcname, path.Base(file), line)
		callStacks = append(callStacks, callfunc)
	}

	return strings.Join(callStacks, " ")
}
