# goLang
Golang notes from the course : https://www.coursera.org/learn/golang-getting-started/

## Introduction to Go lang 
### object oriented : 
Golang is NOT an object oriented. 
The structs in Golang are very simplified, it doesn't have **Inheritance**, **constructors** nor **generics**. Which make the golang code simple and quick to run. 
It compiles directly to machine code : 

![image](https://user-images.githubusercontent.com/42012627/183110023-17b73dac-003f-429b-b889-38efd1d073e6.png)




### Data types :

#### variables : 
constants : ```const CONSTANT = "constant"```

*it is possible to declare and initialize a variable like this var a = "a". If we want to declare a variable and attribue a value to it later we need to specify the type of the variable. Golang is a statically typed language*

#### type casting : 
The type casting is the convertion of a data type to another. 
Unlike Java that has an implicite type conversion, Golang requires an explicite type conversion. We do this like this 
```
var simpleInt int64 = 13
var simpleFloat float64 = float64(simpleInt)
```

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
the varaible's scope is organized in go with "{}" they constitute a **block** 
We can declare package level variables that are accessible to the hole package. **this means that all the files within the same package have access to these variables** 
But it is always better to declare variables as local as possible. This is because of the memory deallocation and the readiblity of the code 

To summarise we have three types of varible scopes : 
* Local variables : Within functions or blocks 
* Package variables : Shared between all the files of the same package 
* Global variables : Exported variables used between packages. To exporte a variable we start by a capitale letter.

*note: Strings are immutable*

### Packages : 
A package contains multiple files. For example a the package main contains many files. The file that has the main() function is the entrypoint. To run the go script we usually use the command go run main.go, we need to add all the files of the package in command. an alternative is to use 
```
go run .
```

### common operations :

#### sort array :
*requires the sort package*
```
sort.Ints(numbers)
```

#### merge to slices 
*it is important to add the "..."* 
```
slice = append(slice, slice2...)
```

### Deallocation of memory : 
image.png
For example this function f(), can be run 1M times. And if the variable x doesn't get deallocated it will be created 1M times. 
In the traditionnal languages, memory is handeled like this : 
image.png
A variable can be stored in two different spaces :  
- Stack : affected to variables within functions. The memory is deallocated as soon as the function call ends 
- Heap : Affected to other variables, it is persitant and doesn't end as soon as it becomes useless. It needs to be deallocated manually


### Control flow : 
The control flow is the order in which the statements are exeecuted.
But it can change if the developer uses some specific statements such as : *if, for loop, switch case, *

#### Loops : 
Loops in golang are very simple. There is only one type of loops, for.
When used alone like for {}, it is equivalent to a while(true)
When used on a slice and array data structure it is used like this 
```
for index, booking := range bookings {
```
Sometimes we don't need the index variable. And since Go generates an error over non used variables, we can ask Go to ignore it by replacing with '_'.



#### Arrays :
Similat to all the languages 
Except :
* it is initialized to 0 value initially (0 of the expected type)
* it has a fixed size but is not immutable (we can change its values)

**to get the length of the array**
len(arrayName)

![image](https://user-images.githubusercontent.com/42012627/180615715-70e8d624-79ce-472a-a80c-7916e2725a43.png)

this is how we declare an array of size 5 

![image](https://user-images.githubusercontent.com/42012627/180615725-4f82ab34-ec28-475c-b515-389fc466b3c5.png)

here the "..." mean that we dont specify an exact size to the array 

to iterate over an array in golang : 

![image](https://user-images.githubusercontent.com/42012627/180615841-c0438daa-5cc8-4d76-993f-a5d735953f93.png)

with i the index and v the value 


#### Slice : 
It is a window of an array. It has a variable size up to the whole array. 

It is declared like this : 

```
var slice []int
```

**Slices are windows on an underlying array** 

It is different from arrays because we don't specify the length between the brackets

It has three properties : 
* **Pointer** thaat indifcates the start of the slice 

![image](https://user-images.githubusercontent.com/42012627/189490365-4a1f1f8b-b5dd-42ee-8b00-1f3649c3e35b.png)


* **The length** of the slice 
* **The capacity**, the maximum number of elements in the slice. (length - beguining of the slice)

![image](https://user-images.githubusercontent.com/42012627/180997736-03fd21ef-177a-4cf1-872a-c24e7e87e48a.png)

![image](https://user-images.githubusercontent.com/42012627/180998231-387b66d9-1392-42d6-8248-6e486414db13.png)

**Very important**

![image](https://user-images.githubusercontent.com/42012627/180998350-1b16e5d7-1e6d-4f72-a82c-b36d684a3fec.png)

![image](https://user-images.githubusercontent.com/42012627/180998671-19ef56cd-5a05-4b96-b9ca-b4e12d994b78.png)

Useful functions 
* make() that initialize arrays.
* append() that increases the size of the slice 

#### functions : 
**Super important** : 
In go lang the parameters are passed as call by value. So the parameters passed to the functions are deep copied to its arguments. 
this will return 2 : 

![image](https://user-images.githubusercontent.com/42012627/189536558-aa1b6fca-b78d-47a9-ba0f-33b6577bf8aa.png)

This is great but this call by value solution has a drawback. Making a copy of each argment takes a certain amout of type. For complex structs it can be significative. To avoid this lapse of time we can use call by referance arguments. like this : 

![image](https://user-images.githubusercontent.com/42012627/189536888-aa07669b-d298-42a4-9e62-fee6794cf4e8.png)

**important: no need to do this for slices they already are pointers.**
outside of the main function, declared like this : 

```
func greetUser(theaterName string, remainingTickets uint, theaterTicekets uint) {
	fmt.Println("Welcome to ", theaterName, " !")
	fmt.Printf("we have total of %v places and %v avaiable\n", theaterTicekets, remainingTickets)
}
```
to return a value we can use : 
```
func printFirstNames(bookings []string, firstNames []string) []string{
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
```
It is also possible to return multiple values : 

```
func validateUser(firstName string, lastName string, email string, nbTickets int, remainingTickets uint) (bool, bool, bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := nbTickets < int(remainingTickets) && nbTickets > 0
	return isValidName, isValidEmail, isValidTicketNumber
}
```

##### Well written functions : 
* Understandability : If someOne is looking for a feature (for example the bluring) he can find it easly
* Give functions a good name 
* Function do only one action 
* Keep functions short

##### first class values : 
From functional programming and it means treating a function as a normal variable. 

#### Hash tables : 

Key value pair just like in other languages. See here https://github.com/ShameGod/HashMaps

![image](https://user-images.githubusercontent.com/42012627/181000977-f37b7a4d-02b4-454c-801f-42fbe9e15413.png)


##### Maps : 
A map is the implementation of hashtables in GoLang.
A map can be delcared by type like this : 
```
var myMap map[string]string
```

or creating an empty map directly 
```
myMap := make(map[string]string)
```
**To iterate over a map :** 

```
for key,value := rang map{
	fmt.Printf("key %v and value %v", key, value)
}

```


#### Structs : 
Groups related variables together (like name, address, email) and can have variables of multiple types 
We define structs on a **package level** like this : 
```
type UserData struct {
	firstName string
	lastName  string
	email     string
	nbTickets int
}
```
To instantiate a Struct we do this : 
```
var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		nbTickets: nbTickets,
	}
```
Or we can initialise the structure using default field values with the function new()
```
userData := new(UserData)
```

### Trading Data : 
Go lang has a way of organising traded data between services. 
**RFC : Request For Comment** defines a protocol of data format
Examples of RFCs are : Json, Avro, Xml... 
golang has packages to encode and decode data using these protocols. 
#### Marshalling and Unmarshalling Json : 
Mashalling means converting golang struct to json object. 
we can do this like this : 
```
barr, err := json.Marshal(userData)
```
barr for binary array will contain the binaries for a json. Err If something wrong happens if nothing happens it takes a Nil value.

We do an unmarshalling like this 
```
var userData UserData 

err := json.Unmarshal(barr, userData)
```
the Json needs to fit the userData, they need to have the same attributes.

### Equivalent of classes in golang :

![image](https://user-images.githubusercontent.com/42012627/189546801-ecf79fba-0cd9-404a-8616-abe09bcf7dfa.png)

by adding the (mi MyInt) in my function, I am saying that the function Double is only available for the MyInt type.

In the picture above there is an implicit parameter, mi. We do not add a parametre before the v.Double() but the real parameter is the " v ". 
```
func (user UserData) sendEmail() {
	time.Sleep(1 * time.Second)
	fileName := fmt.Sprintf("bill of %v %v", user.firstName, user.lastName)
	err := ioutil.WriteFile(fileName, []byte(user.email), 0777)
	fmt.Print(err)
	fmt.Println("#####################################")
	fmt.Printf("Sending the email to the user %v %v", user.firstName, user.lastName)
	fmt.Println("#####################################")
}
```
#### Error handling : 
Many go function return an error as a second argument. It is necessary to handle them like this : 

```
f, err := os.Open("file.txt")
if err != nil {
	fmt.print(err)
	return
}
```

#### Encapsulation : 
It is possible to do encapsulation by exporting functions giving access to variables, by setting a capital letter. 

#### Polymorphisme : 
A polymorphique function is a function that **does two different operations according to the input argument**. For example a function area() that calculates the area of a form. It is going to do different operations according to its input (a triangle, a square, a circle ...) 
This aspect that is very present in OOP requires inheritance and overriding. 
In order to get Polymorphisme we use **interfaces**

#### Interfaces : 
they have the same definition as the Java interfaces. **Used to express conceptual similarity between types**
They define a list of methods. All the types that have **ALL** these methods are concidered of the type of the interface. 
If a define an interface called shape2D with two methods, area and perimeter. 
if the types triangle and circle implemente these two methods they are concidered of type shape2D 

```
type shape2D interface {
	Area() float64
	Perimeter() float64
}
type Triangle{....}
func (t Triangle) Area() float 64 {....}
func (t Triangle) Perimeter() float 64 {....}
```
That is it, the class Triangle satisfies the shape2D interface.
This is not declared explicitly golang just understands that you meant to set Triangle as a shape2D 

##### Concrete vs Interface types : 
An interface has two unknown thigs, its exact type and its value. For example shape2D has an unknown type and an unknown value. For example : 
```
type shape2D interface {
	Area() float64
	Perimeter() float64
}
var s shaped2D
t :=  new(Triangle)
s = t
```
the interface s has now a type but No value. **it is still possible to call the speak() function**

##### Interfaces usecases : 
* a function which takes multiple types of parameter
```
func IsTooBig(s Shape2D) bool{
	if(s.Area()>100 && s.Perimeter() > 100) {
		return true
	}
	return false 
}
``` 
This function takes an interface as an argument and uses its functions.
* Empty interface is an interface that specifies no methods and that every type satisfies for example **rare syntax** :
```
func PrintMe(val interface{}){
	fmt.PringIn(val)
}
```

###### Type assersion : 
Sometimes we need to know the exact type of an interface in a function not like the IsTooBig() function. Like I want to take any shape2D but I will call type specific functions 

```
func DrawShape(s Shape2D){
	rect, ok := s.(Rectangle)
	if ok{
		DrawRectangle()
	}
	tri, ok := s.(Triangle)
	if ok{
		DrawTriangle()
	}
}
```

But there are eventually too much strcuts that satisfy the interface. We can use the switch method to make it easier

```
func DrawShape(s Shape2D){
	swich := sch := s.(type){
	case Rectangle :
		DrawRectangle()
	case Triangle : 
		DrawTriangle()
	}
}
```

##### Commun example of interface use 
Many functions built in go packages, return an Error interface defined like this : 
```
type error interface {
	Error() string 
}
```

### Concurrency : 
The motivation of concurrency is the need of speed. Concurrency is here to over come the **performance limitation**. 
Moure's law says : Number of transistors doubles every 18months. But that's not the case anymore. We arrived to the limits of hardware optimisation (except for the case of quantum computers). 
In order to go faster with the same hardware we have two solutions : 
*  parallelism : Depends on the hardware, 4 cores => 4 parallel tasks. It is very difficult to program (conflicts, sharing resources between threads ..) 
*  Concurrency : management of many tasks in the same time. It enables parallelism. It is exactly like parallelism but without the drawbacks. So concurrency would have : *synchronization and communication between tasks*  

![image](https://user-images.githubusercontent.com/42012627/189488950-b3e56c2a-37f8-49b2-8428-7ad89c2778ed.png)

ti charge a new thread to do something we do this 
```
go concurrentFunction()
```
if the main thread finishes executing its lines of code it exits and the goroutine we started for the concurrentFunction never gets to finish being executed.
to do that we do : 
```
var wg = sync.WaitGroup{}
main {
	..... code .....
	wg.Add(1)
	go concurrentFunction()
	..... code .....
	wg.Wait()
}

func concurrentFunction(){
	..... logic code .....
	wg.Done()
}
```

**the difference between golang and other languages with concurrency, is that Go lang uses an abstraction of real OS threads called go routines.**
**If the Computer has 4 cores it has 4 threads, but we can create thousands of gorroutines**
## ___________________________ Concurrency 2 : depper _____________

### Definitions and basics : 
**a process** : is an instance of a running program. Every process has **its own** memory, some code, registers(super fast memories), program counter (tell the process what's the next action he is going to execute )

**Operating System** : on of the main goal of an OS is to allow many processes to execute concurrently. To do that the OS switches quickly between process. For example linux allows a process to acces the CPU for an average of **20ms** it switches so quickly that we feel like it is parallel 

**scheduling** : it is the order in which the processes are run, there many scheduling algrithmes. The basic one is the round and robin. Other algorithms take priority in count for the scheduling. **Scheduling is the main task in OS**

**context switch** : it is the action of changing the process being executed. When we stop a process A to start a process B, we save its state(memory, code, register..) where we arrived that is called a **process Context** in the memory. The OS does the context switch.  

![image](https://user-images.githubusercontent.com/42012627/191314619-fbe1ee62-4dce-47b7-92e9-f1c0c5367a20.png)

**Threads** : Before we only had prosses and the problem with them is that swiching the context takes too much time (reading from memory). A process contains multiple threads. the threads share some of their context, which makes thread switching quicker than process switching. The yellow part is shared between threads and the green part is specific to the thread 

![image](https://user-images.githubusercontent.com/42012627/191316362-b170c9cd-2dc0-4ba4-95f3-db1f05b7a4df.png)

**Goroutines** : 
we can have multiple goroutines in the same thread. We put many goroutines inside the main thread(when we only have one thread in a process). It is the Go Runtime Scheduler that switches between goroutines.

![image](https://user-images.githubusercontent.com/42012627/191316934-11e40f74-8ac8-45c7-966a-32342e8f9d9c.png)

The scheduling of goroutines : 
We have a main thread (1 thread in the processor). So the OS schedules the main threads. Every main thread has a Logical Processor. The the GRS(Go Runtime Scheduler) runs on the Logical processor to schedules the Goroutines. In order to do parallelism + concurrency, we add more logical processors according to the number of Cores that we have. **The programmer can change the number of Logical Processor (by default 1)** 

![image](https://user-images.githubusercontent.com/42012627/191318314-b43bf5ad-b6d3-40fd-870b-273abc021f7a.png)


Goroutines are better than threads for the following reasons :

![image](https://user-images.githubusercontent.com/42012627/191325098-c83d9b82-a5c7-498a-ba04-2b7ed0b50d06.png)

(source : [geeks4geeks](https://www.geeksforgeeks.org/golang-goroutine-vs-thread/))

### Interleaving and race condition : 
When something crashes in the app, we can refer to the **control flow** to know where it crashed exactly and the lines that were runned before that. But when we use concurrency it gets complicated. **We can't predict the order of multiple concurrent operations**. This problem is called **interleaving**.
A program has to not depend on the order of concurency, otherwise it will be in **race condition**. The code below is a race condition

![image](https://user-images.githubusercontent.com/42012627/191327049-cf71bcb6-a362-4fe9-b2c7-498ead639f0d.png)

This happens because the tasks don't communicate. Goroutines communicate easier than threads. 

### Wait groups : 
There is a package called sync that provides functions to synchronize goroutines. *sync.WaitGroup* forces a goroutine to wait for other goroutines to finish. It is possible to wait on more than one goroutine. 

```
func createFile(wg *sync.WaitGroup){ 
   err := ioutil.WriteFile(fileName, []byte(user.email), 0777)
   fmt.Print(err)
   wg.Done()
}

func main(){
	var wg Synch.WaitGroup
	wg.Add(1)
	go createFile(&wg)
	wg.Wait()
	fmt.Println("the file was created")

}
```

### Channels : 
So far we have noticed that goroutines can only receive data in the beguining (passed as arguments). But thanks to channels, it is possible for goroutines to send and receive data during their execution. **channels are typed**. 

```
func prod(a int, b int, c chan int){
	c <- a * b 
}
func main(){
	c := make(chan int)
	go prod(1,2,c)
	go prod(4,5,c) 
	a := <- c
	b := <- c
	fmt.Println(a * b)
}
```

**When channels are executing, a synchronization is done implicitly. when a goroutine sends somethink to a channel, it is blocked until an other goroutine receives data from that channel.**

```
func prod(a int, b int, chan c){
	c <- a * b 
	//This goroutine will be blocked untill another goroutine receives the int sent into the channel
	fmt.Printf("the result of the product is %v", a*b)
}
func main(){
	c := make(chan int)
	go prod(1,2,c)
	a := <- c // the main goroutine will be blocked untill it receives some data 
	fmt.Println(a)
}
```

### Buffered channels : 
It it possible to make goroutines communicate without synchronization. By fixing a buffer size, the goroutine sending values to channels will continue to run. It will be blocked if the number of values it sent has reached the buffer's capacity. The capacity of the buffer is precised when calling the *make()* function

### Receive from multiple channels whith select 
abort
```
for{
  select{
    case a <- c:
      fmt.Print(a)
    case <-abort:
      return
  }
}
```

default pattern
```
for{
  select{
    case a <- c1:
      fmt.Print(a)
    case b <- c2:
      fmt.Print(b)
    default:
      fmt.Print("still waiting for a result")
  }
}
```

## Common erros : 

### fatal error: all goroutines are asleep - deadlock
This happens when the main goroutine finishes executing but the other goroutines didn't finish. 
It happened to me because I forgor to add the **&** to the **wg** variable of type WaitGroup




## Booking app : 
I will be explain the development process of a booking app with Golang : 

### 1 : the main file : 

When we run a go lang application with many files, go looks for a file where it will start running (**entrypoint**) .
We recognize this file with the function ``` func main() ```
**Every go application can only have one main function, one entrypoint** 

#### printf : 
Thanks to this function we can write print code like this : 
fmt.Printf("we have total of %v and we have %v avaiable", theaterTicekets, remainingTickets)
the 'v' in %v stands for variable there are other ways to display values such as %s, with s for string .. 
another usefull tip is to use %T to get the Type of the variable passed as an argument

The fmt package handles input and output
##### Scan function : 
to read the input of the user we use the function scan like this. fmt.Scan(&input)
The scan function takes what the user types and puts in a place in memory. The Scan function then takes the address of this input and put it in the pointer of the variable entered as an argument.
**very important** : **the scan function can only take one word, as soon as the user uses a white space the scan stops** 
it is recommanded to use another method to read the user input 
 ```
 br := bufio.NewReader(os.Stdin)
 input,_,_ := br.ReadLine()
 ```

#### variable types : 
we can either write a variable like : var a = 4 and golang will imply that it is an int. or we can specify the type for more precision. For example var a int32 = 4

![image](https://user-images.githubusercontent.com/42012627/189489919-879d57a8-94e2-4e96-9cd4-8b5b4732feac.png)

We can also declare the variable a like thuis : a := 4











