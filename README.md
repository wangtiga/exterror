
## error 应该包含以下信息
* 错误描述信息
* 错误发生时的调用堆栈
* 错误发生时的环境信息

但是
* go内置的 error 只能表达 错误描述信息
* pkg/errors 只能表达 错误描述信息 和 调用堆栈
* exterror 能同时表达以上三个信息，而且可以和 logrus 日志模块配合，方便输出 结构化日志


## 用法
exterror.New(msg) 就返回满足 error interface{} 接口的 struct
```go
func SaveFile(name string, data []byte) (error) {
	tmpDir := os.TempDir()
	tmpfile, err := ioutil.TempFile(tmpDir, name)
	if nil!= err {
		return exterror.New(err)
	}

	if _, err := tmpfile.Write(data); err != nil {
		return exterror.New(err)
	}

	if err := tmpfile.Close(); err != nil {
		return exterror.New(err)
	}

	log.Printf("SaveFile Succ! Paht: %s \n", tmpfile.Name())
	return nil
}


```

## 原则
* 生成 exterror.Error 接口，只要涉及保存error参数，都会自动生成 zcallstack 和 zerror 
* 可以和没有使用 exterror 的日志方便融合
* 可以和使用了 logrus 的日志方便融合

## 目的
* 服务端模块通过日志来分析排查问题
* 但记录内容不完整的日志往往需要反复重现，反复增加日志输出才能定位问题
* 对于偶现问题，时间都浪费到重现问题上了
* logrus 建议结构化日志，方便分析，所以将 error 与 log 结合，才是最好的错误处理方案

## 参考:
* [断点单步跟踪是一种低效的调试方法](https://blog.codingnow.com/2018/05/ineffective_debugger.html)
* [pkg/errors 作者的博客](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)
* [logrus 建议结构化日志，方便分析](https://github.com/Sirupsen/logrus#fields)
* [design-philosophy-on-logging](https://www.ardanlabs.com/blog/2017/05/design-philosophy-on-logging.html)

## TODO
* 1.zerror 应该保存原始 error 的 interface{} 引用？还是仅仅保存 Error() 返回的 string 就行呢？
* 2.哪种错误处理更好？ recover panic()  还是 if nil != err ?
* 3.如何继承package，以方便扩展 log 包
* 4.仅输出文件名称，如何解决重名文件
	目前调整编码习惯，尽量保持每个文件名，函数名唯一。
    另外，其实根据调用堆栈和文件名来判断，即使文件重名，也不难识别。
