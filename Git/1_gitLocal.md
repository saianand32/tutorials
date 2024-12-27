1. $git init //initializes a git repo

2. $git status  //check the status of repo
On branch master

No commits yet

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        sai.txt

nothing added to commit but untracked files present (use "git add" to track)

3. $git add . //puts all untracked files in staging area

now again type git status

4. $ git status
On branch master

No commits yet

Changes to be committed:
  (use "git rm --cached <file>..." to unstage)
        new file:   instructions.txt
        new file:   sai.txt


5. $git reset sai.txt or $git reset 
// unstages all the staged items
// can also use $git restore --staged <file> or $git restore --staged . for all


6. $git commit -m "first commit" //to commit all changes

7. $git log //prints a log of all commits

8. $git reset ebb82265c5b48bc05015c9d4cef85f0d652bfae1    //reverts to this commit id and removes all the above commits
// and the reverted items are moved to unstaged area not even on the staged area

9. If u  want to clean ur repo to the latest commit but  want to preserve work temporalily
 u can us the stash option

 $ git stash  // stashes any unstaged or even staged items but doesnt touch the commited items 
 $ git status // On branch master nothing to commit, working tree clean

 // to get back the stashed items to the unstaged area
 $ git stash pop

 //to clear the stash permanently
 $ git stash clear

 10. Amend most recent commit
    $ git add .
    $ git commit -m "new message" --amend