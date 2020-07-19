# Looping

Iteration is the process of repeating the execution of blocks of code without interruption, as long as the conditions used as references are met. Usually a variable is prepared for iteration or a marker variable when the loop will be stopped.

In Go looping keywords are only ```for```, but even so, the ability is a combination of ```for, foreach, and while``` like other programming languages.

# Iteration Using Keyword for

There are several standard ways to use for. The first way is to enter the loop counter variable and its condition after the keyword. Pay attention and practice the following code.

```
for i := 0; i < 5; i++ {
    fmt.Println("Angka", i)
}
```

Perulangan di atas hanya akan berjalan ketika variabel ```i``` bernilai dibawah ```5```, dengan ketentuan setiap kali perulangan, nilai variabel ```i``` akan di-iterasi atau ditambahkan ```1``` (```i++``` artinya ditambah satu, sama seperti ```i = i + 1```). Karena ```i``` pada awalnya bernilai ```0```, maka perulangan akan berlangsung ```5``` kali, yaitu ketika ```i``` bernilai ```0, 1, 2, 3, dan 4```.

# Use of Keywords for Arguments Only Conditions


The second way is to write the conditions after the keyword for (conditions only). Declarations and iterations of the counter variable are not written after the keyword, only the iteration condition. The concept is similar to ```while``` in other programming languages.

The following code is an example of for with only condition arguments (such as if), the resulting output is the same as applying for the first way.

```
var i = 0

for i < 5 {
    fmt.Println("Angka", i)
    i++
}
```

# Use of Keyword for No Arguments

The third way is for written without conditions. This will produce a loop without stopping (equal to true). Stop looping is done by using the keyword ```break```.

```
var z = 0

for {
    fmt.Println("nums", z)

    i++
    if i == z {
        break
    }
}
```

In the non-stop loop above, variable i whose initial value is 0 is incremented. When the value of i has reached 5, the keyword break is used, and the loop will stop.d

# Use of Keyword for - range

```for-range``` is a loop using a combination of keywords for and range. This method is commonly used for looping array type data

```	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
		fmt.Println("sum:", sum)
	}

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}```

range on arrays and slices provides both the index and value for each entry. Above we didn’t need the index, so we ignored it with the blank identifier _. Sometimes we actually want the indexes though.

range on map iterates over key/value pairs.

range can also iterate over just the keys of a map

range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself.






# use Keyword ```break``` & ```continue```

The keyword ```break``` is used to forcibly stop a loop, while ```continue``` is used to force it forward to the next loop.

The following is an example of the application of ```continue``` and ```break```. Both of these keywords are used to display sequential even numbers greater than 0 and below 8.

```
for i := 1; i <= 10; i++ {
    if i % 2 == 1 {
        continue
    }

    if i > 8 {
        break
    }

    fmt.Println("Angka", i)
}
```
The code above will be easier to digest if explained in sequence. The following is the explanation.

* Repetition starts from numbers 1 to 10 with i as an iteration variable.
* When i is odd (it can be seen from i% 2, if the result is 1, it means odd), it will be forced to continue to the * * next iteration.
* When i is greater than 8, the loop will stop.
* The value of i is displayed.



# Perulangan Bersarang

Not only condition selection can be nested, iteration can also. The method of application is more or less the same, all you have to do is write the loop statement block inside the loop.

```
for i := 0; i < 5; i++ {
    for j := i; j < 5; j++ {
        fmt.Print(j, " ")
    }

    fmt.Println()
}
```

In the above code, for the first time the fmt.Println () function is called without the parameter inserted. This method can be used to display new lines. Its usefulness is the same as the output of the fmt.Print ("\ n") statement.


# Utilization of Labels in Loop



In a nested loop, the break and continue will apply to the loop block where it is used only. There are ways that these two keywords can be aimed at the outermost repetition or a specific iteration, namely by utilizing the labeling technique.

The program for generating the following matrix is ​​an example of applying a loop label.


```
outerLoop:
for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
        if i == 3 {
            break outerLoop
        }
        fmt.Print("matriks [", i, "][", j, "]", "\n")
    }
}
```

Just before the outer for keyword, there is a line of outerLoop: code. The purpose of the code is to prepare a label named outerLoop for for underneath. The label name can be replaced by another name (and must end with a colon or colon (:)).

On the inside, there is a condition selection for checking the value of i. When the value is equal to 3, then the break called by the target is a loop that is labeled outerLoop, the loop will be stopped