# go-benchmark

Small benchmarks about Golang functions and packages.

To run simply execute the tests like this:

```bash
go test -bench=. -benchmem
```

The "Def" in the test function name means "Default", as in "using the go builtin package".

# Result

```
$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: mabench
cpu: Intel(R) Core(TM) i5-8365U CPU @ 1.60GHz
BenchmarkChannels-8                              2160372               572.2 ns/op             0 B/op          0 allocs/op
BenchmarkChannelsBuffered-8                      7838100               161.6 ns/op             0 B/op          0 allocs/op
BenchmarkMutexDeferUnlock-8                     43951117                25.67 ns/op            0 B/op          0 allocs/op
BenchmarkMutexUnlock-8                          62244088                19.09 ns/op            0 B/op          0 allocs/op
BenchmarkJsonDefMarshalBig-8                        3837            538646 ns/op           49319 B/op          2 allocs/op
BenchmarkJsonJsoniterMarshalBig-8                   2816            355334 ns/op           49406 B/op          2 allocs/op
BenchmarkJsonJzonMarshalBig-8                       3747            312898 ns/op           49422 B/op          2 allocs/op
BenchmarkJsonGojsonMarshalBig-8                     4014            277369 ns/op           49683 B/op          2 allocs/op
BenchmarkJsonDefUnmarshalBig-8                       763           1562475 ns/op           10116 B/op       3004 allocs/op
BenchmarkJsonJsoniterUnmarshalBig-8                 3153            342729 ns/op            9646 B/op       2997 allocs/op
BenchmarkJsonJzonUnmarshalBig-8                     1218           1123726 ns/op          242667 B/op       3035 allocs/op
BenchmarkJsonGojsonUnmarshalBig-8                   2739            373004 ns/op           49665 B/op          1 allocs/op
BenchmarkJsonDefMarshalSmall-8                   1000000              1095 ns/op             144 B/op          2 allocs/op
BenchmarkJsonJsoniterMarshalSmall-8              1337359               886.9 ns/op           144 B/op          2 allocs/op
BenchmarkJsonJzonMarshalSmall-8                  1437928               830.0 ns/op           144 B/op          2 allocs/op
BenchmarkJsonGojsonMarshalSmall-8                1672812               703.7 ns/op           144 B/op          2 allocs/op
BenchmarkJsonDefUnmarshalSmall-8                  357294              3059 ns/op             239 B/op          7 allocs/op
BenchmarkJsonJsoniterUnmarshalSmall-8            2403166               548.3 ns/op            16 B/op          3 allocs/op
BenchmarkJsonJzonUnmarshalSmall-8                2288898               524.7 ns/op            16 B/op          3 allocs/op
BenchmarkJsonGojsonUnmarshalSmall-8              2722290               449.5 ns/op            48 B/op          1 allocs/op
BenchmarkCreationMap-8                          67732000                17.10 ns/op            0 B/op          0 allocs/op
BenchmarkCreationSyncMap-8                      510674442                2.269 ns/op           0 B/op          0 allocs/op
BenchmarkAddMap-8                                 755418              1974 ns/op             167 B/op          4 allocs/op
BenchmarkAddLockMap-8                             727017              1968 ns/op             153 B/op          4 allocs/op
BenchmarkAddSyncMap-8                             618416              2736 ns/op             259 B/op         10 allocs/op
BenchmarkGetMap-8                                 955684              1199 ns/op              16 B/op          2 allocs/op
BenchmarkGetLockMap-8                             964971              1158 ns/op              16 B/op          2 allocs/op
BenchmarkGetSyncMap-8                             709327              1631 ns/op              16 B/op          2 allocs/op
BenchmarkSnapshotMap-8                                 2         560841877 ns/op        102759192 B/op      5034 allocs/op
BenchmarkSnapshotSyncMap-8                             2         769401794 ns/op        83966424 B/op       4117 allocs/op
BenchmarkRandMathV1intn-8                       50162834                23.75 ns/op            0 B/op          0 allocs/op
BenchmarkRandMathV2intn-8                       77402388                15.43 ns/op            0 B/op          0 allocs/op
BenchmarkRandMathV1int63-8                      82569901                14.47 ns/op            0 B/op          0 allocs/op
BenchmarkRandMathV2int63-8                      93737835                12.76 ns/op            0 B/op          0 allocs/op
BenchmarkRandMathV1sourced-8                    270161906                4.311 ns/op           0 B/op          0 allocs/op
BenchmarkRandMod-8                              484884189                2.310 ns/op           0 B/op          0 allocs/op
BenchmarkNetHttpGetFMT-8                         3958876               280.3 ns/op             4 B/op          0 allocs/op
BenchmarkNetHttpGetIO-8                          5296562               210.6 ns/op             6 B/op          0 allocs/op
BenchmarkNetHttpGetWrite-8                       5745277               199.4 ns/op             7 B/op          1 allocs/op
BenchmarkGinGet-8                                5342622               252.2 ns/op            54 B/op          1 allocs/op
BenchmarkEchoGet-8                               4512482               250.6 ns/op            15 B/op          1 allocs/op
BenchmarkGorillaMuxGet-8                         1000000              1399 ns/op             852 B/op          7 allocs/op
BenchmarkDefPostJson-8                            192050              8505 ns/op            6720 B/op         26 allocs/op
BenchmarkGinPostJson-8                            172776              9194 ns/op            6768 B/op         27 allocs/op
BenchmarkEchoPostJson-8                           190849              8677 ns/op            6720 B/op         26 allocs/op
BenchmarkWebPanic-8                               296112              3403 ns/op             148 B/op          5 allocs/op
BenchmarkWebLog-8                                 263684              4515 ns/op             141 B/op          6 allocs/op
PASS
ok      mabench 63.205s
```

# Learnings

- Using buffered channels is faster than non-buffered ones
- Try to unlock a mutex manually instead of using defer all the time (if no panic can occur)
- The default encoding/json package is quite slow
- A sync/map is slower than using a map with a mutex and read/write function
- Creating a new math/rand source creates random numbers faster
- It is way faster to use w.Write instead of fmt.Fprint in your http HandleFuncs
- Web frameworks are not always faster, just more convenient
- It is faster to log your http errors via panic/recover instead of creating an extra log line manually with a Logger

