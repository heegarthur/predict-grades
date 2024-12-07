#predict grades

I made a bot in golang where you can predict your grades, I didn't compiled it to a .exe btw because my antivirus detecter deleted it.

you can:
  -predict your grades
  
  -generate motivation
  
  -choose a study method
  
  -customize the grades bot

**you have to customize it because it is now by it's default customized on MY grades**


#how to use

first type this in your powershell:
```
go mod tidy
```
and
```
go mod init predict_grades
```

and this:
```
go get github.com/heegarthur/predict-grades
```


  in your golang project:
  ```
    package main
    
    import (
        "fmt"
        "github.com/heegarthur/predict-grades"  // the go grades predicter
    )
    
    func main() {
        grades_bot.main_grades_predict() //the main function of my grades bot
    }
  ```
this are all the functions that are usefull to run apart:
```
    package main
    
    import (
        "fmt"
        "github.com/heegarthur/predict-grades"  // the go grades predicter
    )
    
    func main() {
        grades_bot.main_grades_predict() //main
        grades_bot.generate_motivation() //generates motivation quotes
        grades_bot.choose_studyway() //choose a study way
        grades_bot.main_predict() //where you can predict your grades and edit them
        
    }
```

if you want to edit the raw grades you can open the customizebot.txt file
