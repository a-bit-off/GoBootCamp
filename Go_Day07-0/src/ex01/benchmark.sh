GOGC=off go test -bench=. -cpuprofile=benchmark.prof > top10.txt
go tool pprof benchmark.prof 