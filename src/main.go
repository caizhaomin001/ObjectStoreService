package main

import (
	"fmt"
	"os"
	"runtime"
)

func getStorePathpath() string {
	path := UserHomeDir() + string(os.PathSeparator) + "object_storage_service" + string(os.PathSeparator) + "data"
	_, err := os.Stat(path)
	if err != nil {
		if !os.IsExist(err) {
			err := os.MkdirAll(path, 0755)
			if err != nil {
				fmt.Printf("create store path failed, error:%s", err.Error())
				os.Exit(-1)
			}
			fmt.Print("create store path successfully.")
		}
	}
	return path
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func main() {
	path := getStorePathpath()
	server := &Server{StoreAddr: "0.0.0.0:9000", DataPath: path}
	server.start()
}
