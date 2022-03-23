package main

import (
    "archive/zip"
    "fmt"
    "io"
    "os"
    "path/filepath"
)

func UnzipMas(zipFile, dest string) error {
    r, err := zip.OpenReader(zipFile)
    if err != nil {
        return err
    }
    defer r.Close()
    err = os.MkdirAll(dest, 0755)
    if err != nil {
        return err
    }
    
    extractFile := func(f *zip.File) error {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer rc.Close()
        
        path := filepath.Join(dest, filepath.Base(f.Name))
        
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
            fmt.Println(f.Name)
            err := extractFile(f)
            if err != nil {
                return err
            }
        }
    }
    
    return nil
}
