

Docker components---

1. "Docker Runtime" ---- mainly acts with host OS to manage containers
   -- "runc": Responsible for spawning and running containers based on OCI (Open Container Initiative) specifications. It manages container lifecycle       operations like start, stop, and delete.

   -- "containerd": Acts as an industry-standard core container runtime, managing container lifecycle, image distribution, and storage. It interfaces with runc and handles low-level container operations such as networking and storage.
Docker Engine

2. "Docker Engine" ---- 
    - Daemon: The core background service running on the host system, responsible for managing Docker objects such as images, containers, networks, and volumes.
    - REST API: Provides a secure and accessible interface for communicating with the Docker daemon, enabling remote management of Docker containers and services.
    - Docker CLI: Command-line interface used by users to interact with Docker, issuing commands to the Docker daemon via the REST API.
    Orchestration

3. Orchestration - 
   -- scaling multiple container application
Involves managing and scaling multiple containers that collectively form an application or service.
Tools like Docker Swarm (native to Docker) or Kubernetes (developed by Google) are used for container orchestration. They automate deployment, scaling, and management of containerized applications, ensuring efficient resource utilization and high availability.
These notes provide a clearer structure and explanation of Docker's key components and their roles in containerization and orchestration.



---- Docker Image is just a file that contains all instructions etc
---- Docker container is a running instance of an image
 

DOCKER_FILE -----> DOCKER IMAGE -----> DOCKER CONTAINER

-- When u run a docker file u get a docker image ... run the image u get a container 



---- Shim
if docker daemon goes down initially containers also go down but now its changed
with intro to containerd & shim we now can have daemonless containers means even if daemon is down containers keep running as shim and containerd handle all ops. runc has only one job create delete containers it does it and shuts down then onward shim handles container communincation with containerd.

in linux binary names are as follows --
1. Docker Daemon --- dockerd
2. containerd    --- docker-containerd
3. shim          --- docker-containerd-shim
4. runc          --- docker-runc



---Docker server & client (advanced topics)
-- client --- port 2375 ---> server
must be TLS(transport layer security) secure
