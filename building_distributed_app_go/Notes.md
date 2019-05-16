## 1. Distributed programming

## 2. Introduction to RabbitMQ

## 3. Publish & Subscribe to messages

## 4. Using WebSockets

## Distributed System: 
A distributed system is a software system in which components located 
on networked computers communicate and coordinate their actions by passing message.

Target Audience: Concurrency, Web Applications, Open Mind.

Course Layout: RabbitMQ, Publishing Messages, Consuming Messages, Persisting Data, Updating Web Clients.

Tools and Libraries: RabbitMQ, Postgres, Pure Go.

## Introduction to RabbitMQ
- Message Broker: Message broker is an intermediary program module which translates a message from
    the formal messaging protocol of the sender to the formal messaging protocol of the receiver.
- Message Broker (RabbitMQ): Message brokers receive messages from publishers and route them to consumers.
Pub ------> Message broker (protocol - 11010...) -------> Sub

- RabbitMQ Cli: Interacting with RabbitMQ from the Command Line

sudo docker run {container_id}

sudo docker run -d --hostname my-rabbit --name trunglv-rabbit -p 8080:15672 rabbitmq:3-management
=> You can then go to http://localhost:8080 or http://host-ip:8080 in a browser

sudo docker exec -i -t {container_id} /bin/bash #by ID

rabbitmqctl status

rabbitmqctl list_queues

rabbitmqctl cluster_status

- RabbitMQ interface
- Integrating RabbitMQ into Go
