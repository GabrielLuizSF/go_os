package main

import (
 "fmt"
 "os"
 "log"

)



func main(){

	FileCreator()

}

func FileCreator(){


	file := os.Create("express_api.js");
  
	if file != nil {
    
	log.fatal(file);

  }
  defer file.Close();
  _,fileWriter := file.WriteString("\n// GET method route\napp.get('/', function (req, res) {\nres.send('GET request to the homepage');\n});\n// POST method route\n\napp.post('/', function (req, res) {\nres.send('POST request to the homepage');\n});\n");

  if fileWriter != nil {
	log.fatal(fileWriter);
  }

  fmt.Println("Template Create");
}


