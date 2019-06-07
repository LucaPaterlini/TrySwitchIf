## Switch vs If then else if

### Intro
Hello,
this is a test just to show the diffent performances of switch 
vs if then else if.

### Test

I have added the specs of the hw I have used in mySpecs.html,
in short retail hp laptop 2 real cores 4 threads, 4gb memory

```goos: linux
   goarch: amd64
   BenchmarkTryIf-4        	2000000000	         0.41 ns/op
   BenchmarkTry2If-4       	2000000000	         0.76 ns/op
   BenchmarkTrySwitch-4    	2000000000	         0.38 ns/op
   BenchmarkTrySwitch2-4   	2000000000	         0.78 ns/op
   PASS
```

### Conclusion

I was expecting a slower speed using switch instead of using if then else if
and Instead they are compareble and switch marginally faster,
The unexpected result is that re executing each function gives a benchmark
with the double of the execution time.

### Wish
Could someone advice on how there is this behave?
