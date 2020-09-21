package s3

import (
    "os"
	"fmt"
    "time"
    "golang.org/x/net/context"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
    //"github.com/aws/aws-sdk-go/service/costexplorer"
)

// Base S3Bucket struct.
type S3Bucket struct {
    Name string
    Created *time.Time
    size int64
    LastUpdate int64
    FileCount int64
    pages int64
}

/*
    Builds a slice of S3Buckets  
*/
func ListBuckets() ([]S3Bucket,error) {
    var ret []S3Bucket
    svc := s3.New(session.New(),aws.NewConfig())
    result, err := svc.ListBuckets(&s3.ListBucketsInput{})

    for _,b := range result.Buckets {
        ret = append(ret,NewS3Bucket(b.Name,b.CreationDate))
    }
    return ret,err

}

/*
    A very simple generator that properly populates all S3Bucket struct
    fields.
*/
func NewS3Bucket(name *string, created *time.Time) S3Bucket {
    return S3Bucket{
        *name,
        created,
        0,
        0,
        0,
        0,
    }

}

/*
    Determines the region/location of the bucket in question
*/
func (b *S3Bucket) Location() string {
    sess := session.Must(session.NewSession())
    ctx := context.Background()
    region, err := s3manager.GetBucketRegion(ctx, sess, b.Name, "us-east-1")
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "NotFound" {
             fmt.Fprintf(os.Stderr, "unable to find bucket %s's region not found\n", b.Name)
        }
        return ""
    }
    if region == "" {
        return os.Getenv("AWS_REGION")
    }
    return region
}

/*
    Returns the size of the bucket in bytes
*/
func (b *S3Bucket) Size() int64 {
    return b.size
}

/*
    Returns the size of the bucket in the unit of your choice
*/
func (b *S3Bucket) PrettySize(unit string) string {
    switch unit {
    case "KB":
        return fmt.Sprintf("%.3f KB",float64(b.size)/1024)
    case "MB":
        return fmt.Sprintf("%.3f MB",float64(b.size)/1024/1024)
    case "GB":
        return fmt.Sprintf("%.3f GB",float64(b.size)/1024/1024/1024)
    default:
        return fmt.Sprintf("%d B",b.size)
    }
}

/*
    Builds an ARN for the bucket
*/
func (b *S3Bucket) ARN(location string) string {
    region := b.Location()
    if region == "" {
        region = os.Getenv("AWS_REGION")

    }
    if (location == "" ) {
        location = "*"
    }
    cid,_ := b.CallerID()
    return fmt.Sprintf(
        "arn:aws:%s:%s:accesspoint/%s",
        region,
        aws.StringValue(cid.Account),
        location,
    )
}

/*
    Retreives the current users identification information, which is 
    necenssary when building an ARN.
*/
func (b *S3Bucket) CallerID() (*sts.GetCallerIdentityOutput,error) {
    svc := sts.New(session.New())
    input := &sts.GetCallerIdentityInput{}

    result, err := svc.GetCallerIdentity(input)
    if err != nil {
        return nil, err
    }
    return result, nil
}

/*
    Invntories all objects in the bucket.  This generates the object count, 
    size, and last update stamp.  
*/
func (b *S3Bucket) InventoryObjects() {
    var size int64 = 0
    var count int64 = 0
    var latest int64 = 0

    svc := s3.New(session.New(),aws.NewConfig().WithRegion(b.Location()))
    input := &s3.ListObjectsV2Input{
        Bucket:  aws.String(b.Name),
        MaxKeys: aws.Int64(10000),
    }
    err := svc.ListObjectsV2Pages(
        input,
        func(page *s3.ListObjectsV2Output, lastPage bool) bool {
            for _, o := range page.Contents {
                count = count + 1
                size = size + *o.Size
                if o.LastModified.Unix() > latest {
                    latest = o.LastModified.Unix()
                }
            }
            b.pages++
            return true
        },
    )
    if err != nil {
        panic(err)
    }

    b.size = size
    b.FileCount = count
    b.LastUpdate = latest
}
