1.配置golang环境

2.克隆本项目：`git clone https://github.com/Ailzr/Navidrome-Manager-go.git `

3.进入项目目录：`cd Navidrome-Manager-go`

4.安装依赖：`go mod tidy`

5.修改配置文件，位于`conf/config.yaml`，其中Identity可以随意填写，Name为管理员登录用户名，Password为管理员登录密码，savePath为Navidrome的音乐存储路径

6.编译项目`go build`

7.运行本项目`./MusicManager`(推荐使用后台运行：`nohup ./MusicManager &`)

8.在nginx反向代理路径location /manage 中 proxy_pass到localhost:8080/manage，可以通过https://server/manage/login来访问管理登录页面，如果想用其他访问路径，请自行修改代码（主要是把router的路径给改了，还有前端访问api的路径以及manage文件夹）

9.注意项：请注意修改nginx配置文件中传输文件大小限制，不然可能上传不了文件

10.本人没有审美，所以页面做的很稀碎，大概看着能用就行了
