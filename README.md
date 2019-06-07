## Switch vs If then else if

### Intro
Hello,
this is a test just to show the diffent performances of switch 
vs if then else if.

### Test

I have added the specs of the hw I have used in mySpecs.html,
in short retail hp laptop 2 real cores 4 threads, 4gb memory

####Repeating the benchmarks 2 times

```goos: linux
   goarch: amd64
   BenchmarkTryIf-4        	2000000000	         0.41 ns/op
   BenchmarkTry2If-4       	2000000000	         0.76 ns/op
   BenchmarkTrySwitch-4    	2000000000	         0.38 ns/op
   BenchmarkTrySwitch2-4   	2000000000	         0.78 ns/op
   PASS
```

#### Rebating the benchmarks 3 times

```
goos: linux
goarch: amd64
BenchmarkTryIf-4        	2000000000	         0.40 ns/op
BenchmarkTry2If-4       	2000000000	         0.78 ns/op
BenchmarkTry3If-4       	2000000000	         0.39 ns/op
BenchmarkTrySwitch-4    	2000000000	         0.76 ns/op
BenchmarkTrySwitch2-4   	2000000000	         0.38 ns/op
BenchmarkTrySwitch3-4   	2000000000	         0.77 ns/op
```

### Conclusion

I was expecting a slower speed using switch instead of using if then else if
and Instead they are compareble and switch marginally faster,
The unexpected result is that re executing each function gives a benchmark
with the double of the execution time.

### Wish
Could someone advice on how there is this behave?

## Idea

Maybe the benchmarks in odd position are executed by the cpu with no load due to the
operative system while the other one is executed using the cpu that is also running the os

Tested, 
```
goos: linux
goarch: amd64
BenchmarkTryIf      	2000000000	         0.40 ns/op
BenchmarkTry2If     	2000000000	         0.78 ns/op
BenchmarkTry3If     	2000000000	         0.38 ns/op
BenchmarkTrySwitch  	2000000000	         0.80 ns/op
BenchmarkTrySwitch2 	2000000000	         0.42 ns/op
BenchmarkTrySwitch3 	2000000000	         0.79 ns/op
```
same result with a single processor