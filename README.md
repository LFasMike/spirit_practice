

# Start

下载第三方依赖包:
~~~
    go get github.com/gogap/spirit
    go get github.com/spirit-contrib/inlet_http_api
~~~ 

## 配置配置文件:

将项目中的{{your_access_key_id}}:{{your_access_secret}}@{{your_account}}替换成个人的阿里云账号

## 运行apioutside

对github.com/spirit-contrib/inlet_http_api进行build,生成inlet_http_api可执行文件,移动到api_outside项目中。

➜  api_outside ./inlet_http_api run -c spirit.conf inlet_http_api

## 运行todo

➜  todo ./todo run -c spirit.conf todo

# 数据库
```
db: todo

Create Table: CREATE TABLE `task` (
  `id` varchar(45) NOT NULL,
  `owner_id` varchar(45) DEFAULT NULL,
  `title` varchar(45) DEFAULT NULL,
  `description` varchar(45) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `version` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
```
