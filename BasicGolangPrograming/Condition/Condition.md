# Conditions

Condition selection is used to control program flow. The analogy is similar to the function of traffic signs on the highway. When vehicles are allowed to go and when to stop is regulated by these signs. The condition selection in the program is also more or less the same, when a block of code will be executed is controlled.

The reference for condition selection is a bool type value, it can come from variables, or the results of comparison operations. This value determines which code block will be executed.

Go has 2 kinds of keywords for condition selection, ```if```, ```else``` and ```switch```. In this chapter we will study them one by one.

# Condition Selection Using the Keyword ```if```, ```else if```, & ```else```

How to implement if-else in Go is the same as in other programming languages. What distinguishes only parentheses (parentheses), in Go does not need to be written. The following code is an example of applying another condition selection, with the number of conditions 4 pieces.

```
var point = 8

if point == 10 {
    fmt.Println("Perfect")
} else if point > 5 {
    fmt.Println("Oke")
} else if point == 4 {
    fmt.Println("Hmmmm")
} else {
    fmt.Printf("Get Out !! %d\n", point)
}
```

Of the four conditions above, what is met is if point> 5, because the value of the point variable is indeed greater than 5. Then the code block just below the condition will be executed (the code block is marked with open and close curly braces), the result is the text "pass "appears as output.

The if else Go scheme is the same as in general programming. Namely at the beginning of the condition selection using if, and when the condition is not met will go to else (if any). When there are many conditions, use else if.

# Temporary Variables in if - else

The temporary variable is a variable that can only be used in the condition selection block in which it is placed. The use of this variable brings several benefits, including:

* Scope or variable scope is clear, can only be used in the condition selection block
* The code becomes neater
* When the value of the variable is obtained from a computation, the calculation does not need to be done in the
* block of each condition.


```
var pnt = 8840.0

if percent := pnt / 100; percent >= 100 {
    fmt.Printf("%.1f%s perfect!\n", percent, "%")
} else if percent >= 70 {
    fmt.Printf("%.1f%s good\n", percent, "%")
} else {
    fmt.Printf("%.1f%s not bad\n", percent, "%")
}
```

The percent variable value is derived from the calculation results, and can only be used in a row of condition selection blocks.

Declaration of temporary variables can only be done via the type inference method that uses the sign: =. The use of the var keyword there is not permitted because it will cause an error

# Condition Selection Using the Keyword switch - case

Switch is a selection of conditions that are focused on one variable, then the value is checked. Examples are as simple as determining whether the value of the variable x is: 1, 2, 3, or others

```
    var point = 6

    switch point {
    case 8:
    fmt.Println("perfect")
    case 7:
    fmt.Println("awesome")
    default:
    fmt.Println("not bad")
}

```

In the above code, no conditions or cases are met because the value of the fixed point variable is 6. When something like this happens, the default condition block is called. You could say that default is else in a switch.

Keep in mind, the switch in Go programming has differences compared to other languages. In Go, when a case is fulfilled, it will not proceed to checking the next case, even though there is no keyword break there. This concept is the opposite of switches in general, which when a case is fulfilled, it will continue to check the next case unless there is a keyword break


# Use cases for many conditions

```case``` can accommodate many conditions. How to apply that is by writing the values ​​of the comparison variables that are switched after the keyword ```case``` separated by commas (,)

```
var point = 6

switch point {
case 8:
    fmt.Println("perfect")
case 7, 6, 5, 4:
    fmt.Println("awesome")
default:
    fmt.Println("not bad")
}
```


Case conditions 7, 6, 5, 4: will be met when the value of the point variable is 7 or 6 or 5 or 4

# Curly Brackets in Keyword & default cases

Curly braces ({}) can be applied to the keyword case and default. This mark is optional, may or may not be used. It is good if used in a block of conditions in which there are many statements, the code will look neater and easier to maintain.

Note the following code, can be seen on the default keyword there are curly braces that surround the 2 statements in it.

```
var point = 6

switch point {
case 8:
    fmt.Println("perfect")
case 7, 6, 5, 4:
    fmt.Println("awesome")
default:
    {
        fmt.Println("not bad")
        fmt.Println("you can be better!")
    }
}
```

# Switch with if - else Style

Uniquely on Go, the ```switch``` can be used in the style of ```if-else```. The value to be compared is not written after the keyword ```switch```, but will be written directly in the form of a comparison in the keyword ```case```.

In the code below, the switch program code above is changed to the if-else style. The point variable is removed from the keyword ```switch```, then the conditions are written in each ```case```.

```
var num = 6

switch {
case num == 8:
    fmt.Println("perfect")
case (num < 8) && (num > 3):
    fmt.Println("awesome")
default:
    {
        fmt.Println("not bad")
        fmt.Println("you need to learn more")
    }
}
```

# Use the Keyword ```fallthrough``` in a ```switch```

As previously explained, that the ```switch``` on Go has a difference with other languages. When a case is fulfilled, checking the condition will not be forwarded to the following ```cases```.

The ```fallthrough``` keyword is used to force the checking process to be forwarded to the next case regardless of the condition value, so the case in the next check is always assumed to be correct (even if the original is false).

```
var number = 6

switch {
case nums == 8:
    fmt.Println("perfect")
case (nums < 8) && (nums > 3):
    fmt.Println("awesome")
    fallthrough
case nums < 5:
    fmt.Println("you need to learn more")
default:
    {
        fmt.Println("not bad")
        fmt.Println("you need to learn more")
    }
}
```

After checking the case (point <8) && (point> 3) is complete, it will proceed to checking case point <5, because there is a fallthrough there.

# Nested Condition Selection

Nested condition selection is condition selection, which is in condition selection, which may also be in condition selection, and so on. Nested condition selection can be done on ```if - else```, ```switch```, or a combination of both.
