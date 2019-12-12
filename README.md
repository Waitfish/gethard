# getHard
查看硬盘利用率和内存利用率,如果超过指定阀值就把消息推送到钉钉机器人.
## Download bin
`https://github.com/Waitfish/gethard/releases`
## Usage
```crontab
*/5 * * * * /usr/local/bin/gethard  -m disk -p /data -w 90 -d xxxx -h MARIADB_105
*/5 * * * * /usr/local/bin/gethard  -m disk -p / -w 90 -d xxx -h MARIADB_105
*/5 * * * * /usr/local/bin/gethard  -m mem  -w 90 -d xxx -h MARIADB_105
```

```bash
[root@xxx ~]# gethard --help
Usage of gethard:
  -d string
    	The DingDing token (default "dingcode")
  -h string
    	The host name you want to push (default "myServer")
  -m string
    	The mode you want to show (default "disk")
  -p string
    	The disk mount you want to know (default "/")
  -w float
    	UsedPercent to warn dingding (default 90)

```
钉钉上收到的消息
```
MARIADB_105  /data 硬盘利用率已经超出 90%
MARIADB_105  / 硬盘利用率已经超出 90%
MARIADB_105    内存利用率已经超出 90%

```
## build for linux amd64
```bash
env GOOS=linux GOARCH=amd64 go build gethard.go
```
