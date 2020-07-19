# Operator

In general there are 3 categories of operators: arithmetic, comparison, and logic.

# Operator Aritmatika

Arithmetic operators are operators that are used for operations that are calculated in nature. Go supports several standard arithmetic operators, the list can be seen in the following table.

| Sign |                   explanation                  |
|------|:----------------------------------------------:|
|   +  |                    addition                    |
|   -  |                   subtraction                  |
|   *  |                 multiplication                 |
|   /  |                    division                    |
|   %  | modulus / remainder of the distribution result |

example usage

```
var value = (((2 + 6) % 3) * 4 - 2) / 3
fmt.Println(value)

```

# Comparison Operators

Comparison operators are used to determine the truth of a condition. The result is a boolean, ```true``` or ```false``.

The table below contains comparison operators that can be used on Go.

| sign |                          explanation                          |
|------|:-------------------------------------------------------------:|
|  ==  |     whether the left value is the same as the right value     |
|  !=  |   whether the left value is not the same as the right value   |
|   <  | whether the left value is smaller than the right value        |
|  <=  | whether the left value is smaller or equal to the right value |
|   >  |     whether the left value is greater than the right value    |
|  >=  | whether the left value is greater or equal to the right value |

example usage:


```
var value = (((2 + 6) % 3) * 4 - 2) / 3
var isEqual = (value == 2)

fmt.Printf("nilai %d (%t) \n", value, isEqual)
```



# Logic Operator

This operator is used to find whether or not a combination of data of type Bool (can be a variable of type Bool, or the result of a comparison operator).

Some standard logic operators that can be used:

| Sign |       explanation      |
|:----:|:----------------------:|
|  &&  |  And (Left and Right)  |
| \|\| | Or (Left Or Right)     |
|   !  | Negation /the opposite |

Example usage :

```
var left = false
var right = true

var leftAndRight = left && right
fmt.Printf("left && right \t(%t) \n", leftAndRight)

var leftOrRight = left || right
fmt.Printf("left || right \t(%t) \n", leftOrRight)

var leftReverse = !left
fmt.Printf("!left \t\t(%t) \n", leftReverse)
```
Template \t digunakan untuk menambahkan indent tabulasi. Biasa dimanfaatkan untuk merapikan tampilan output pada console.
