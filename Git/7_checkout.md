Checkout has many usecases:

1. switch branches
$ git checkout -b test  //create a new branch an move to that
$ git checkout test1    //move to an existing


2. Clear/discard changes in working area(unstaged changes only)(doesnt remove created files)
$ git checkout -- .             //discard all
$ git checkout -- config.js     //discard particular file