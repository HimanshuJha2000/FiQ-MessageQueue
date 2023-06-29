# producer-consumer

## Assumptions : 

1. Key is being considered as a topic in kafka.

2. All the messages are saved in this hierarchy that first topic directory is created inside then the message is being passed and then the numbe rof count of messages


## How To Run :

1. Start the main.go by command `go run cmd/main.go`

2. Choose what to operate :`0` for consumer and `1` for producer

3. For Producer, the message format should be : <key>:<message>:<processing_time>:<count>

Sample message : bjsjdw:jhsasajsajkhsaS:5s:2

4. For consumer, pass the number of concurrent workers you want to operater: <concurrent worker>

5. Thats't it! Enjoy the operations!
