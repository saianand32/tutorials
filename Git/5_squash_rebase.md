1. Squashing commits means say A<-B<-C<-D<-HEAD 

make commits bcd into one commit then  becomes A<-newNamedComboCommit<-HEAD

$git rebase -i A 
//A=commit_id of A so A is not included but B to Head is included 

//a editor opens like 
pick 896686 B
pick 897897 C
pick 897897 D
pick 897897 E(this is head commit name just say)

// u can squash by puttng s ... each s is squashed into nearest above pick
pick 896686 B
s 897897 C
pick 897897 D
s 897897 E

now C will be squashed with B as one commit & E with D as one

// now another editor opens for commit name add name nd exit

$git push -f origin main //need to force push



// u can squash all into one by puttng 
pick 896686 B
s 897897 C
s 897897 D
s 897897 E


