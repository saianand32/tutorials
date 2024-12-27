
1. Adding a remote repo

$ git remote add <url_name> <url>
//example
$ git remote add origin https://github.com/repo.git

//origin is just the name u r giving to the url

2. U can have multiple urls associated with one repo

$ git remote add <url_name1> <url1>
$ git remote add <url_name2> <url2>
//example
$ git remote add origin https://github.com/trash.git
$ git remote add origin2 https://github.com/trash2.git

3. List all the remote urls associated with your project ( below after i have added two remote repos names origin and origin2 )

$git remote -v
origin  https://github.com/saianand32/trash.git (fetch)
origin  https://github.com/saianand32/trash.git (push)
origin2 https://github.com/saianand32/trash2.git (fetch)
origin2 https://github.com/saianand32/trash2.git (push)

4. Remove a remote repo
$  git remote remove <url_name>
//example
$  git remote remove origin2


5. push to a remote repo
$ git push <url_name> <branch_name>
//example
$ git push origin main



6. *********  undo the comit to remote repo (https://youtu.be/H2DuJNWbqLw?si=A56h67YV89K8fFuf) **********

stack overflow - https://stackoverflow.com/questions/1463340/how-can-i-revert-multiple-git-commits 

2 ways ----

******* 1. Make a new comit of reverting (use git revert) ******

s1 - $git log  
//and copy the comit id of the comit that is wrong and u want to revert from (Note.. not the comit u want to revert to but the commit that is wrong that ones id u need to copy)
//editor exit with 'q' key

s2 - $git revert <wrong_commit_id> 
or - $git revert HEAD //if wrong commit id is head  //Head is reverted nd changes removed
   
//ex - git revert 17578e757e906ef5f61723ab83640d78480c8b74
//this will open an editor to put commit message exit with ':x<enter>'

s3 - $git push
//actually this reverting of commit is also a new commit u need to push

//Note revert multiple commits    A <- B <- C <- D <- HEAD  i want to revert to B (dont chnge B)

$git revert B...HEAD  //Bis commit_id of B nd B changes r kept but C, D, HEAD reverted
$git  push //push this revert as a new change

//if u want only one commit message
$git revert --no-commit B...HEAD 
$git commit -m "reverted C, D, HEAD"
$git  push //push this revert as a new change



***** 2. delete the comit from github like it never existed (using git reset) ***** 

s1 - $git reset HEAD~2
//goes to last 3rd commit = HEAD-2 or u can also place the commit id as well
$git reset <commit_id>

s2 - $git push -f origin main //need to do a force push

note - wen u do $git reset commit_id ,the remaining commits changes go to staging area....
if u want to delete them off use git stash and then git stash release
but can do in one step
$ git reset --hard commit_id //this doesnt get changes to staging area