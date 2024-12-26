

- Find a file by name "sai.txt" (case sensitive and does subdir check also)
$ find var/lib/neo -name sai.txt


- what if u dont know path to search in 
 Can use "locate" command -- one drawback u have to run "update db" command before running locate

 in linux:
 $ sudo updatedb  //need root access
 $ locate sai.txt

 in mac:
 $  sudo /usr/libexec/locate.updatedb
 $ locate sai.txt

