

## 特性
```
1.问题：备份前未检查磁盘空间，导致空间被占满。解决：备份前判断磁盘空间使用情况

```



## Features
```
1.问题：备份文件在升级后未删除，或则是保留的备份文件较多，造成运维人员不敢轻易删除历史遗留文件。解决：备份前检查目前备份文件数量，例如，可以设置备份文件只保留最近的几个版本或保留一定的时间范围内的备份文件。一旦备份文件超过保留期限，就自动删除。

2.添加备份文件元数据记录：在备份文件中添加更丰富的元数据(和备份文件成对出现的元数据文件)，例如执行备份的人、备份原因、相关信息等。这样可以更好地追踪和管理备份文件，帮助运维人员确定哪些备份文件可以安全删除。
--operator 'Yidon'
--desc '3.5.1版本升级到3.5.2版本'
3.指定备份目录，并记录备份目录
```

