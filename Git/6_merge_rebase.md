//Rebase can be used even to merge two branches in a different way

1. **** using merge ****

say u have 2 branches main and br1 which separated at comit id 120

now in main i make a commit with commit id 130

and in br1 i make a commit with commit id 140

now when i merge
$ git checkout main
$ git merge br1 
$git push origin main

then this is the order of the commits

120 <- 130 <- 140 <-1xx(some id saying commit for merged branches)


1. **** using rebase ****

say u have 2 branches main and br1 which separated at comit id 120

now in main i make a commit with commit id 130

and in br1 i make a commit with commit id 140

now when i rebase 
$ git checkout main
$ git rebase br1 
$ git push -f origin main

then this is the order of the commits

120<-140<-130   //Note no new commit is there stating a rebase/merge

see how the main branches commits got rebased from base being 120 to adding commit 140 from br1 and making it new base
i have attached pis as well (Note my branch name in image was n1 ot br1...)