/*
Copyright © 2020 Scott Cudney <scott@cudneys.net>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
    //"os"
	"fmt"
	"github.com/spf13/cobra"
    s3 "github.com/cudneys/s3inator/s3"
)

var bucketsCmd = &cobra.Command{
	Use:   "buckets",
	Short: "Lists AWS S3 Buckets",
	Long: `Lists AWS S3 buckets and relevant details, such as 
the creation date, nubmer of files, space used, etc.  
`,
	Run: func(cmd *cobra.Command, args []string) {
		listBucketsCommand()
	},
}

/*
    Generates the list buckets output.
*/
func listBucketsCommand() {
    buckets,err := s3.ListBuckets()
    if err != nil {
        panic(err)
    }

    for _, bucket := range(buckets) {
        bucket.InventoryObjects()
        fmt.Println("")
        fmt.Printf("Bucket Name:                 %s\n", bucket.Name)
        fmt.Printf("Bucket Created:              %s\n", bucket.Created)
        fmt.Printf("Bucket Region:               %s\n", bucket.Location())
        fmt.Printf("Bucket Total Size:           %d\n", bucket.Size())
        fmt.Printf("Bucket Total Size (Pretty):  %s\n", bucket.PrettySize("MB"))
        fmt.Printf("Bucket Object Count:         %d\n", bucket.FileCount)
        fmt.Println("")
    }
}

func init() {
	listCmd.AddCommand(bucketsCmd)
}
