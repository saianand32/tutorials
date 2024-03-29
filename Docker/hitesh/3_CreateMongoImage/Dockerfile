

#step1 (get a base image -- alpine is a os .. so we define which os we need)
FROM alpine 


#step 2 install software

RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/main' >> /etc/apk/repositories
RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/community' >> /etc/apk/repositories

RUN apk update 
RUN apk add mongodb=3.4.4-r0

#step 2.5 configure the software

VOLUME ["/data/db"]
WORKDIR /data
EXPOSE 27017

#step3 set default commands
CMD ["mongod"]


# to run this 
# $ doker build .
# $ docker run <ImageId> //can get image id from $docker images
# Open another terminal
# $ docker exec -it <runningContainerId> sh 
# $ mongo //to start the mongoshell


# if u want to name docker containers (name = saish/mongo, tag = latest)
#  $ docker build -t saish/mongo:latest .
#   $docker run saish/mongo 
 






