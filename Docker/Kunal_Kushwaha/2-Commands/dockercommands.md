

<!-- HelloWorld  -->
-- $ docker run hello-world


------------------------------------------------------------------------- 

<!-- Pulling images from docker hub  -->
-- $ docker pull image_name

<!-- remove one image with id  -->
$ docker rmi <ID>

%  Lists  all image
$ docker images

%  Lists  all image id's
$ docker images -q 

% <!-- remove all images  -->
$ docker rmi $(docker images -q)

<!-- run a container from image -->
-- $ docker run image_name
-- $ docker run IMG_ID

-------------------------------------------------------------------------

<!-- Print all running containers  -->
-- $ docker ps 
-- $ docker container ls

<!-- Print all existing(running/non running) containers  -->
-- $ docker ps -a

<!-- stop a container with id=ID -->
-- docker stop ID
-- docker kill ID    //this immediatly stops container

<!-- start a container with id=ID -->
-- docker start ID

<!-- remove a container  -->
-- docker rm ID

<!-- delete all stopped cntainers -->
docker container prune -f

<!-- inspect container -->
-- docker inspect ID

<!-- logs -->
-- docker logs ID

-------------------------------------------------------------------------

<!-- Running container in background using -d flag  (detach mode) -->
-- $ docker run -d alpine ping www.google.com     //returns only the container id say - 235456433

<!-- logs since 5 sec of a detached container with id = 235456433 -->
-- $ docker logs --since 5s 235456433


-------------------------------------------------------------------





