package qiniu

import (
	"fmt"

	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
)

var (
	// 设置上传到的空间
	bucket = "yourbucket"

	// 设置上传文件的key
	key = "yourdefinekey"

	// 指定私有空间的域名
	domain = "xxxx.com2.z0.glb.qiniucdn.com"

	// 指定需要抓取的文件的url，必须是公网上面可以访问到的
	target_url = "xxxx"

	movekey = "movekey"

	copykey = "yourcopykey"
)

// 构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

//简单上传
func ABCTestmain() {
	// 初始化AK，SK
	conf.ACCESS_KEY = "ACCESS_KEY"
	conf.SECRET_KEY = "SECRET_KEY"

	// 创建一个Client
	c := kodo.New(0, nil)

	// 设置上传的策略
	policy := &kodo.PutPolicy{
		Scope: bucket,
		//设置Token过期时间
		Expires: 3600,
	}
	// 生成一个上传token
	token := c.MakeUptoken(policy)

	// 构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	var ret PutRet
	// 设置上传文件的路径
	filepath := "/Users/dxy/sync/sample2.flv"
	// 调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名
	res := uploader.PutFileWithoutKey(nil, &ret, token, filepath, nil)
	// 打印返回的信息
	fmt.Println(ret)
	// 打印出错信息
	if res != nil {
		fmt.Println("io.Put failed:", res)
		return
	}

}

//覆盖上传
func ABCTestmainU() {
	// 初始化AK，SK
	conf.ACCESS_KEY = "ACCESS_KEY"
	conf.SECRET_KEY = "SECRET_KEY"

	// 创建一个Client
	c := kodo.New(0, nil)

	// 设置上传的策略
	policy := &kodo.PutPolicy{
		Scope: bucket + ":" + key,
		// 设置Token过期时间
		Expires: 3600,
	}
	// 生成一个上传token
	token := c.MakeUptoken(policy)

	// 构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	var ret PutRet
	// 设置上传文件的路径
	filepath := "/Users/dxy/sync/sample2.flv"
	// 调用PutFile方式上传，这里的key需要和上传指定的key一致
	res := uploader.PutFile(nil, &ret, token, key, filepath, nil)
	// 打印返回的信息
	fmt.Println(ret)
	// 打印出错信息
	if res != nil {
		fmt.Println("io.Put failed:", res)
		return
	}

}

// 调用封装好的downloadUrl方法生成一个下载链接
func downloadUrl(domain, key string) string {
	// 调用MakeBaseUrl()方法将domain,key处理成http://domain/key的形式
	baseUrl := kodo.MakeBaseUrl(domain, key)
	policy := kodo.GetPolicy{}
	// 生成一个client对象
	c := kodo.New(0, nil)
	// 调用MakePrivateUrl方法返回url
	return c.MakePrivateUrl(baseUrl, &policy)
}

func ABCTestmainDownload() {
	// 初始化AK，SK
	conf.ACCESS_KEY = "ACCESS_KEY"
	conf.SECRET_KEY = "SECRET_KEY"
	// 打印出下载链接
	println(downloadUrl(domain, key))
}

//获取文件信息
func ABCTestmainGet() {

	conf.ACCESS_KEY = "ACCESS_KEY"
	conf.SECRET_KEY = "SECRET_KEY"

	// new一个Bucket管理对象
	c := kodo.New(0, nil)
	p := c.Bucket(bucket)

	// 调用Stat方法获取文件的信息
	entry, err := p.Stat(nil, key)
	// 打印列取的信息
	fmt.Println(entry)
	// 打印出错时返回的信息
	if err != nil {
		fmt.Println(err)
	}
}

//移动文件
func ABCTestmainmove() {

	conf.ACCESS_KEY = "ACCESS_KEY"
	conf.SECRET_KEY = "SECRET_KEY"

	// new一个Bucket管理对象
	c := kodo.New(0, nil)
	p := c.Bucket(bucket)

	// 调用Move方法移动文件
	res := p.Move(nil, key, movekey)

	// 打印返回值以及出错信息
	if res == nil {
		fmt.Println("Move success")
	} else {
		fmt.Println("Move failed:", res)
	}
}

//复制文件
func ABCTestmaincopy() {

	conf.ACCESS_KEY = "ACCESS_KEY"
	conf.SECRET_KEY = "SECRET_KEY"

	// new一个Bucket管理对象
	c := kodo.New(0, nil)
	p := c.Bucket(bucket)

	// 调用Copy方法移动文件
	res := p.Copy(nil, key, copykey)

	// 打印返回值以及出错信息
	if res == nil {
		fmt.Println("Copy success")
	} else {
		fmt.Println("Copy failed:", res)
	}
}

//删除文件
func ABCTestmaindelete() {

	conf.ACCESS_KEY = "ACCESS_KEY"
	conf.SECRET_KEY = "SECRET_KEY"

	// new一个Bucket管理对象
	c := kodo.New(0, nil)
	p := c.Bucket(bucket)

	// 调用Delete方法删除文件
	res := p.Delete(nil, key)
	// 打印返回值以及出错信息
	if res == nil {
		fmt.Println("Delete success")
	} else {
		fmt.Println(res)
	}
}

//列举文件条目
func ABCTestmainlist() {

	conf.ACCESS_KEY = "xxxx"
	conf.SECRET_KEY = "xxxx"

	// new一个Bucket对象
	c := kodo.New(0, nil)
	p := c.Bucket("xxx")

	// 调用List方法，第二个参数是前缀,第三个参数是delimiter,第四个参数是marker，第五个参数是列举条数
	// 可以参考 https://github.com/qiniu/api.v7/blob/f956f458351353a3a75a3a519fed4e3069f14df0/kodo/bucket.go#L131
	ListItem, _, _, err := p.List(nil, "photo/", "", "", 100)

	if err == nil {
		fmt.Println("List success")
	} else {
		fmt.Println("List failed:", err)
	}

	// 循环遍历每个操作的返回结果
	for _, item := range ListItem {
		fmt.Println(item.Key, item.Fsize)
	}
}

//抓取文件
func ABCTestmainCat() {

	conf.ACCESS_KEY = "xxxx"
	conf.SECRET_KEY = "xxxx"

	// new一个Bucket对象
	c := kodo.New(0, nil)
	p := c.Bucket(bucket)

	// 调用Fetch方法
	err := p.Fetch(nil, key, target_url)
	if err != nil {
		fmt.Println("bucket.Fetch failed:", err)
	} else {
		fmt.Println("fetch success")
	}
}

//批量操作文件
func ABCTestmainpil() {

	// new一个数组，需要批量操作的数组
	entryPairs := []kodo.KeyPair{
		kodo.KeyPair{
			Src:  "xxx.jpg",
			Dest: "xxxx.jpg",
		}, kodo.KeyPair{
			Src:  "xxxxx.jpg",
			Dest: "xxxxx.jpg",
		},
	}
	conf.ACCESS_KEY = "xxxx"
	conf.SECRET_KEY = "xxxx"
	// new一个Bucket对象
	c := kodo.New(0, nil)
	p := c.Bucket("xxxx")

	// 调用BatchCopy方法
	batchCopyRets, err := p.BatchCopy(nil, entryPairs...)

	if err == nil {
		fmt.Println("Move success")
	} else {
		fmt.Println("Move failed:", err)
	}

	// 循环遍历每个操作的返回结果
	for _, item := range batchCopyRets {
		fmt.Println(item.Code, item.Error)
	}
}
