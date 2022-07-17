# goLang
Golang notes from tutorials 

## Introduction to Go lang 
### object oriented : 
Golang is an object oriented. Golang doesn't use the term *class* it uses *structs* 
The structs in Golang are very simplified, it doesn't have **Inheritance**, **constructors** nor **generics**. Which make the golang code simple and quick to run. 
### Concurrency : 
The motivation of concurrency is the need of speed. Concurrency is here to over come the **performance limitation**. 
Moure's law says : Number of transistors doubles every 18months. But that's not the case anymore. We arrived to the limits of hardware optimisation (except for the case of quantum computers). 
In order to go faster with the same hardware we have two solutions : 
*  parallelism : add more cores but it is very difficult to program (conflicts, sharing resources between threads ..)
*  Concurrency : management of many tasks in the same time. It enables parallelism. It is exactly like parallelism but without the drawbacks. So concurrency would have : *synchronization and communication between tasks*  [needs more explenation]

### Data types :
#### pointers : 
Which a pointer to an address to data in memory 
it has two operators : 
- & operator : returns the address of a variable/function  
- * operator : returns the data at an address (dereferencing)
image.png
In this piece of code, we have : 
x as an integer who's value is 1, y is not initialized(gets 0 by default) and we declare a variable ip. The variable ip is a pointer to an integer. 
- When we do ```ip = &x```, we give ip the value of the address of x, so ip points now to x
- When we do ```y = *ip```, we give y the data of its address

##### the function new()
this function creats a variable and returns a pointer to this variable instead of the data in it.
``` ptr := new(int)
    *ptr = 3    
```
*note: the type of space allocated to a variable depends on its type. For example int has a range of values*

### Scope of variables : 
The scope of a variable is the places in the code where the variable can be accessed 
image.png
the varaible's scope is organized in go with "{}" they constitute a **block** 

### Deallocation of memory : 
image.png
For example this function f(), can be run 1M times. And if the variable x doesn't get deallocated it will be created 1M times. 
In the traditionnal languages, memory is handeled like this : 
image.png
A variable can be stored in two different spaces :  
- Stack : affected to variables within functions. The memory is deallocated as soon as the function call ends 
- Heap : Affected to other variables, it is persitant and doesn't end as soon as it becomes useless. It needs to be deallocated manually
