

<!-- create ubuntu image -->
$ docker run -it ubuntu
$ echo "hii" >>sai.txt
$ docker stop <ID>

% How to restart 
$ docker start <ID>
$ exec -it <ID> bash

% now u will still see the sai.txt with hi 
% but how can i share this in a commit 

% so you can create a copy of your container with state 

$ docker commit -m "added sai.txt" <ID> ubuntu_sai:1.01
$ docker images 
% here now u can see a new image ubuntu_sai with tag 1.01

% this will run your new image
$ docker run -it ubuntu_sai:1.01
