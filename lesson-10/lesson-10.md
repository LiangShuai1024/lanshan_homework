
## Linux 基础

### linux是什么玩意?

> Linux是一套免费使用和自由传播的类Unix操作系统，是一个基于POSIX和UNIX的多用户、多任务、支持多线程和多CPU的操作系统。它能运行主要的UNIX工具软件、应用程序和网络协议。
> 它支持32位和64位硬件。Linux继承了Unix以网络为核心的设计思想，是一个性能稳定的多用户网络操作系统。 Linux操作系统诞生于1991 年10 月5 日（这是第一次正式向外公布时间）。
> Linux存在着许多不同的Linux版本，但它们都使用了Linux内核。Linux可安装在各种计算机硬件设备中，比如手机、平板电脑、路由器、视频游戏控制台、台式计算机、大型机和超级计算机。 
> 严格来讲，Linux这个词本身只表示**Linux内核**，但实际上人们已经习惯了用Linux来形容整个基于Linux内核，并且使用GNU 工程各种工具和数据库的操作系统。

（全世界 99% 服务器的操作系统都是 Linux）

### linux 的各个发行版

![linux发行版本](/lesson-10/v2-linux发行版本.png)

Linux发行版主要有三个分支：Debian、Slackware、Redhat。

(1)Debian:（以社区的方式运作）

1. Ubuntu:基于Debian开发的开源Linux操作系统，主要针对桌面和服务器；
2. Linux Mint:基于Debian和Ubuntu的Linux发行版，致力于桌面系统对个人用户每天的工作更易用，更高效，且目标是提供一种更完整的即刻可用体验。

(2)Redhat

1. RHEL(red hat enterprise Linux):Red Hat公司发布的面向企业用户的Linux操作系统。早起版本主要用于桌面环境，免费：
2. CentOS:基于Red hat Linux提供的可自由使用源代码的企业级Linux发行版本。每个版本的Centos都会获得十年的支持（通过安全更新的方式）。新版本的Centos大约每两年发行一次，而每个版本的Centos会定期（大概6个月）更新一次，以支持新的硬件。这样，建立一个安全、低维护、稳定、高预测性、高重复性的Linux环境。
3. Fedora:基于Red Hat Linux终止发行后，红帽公司计划以Fedora来取代Red Hat Linux在个人领域的应用，而另外发行的Red Hat Enterprise Linux取代Red Hat Linux在商业应用的领域。Fedora的功能对于用户而言，它是一套功能完备、更新快速的免费操作系统，而对赞助者Red Hat公司而言，它是许多新技术的测试平台，被认为可用的技术最终会加入到Red Hat Enterprise Linux中。Fedora大约每六个月发布新版本。

(3)Slackware

