# GRPCvsTTRPC

Benchmarks
---
SmallPayload Test: Is a small payload with each atomic type of string, int, double, and bool being set

LargeArrayPayload Test: Contains large byte slices of 10kb each where there are 4 elements.

LargeNArrayPayload Test: Contains an array of size 48 with an integer defined.

LargePayload Test: Contains a 10mb payload


```
BenchmarkGRPCSmallPayload-8          	   10000	    178074 ns/op	   18168 B/op	     336 allocs/op
BenchmarkGRPCLargeArrayPayload-8     	    3000	    509941 ns/op	  415507 B/op	     370 allocs/op
BenchmarkGRPCLargeNArrayPayload-8    	    2000	    845890 ns/op	  731696 B/op	     452 allocs/op
BenchmarkGRPCLargePayload-8          	     500	   3399196 ns/op	13480294 B/op	     315 allocs/op
BenchmarkTTRPCSmallPayload-8         	   20000	     63346 ns/op	    2926 B/op	      69 allocs/op
BenchmarkTTRPCLargeArrayPayload-8    	    5000	    257488 ns/op	  417835 B/op	      80 allocs/op
BenchmarkTTRPCLargeNArrayPayload-8   	    3000	    382780 ns/op	  750700 B/op	     157 allocs/op
BenchmarkTTRPCLargePayload-8         	     200	   7897544 ns/op	20992103 B/op	      64 allocs/op
```
