Checkout has many usecases:

1. switch branches
$ git checkout -b test  //create a new branch an move to that
$ git checkout test1    //move to an existing


2. Clear/discard changes in working area(unstaged changes only)(doesnt remove created files)
$ git checkout -- .             //discard all
$ git checkout -- config.js     //discard particular file

3. To go to a prev commit of code
$ git checkout <commit hash>

4. Going to prev commit like in example 3 and after that if u add a commit a new branch gets created. (try $git log --all --graph)
To prevent this branching u can do 
$ git checkout <commit hash> .
Now head will point to master 

Note: So if u want to just view the prev version of code use example 3.. but if u want to restore the prev version of code
Use example 4. U can also use git reset --hard but that will remove any of the above commit history