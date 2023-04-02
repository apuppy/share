# shell

- [shell](#shell)
  - [basis](#basis)
    - [define variables](#define-variables)
    - [Arithmetic Expressions](#arithmetic-expressions)
    - [numberical expressions](#numberical-expressions)
    - [read user input](#read-user-input)
    - [Numeric Comparison logical operators](#numeric-comparison-logical-operators)
    - [Conditional Statements (Decision Making)](#conditional-statements-decision-making)
    - [loop with numbers](#loop-with-numbers)
    - [loop with strings](#loop-with-strings)
    - [while loop](#while-loop)
    - [reading file](#reading-file)
    - [execute commands with back ticks](#execute-commands-with-back-ticks)
    - [get arguments for scripts from the command line](#get-arguments-for-scripts-from-the-command-line)

## basis
### define variables
```bash
#!/bin/bash
greeting=Hello
name=John
echo $greeting $name
```

### Arithmetic Expressions
| OPERATOR | USAGE          |
|----------|----------------|
| +        | addition       |
| -        | subtraction    |
| *        | multiplication |
| /        | division       |
| **       | exponentiation |
| %        | modulus        |

### numberical expressions
```bash
# built-in
var=$((expression))
# use bc
echo "scale=2;22/7" | bc
```

### read user input
```bash
#!/bin/bash

# read variable_name
# read -p "Enter your age" variable_name

echo "Enter a numner"
read a

echo "Enter a numner"
read b

var=$((a+b))
echo $var
```

### Numeric Comparison logical operators
| OPERATION             | SYNTAX        | EXPLANATION                        |
|-----------------------|---------------|------------------------------------|
| Equality              | num1 -eq num2 | is num1 equal to num2              |
| Greater than equal to | num1 -ge num2 | is num1 greater than equal to num2 |
| Greater than          | num1 -gt num2 | is num1 greater than num2          |
| Less than equal to    | num1 -le num2 | is num1 less than equal to num2    |
| Less than             | num1 -lt num2 | is num1 less than num2             |
| Not Equal to          | num1 -ne num2 | is num1 not equal to num2          |

### Conditional Statements (Decision Making)
- `if...then...fi` statements
- `if...then...else...fi` statements
- `if..elif..else..fi`
- `if..then..else..if..then..fi..fi..` (Nested Conditionals)

### loop with numbers
```bash
#!/bin/bash

for i in {1..5}
do
    echo $i
done
```

### loop with strings
```bash
#!/bin/bash

for s in cyan magenta yellow
do
	echo $s
done
```

### while loop
```bash
#!/bin/bash
i=1
while [[ $i -le 10 ]] ; do
   echo "$i"
  (( i += 1 ))
done
```

### reading file
```bash
#!/bin/bash

LINE=1

while read -r CURRENT_LINE
	do
		echo "$LINE: $CURRENT_LINE"
    ((LINE++))
done < "sample_file.txt"
```

### execute commands with back ticks
```bash
#!/bin/bash

var=`df -h | grep tmpfs`
echo $var
```

### get arguments for scripts from the command line
```bash
#!/bin/bash

for x in $@
do
    echo "Entered arg is $x"
done
```