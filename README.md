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

## 4. Number lines

I need to parse flags to add line numbers to output.
I'll put my flags into a struct and pass that around to determine what do.
I think for args that involve formatting my output, I'll create a new function to format each line
as it prints according to the arguments.

## 5. Number lines pt 2: Include and Exclude blank lines

Based on what I did in part 4, it already prints line numbers for blank lines.

Since we don't want numbers added an extra conditional to include or exclude the number count.
