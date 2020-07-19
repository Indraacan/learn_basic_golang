# Array

An array is a collection of data of the same type, which is stored in a variable. Array has a capacity whose value is determined at the time of creation, making the number of elements / data stored in the array not exceed that which has been allocated. The default value of each array element initially depends on the data type. If int, then each zero value element is 0, if it is bool, then false, and so on. Each array element has an index in the form of a number that represents the order in which the elements are arranged. Array index starts at 0.

```
var names [4]string
names[0] = "trafalgar"
names[1] = "d"
names[2] = "water"
names[3] = "law"

fmt.Println(names[0], names[1], names[2], names[3])
```

The names variable is declared as a string array with 4 slot element allocation. How to fill the array element slots can be seen in the code above, namely by directly accessing the elements using an index, then filling it

# Initializing Initial Array Values ​​Without Number of Elements

Declaration of an array whose value is set at the beginning, may not be written down the number of the width of the array, simply replace it with a 3-point mark (```...```). The number of elements will be calculated automatically according to the data of the element that is loaded.

```
var numbers = [...]int{2, 3, 2, 4, 3}

fmt.Println("data array \t:", numbers)
fmt.Println("jumlah elemen \t:", len(numbers))

```

The numbers variable will automatically have element number 5, because at the time of the declaration 5 elements are prepared.

# Multidimensional Array

Multidimensional arrays are arrays where each element is also an array (and can be so on, depending on the depth of the dimensions).

How to declare a multidimensional array in general the same as the usual array declaration, by writing the next dimension array data as an element of the previous dimension array.

Especially for arrays which are sub dimensions or elements, the amount of data may not be written. An example can be seen in the declaration of the numbers2 variable in the following code.


```
var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
fmt.Println("numbers1", numbers1)
fmt.Println("numbers2", numbers2)
```

The two arrays above have the same elements.

# Looping Array Elements Using Keyword for

Keyword for and array have a very close relationship. By utilizing looping using these keywords, elements in the array can be obtained.

There are several ways you can loop data arrays, the first is to use iteration iteration variables to access elements based on their indexes.

```
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}
```

The iteration is run as many as the elements of the fruits array (it can be seen from the condition ```i <len (fruits)```. In each iteration, the ```array``` element is accessed via the iteration variable ```i```

# Looping Array Elements Using Keyword for - range

There is a simpler way to loop data arrays, using the for-range keyword. Examples of application can be seen in the following code.

```
var langue = [4]string{"C", "Python", "JavaScript", "Golang"}

for i, langue := range langue {
    fmt.Printf("elemen %d : %s\n", i, langue)
}

```

Arrays of ```langue``` are taken in order. The value of each element is variable by ```langue``` (without the letter s), while the index is for variable ```i```.

The above program output, the same as the previous program output, only the method used is different

# Use of Underscore Variables _ In the for-range

Sometimes when looping uses ```for-range```, it is possible that the data needed is only the element, the index is not. While the code above, range returns 2 data, namely index and element.

As you well know, that in Go does not allow any variable that lags or is not used. If forced, an error will appear,

This is one of the uses of the unemployment variable, or underscore (```_```). Just hold values ​​that don't want to be used to underscore.

```
var country = [4]string{"indonesia", "singapore", "Rusia", "Korea"}

for _, country := range country {
    fmt.Printf("name country : %s\n", country)
}
```

If you only need the element index, you can use one variable after the keyword for.

```
for i, _ := range country { }
// or
for i := range country { }

```

# Array Element Allocation Using Keyword make

Declaration as well as array data allocation can also be done via the keyword make.

```
var name = make([]string, 2)
fruits[0] = "sony"
fruits[1] = "nurdianto"

fmt.Println(name)
```

The first parameter make keyword is filled with the desired array element data type, the second parameter is the number of elements. In the above code, the variable name is printed as a string array with an allocation of 2 slots.