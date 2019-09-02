package exterror

import (
	"fmt"
	"testing"
	"gitee.com/wangtiga/ytxpkg/log"
	"gitee.com/wangtiga/exterror"
)

func TestExtErrorNew(t *testing.T) {
	// -- support --
	// string
	// Error
	// *Error

	log.WithMaps(exterror.ToMaps(
		exterror.New("string"),
	)).Debugf("test string")

	log.WithMaps(exterror.ToMaps(
		exterror.New(exterror.Error{"Type": "exterror.Error"}),
	)).Debugf("test exterror.Error")

	log.WithMaps(exterror.ToMaps(
		exterror.New(&exterror.Error{"Type": "exterror.Error"}),
	)).Debugf("test &exterror.Error")

	// TODO
	// -- not support --
	// *string
	// nil
	// int

}

func TestExtError0(t *testing.T) {
	err := subFunc0()
	log.WithMaps(exterror.ToMaps(err)).Info("TestExtError0")
}

func TestExtError1(t *testing.T) {
	err := subFunc1()

	errField := exterror.Error{
		"akey1": "arg1",
		"bkey2": "arg2",
		"ckey3": "arg3",
		"dkey4": "arg4",
		"1key1": "arg1",
		"2key2": "arg2",
		"3key3": "arg3",
		"4key4": "arg4",
		"_key4": "arg4",
		"-key4": "arg4",
		":key4": "arg4",
		".key4": "arg4",
	}
	err.WithFields(errField)

	log.WithMaps(exterror.ToMaps(err)).Debugf("TestExtError1")
	log.WithMaps(subFunc1().ToMaps()).Debugf("TestExtError1")
	log.WithMaps(exterror.New("xxx").ToMaps()).Debugf("TestExtError1")
}

func TestExtError2(t *testing.T) {
	errField := exterror.Error{
	}
	errField.WithErr(fmt.Errorf("this is error type"))

	log.WithMaps(exterror.ToMaps(errField)).Debugf("TestExtError2")
}

func TestExtError3(t *testing.T) {
	errField := exterror.Error{
	}
	errField.WithErr("this is string type")

	log.WithMaps(exterror.ToMaps(errField)).Debugf("TestExtError3")
}

func TestExtError4(t *testing.T) {
	errField := exterror.Error{
	}
	errField.WithErr(exterror.Error{"tips":"this is exterror.Error type"})

	log.WithMaps(exterror.ToMaps(errField)).Debugf("TestExtError4")
}

func TestExtError5(t *testing.T) {
	err := orginErr()
	errField := exterror.New(err)
	log.WithMaps(exterror.ToMaps(errField)).Debugf("TestExtError5")

	if err != errField.Cause() {
		t.Errorf("errField.Cause() unexception")
	}
}

func TestExtError6(t *testing.T) {
	err := subFunc0()
	errField := exterror.New(err)
	log.WithMaps(exterror.ToMaps(errField)).Debugf("TestExtError6")

	if err != errField.Cause() {
		switch vType := err.(type) {
		case exterror.Error:
			{
				if vType.Cause() != errField.Cause() {
					t.Errorf(
						"errField.Cause() unexception vType.Cause=%#v, errField.Cause=%#v",
						vType.Cause(),
						errField.Cause(),
					)
				}
				return
			}
		}
	}

	t.Errorf("errField.Cause() unexception err=%#v, cause=%#v", err, errField.Cause())
}

func subFunc0() error {
	return extErr()
}

func subFunc1() exterror.Error {
		// 不建议返回 exterror.Error ，会踩坑。
		// 比如本函数 return  nil 时，上层判断返回值 nil != err 会得到 true 
	return extErr()
}

func extErr() exterror.Error {
	return exterror.New("extErr")
}

func orginErr() error {
	return fmt.Errorf("orginErr")
}
