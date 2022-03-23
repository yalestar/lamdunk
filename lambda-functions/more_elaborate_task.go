package main

import (
    "bufio"
    "fmt"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "go-lambda/db"
    "io"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
)

const (
    keyName          = "2021-10-29-NCPDP_Monthly_Master_20211001.zip"
    bucketName       = "ncpdp"
    internalEndpoint = "http://172.17.0.1:4566"
)

func doTheNeedful() (string, error) {
    s3cfg, err := getS3Config(internalEndpoint)
    if err != nil {
        log.Println(err)
    }
    
    s3Client := getS3Client(s3cfg)
    
    log.Printf("-------- looking at %s in %s", keyName, bucketName)
    globject, err := getObject(*s3Client, bucketName, keyName)
    
    log.Println("--------- GLOBJECT?")
    log.Println(">>>>>>>>>>", *globject.ContentType)
    
    log.Printf("->>>>>mbout to list files")
    listObjects(*s3Client, bucketName)
    
    fp, err := createLocalFile(globject.Body)
    fmt.Println("FP HAS TURNED OUT TO BE: ", fp)
    if err != nil {
        log.Println(err)
    }
    
    unzippedPath, unzipErr := UnzipMas(fp, "wangoztango")
    if unzipErr != nil {
        log.Println(unzipErr)
    }
    
    // masFile := filepath.Join(unzippedPath, "mas.txt")
    // bytes, readErr := os.ReadFile(masFile)
    // if readErr != nil {
    //     log.Println("COULD IN NO WAY READ FILE", readErr)
    // }
    log.Printf("Unzipped path: %s", unzippedPath)
    
    dbHealthCheck()
    return fp, nil
}

func dbHealthCheck() {
    tableCheck, err := db.DescribeTable()
    if err != nil {
        log.Println(err)
    }
    
    log.Println("---------------- TABLECZECH: ", tableCheck)
}

func catFile(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        log.Println(err)
    }
    fileStat, err := os.Stat(file.Name())
    fileSize := fileStat.Size()
    lineCount := fileSize / 1000
    
    log.Printf("LInes in mas.txt: %d", lineCount)
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        line := scanner.Text()
        log.Println(line)
    }
}
func createLocalFile(blob io.Reader) (string, error) {
    fudBytes, err := ioutil.ReadAll(blob)
    if err != nil {
        log.Println("_>>>>>>>>>>>>>> ReadAll error")
        log.Println(err)
    }
    where, err := os.Getwd()
    if err != nil {
        
        log.Println("UNABLE TO GET WHERE WE ARE", err)
    }
    log.Printf("------- About to create %s", keyName)
    log.Printf("------- Where are we anyway?%s", where)
    newFileName := filepath.Join("/tmp", "ship.zip")
    newFile, err := os.Create(newFileName)
    defer newFile.Close()
    fullAssPath, fpErr := filepath.Abs(newFile.Name())
    if fpErr != nil {
        log.Println("NO GET FULL PATH 4 U")
        log.Println(fpErr)
    }
    
    bytesWrote, err := newFile.Write(fudBytes)
    if err != nil {
        log.Println(">>>>>>>>>>>>>>>>>>>>>>>> WRITIN' ERROR")
        log.Println(err.Error())
    }
    log.Printf("->>>>> Supposedly wrote %d bytes to %s\n", bytesWrote, newFileName)
    fileExists, feErr := os.Stat(fullAssPath)
    if feErr != nil {
        log.Println(feErr)
    }
    log.Println("---------do it exists?", fileExists.Name())
    log.Println("---------do it exists Mtime?", fileExists.ModTime())
    log.Println("---------do it exists Size?", fileExists.Size())
    log.Println("---------full ass path", fullAssPath)
    
    return fullAssPath, nil
    
}
func handler(event events.S3Event) error {
    // ASS-ume a file named 2021-10-29-NCPDP_Monthly_Master_20211001.zip
    // is in S3 bucket named ncpdp running on localstack with an
    // internalEndpoint localhost:4566
    
    good, err := doTheNeedful()
    
    if err != nil {
        log.Println(err)
    }
    fmt.Println(good)
    return err
}

func main() {
    lambda.Start(handler)
}
