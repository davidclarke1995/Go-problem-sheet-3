//David Clarke G00335563
package main

//imports
import (
    "fmt"
    "math/rand"
    //"bufio"
    "time"
    "regexp"
)


    //function taking in a single string as an input
    func ElizaResponse() string {

        var userInput string

        fmt.Println("Please message me")
        fmt.Scanf("%s", &userInput)

        //matching the word father
        if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, userInput); matched {
            //returning the string below
            return "Why don’t you tell me more about your father?"

        }

        re := regexp.MustCompile("I am ([^.!?]*)[.!?]?")
        if re.MatchString(userInput){
        return re.ReplaceAllLiteralString(userInput, "How do you know you are $1")
        }

        answers := []string{//responses
        "I’m not sure what you’re trying to say. Could you explain it to me?",
        "How does that make you feel?",
        "Why do you say that?",
        }

        //returning the single string response
    return answers[rand.Intn(len(answers))]
    }


func main() {
rand.Seed(time.Now().UTC().UnixNano())//random number
ElizaResponse()
}