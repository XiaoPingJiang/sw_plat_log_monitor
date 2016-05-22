# sw_plat_log_monitor

* 首先需要安装golang包管理工具 glide
* 安装好后，git clone github.com/XiaoPingJiang/sw_plat_log_monitorit
* cd sw_plat_log_monitorit
* glide install
* bee run
* ps: bee 工具是beego框架里的一个热编译工具，需要先安装bee, 传送门：http://beego.me/quickstart

* 题外话：如果让beego程序，以windows服务的方式运行呢?
* 使用nssm注册会非常方便
* nssm下载地址：http://nssm.cc/download
* 注意要采用管理员身份运行
* ./nssm.exe install PlatLogMonitor sw_plat_log_monitorit.exe
* net start PlatLogMonitor
* 如果要删除服务：./nssm.exe remove PlatLogMonitor
