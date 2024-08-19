

1. Using 3rdparty containers

# Create a container from the ubuntu image
docker run --interactive --tty --rm ubuntu:22.04

# Try to ping google.com
ping google.com -c 1 # This results in `bash: ping: command not found`

# Install ping
apt update
apt install iputils-ping --yes

ping google.com # This time it succeeds!
exit

Let's try that again:

docker run -it --rm ubuntu:22.04
ping google.com -c 1 # It fails! 🤔
It fails the second time because we installed it into that read/write layer specific to the first container, and when we tried again it was a separate container with a separate read/write layer!


2. We can give the container a name so that we can tell docker to reuse it:

# Create a container from the ubuntu image (with a name and WITHOUT the --rm flag)
docker run -it --name my-ubuntu-container ubuntu:22.04

# Install & use ping
apt update
apt install iputils-ping --yes
ping google.com
exit

# List all containers
docker container ps -a 

# Restart the container and attach to running shell
docker start my-ubuntu-container
docker attach my-ubuntu-container

# Test ping
ping google.com# It should now succeed! 🎉
exit


3. We generally never want to rely on a container to persist the data, so for a dependency like this, we would want to include it in the image:

# Build a container image with ubuntu image as base and ping installed
docker build --tag my-ubuntu-image -<<EOF
FROM ubuntu:22.04
RUN apt update && apt install iputils-ping --yes
EOF

# Run a container based on that image
docker run -it --rm my-ubuntu-image

# Confirm that ping was pre-installed
ping google.com -c 1 # Success! 🥳