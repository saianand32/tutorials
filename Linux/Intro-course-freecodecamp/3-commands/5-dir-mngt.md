

1. make directory 
$ mkdir sai

2. rename folder (note if u mv to existing dir it will go to a subdir of existing dir else it will rename to new one)
$ mv sai sai2

3. delete dir
$ rmdir sai         //works on empty dir
$ rm -rf sai        //can delete dir with contents

4. change dir
$ cd ./Desktop

$ cd ..             // move to parent

$ cd ../..          // move to parent of parent

$ cd /              // move to root path

$ cd ~              // move to home path

5. Copy file to folder or folder to folder
$ cp sai.txt ./sai

$ cp -r sai sai2    // -r flag is for recursive copy for folder to folder

6. moving folders
$ mv sai sai2      // no need for -r in mv command as its default behaviour