# s3inator

Lists information related to S3 buckets

## Usage

```
Usage:
  s3inator [flags]
  s3inator [command]

Available Commands:
  help        Help about any command
  list        Lists S3 resources, such as buckets

Flags:
      --authkey string     AWS Auth Key
      --config string      config file (default is $HOME/.s3inator.yaml)
  -h, --help               help for s3inator
      --region string      AWS Region (default "us-east-1")
      --secretkey string   AWS Auth Key

Use "s3inator [command] --help" for more information about a command.
```

## Authentication

The s3inator requires that you specify your AWS Auth Key, AWS Secret key, and AWS Region.
This can be done using the provided CLI flags or you can set the following environment
variables:

    AWS_REGION
    AWS_ACCESS_KEY_ID
    AWS_SECRET_ACCESS_KEY


## Pre-Compiled Binaries

You can find pre-compiled binaries in the dist/ directory.  

## Compiling

You can use the "make" command to build new binaries.  All binaries are built using
GNU Make on a Linux host.  

