1. fork a git repo //creates a copy of it in ur account

2. create a separate branch and add a feature

3. push the branch git push origin branch_name

4. go to github u can see an option to create a pull req do it

5. Always create separate branches for a pull request as in one branch only one pull req can be raised at a time and any further commits will just get added to the same pull req which is not good


******* syncing ur forked repo with main repo *****

1. u can do it on github ... there is a button Sync fork

2. add the original repo url from where u forked as upstream
$ git remote add upstream https://github.com/react //this is ur upstream  url 

$ git pull upstream main  //this is the main command to sync main from forked