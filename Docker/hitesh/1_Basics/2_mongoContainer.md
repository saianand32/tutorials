1. pull mongo container

$ docker pull mongo

2. run mongo 
$ docker run mongo 

after this just close the running terminal

3. check which al containers were running 
$ docker ps --all

copy the containerId

4. start mongodb container using the id
$ docker start <ContainerId>  //ex docker start 444863bfaf34

5. run shell
$ docker exec -it <ContainerId> bash   //now u will see u r in a linux like terminal (try ls command etc)

6. start mongo shell
 $ mongosh   //now mongosh starts (try "show dbs" commmand)


7. to kill/stop a container from running use

$ docker stop <Image>  //always use this not kill
$ docker kill <Image>

eg docker kill 444863bfaf34