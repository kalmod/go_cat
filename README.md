# go_cat

## 1. First I wanna be able to read file from input

## 2. Read input from standard in

- I can generate a scanner from both stdin and a file. I should separate this
part from the `readfile` function
- If valid file: read file
- If not a file and arg != "-": return error
- if arg == "-" or empty (empty args get "-") read from stdin


## 3. Concatenate files
I had added a for loop to go through all arguments passed into the program
