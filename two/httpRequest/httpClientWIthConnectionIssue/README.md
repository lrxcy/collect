# quick check
> watch -n1 "netstat -n | grep -i 9999 | grep -i time_wait | wc -l"

# refer:
- http://tleyden.github.io/blog/2016/11/21/tuning-the-go-http-client-library-for-load-testing/


# extend-refer:
> use netstat and lsof on macOS
- https://www.cnblogs.com/blackay03/p/8836135.html