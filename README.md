

# RUN 

## 运行apioutside

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
