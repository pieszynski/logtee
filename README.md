# logtee
Opinionated app for writing log files `tee`-like from stdin/pipe.

Runs like this:

```shell
app_generating_logs | logtee app.log
```

and writes to daily tagged files `app.YYYYMMDD.log` based on passed log file name parameter.

## Done

- [&check;] like `tee` but for log files
- [&check;] auto insert date in log file name

## Planned

- [&cross;] rotate logs

## Development

```shell
# test run from json log lines
./generator.sh | go run logtee.go access.log

# compile as debugger attachable
go build -gcflags="all=-N -l" -o logtee logtee.go
./generator.sh | ./logtee access.log

# publish multitarget
go mod tidy
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o dist/logtee.amd64 logtee.go
env GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o dist/logtee.arm64 logtee.go

# list possible go ARM targets
go tool dist list | grep arm 
```