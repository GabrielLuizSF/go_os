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
    
    err := os.Mkdir("../" + dirname, 0755);
    
    if err != nil {
        log.Fatal(err)
    }

   fmt.Println("Folder Created");
}
