
OVERVIEW:
The SCP (Secure Copy Protocol) command is used to securely transfer files and directories between local and remote hosts, or between two remote hosts. SCP uses the same underlying technology as SSH (Secure Shell) for secure data transfer. It is widely used because it is simple, fast, and encrypted, ensuring that files are transferred securely over a network.


1. Basic Syntax 
$ scp [options] source destination

2. Copy files from local to server
$ scp ./file.txt paul@192.122.124.99:~/Desktop/sai

3. Copy file from server to local
$ scp paul@192.122.124.99:~/Desktop/sai/file.txt .

4. Copy a dir - -r flag for recursive copy
$ scp -r ./sai paul@192.122.124.99:~/Desktop/sai

5. Copy using -i flag pem file
$ scp -i ./private_key.pem ./file.txt paul@192.122.124.99:~/Desktop/sai

6. Copy a file between two remote servers: You can also use SCP to copy files directly between two remote systems without involving the local machine:
$scp user1@192.122.124.99:/path/to/file.txt user2@192.122.124.100:/home/user2/

7. -C: Enable compression to speed up the transfer of large files.
$ scp -C large_file.tar.gz user@remote_host:/path/to/target/

