# Super Fast Time

In some cases, we don't need to use time.Now() for getting Second/Minute, for example, getting time in a fair queue.

For performance, we only need the "increment" timestamp.

## Usage

```
fastime.Now()
```

Now() returns "increment" timestamp in ns, which may be walltime or not a walltime, in other words, **_the time is not corrent all the time._**

## Benchmark

```
goos: darwin
goarch: arm64
pkg: github.com/MeteorsLiu/fastime
cpu: Apple M2
BenchmarkTimeNow
BenchmarkTimeNow-8    	35143980	        32.01 ns/op
BenchmarkFastTime
BenchmarkFastTime-8   	80893653	        14.99 ns/op
PASS
ok  	github.com/MeteorsLiu/fastime	3.556s
```
