# Constant

Constants are types of variables whose value cannot be changed. Initialization of values is only done once at the beginning, after that the variable cannot be changed in value.

# Constant use

Cara penerapan konstanta sama seperti deklarasi variabel biasa, selebihnya tinggal ganti keyword var dengan ```const```.

````
const firstName string = "john"
fmt.Print("halo ", firstName, "!\n")
````

The type inference technique can be applied to constants, by simply eliminating the data type at the time of the declaration.

```
const lastName = "wick"
fmt.Print("nice to meet you ", lastName, "!\n")
```

# function usage ```fmt.Print()```

This function has the same role as the ```fmt.Println ()``` function, the differentiation of the ```fmt.Print ()``` function does not generate a new line at the end of its output.

Another difference is, the values of the parameters entered into the function are combined without separators. Unlike the ```fmt.Println ()``` function whose parameter values are combined using a space connector.

```

fmt.Println("john wick")
fmt.Println("john", "wick")

fmt.Print("john wick\n")
fmt.Print("john ", "wick\n")
fmt.Print("john", " ", "wick\n")

```

The code above shows the difference between ```fmt.Println ()``` and ```fmt.Print ()```. The output produced by the 5 statements above is the same, although the method used is different.

When using ```fmt.Println ()``` there is no need to add spaces to each word, because the function will automatically add it between values. In contrast to ```fmt.Print ()```, spaces need to be added, because this function does not add spaces between the merged parameters.