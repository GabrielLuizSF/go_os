package dir

import (
    "log";
    "os";
    "fmt";
)

func DirCreate() {


 var dirname string=""
    fmt.Println("DIR NAME:\t");
    fmt.Scanln(&dirname)
    
    folder := os.Mkdir("../" + dirname, 0755);
    
    if folder != nil {
        log.Fatal(folder)
    }

   fmt.Println("Folder Created");
}
