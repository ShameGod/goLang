# goLang
Golang notes from the course : https://www.coursera.org/learn/golang-getting-started/

## Introduction to Go lang 
### object oriented : 
Golang is an object oriented. Golang doesn't use the term *class* it uses *structs* 
The structs in Golang are very simplified, it doesn't have **Inheritance**, **constructors** nor **generics**. Which make the golang code simple and quick to run. 
It compiles directly to machine code : 

![image](https://user-images.githubusercontent.com/42012627/183110023-17b73dac-003f-429b-b889-38efd1d073e6.png)



### Concurrency : 
The motivation of concurrency is the need of speed. Concurrency is here to over come the **performance limitation**. 
Moure's law says : Number of transistors doubles every 18months. But that's not the case anymore. We arrived to the limits of hardware optimisation (except for the case of quantum computers). 
In order to go faster with the same hardware we have two solutions : 
*  parallelism : add more cores but it is very difficult to program (conflicts, sharing resources between threads ..)
*  Concurrency : management of many tasks in the same time. It enables parallelism. It is exactly like parallelism but without the drawbacks. So concurrency would have : *synchronization and communication between tasks*  

![image](https://user-images.githubusercontent.com/42012627/189488950-b3e56c2a-37f8-49b2-8428-7ad89c2778ed.png)


### Data types :

#### variables : 
constants : ```const CONSTANT = "constant"```

*it is possible to declare and initialize a variable like this var a = "a". If we want to declare a variable and attribue a value to it later we need to specify the type of the variable. Golang is a statically typed language*



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
### Packages : 
A package contains multiple files. For example a the package main contains many files. The file that has the main() function is the entrypoint. To run the go script we usually use the command go run main.go, we need to add all the files of the package in command. an alternative is to use 
```
go run .
```


### Deallocation of memory : 
image.png
For example this function f(), can be run 1M times. And if the variable x doesn't get deallocated it will be created 1M times. 
In the traditionnal languages, memory is handeled like this : 
image.png
A variable can be stored in two different spaces :  
- Stack : affected to variables within functions. The memory is deallocated as soon as the function call ends 
- Heap : Affected to other variables, it is persitant and doesn't end as soon as it becomes useless. It needs to be deallocated manually

### Type conversion in Golang 
for example 
![image](https://user-images.githubusercontent.com/42012627/179506443-c14cf386-942f-4e0f-911a-d1d7b1ce7e1d.png)
This will fail. 
We need to convert their types with the T() operation : 
![image](https://user-images.githubusercontent.com/42012627/179506607-796b816d-f411-483b-b5f6-a871bc84a6a7.png)

*note: Strings are immutable*

### Control flow : 
The control flow is the order in which the statements are exeecuted.
But it can change if the developer uses some specific statements such as : *if, for loop, switch case, *
![image](https://user-images.githubusercontent.com/42012627/179508932-5a18929f-d0ce-4a63-a551-b21ae9c61074.png)

![image](https://user-images.githubusercontent.com/42012627/179509237-16959b34-a6f8-4fd2-b26a-47b5722dee25.png)

![image](https://user-images.githubusercontent.com/42012627/179509341-3927d4c9-b0d1-471d-bae5-83abb8d8785d.png)

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
![image](https://user-images.githubusercontent.com/42012627/180998811-d3441b4c-ee30-41fd-a042-ed937962dc3b.png)

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


#### Hash tables : 

Key value pair just like in other languages. See here https://github.com/ShameGod/HashMaps

![image](https://user-images.githubusercontent.com/42012627/181000977-f37b7a4d-02b4-454c-801f-42fbe9e15413.png)


#### Maps : 
A map can be delcared by type like this : 

```

var myMap map[string]string
```

or creating an empty map directly 
```
myMap := make(map[string]string)
```

![image](https://user-images.githubusercontent.com/42012627/182373927-c21497a8-b4df-4402-ab7f-9f95a42c18ee.png)

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

#### variable types : 
we can either write a variable like : var a = 4 and golang will imply that it is an int. or we can specify the type for more precision. For example var a int32 = 4

![image](https://user-images.githubusercontent.com/42012627/189489919-879d57a8-94e2-4e96-9cd4-8b5b4732feac.png)

We can also declare the variable a like thuis : a := 4














