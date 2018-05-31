# Log sampling

Log sampling tries to reduce CPU and I/O pressure by only recording a subset of entries and 
dropping duplicate log entries. A log entry having the same log level and message content 
are considered duplicate. 

The logger maintains a separate bucket for each log entry. 
At each tick, the sampler will emit the first N  initial logs in each bucket and every Mth log 
therafter. Sampling loggers are safe for concurrent use.
In this example, we will emit the first 5 messages then one every each 100 messages thereafter.

```console
go run src/sampling/main.go

Log sampling to reduce the presure on I/O and CPU by reducing logging
Using the built in sampler. You probably need to wrap the whole zapcore Sampler public methods if you need to write our own custom sampler
We will first emit the first 5 messages then one every each 100 messages thereafter
{"level":"info","ts":1527544371.1515696,"msg":"test at info","n":1}
{"level":"info","ts":1527544371.1515965,"msg":"test at info","n":2}
{"level":"info","ts":1527544371.1516032,"msg":"test at info","n":3}
{"level":"info","ts":1527544371.1516066,"msg":"test at info","n":4}
{"level":"info","ts":1527544371.15161,"msg":"test at info","n":5}
{"level":"info","ts":1527544371.1518133,"msg":"test at info","n":105}

```


