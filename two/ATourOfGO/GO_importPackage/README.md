### default $GOPATH is /Users/mac/go
change import "path" if needed

### notice with test code
use rune instead of byte to avoid unix input error cause test failure

### funciton test (Refer to)
- https://openhome.cc/Gossip/Go/Testing.html
1. test code's file name should include 'test.go'
2. import package must be the last element of the import 'path'