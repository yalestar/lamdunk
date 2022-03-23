package main

import (
    "archive/zip"
    "io"
    "log"
    "os"
    "path/filepath"
)

func UnzipMas(src, dest string) (string, error) {
    log.Printf("I was asked to unzip %s to %s", src, dest)
    r, err := zip.OpenReader(src)
    if err != nil {
        return "", err
    }
    defer r.Close()
    realDestDir := filepath.Join("/tmp", dest)
    err = os.MkdirAll(realDestDir, 0755)
    if err != nil {
        return "", err
    }
    
    extractFile := func(f *zip.File) error {
        log.Printf("========Extracting %v", f.Name)
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer rc.Close()
        
        path := filepath.Join(realDestDir, filepath.Base(f.Name))
        log.Printf("Destination dir was %s", path)
        if f.FileInfo().IsDir() {
            os.MkdirAll(path, f.Mode())
        } else {
            
            f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, f.Mode())
            if err != nil {
                return err
            }
            defer f.Close()
            
            _, err = io.Copy(f, rc)
            if err != nil {
                return err
            }
        }
        return nil
    }
    
    for _, f := range r.File {
        if filepath.Base(f.Name) == "mas.txt" {
            log.Println(f.Name)
            err := extractFile(f)
            if err != nil {
                log.Println("========================= SHIT WENT TO SHIT")
                log.Println(err)
                return "", err
            }
            break
        }
    }
    log.Printf(">>>>>>>>>>>>>> Returning %s from unzip method", realDestDir)
    return realDestDir, nil
}
