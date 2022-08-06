package main

import (
    "log";
    "os";
    "fmt";
)

func main(){
     UserChoice()
}

func UserChoice(){
    var userinput int;
    fmt.Println("\n\nOptions:\n 1-DirCreator\t 2-FileCreator\t 3-FileEraser\n");
    fmt.Scanln(&userinput)

    switch userinput{
        case 1:
        DirCreator()
        case 2:
        FileCreator()
        case 3:
        DeleteFile()
        default:
       fmt.Println(" NAN\n\n")
    }
    
}
func DirCreator() {

    var dirname string = ""
    fmt.Println("DIR NAME:\t");
    fmt.Scanln(&dirname)
    
    folder := os.Mkdir("./" + dirname, 0755);
    
    if folder != nil {
        log.Fatal(folder)
    }else{
       fmt.Println("\nFolder Created\n"); 
    }

   
}

func FileCreator(){
    file, createFile := os.Create("express_api.js");
  
	if createFile != nil {
    
	log.Fatal(createFile);

  }

  defer file.Close();


  _,fileWriter := file.WriteString("\n// GET method route\napp.get('/', function (req, res) {\nres.send('GET request to the homepage');\n});\n// POST method route\n\napp.post('/', function (req, res) {\nres.send('POST request to the homepage');\n});\n");

  if fileWriter != nil {
	log.Fatal(fileWriter);
  }else{
    fmt.Println("\nFile Created\n")
}
    }

type DeleteConfig struct {
    fileName   string
}
func DeleteFile(){
var del DeleteConfig


fmt.Println("Escolha o nome do arquivo junto com a extensão\n exemplo: index.js");
fmt.Scanln(&del.fileName)

 ideleteFile := os.Remove(del.fileName)
    if ideleteFile != nil {
        log.Fatal(ideleteFile)
    }else{
    
    fmt.Println("Foi Deletado o",del.fileName)
        }
}


