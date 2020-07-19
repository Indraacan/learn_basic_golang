# Data Types

Go recognizes several types of data types, including numeric data types (decimal & non-decimal), strings, and booleans.

In the previous chapters we unconsciously applied several data types, such as ```string``` and ```int``` numeric types.

This chapter will explain some of the standard data types provided by Go, along with how to use them.

# Non-Decimal Numeric Data Type

There are several types of non-decimal or non-floating point numeric data types in Go. In general there are 2 types of data in this category that need to be known.

- uint, data type for count numbers (positive numbers).
- int, data type for integers (negative and positive numbers).

The two data types above are further divided into several types, with division based on the width of the value range, details can be seen in the following table.

|Data Types  |number coverage                                             |
|------------|------------------------------------------------------------|
|```uint8``` | 0 ↔ 255                                                    |
|```uint16```| 0 ↔ 65535                                                  |
|```uint32```| 0 ↔ 4294967295                                             |
|```uint64```| 0 ↔ 18446744073709551615                                   |
|```uint```  | same as ```uint32``` or ```uint64``` (depending on value)  |
|```byte```  | same as ```uint8```                                        |
|```int```   | -128 ↔ 127                                                 |
|```int16``` | -32768 ↔ 32767                                             |
|```int32``` | -2147483648 ↔ 2147483647                                   |
|```int64``` | -9223372036854775808 ↔ 9223372036854775807                 |
|```int```   | same as ```int32``` or ```int64``` (depend on value)       |
|```rune```  | same as ```int32```                                        |

It is recommended not to arbitrarily determine variable data types, as much as possible the selected type must be adjusted to its value, because the effect is to the variable memory allocation. The selection of the right data type will make memory usage more optimal, not excessive.

```
var positiveNumber uint8 = 89
var negativeNumber = -1243423644

fmt.Printf("bilangan positif: %d\n", positiveNumber)
fmt.Printf("bilangan negatif: %d\n", negativeNumber)
```
The positiveNumber variable is of type ```uint8``` with an initial value of ```89```. While the negativeNumber variable is declared with an initial value of ```-1243423644```. The compiler will intelligently determine the variable's data type as ```int32``` (because the number falls within the scope of the ```int32``` data type).

The ```%d``` template in ```fmt.Printf ()``` is used to format non-decimal numeric data.

# Decimal Numeric Data Type

Decimal numeric data types that need to be known are 2, ```float32``` and ```float64```. The difference between the two data types is in the wide range of decimal values that can be accommodated

````
var decimalNumber = 2.62

fmt.Printf("bilangan desimal: %f\n", decimalNumber)
fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)
````


# ```bool``` Data Type

BOOL data types contain only 2 variance values, true and false. This data type is commonly used in condition selection and looping

```
    var exist bool = true
    fmt.Printf("exist? %t \n", exist)

```

Use ```%t``` to format the bool data using the ```fmt.Printf ()``` function.


# ```string``` Data Type

A distinctive feature of a string data type is that its value is enclosed in quotation marks or quotation marks ("). Examples of its application:

```
var message string = "Halo"
fmt.Printf("message: %s \n", message)

```

In addition to using quote marks, string declarations can also be marked with grave accent / backticks (`), this sign is located to the left of key 1. The specialty of strings declared using backtics is to make all characters in it not escape, including \ n, double quotation marks and quotation marks one, new lines, and so on. All will be detected as strings.



# ```nil``` value & Zerro Value

``` nil``` is not a data type, but a value. A variable whose value is nil means it has an empty value(same as ```null`` in JavaScript).

All data types discussed above have zero values (default data type values). This means that even if the variable is declared with no initial value, there will still be a default value.

* The zero value of the string is ```""``` (empty string).
* The zero value of BOOL is ```false```.
* The zero value of the non-decimal numeric type is ```0```.
* The zero value of the decimal numeric type is ```0.0```.

Zero value is different from nil. Nil is an empty value, completely empty. nil cannot be used with the data types discussed above. There are several data types that can be set to nil values, including:

* pointer
* tipe data fungsi
* slice
* map
* channel
* interface kosong atau interface{}