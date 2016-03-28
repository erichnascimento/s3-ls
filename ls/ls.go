package ls

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// List files
func List(file string, config *Config) ([]*FileInfo, error) {
	client := s3.New(&aws.Config{Region: config.Region})

	if len(file) > 0 && file[0] == '/' {
		file = file[1:]
	}

	// Obtem os arquivos diretorio
	params := &s3.ListObjectsInput{
		Bucket: aws.String(config.Bucket),
		Prefix: aws.String(file),
	}
	response, err := client.ListObjects(params)

	if err != nil {
		return nil, err
	}

	if *response.IsTruncated {
		return nil, fmt.Errorf(
			"Sorry, the max limit for files in list is %d. The solution is comming soon :(",
			*response.MaxKeys)
	}

	files := make([]*FileInfo, len(response.Contents))
	for i, obj := range response.Contents {
		files[i] = &FileInfo{
			Name:         *obj.Key,
			LastModified: *obj.LastModified,
			Owner:        *obj.Owner.DisplayName,
			Size:         uint64(*obj.Size),
		}
	}

	return files, nil
}

/*


package main

import (
	"fmt"
	"log"

	"youse/yc-bkp-pickup/storage"

	"github.com/tj/docopt"
	"github.com/dustin/go-humanize"
)

const VERSION = "0.0.1"

const Usage = `
  Usage:
    yc-bkp-pickup [--list] <dbname>
    yc-bkp-pickup -h | --help
    yc-bkp-pickup -v | --version

  Options:
  	-l, --list     list Backup files
    -h, --help     output help information
    -v, --version  output version

`

func main() {
	args, err := docopt.Parse(Usage, nil, true, VERSION, false)
	if err != nil {
		log.Fatal(err)
	}

	dbname := args["<dbname>"].(string)

	err, files := getFiles(dbname)
	if (err != nil) {
		log.Fatal(err)
	}

	if args["--list"].(bool) {
		listFiles(files)
		return
	}

	file := showChoice(files)
	if file == nil {
		log.Fatal("Arquivo inválido")
	}

	fmt.Println(file)
	return

	//p := picker.NewPicker(picker.NewDefaultConfig())
	//p.Pickup(dbname, "/tmp/" + dbname, "")

	fmt.Print(Usage)
}

func showChoice(files []*storage.FileInfo) *storage.FileInfo {
	for k, f := range files {
		fmt.Printf(" %4d: %6s %15s %s\n", k + 1, humanize.Bytes(f.Size), humanize.Time(f.LastModified), f.Name)
	}

	fmt.Print(" Selecione um arquivo (ENTER para sair): ")

	choice := 0
	fmt.Scanf("%d", &choice)

	switch {
	case choice == 0:
		return nil
	case choice > len(files) || choice < 0:
		fmt.Println("Escolha inválida!!!")
		return nil
	}

	return files[choice - 1]
}

func listFiles(files []*storage.FileInfo) {
	fmt.Printf("total %d\n", len(files))

	for _, f := range files {
		fmt.Printf(" %10s %6s %15s %s\n", f.Owner, humanize.Bytes(f.Size), humanize.Time(f.LastModified), f.Name)
	}
}

func getFiles(dbname string) (error, []*storage.FileInfo) {
	return storage.List(dbname, storage.NewConfig())
}

*/
