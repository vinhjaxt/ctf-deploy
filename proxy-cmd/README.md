# Proxy CMD
```
Net <==> CMD
```

# Usage
```
Usage: ./proxy-cmd [OPTIONS] -- CMD [ARGUMENTS...]
  -b string
        Bind address (default ":9999")
  -read-timeout duration
        Connection read timeout (default 5s)
  -run-timeout duration
        Run timeout (default 2m0s)
  -stderr-print
        Print stderr to connection
  -stdin-as-argument
        Use stdin as an argument pass to the CMD
  -stdin-maxsize int
        Maximum number of bytes stdin as argument (default 51200)
  -w string
        Working directory (default "/opt")
Example: ./proxy-cmd -w /opt -b :9999 -- cat /etc/passwd
```
# Example
```
./proxy-cmd -w /opt -b :9999 -- cat /etc/passwd
./proxy-cmd -w /opt -b :9999 -stdin-as-argument -- cat
./proxy-cmd -w /opt -b :9999 -- cat
./proxy-cmd -w /opt -b :9999 -- pwd

```
