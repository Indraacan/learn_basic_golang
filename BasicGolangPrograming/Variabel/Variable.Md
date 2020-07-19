#Variable

Go adopts two types of variable writing, those that write the data type, and those that don't. Both methods are valid and the purpose is the same, the difference is just the way of writing it.

#Declaration of Variables With Data Types

Go has quite strict rules in terms of writing variables. When declaration, the data type used must be written down as well. Another term for this concept is **manifest typing**.

``` 
package main

import "fmt"

func main() {
    var firstName string = "john"

    var lastName string
    lastName = "wick"

    fmt.Printf("halo %s %s!\n", firstName, lastName)
}
````

The ```var``` keyword above is used for variable declarations, for example, can be seen in ```firstName``` and ```lastName```.

The value of the ```firstName``` variable is filled directly when the declaration, different from ```lastName``` whose value is filled after the declaration code line, something like this is allowed on Go

#Use of Functions ```fmt.Printf```

This function is used to display output in a certain form. It has the same function as the ```fmt.Println ()``` function, only the output structure is defined at the beginning.

Note the ```"hello% s% s! \ n"``` section, the ```%s``` character there will be replaced by the string data that is in the 2nd, 3rd, and so on parameters.

Another example, the following three lines of code will produce the same output, even though the way of writing is different.

```
fmt.Printf("halo john wick!\n")
fmt.Printf("halo %s %s!\n", firstName, lastName)
fmt.Println("halo", firstName, lastName + "!")

```

The plus sign (+) if placed between strings, the function is to concatenate strings or string concatenation.

The ```fmt.Printf ()``` function does not generate a new line at the end of the text, therefore a newline literal,``` \ n```, is used to bring up a new line at the end. This is very different compared to the ```fmt.Println ()``` function which automatically generates a new line at the end.

#Declaration of Variables Without a Data Type

In addition to manifest typing, Go also adopts the concept of type inference, which is a variable declaration method whose data type is determined by the value data type, contradictory when compared to the first method. With this type of method, the var keyword and data type don't need to be written.

```
var fName string = "john"
lName := "wick"

fmt.Printf("halo %s %s!\n", firstName, lastName)
```

#Multi Variable Declaration

Go supports the declaration method of many variables simultaneously, the way to write the variables with a comma delimiter (,). For charging the value is allowed simultaneously.

```
var first, second, third string
one, two, three = "satu", "dua", "tiga"

````

#Variabel UnderScore ```_```

Go has unique rules that are rarely owned by other languages, i.e. none of the variables can be idle. That is, all variables declared must be used. If there are variables that are not used but declared, an error will appear during compilation and the program will not be able to run.

Unemployment variable

Underscore (_) is a reserved variable that can be used to hold values that are not used. You could say this variable is a wastebasket

```
    _ = "learn Golang"
    _ = "Golang is something"
    name, _ := "john", "wick"

```

In the example above, the name variable will contain john text, while the wick value is contained by the underscore variable, indicating that the value will not be used.

The underscore variable is predefined, so there is no need to use ```:=``` to fill in the value, just enough ```=```. However, specifically for filling multi-variable values performed by the type inference method, there may be an underscore variable inside.

Usually the underscore variable is often used to accommodate the return value of functions that are not used.

Please note, that the contents of the underscore variable cannot be displayed. Data that has entered the variable will be lost. Like an underscore variable such as a blackhole, any object that enters into it will be trapped forever in the singularity and will not be able to get out.

#Declaration of Variables Using the ```New``` Keyword

    