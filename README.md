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
github.com/heegarthur/predict-grades
```

  in your golang project:
  ```
    package main
    
    import (
        "fmt"
        "jouwmodule/grades_bot"  // the go grades predicter
    )
    
    func main() {
        grades_bot.main_grades_predict() //the main function of my grades bot
    }
  ```
