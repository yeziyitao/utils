# 1 shell.exec("")
+ shell.exec("ls -lh") --output result
+ shell.exec("lss -lh") --output err

# go run main.go
```
total 24
-rw-r--r--  1 yezi  staff    93B 11 26 00:17 Readme.md
-rw-r--r--  1 yezi  staff    43B 11 26 00:04 go.mod
-rw-r--r--  1 yezi  staff   155B 11 26 00:16 main.go
drwxr-xr-x  3 yezi  staff    96B 11 26 00:16 shell
 <nil>
/bin/bash: lss: command not found
 exit status 127
```
