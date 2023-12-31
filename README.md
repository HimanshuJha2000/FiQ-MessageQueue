# FiQ-MessageQueue

## Features :


- Supports multiple consumers(consumer group) functionality.
- Supports a CLI based interface for easy operation.
- Have relevant validations for input message.
- Can run multiple worker threads within single consumer to process messages.
- Retains messages even if the producer/consumer stops working.
- Supports topic based message proceassing and storing.


## Assumptions : 

1. Key is being considered as what topic is in kafka.

2. All the messages are saved in this hierarchy that first main(fig_queue) directory is created inside then the subdirectry of topic will be created and then the message is being passed in a file post that the count is considered and whatever the count is those many messages will be replicated inside that file.

## How To Run 

### In Local:

```shell
   git clone <repository_url>
   cd FiQ-MessageQueue
   go build -o fiq_message_queue cmd/main.go
    ./fiq_message_queue
```
- `0` for consumer and `1` for producer

- For Producer, the message format should be key:message:processing_time:count
 Sample message will be : my_topic:my_message:5s:message_count

- For consumer, pass the number of concurrent workers you want to operater- <concurrent_worker>

### Contributing
Contributions to FiQ are welcome! If you find a bug or want to add a new feature, please open an issue or submit a pull request.

### Contact
For any questions or inquiries, please contact at himanshu.jha1702@gmail.com.


