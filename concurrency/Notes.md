## Concurrency Models: 
# 1. Processor Threads: Process, Thread, Mutex, ...
- Advantages: Control, Performance, Responsive User Interface
- Disadvantages: Poor Performance, Memory Consumption, Shared Memory, Race Conditions

#2. Events: Emitter, Receiver, Event Object
- Advantages: Memory Isolation, Separate Callee from Caller
- Disadvantages: Traceability, Synchronizing Receivers

#3. Callbacks and Promises: Callback, Promise
- Advantages: Memory Isolation, Simplify Async Operations
- Disadvantages: Pyramid of Doom, Handling Multiple Receivers

#4. Communicating Sequential Processes: Actor, Message
- Advantages: Fully Decoupled, Multiple Handlers, Memory Isolation
- Disadvantages: Complicated Mental Model, Traceability

## Concurrency in Go:
# 1. No Thread Primitives

# 2. Goroutines: Light, Scale, Schedule
- Memory Requirements: Processor Thread 1MB ===> Green Thread 2KB
- Parallelism in Go: GOMAXPROCS - sets the maximum number of CPUs that can be executing simultaneously 
  and returns the previous setting 

# 3. Channels 