1. SUSE：基于Slackware二次开发的一款Linux,主要用于商业桌面、服务器。
2. SLES(SUSE Linux Enterprise Server(SLES):企业服务器操作系统，是唯一与微软系统兼容的Linux操作系统。
3. OpenSUSE:由suse发展而来，旨在推进linux的广泛使用，主要用于桌面环境，用户界面非常华丽，而且性能良好。


(4)其他发行版本：

1. Gentoo:基于linux的自由操作系统，基于Linux的自由操作系统，它能为几乎任何应用程序或需求自动地作出优化和定制。追求极限的配置、性能，以及顶尖的用户和开发者社区，都是Gentoo体验的标志特点， Gentoo的哲学是自由和选择。得益于一种称为Portage的技术，Gentoo能成为理想的安全服务器、开发工作站、专业桌面、游戏系统、嵌入式解决方案或者别的东西--你想让它成为什么，它就可以成为什么。由于它近乎无限的适应性，可把Gentoo称作元发行版。
2. Aech Linux(或称Arch):以轻量简洁为设计理念的Linux发行版。其开发团队秉承简洁、优雅和代码最小化的设计宗旨。

### linux 目录结构

登录系统后，在当前命令窗口下输入命令：
```
 ls / 
```
你会看到如下图所示:

![linux目录结构01](/lesson-10/linux目录结构01.png)

根目录 `/` 是整个文件系统的起始点，但实际上，Linux 文件系统采用了一种层次结构来组织文件和目录。

树状目录结构：
![linux目录结构02](/lesson-10/linux目录结构02.jpg)

根目录 `/` 下有许多子目录，这些子目录用于组织不同类型的文件和系统资源。以下是一些常见的子目录及其用途：

- `/bin` Binaries的缩写,包含可执行文件，用于存放基本的系统命令。
- `/boot`：这里存放的是启动 Linux 时使用的一些核心文件，包括一些连接文件以及镜像文件，例如内核和引导加载程序。
- `/dev`: Device的缩写, 该目录下存放的是 Linux 的外部设备，在 Linux 中访问设备的方式和访问文件的方式是相同的。
- `/etc`：Etcetera的缩写,包含所有的系统管理所需要的配置文件和子目录。
- `/home`：用户的主目录，每个用户都有一个对应的子目录，一般以用户的账号命名。(如上图中的 alice、bob 和 eve)
- `/lib`：Library的缩写，包含系统最基本的动态连接共享库，用于支持系统和应用程序，类似于 Windows 里的 DLL 文件。
- `/opt`：用于安装可选软件包的目录。
- `/proc`：Processes的缩写，/proc 是一种伪文件系统（即虚拟文件系统），存储的是当前内核运行状态的一系列特殊文件，这个目录是一个虚拟的目录，它是系统内存的映射，我们可以通过直接访问这个目录来获取系统信息。
  这个目录的内容不在硬盘上而是在内存里，我们也可以直接修改里面的某些文件，比如可以通过下面的命令来屏蔽主机的ping命令，使别人无法ping你的机器：
```
echo 1 > /proc/sys/net/ipv4/icmp_echo_ignore_all
```
- `/root`：系统管理员，超级用户（root）的主目录。
- `/sbin`：Superuser Binaries的缩写，包含系统管理员使用的系统命令。
- `/selinux`：这个目录是 Redhat/CentOS 所特有的目录，Selinux 是一个安全机制，类似于 windows 的防火墙，这里存放selinux的相关文件。
-`/tmp`：temporary的缩写,用来存放一些临时文件。
- `/usr`：unix shared resources的缩写，包含用户安装的应用程序和文件，类似于 windows 下的 program files 目录。
- `/var`：variable的缩写，包含可变的数据文件，例如日志文件和缓存文件。

值得提出的是 /bin、/usr/bin 是给系统用户使用的指令（除 root 外的通用用户），而/sbin, /usr/sbin 则是给 root 使用的指令。

***

/usr/bin： 系统用户使用的应用程序。

/usr/sbin： 超级用户使用的比较高级的管理程序和系统守护程序。

/usr/src： 内核源代码默认的放置目录。

/media: linux 系统会自动识别一些设备，例如U盘、光驱等等，当识别后，Linux 会把识别的设备挂载到这个目录下

/lost+found: 这个目录一般情况下是空的，当系统非法关机后，这里就存放了一些文件。

/srv: 该目录存放一些服务启动之后需要提取的数据。

/run：
一个临时文件系统，存储系统启动以来的信息。当系统重启时，这个目录下的文件应该被删掉或清除。如果你的系统上有 /var/run 目录，应该让它指向 run。

/sys：
这是 Linux2.6 内核的一个很大的变化。该目录下安装了 2.6 内核中新出现的一个文件系统 sysfs 。
sysfs 文件系统集成了下面3种文件系统的信息：针对进程信息的 proc 文件系统、针对设备的 devfs 文件系统以及针对伪终端的 devpts 文件系统。
该文件系统是内核设备树的一个直观反映。
当一个内核对象被创建的时候，对应的文件和目录也在内核对象子系统中被创建。

在 Linux 系统中，有几个目录是比较重要的，平时需要注意不要误删除或者随意更改内部文件。

/etc： 上边也提到了，这个是系统中的配置文件，如果你更改了该目录下的某个文件可能会导致系统不能启动。

/bin, /sbin, /usr/bin, /usr/sbin: 这是系统预设的执行文件的放置目录，比如 ls 就是在 /bin/ls 目录下的。




### linux 文件路径

在 Linux 中，简单的理解一个文件的路径，指的就是该文件存放的位置，
例如Linux文件系统的层次结构中提到的 /home/usr就表示的是 usr目录所在的路径。只要我们告诉 Linux 系统某个文件存放的准确位置，那么它就可以找到这个文件。

>Linux中的路径可以分为**绝对路径**和**相对路径**，因为根据档名写法的不同，也可以将所谓的路径（path）定义为绝对路径（absolute）和相对路径（relative）。这两种文件名/路径的写法根据是这样的：
>
> - 绝对路径：由跟目录（/）开始起的文件或者目录名称，例如 /home/dmtais/.bashrc：**（绝对路径的写法一定是由 / 目录写起的）**。
> - 相对路径：相对于目前路径的文件名写法。例如 ./home/dmtsai  或  ../ ../home/dmtsai  等。反正开头不是 / 属于相对路径的写法。**（相对路径的写法不是由 / 目录写起的）**。

 - ` . ` ：代表当前的目录，也可以使用  ./  来表示；
 - ` .. ` ：代表上一层目录，也可以 ../ 来表示；
```
$ cd /var/spool/mail/            ##进入到/var/spool/mail/下
$ cd ../lpd/                     ##返回上一层目录并进入lpd目录
$ cd ../..                       ##返回上边的两级目录回到/var目录下
```
![linux目录结构01](/lesson-10/linux目录结构02.png)


##### 总而言之

绝对路径是相对于根目录 / 的，只要文件不移动位置，那么**它的绝对路径是恒定不变的**；

而相对路径是相对于当前所在目录而言的，随着程序的执行，当前所在目录可能会改变，

因此文件的相对路径**不是**固定不变的。


### linux常用命令(以Ubuntu为主)

##### 安装

###### **Ubuntu（apt）**

- 安装软件包

```
sudo apt-get update  # 更新软件包列表
sudo apt-get install packageName  # 安装软件包
```

- 删除软件包

```
sudo apt-get remove packageName  # 删除软件包（保留配置文件）
sudo apt-get purge packageName   # 删除软件包及其配置文件
```

- 搜索软件包

```
sudo apt search packageName  # 搜索软件包
```

- 更新软件包列表

```
sudo apt-get update  # 更新软件包列表
```

###### **CentOS（yum）**

- 安装软件包

```
sudo yum update  # 更新软件包列表（yum）
sudo yum install packageName  # 安装软件包（yum）
```

- 删除软件包

```
sudo yum remove packageName  # 删除软件包
```

- 搜索软件包

```
sudo yum search packageName  # 搜索软件包（yum）
```

- 更新软件包列表

```
sudo yum update  # 更新软件包列表（yum）
```

###### **Alpine（apk）**

- 安装软件包

```
apk add packageName  # 安装软件包
```

- 删除软件包

```
apk del packageName  # 删除软件包
```

- 搜索软件包

```
apk search packageName  # 搜索软件包
```

- 更新软件包列表

```
apk update  # 更新软件包列表
```
***
### 命令格式

```
$command    [option(s)]   [argument(s)]

命令名   空格   选项    空格     参数
```
### **注意：**

- `每个命令行可使用的最多的命令字符是256个`
- `命令区分大小写`
- `不同的命令提示符使用分隔符号 “/”`
- `命令中的参数/选项可以是多个，并且参数其实就是要传入命令程序主函数main的参数`
- `[ ] 表示这个内容可以不包含，比如 [argument(s)] ，输入命令时可以不加参数`

### 一、基础操作
##### 1.1 关闭系统

```
  shutdown -h now    #或者 powerof，立刻关机
  shutdown -h 2      #两分钟后关机
```

##### 1.2 重启
```
  shutdown -r now    #或者 reboot，立刻重启
  shutdown -r 2      #两分钟后重启
```

##### 1.3 帮助命令（help）
```
  ifconfig  --help     #查看 ifconfig 命令的用法
```

##### 1.4 命令说明书（man）
(遇到一些函数和命令，不知道含义时可以是用命令man来查看帮助信息)
```
  man xxx    #打开命令说明后，可按"q"键退出
```

##### 1.5 切换用户（su）
```
  su yao    #切换为用户"yao",输入后回车需要输入该用户的密码
  exit      #退出当前用户
```

### 二、目录操作

##### 2.1 切换目录（cd）
```
  cd /                #切换到根目录
  cd /bin             #切换到根目录下的bin目录 (绝对路径)
  cd ../              #切换到上一级目录 或者使用命令：cd ..
  cd ~                #切换到home目录
  cd -                #切换到上次访问的目录
  cd xx(文件夹名)       #切换到本目录下的名为xx的文件目录，如果目录不存在报错 (相对路径)
  cd /xxx/xx/x        #可以输入完整的路径，直接切换到目标目录，输入过程中可以使用tab键快速补全 (绝对路径)
```

##### 2.2查看当前目录下所有文件

```
  ls                   #查看当前目录下的所有目录和文件
  ls -a                #查看当前目录下的所有目录和文件（包括隐藏的文件）
  ls -l                #列表查看当前目录下的所有目录和文件（列表查看，显示更多信息），与命令"ll"效果一样
  ls /bin              #查看指定目录下的所有目录和文件 
```

##### 2.3 创建目录（mkdir）
```
  mkdir tools          #在当前目录下创建一个名为tools的目录
  mkdir /bin/tools     #在指定目录下创建一个名为tools的目录
```

##### 2.4 删除目录与文件（rm）
```
  rm 文件名              #删除当前目录下的文件
  rm -f 文件名           #删除当前目录的的文件（不询问）
  rm -r 文件夹名         #递归删除当前目录下此名的目录
  rm -rf 文件夹名        #递归删除当前目录下此名的目录（不询问）
  rm -rf *              #将当前目录下的所有目录和文件全部删除
  rm -rf /*             #将根目录下的所有文件全部删除【慎用！相当于格式化系统】

```

##### 2.5 修改目录（mv）
```
  mv 当前目录名 新目录名          #修改目录名，同样适用与文件操作
  mv /usr/tmp/tool /opt       #将/usr/tmp目录下的tool目录剪切到 /opt目录下面
  mv -r /usr/tmp/tool /opt    #递归剪切目录中所有文件和文件夹
```

##### 2.6 拷贝目录（cp）
```
  cp /usr/tmp/tool /opt       #将/usr/tmp目录下的tool目录复制到 /opt目录下面
  cp -r /usr/tmp/tool /opt    #递归剪复制目录中所有文件和文件夹
```

##### 2.7 搜索目录（find）
```
  find /bin -name 'a*'        #查找/bin目录下的所有以a开头的文件或者目录
```

##### 2.8 查看当前目录（pwd）
```
  pwd                         #显示当前位置路径
```

### 三、文件操作

#### **3.00 传输文件(scp)**
```
	scp /opt/data.txt  192.168.1.101:/opt/    //将本地opt目录下的data文件发送到192.168.1.101服务器的opt目录下
```
[scp相关文章]：

https://www.runoob.com/linux/linux-comm-scp.html

https://blog.csdn.net/u012964600/article/details/135078489

##### 3.1 新增文件（touch）
```
   touch  a.txt         #在当前目录下创建名为a的txt文件（文件不存在），如果文件存在，将文件时间属性修改为当前系统时间
   vim a.txt            #vim工具
```

##### 3.2 删除文件（rm）
```
  rm 文件名              #删除当前目录下的文件
  rm -f 文件名           #删除当前目录的文件（不询问）
```

##### 3.3 编辑文件（vi、vim）
```text
  vi 文件名              #打开需要编辑的文件
  --进入后，操作界面有三种模式：命令模式（command mode）、插入模式（Insert mode）和底行模式（last line mode）
  命令模式
  -刚进入文件就是命令模式，通过方向键控制光标位置，
  -使用命令"dd"删除当前整行
  -使用命令"/字段"进行查找
  -按"i"在光标所在字符前开始插入
  -按"a"在光标所在字符后开始插入
  -按"o"在光标所在行的下面另起一新行插入
  -按"："进入底行模式
  插入模式
  -此时可以对文件内容进行编辑，左下角会显示 "-- 插入 --""
  -按"ESC"进入底行模式
  底行模式
  -退出编辑：      :q
  -强制退出：      :q!
  -保存并退出：    :wq
  ## 操作步骤示例 ##
  1.保存文件：按"ESC" -> 输入":" -> 输入"wq",回车     //保存并退出编辑
  2.取消操作：按"ESC" -> 输入":" -> 输入"q!",回车     //撤销本次修改并退出编辑
  ## 补充 ##
  vim +10 filename.txt                   //打开文件并跳到第10行
  vim -R /etc/passwd                     //以只读模式打开文件
```

##### 3.4 查看文件
```
  cat a.txt          #查看文件最后一屏内容
  less a.txt         #PgUp向上翻页，PgDn向下翻页，"q"退出查看
  more a.txt         #显示百分比，回车查看下一行，空格查看下一页，"q"退出查看
  tail -100 a.txt    #查看文件的后100行，"Ctrl+C"退出查看
```

### 四、文件权限

##### 4.1 权限说明
```
  文件权限简介：'r' 代表可读（4），'w' 代表可写（2），'x' 代表执行权限（1），括号为"8421法"
  ##文件权限信息示例：-rwxrw-r--
  -第一位：'-'就代表是文件，'d'代表是文件夹
  -第一组三位(rex)：拥有者的权限
  -第二组三位(rw-)：拥有者所在的组，组员的权限
  -第三组三位(r--)：代表的是其他用户的权限
```

##### 4.2 文件权限
```
  普通授权    chmod +x a.txt    
  8421法     chmod 754 a.txt     //1+2+4=7，"7"说明授予所有权限
```

### 五、打包与解压

##### 5.1 说明
```text
  .zip、.rar        //windows系统中压缩文件的扩展名
  .tar              //Linux中打包文件的扩展名 
  .gz               //Linux中压缩文件的扩展名
  .tar.gz           //Linux中打包并压缩文件的扩展名
```

##### 5.2 打包文件
```text
  tar -zcvf 打包压缩后的文件名 要打包的文件
  参数说明：z：调用gzip压缩命令进行压缩; c：打包文件; v：显示运行过程; f：指定文件名;
  示例：
  tar -zcvf a.tar file1 file2,...      //多个文件压缩打包
```

##### 5.3 解压文件
```text
  tar -zxvf a.tar                      //解包至当前目录
  tar -zxvf a.tar -C /usr------        //指定解压的位置
  unzip test.zip             //解压*.zip文件 
  unzip -l test.zip          //查看*.zip文件的内容 
```

### 六、其他常用命令

##### 6.1 find
```
  find . -name "*.c"     #将目前目录及其子目录下所有延伸档名是 c 的文件列出来
  find . -type f         #将目前目录其其下子目录中所有一般文件列出
  find . -ctime -20      #将目前目录及其子目录下所有最近 20 天内更新过的文件列出
  find /var/log -type f -mtime +7 -ok rm {} \;     #查找/var/log目录中更改时间在7日以前的普通文件，并在删除之前询问它们
  find . -type f -perm 644 -exec ls -l {} \;       #查找前目录中文件属主具有读、写权限，并且文件所属组的用户和其他用户具有读权限的文件
  find / -type f -size 0 -exec ls -l {} \;         #为了查找系统中所有文件长度为0的普通文件，并列出它们的完整路径
```

##### 6.2 whereis
```
  whereis ls             //将和ls文件相关的文件都查找出来
```

##### 6.3 which
```
  #说明：which指令会在环境变量$PATH设置的目录里查找符合条件的文件。
  which bash             #查看指令"bash"的绝对路径
```

##### 6.4 sudo

```
  #说明：sudo命令以系统管理者的身份执行指令，也就是说，经由 sudo 所执行的指令就好像是 root 亲自执行。需要输入自己账户密码。
  #使用权限：在 /etc/sudoers 中有出现的使用者
  sudo -l                              #列出目前的权限
  $ sudo -u yao vi ~www/index.html     #以 yao 用户身份编辑  home 目录下www目录中的 index.html 文件
```

##### 6.5 grep
```
  grep -i "the" demo_file              #在文件中查找字符串(不区分大小写)
  grep -A 3 -i "example" demo_text     #输出成功匹配的行，以及该行之后的三行
  grep -r "ramesh" *                   #在一个文件夹中递归查询包含指定字符串的文件
```

##### 6.6 service
```
  说明：service命令用于运行System V init脚本，这些脚本一般位于/etc/init.d文件下，这个命令可以直接运行这个文件夹里面的脚本，而不用加上路径
  service ssh status      #查看服务状态 
  service --status-all    #查看所有服务状态 
  service ssh restart     #重启服务 
```

##### 6.7 free
```
  #说明：这个命令用于显示系统当前内存的使用情况，包括已用内存、可用内存和交换内存的情况 
  free -g            #以G为单位输出内存的使用量，-g为GB，-m为MB，-k为KB，-b为字节 
  free -t            #查看所有内存的汇总
```

##### 6.8 top
```
  top               #显示当前系统中占用资源最多的一些进程, shift+m 按照内存大小查看
```

##### 6.9 df
```
  #说明：显示文件系统的磁盘使用情况
  df -h            #一种易看的显示
```

##### 6.10 mount
```
  mount /dev/sdb1 /u01              #挂载一个文件系统，需要先创建一个目录，然后将这个文件系统挂载到这个目录上
  dev/sdb1 /u01 ext2 defaults 0 2   #添加到fstab中进行自动挂载，这样任何时候系统重启的时候，文件系统都会被加载 
```

##### 6.11 uname
```
  #说明：uname可以显示一些重要的系统信息，例如内核名称、主机名、内核版本号、处理器类型之类的信息 
  uname -a
```

##### 6.12 yum
```
  #说明：安装插件命令
  yum install httpd      #使用yum安装apache 
  yum update httpd       #更新apache 
  yum remove httpd       #卸载/删除apache 
```

##### 6.13 rpm
```text
  #说明：插件安装命令
  rpm -ivh httpd-2.2.3-22.0.1.el5.i386.rpm      #使用rpm文件安装apache 
  rpm -uvh httpd-2.2.3-22.0.1.el5.i386.rpm      #使用rpm更新apache 
  rpm -ev httpd                                 #卸载/删除apache 
```

##### 6.14 date
```
  date -s "01/31/2010 23:59:53"   #设置系统时间
```

##### 6.15 wget
```
  #说明：使用wget从网上下载软件、音乐、视频 
  wget http://xxx
  #下载文件并以指定的文件名保存文件
  wget -O nagios.tar.gz http://xxx
```

##### 6.16 ftp
```
   ftp IP/hostname    #访问ftp服务器
   mls *.html -       #显示远程主机上文件列表
```

##### 6.17 管道 ｜
竖线符号 `|` 是管道符号（Pipe Symbol）。该符号用于将一个命令的输出传递给另一个命令作为输入，以便二者之间进行数据传输和处理。

```
command1 | command2
ls | grep "pattern"
```

### 七、系统管理

##### 7.1 防火墙操作
```
  service iptables status      #查看iptables服务的状态
  service iptables start       #开启iptables服务
  service iptables stop        #停止iptables服务
  service iptables restart     #重启iptables服务
  chkconfig iptables off       #关闭iptables服务的开机自启动
  chkconfig iptables on        #开启iptables服务的开机自启动
  ##centos7 防火墙操作
  systemctl status firewalld.service     #查看防火墙状态
  systemctl stop firewalld.service       #关闭运行的防火墙
  systemctl disable firewalld.service    #永久禁止防火墙服务
```

##### 7.2 查看网络
```
  ifconfig
```

##### 7.3 查看端口号占用情况
```
lsof -i:8080
```

##### 7.4 修改IP
```text
  修改网络配置文件，文件地址：/etc/sysconfig/network-scripts/ifcfg-eth0
  ------------------------------------------------
  主要修改以下配置：  
  TYPE=Ethernet               //网络类型
  BOOTPROTO=static            //静态IP
  DEVICE=ens00                //网卡名
  IPADDR=192.168.1.100        //设置的IP
  NETMASK=255.255.255.0       //子网掩码
  GATEWAY=192.168.1.1         //网关
  DNS1=192.168.1.1            //DNS
  DNS2=8.8.8.8                //备用DNS
  ONBOOT=yes                  //系统启动时启动此设置
  -------------------------------------------------
  修改保存以后使用命令重启网卡：service network restart
```

##### 7.5 配置映射
```text
  修改文件： vi /etc/hosts
  在文件最后添加映射地址，示例如下：
   192.168.1.101  node1
   192.168.1.102  node2
   192.168.1.103  node3
  配置好以后保存退出，输入命令：ping node1 ，可见实际 ping 的是 192.168.1.101。
```

##### 7.6 进程操作
```
  ps -ef         #查看所有正在运行的进程
  kill pid       #杀死该pid的进程
  kill -9 pid    #强制杀死该进程   
```

##### 7.7 查看链接
```
  ping IP        #查看与此IP地址的连接情况
  netstat -an    #查看当前系统端口
  netstat -an | grep 8080     #查看指定端口8080
```

##### 常用快捷键
```
 #停止进程
 ctrl + c
 #清屏
 ctrl + l
```

[Linux命令大全]：
http://www.runoob.com/linux/linux-command-manual.html


[一个linux网课]:
https://www.lanqiao.cn/courses/1


#### 上传二进制文件到服务器

这就是编译好的二进制文件，可以在任何linux上运行

我们可以通过以下命令，将此文件上传到服务器

```
#将当前目录下main文件拷贝到127.0.0.1的/home/路径下
scp main root@127.0.0.1:/home/
```

如果想要免密登录 配置SSH即可

1.在主机A创建密钥对

```
ssh-keygen #创建证书
```

然后均回车（选择默认）

2、将文件上传至免登录主机B的authorized_keys

```
/root/.ssh/authorized_keys
```

#### 上传源代码到服务器

通常会先将源码上传到github再拉下来，或者使用goland一键上传

[服务器go安装](https://blog.csdn.net/qq_43098070/article/details/126075629)

然后进入项目目录直接编译即可

```
go build main.go
```

#### 运行程序

现在服务器上有编译好的二进制文件了，那我们如何运行它？

只需要

```
./main
```

非常的简单

但是我们会发现无法运行，因为文件没有权限

这时候只需要

```
chmod +x main
chmod 733 main
```

给它高权限再次运行即可

这时候，使用云服务器的同学们需要去控制台找到安全组的配置规则，在入方向添加相应端口

##### 后台运行

如果按以上方式运行，当我们关闭窗口时会发现程序被终止了

那有什么办法可以让程序在后台运行呢？

我们可以使用 tmux，systemctl 或者 nohup+& 命令

###### nohup+&

```
nohup ./main &
```

###### tmux

[tmux](https://www.ruanyifeng.com/blog/2019/10/tmux.html)

###### systemctl

[systemctl设置自己的systemd.service服务设置守护进程](https://blog.csdn.net/lonnng2004/article/details/88964763)

[systemctl 详解](https://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-commands.html)

