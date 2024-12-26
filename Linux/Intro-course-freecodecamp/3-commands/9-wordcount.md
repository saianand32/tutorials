
The wc (word count) command in Linux is used to display the number of lines, words, characters, and bytes in a file or input. Here's the basic syntax and how to use it:

Basic:
wc [options] [file...]

Common Options:
-l: Print the number of lines.
-w: Print the number of words.
-c: Print the number of bytes.
-m: Print the number of characters (same as bytes, but different in some multi-byte encodings like UTF-8).
-L: Print the length of the longest line.

1. Count lines, words, and characters in a file:
$ wc sai.txt    //output: 12  85  500 sai.txt

2. Count only lines:
$ wc -l sai.txt

3. Count only words:
$ wc -w sai.txt

4. Count only the number of characters (or bytes)
$ wc -c sai.txt

5. Count the longest line length:
$ wc -L sai.txt

6. Use wc with input from a pipe:
$ echo "Hello World" | wc -w


