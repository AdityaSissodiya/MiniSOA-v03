# MiniSOA-v03

A minimalistic Service-Oriented Architecture (SOA) solution in Go/Golang

Components of the framework:

● A Service Provider (A random number generator)

● A Service Consumer (A periodic random number displayed in a terminal)

● An Orchestrator (A look-up Service providing information about how to connect to a 
service)

The system works as follows: 
The Consumer sends a request over TCP/IP to a service 
registry (IP and port), in this case, called the Orchestrator, asking for a random number service
(e.g., send a string value such as “serviceRequest=randomNumberGenerator”). 
The Orchestrator responds with the IP and port of the Provider. This information can be hard
coded in the Orchestrator. The Consumer should then use the IP and port received from the 
Orchestrator to request a random number from the Provider once per second and print the 
results to screen.
The Provider is a simple service that should provide the consumer with current time (e.g., Unix
timestamp or format of choice) and a random number between 1-100 each time the Consumer 
requests it.

![image](https://github.com/AdityaSissodiya/MiniSOA-v03/assets/19986905/4835cb4a-9d9a-465f-b0da-265755c9bd65)

Console Screenshots -

random_number_provider.go : 

![image](https://github.com/AdityaSissodiya/MiniSOA-v03/assets/19986905/1fc3ff8f-580e-4616-8534-665226f6557a)

orchestrator.go : 

![image](https://github.com/AdityaSissodiya/MiniSOA-v03/assets/19986905/6de182f5-8a50-441c-8089-191510dc85f0)

random_number_consumer.go : 

![image](https://github.com/AdityaSissodiya/MiniSOA-v03/assets/19986905/c0301b32-4e51-4634-920d-1fc3fd9e347f)
