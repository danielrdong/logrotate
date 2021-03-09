# logrotate

Logrotate: A Go package for writing logs to rolling files, Which forked from [natefinch/lumberjack](https://github.com/natefinch/lumberjack), and did
some customization.

### Using

```go
    import "github.com/danielrdong/logrotate"

    log.SetOutput(&logrotate.Logger{
        Filename:   "xxx.log",
        MaxSize:    500,  // megabytes
        MaxBackups: 3,
        MaxAge:     28,   // days
        Compress:   true, // disabled by default
    })
```

### Version
v1.0.0
