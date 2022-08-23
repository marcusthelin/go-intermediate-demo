# Producer-Consumer Scenario

The producer reads in messages from a mockstream and a consumer is processing the data to find out whether the message is about golang or not. The task is to modify the code inside `main.go` so that producer and consumer can run concurrently to increase the throughput of this program.

## Expected results:
Before:
```
message about golang
message not about golang
message about golang
message about golang
message about golang
Process took 3.580866005s
```

After:
```
message about golang
message not about golang
message about golang
message about golang
message about golang
Process took 1.977756255s
```
