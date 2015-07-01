# s3-ls

Utility and lib for list files in Amazon S3 service.

## Install
```
go get "github.com/erichnascimento/s3-ls"
```

## Usage command line
```
$ s3-ls -h

  Usage:
    s3-ls --region <region> <bucket> <file>

  Options:
    -r, --region region   region of bucket
    -h, --help            output help information
    -v, --version         output version
```

### List
List `myfolder` dir on bucket `mybucket` over region `sa-east-1`
```
$ s3-ls --region sa-east-1 mybucket /myfolder/
total 3
     erichnascimento  5.6GB      3 days ago myfolder/0-Sun.sql.gz
     erichnascimento  5.6GB      2 days ago myfolder/1-Mon.sql.gz
     erichnascimento  5.6GB       1 day ago myfolder/2-Tue.sql.gz
```
