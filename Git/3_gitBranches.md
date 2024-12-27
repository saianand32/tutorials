1. Create a new branch
$ git branch branch_name

2. move to branch
$ git checkout branch_name

3. create and move to a branch
$ git checkout -b branch_name

4. delete branch
$ git branch -d branch_name

5. list all branches
$ git branch

6. merge branches (say u want to merge b1 into main)
$ git checkout main
$ git merge b1 //merges b1 into main

if merge conflict comes resolve and then do
$ git merge --continue
//add a commit name in editor
$git push -u origin main