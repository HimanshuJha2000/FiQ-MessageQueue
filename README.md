# producer-consumer

## Assumptions : 

1. Key is being considered as what topic is in kafka.

2. All the messages are saved in this hierarchy that first main(fig_queue) directory is created inside then the subdirectry of topic will be created and then the message is being passed in a file post that the count is considered and whatever the count is those many messages will be replicated inside that file.


## How To Run :

1. Start the main.go by command `go run cmd/main.go`

2. Choose what to operate :`0` for consumer and `1` for producer

3. For Producer, the message format should be key:message:processing_time:count
 Sample message will be : bjsjdw:jhsasajsajkhsaS:5s:2

4. For consumer, pass the number of concurrent workers you want to operater- <concurrent_worker>

5. Thats't it! Enjoy the operations!
