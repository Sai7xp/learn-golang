## Assginment - 1

You are part of the security team at Webknot and have been given the task to transport some confidential data to a client. To prevent any loss, the data should be encrypted in transit. The data is stored in a text file (Approx 1GB) and each line in the file contains exactly one JSON object representing some important information. The goal is to read the file efficiently using different concepts of concurrency and parallelism, encrypt the contents using an encryption key and store the encrypted contents back to a file. The file will then be transferred to the client where he can use the encryption key to get back the original contents.

### Points to Remember:

The solution should be a CLI driven program (Make can be used here) with following expected behaviour.

#### To encrypt the file

main.go encrypt --file ./imp-data.txt --key bHB&fuw2GD\*I@B --out ./imp-data-enc.txt
Args:
-file : Path to the input file
-key : Encryption key
-out : Path to the output encrypted file

#### To decrypt file

main.go decrypt --file ./imp-data-enc.txt --key bHB&fuw2GD\*I@B --out ./imp-data.txt
Args:
--file : Path to the encrypted file
--key : Encryption key
--out : Path to the output decrypted file

> Note: 
> We can encrypt the text file in a single go, but that’s not the goal here. This assignment will help you in building optimal solutions to real world problems. How to divide large data into smaller chunks and process those chunks to achieve the desired result is our goal here.
