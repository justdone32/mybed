package e

var IniStr = `[runMode]
RunMode = debug

[server]
HTTPPort = 8000
ReadTimeout = 60
WriteTimeout = 60

[app]
PageSize = 9
JwtSecret = 23347$040412
UploadTmpDir = "static/upload"

[database]
DBType = mysql
DBUser = root
DBPwd = 123456
DBHost = 127.0.0.1:3306
DBTableName = mytbed
DBPath = mybed.db

[redis]
RedisAddr = 127.0.0.1:6379
RedisPWD =
RedisDB =
`
