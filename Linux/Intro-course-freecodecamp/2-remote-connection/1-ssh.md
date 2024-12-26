

OVERVIEW: The ssh (Secure Shell) command is used to securely connect to a remote server or machine over a network. It provides a secure channel for communication and allows you to execute commands, transfer files, and manage servers remotely.


1. Basic Syntax:
$ ssh username@hostname_or_ip

2. SSH to a remote server
Connect to a server with the username paul at the IP address 192.122.124.99:
$ ssh paul@192.122.124.99

3. if the server uses a port other than the default (22), use the -p option to specify the port:
$ ssh -p 222 paul@192.122.124.99

4. SSH with an identity file (private key)
If your server requires authentication using an SSH key, you can specify the private key file with the -i option:
$ ssh -i ~/Desktop/private_key.pem paul@192.122.124.99

5. SSH with verbose output
To get more detailed information about the connection process (useful for debugging), you can use the -v option for verbose mode:
$ ssh -v paul@192.122.124.99

