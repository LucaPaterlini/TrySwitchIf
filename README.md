## Switch vs If then else if

### Intro
Hello,
this is a test just to show the different performances of switch 
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

#### Repeating the benchmarks 3 times

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
Looks like thare is something strange, why should re executing the same function
take the double of the time of the even number of attempt

### Mistake
In the first version I have writen a >0 as first condition
and then both where choosing the first condition,anyway
the tests should execute in the same time

## Idea
How can I get a all the benchmark tested on the same vCPU-core?
Maybe the benchmarks in odd position are executed by the cpu with no load due to the
operative system while the other one is executed using the cpu that is also running the os

Code
```
func init(){
	runtime.LockOSThread()
	runtime.GOMAXPROCS(1)
}


```


Tested, 
```
goos: linux
goarch: amd64
BenchmarkTryIf      	2000000000	         0.39 ns/op
BenchmarkTry2If     	2000000000	         0.37 ns/op
BenchmarkTry3If     	2000000000	         0.37 ns/op
BenchmarkTrySwitch  	2000000000	         0.38 ns/op
BenchmarkTrySwitch2 	2000000000	         0.38 ns/op
BenchmarkTrySwitch3 	2000000000	         0.37 ns/op
PASS
```
looks like the problem was the usage of two diffent cores,
one with no load an one that at the same time is running the os

### Fix

I have fixed the test and removed my add to make it single threaded

Tested 
```
goos: linux
goarch: amd64
BenchmarkTryIf-4        	2000000000	         1.72 ns/op
BenchmarkTry2If-4       	2000000000	         1.73 ns/op
BenchmarkTry3If-4       	2000000000	         1.73 ns/op
BenchmarkTrySwitch-4    	2000000000	         1.73 ns/op
BenchmarkTrySwitch2-4   	2000000000	         1.73 ns/op
BenchmarkTrySwitch3-4   	2000000000	         1.74 ns/op
PASS
```


### Conclusion

thanks to advice of  https://twitter.com/empijei?lang=en
I have made it single threaded using the same core thread just adding
to the test file
```
func init() {
	runtime.LockOSThread()
	runtime.GOMAXPROCS(1)
}

```

#### Added a bind of the thread all on the same cpu-thread
```
goos: linux
goarch: amd64
BenchmarkTryIf-4        	2000000000	         1.75 ns/op
BenchmarkTry2If-4       	2000000000	         1.75 ns/op
BenchmarkTry3If-4       	2000000000	         1.74 ns/op
BenchmarkTrySwitch-4    	2000000000	         1.73 ns/op
BenchmarkTrySwitch2-4   	2000000000	         1.74 ns/op
BenchmarkTrySwitch3-4   	2000000000	         1.74 ns/op
```

While I have sorted how to bound a thread to a cpu to compare having the same cpu
capacity looks like the differences in between switch and if then else are
not perceivable at this scale