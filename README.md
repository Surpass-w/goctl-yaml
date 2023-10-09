# goctl-yaml

![go-zero](https://img.shields.io/badge/Github-go--zero-brightgreen?link=https://github.com/tal-tech/go-zero&logo=github)

goctl-yaml是一款基于goctl的插件，用于生成vmp_gateway authenticate yaml file。

> 警告：本插件是对goctl plugin开发流程的指引，切勿用于生产环境。

# 插件使用

* 编译goctl-yaml插件

  ```shell script
  $  CGO_ENABLED=0 GOPROXY=goproxy.cn,direct go build -ldflags "-s -w" -o goctl-yaml
  ```

* 将`$GOPATH/bin`中的`goctl-yaml`添加到环境变量

* 创建api文件

  ```go
  info(
    title: "vmp_leak"
    desc: "vmp_leak"
    author: "wangpeng"
    email: "wangpeng@moresec.cn"
    srv: "vmp_leak"
  )
  
  @server(
    prefix: /v1/leak
    group: leak
  )
  
  service vmp_leak-api {
    @doc(
      summary : "漏洞列表"
	     roles: "1,2"
    )
    @handler GetLeakList
    get /list (LeakListReq) returns (LeakListAPIResp)
  }
  ```

* 生成yaml资源文件

  ```shell script
  $ goctl api plugin -plugin goctl-yaml="authenticate -filename doc/test.yaml" -api api/vmp_leak.api -dir .
  ```

  > 说明： 其中`goctl-yaml`为可执行的二进制文件，`"authenticate -filename doc/test.yaml"`
  为goctl-plugin自定义的参数，这里需要用引号`""`引起来。

# 插件开发流程

* 自定义参数

  ```go
  commands = []*cli.Command{
		    {
			Name:   "authenticate",
			Usage:  "generate authenticate yaml file",
			Action: action.Generator,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "filename",
					Usage: "authenticate save file name",
				    },
			    },
		    },
	 }
  ```

* 获取goctl传递过来的json信息

    * 利用goctl中提供的方法解析

      ```go
      plugin, err := plugin.NewPlugin()
      if err != nil {
          return err
      }
      ```

* 实现插件逻辑

  ```go
  generate.Do(filename, p)
  ```

> 说明：上述摘要代码来自goctl-yaml,完整信息可浏览源码。