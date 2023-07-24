检测使用率不准确，如果目录总空间是50G，使用了25G，剩余使用率为50%，判断条件为剩余空间小于30%不备份。若备份文件25G，条件通过，依旧会备份导致空间被占满。



 SELECT SUM(DATA_LENGTH) AS dataSize, SUM(INDEX_LENGTH) AS indexSize FROM information_schema.TABLES WHERE TABLE_SCHEMA = 'test';


 // 痛点 1
// 备份前未判断磁盘空间
// 解决方案: check磁盘空间、

// 痛点 2
// 备份文件在升级后未删除，或则是保留的备份文件较多，造成运维人员不敢轻易删除历史遗留文件。
// 解决方案: 规范备份文件的保存目录，定时删除。
自动化备份文件管理：编写脚本来自动管理备份文件，包括删除旧的备份文件、保留特定时间范围内的备份文件等。可以使用脚本中的日期和时间函数来确定备份文件的创建时间，并根据设定的保留策略进行清理。
设置备份文件保留期限：根据业务需求和合规要求，确定备份文件的保留期限。例如，可以设置备份文件只保留最近的几个版本或保留一定的时间范围内的备份文件。一旦备份文件超过保留期限，就自动删除。
添加备份文件元数据记录：在备份文件中添加元数据，例如备份日期、版本号、相关信息等。这样可以更好地追踪和管理备份文件，帮助运维人员确定哪些备份文件可以安全删除。
定期审查备份文件：定期审查备份文件，评估其重要性和有效性。与相关团队和业务负责人合作，确定哪些备份文件可以安全删除，以释放存储空间并降低维护成本。
监控备份文件存储空间：设置监控系统，定期检查备份文件存储空间的使用情况。当存储空间接近上限或达到设定的阈值时，触发警报通知运维人员，及时进行备份文件管理和清理。



dbackup --help


命令参数
Common Commands:
    dbackup create --source [ mysql | postgresql ] <username> <password> <database> <table>
    dbackup list
    dbackup delete < 编号 | 文件名 >

分别使用go的参数解析库实现如下功能，对比哪个库更简单直观：
dbackup --help
命令参数
Common Commands:
    dbackup create <username> <password> <database> <table>
    dbackup list
    dbackup delete < 编号 | 文件名 >

用户输入：
--file <文件存储位置>
--meta < operator: nil, action: [ app-update | data-correct | other], desc: nil >

环境变量：



Global Options:


backup -s <数据库命> add -n -m ""
backup list
backup 