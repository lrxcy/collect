# My First Proxy Project
To be more clearly what proxy (especially reverse proxy) do, I decide to write mime. Any refernce or feature would be recorded below...

# pre-requeriement
1. docker
2. golang
3. privilage to use sudo

# quick start
0. build `binary` for reverse proxy with `go build`
1. run a docker conatiner for dummy website `docker run -it -p 8080:80 -d nginx:latest` ... hereby, use 8080 as origin port
2. execute binary with `sudo` previliage


# feature
... on-going

# bug
1. connection keep-alive...

# refer
- https://github.com/jim0409/go-proxy-example
- https://hackernoon.com/writing-a-reverse-proxy-in-just-one-line-with-go-c1edfa78c84b