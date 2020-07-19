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

