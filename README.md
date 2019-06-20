Simplish
========

Simple english, or **simplish** for short, is a dynamic expression based programming language.
Simplish **is not** a general purpose programming language since it lacks many
important features associated to those langauges such as:

* control flow
* loops
* functions
* classes

Instead, **simplish** enables developers to write simple english-like grammar and have it be runnable code. As such, developers can use simplish to write business logic in plain english.

### Examples
```
 ____ ___ __  __ ____  _     ___ ____  _   _
/ ___|_ _|  \/  |  _ \| |   |_ _/ ___|| | | |
\___ \| || |\/| | |_) | |    | |\___ \| |_| |
 ___) | || |  | |  __/| |___ | | ___) |  _  |
|____/___|_|  |_|_|   |_____|___|____/|_| |_|
>> 
>> age is 10
>> age equals 10
true
>> age greater than 20
>> false
>> insurance is "Geico"
>> name not equals "Geico"
false
>> motorcyclist is false
>> age not greater than 21 and not motorcyclist
true
```
### Data Types

* bool
* int
* float
* string
* sequence

### Operators

| Operator  | Description                                    | 
|-----------|------------------------------------------------|
| is        | Assigs a value to a left handside identifier   |
| equals    | Compares left hand expression to right hand expression for equality |
| and       | Logical and |
| or        | Logical or  |
| not       | Inverts a bool value of the right hand expression |
| greater than | Returns wether the left hand value is greater than to the right hand value |
| less than | Returns wether the left hand value is less than to the right hand value |
| greater/equals | Returns wether the left hand value is greater than or equal to the right hand value |
| less/equals | Returns wether the left hand value is less than or equal to the right hand value |
| in | Returns wether the left hand value is found in the right hand sequence |

### Acknowledgements

This project is heavily influenced by learnings from the book [Writing an interpreter in Go](https://interpreterbook.com/) by Mitchell Hashimoto. If you are interested in 
programming language design, interpreters, and compilers I highly recommend his work.

