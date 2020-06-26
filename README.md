
基于gin+gorm开发项目



### 配置MySQL
1. 在你的数据库中执行以下命令，创建本项目所用的数据库：


2. 在`jishiben/conf/config.ini`文件中按如下提示配置数据库连接信息。

### 编译
```bash
go build
```

### 执行

Mac/Unix：
```bash
./jishiben conf/config.ini
```


启动之后，使用浏览器打开`http://127.0.0.1:9090/`即可。

