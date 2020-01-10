# intro
```
-X importpath.name=value
  Set the value of the string variable in importpath named name to
```

# execute command
way to process program directly
```
go run -ldflags "-X main.xyz=`git log|head -1|awk '{print $2}'`" main.go
```

way to build the binary with specific version
```
go build -ldflags "-X main.xyz=`git log|head -1|awk '{print $2}'`" main.go
```

# refer
- https://stackoverflow.com/questions/11354518/application-auto-build-versioning
- https://www.reddit.com/r/golang/comments/4cpi2y/question_where_to_keep_the_version_number_of_a_go/
