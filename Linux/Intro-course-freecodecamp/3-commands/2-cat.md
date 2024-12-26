
Overview:
The cat command in Linux (short for "concatenate") is a versatile tool primarily used for displaying the content of files, combining files, and creating new files. Below are the common use cases of the cat command, along with examples.


1. Print file in terminal
$ cat sai.txt

2. Print multiple files in terminal
$ cat sai.txt sai2.txt sai3.txt

3. create and write or overwrite file with terminal input - 
$ cat > sai.txt

4. create and write or append to file with terminal input - 
$ cat >> sai.txt

5. Redirect output to new file
$ cat sai.txt sai2.txt > sai3.txt

6. Append contents from one/more file to another file
$ cat sai.txt sai2.txt >> sai3.txt

7. Display with line numbers
$ cat -n sai.txt