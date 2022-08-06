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
    fmt.Println("\n\nOptions:\n 1-DirCreator\t 2-FileCreator\n");
    fmt.Scanln(&userinput)

    switch userinput{
        case 1:
        DirCreator()
        case 2:
        FileCreator()
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
    }

   fmt.Println("\nFolder Created\n");
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
  }
    fmt.Println("\nFile Created\n")
}




