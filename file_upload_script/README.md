# Upload a list of files over ssh. 

`upload_file_ssh.sh` will transfer the file from the directory specified in the script over ssh

### How to run

1) first set the `from_dir` from which files need to be transferred, currently it is set to `/home/girish/go/src/github.com/girishg4t/go-grpc-graphql`   
2) set `user`, `host` and `to_path` in which file need to be transferred  

```bash
$ bash ./upload_file_ssh.sh girish localhost
```


Note:
1) The SSH connection need to be established to trasfer the files
2) The variable `from_dir` need to be set to transfer the file from correct folder