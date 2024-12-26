

1. current dir
$ pwd

2. current user
$ whoami

3. System date and time

$ date              ---output: Sun Dec 15 02:06:13 IST 2024 
$ date +%T          ---output: 02:07:49 only time (24hr)
$ date +%D          ---output: 12/15/24 only date


$ date +%F          ---output: 2024-12-15  <Date in YYYY-MM-DD format>
$ date +%r          ---output: 02:06:13 AM  <Time in am pm>
$ date +%s          ---output: 1702523149  <Unix Timestamp>

5. list commands - 

$ ls 
$ ls -a      -- Lists all files, including hidden files (files starting with a dot).
$ ls -A      -- Lists almost all files, except . and ..
$ ls -lt     -- -l long format -t sort by time  -rw-r--r-- 1 user user  4096 Dec 15 02:06 file2.txt
$ ls -ltr    -- same as lt but sorts reverse order of time
$ ls -lh     -- same as lt but for size of file adds B, K, M etc -h is human readable
$ ls -lhr    -- same like lh but reverse time sort
$ ls -S      -- sort by size largest first