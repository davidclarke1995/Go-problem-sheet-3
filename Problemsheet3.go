//David Clarke G00335563
package main

//imports
import (
    "fmt"
    "math/rand"
    //"bufio"
    "time"
    "regexp"
	//"os"
	"strings"
)


    //function taking in a single string as an input
    func ElizaResponse(inputStr string) string {

        //var userInput string

        //fmt.Println("Please message me")
        //fmt.Scanf("%s", &userInput)
		//userInput :=inputStr

		userInput := inputStr
		pattern := `.*father.*`
		//pattern := `(?!).*\bfather\b.*`
		//pattern2 := []string{`.*i am(.*)`, `.*I AM(.*)`, `.*I'm(.*)`, `.*i'm(.*)`, `.*im(.*)`, `.*I am(.*)`}
		output := "Why dont you tell me more about your father?"
		response := ""
		rand.Seed(time.Now().UTC().UnixNano())



		answers := []string{//responses
        	"I’m not sure what you’re trying to say. Could you explain it to me?",
        	"How does that make you feel?",
        	"Why do you say that?",
        }


        //matching the word father
        if matched, _ := regexp.MatchString(pattern, userInput); matched {
        //returning the string below
        	response = "Input: " + userInput + "\nOutput :" + output

		} else {
			response = "Input: " + userInput + " \nOutput :" + answers[rand.Intn(len(answers))]
		}

		//Getting the age pattern
		re3 := regexp.MustCompile("[Ii] (?:T|t)(?:U|u)(?:R|r)(?:N|n)(?:E|e)(?:D|d) ([0-9]+)[.!?]?")
	
		if re3.MatchString(userInput){	
		return re3.ReplaceAllString(userInput, "You are $1 years old?") 

		}


		//for _, checkPattern := range pattern2 {
			re := regexp.MustCompile("(?i)(i[' a]*m) (.*)")
			if re.MatchString(userInput) {
				matched := re.FindStringSubmatch(userInput)[2]
				//match := re.ReplaceAllString(userInput, "How do you know you are $1?")
				matched = matchPronouns(matched)
				response = "Input :" + userInput + " \noutput : How do you know you are " + matched

				return response
			} //if re.MatchString
		//} //for pattern2

		return response

        //re := regexp.MustCompile(`.*I am(.*)`)
        //if re.MatchString(userInput){
       // return re.ReplaceAllLiteralString(userInput, "How do you know you are $1")
       //}

        //returning the single string response
    //return answers[rand.Intn(len(answers))]
    }
	func matchPronouns(inputStr string) string {
		// split inputStr into slice of strings
		splitStr := strings.Fields(inputStr)

		//reflected pronouns contained in a map
		pronouns := map[string]string{
			"i":      "you",
			"was":    "were",
			"i'd":    "you would",
			"i've":   "you have",
			"i'll":   "you will",
			"my":     "your",
			"are":    "am",
			"you've": "I have",
			"you'll": "I will",
			"your":   "my",
			"yours":  "mine",
			"you":    "I",
			"me":     "you",
			"me.":    "you",
			"you're": "I’m",
		}//map

		//loop map to check
		for index, elizaOutput := range splitStr {
			if value, ok := pronouns[strings.ToLower(elizaOutput)]; ok {
				splitStr[index] = value
			}
		}//loop map

		return strings.Join(splitStr, " ")
	}

	func main() {
		//array of inputs
		userInput := [] string{
			"People say I look like both my mother and father.",
			"Father was a teacher.",
			"I was my father’s favourite.",
			"I'm looking forward to the weekend.",
			"My grandfather was French!",
			"I am happy.",
			"I am not happy with your responses.",
			"I am not sure that you understand the effect that your questions are having on me.",
			"I am supposed to just take what you’re saying at face value?",
			"I am 20 years old",
		}
		
		rand.Seed(time.Now().UTC().UnixNano())//get a random number
		//reader := bufio.NewReader(os.Stdin)
		//fmt.Println("Hello, how are you?")

		//	userInput, _ := reader.ReadString('\n')
		//	fmt.Scanf("%s", &userInput)


		//userInput:= "I am happy"
		elizaOutput:= ElizaResponse(userInput[rand.Intn(len("userInput"))])


		fmt.Print("\n")
		fmt.Print(elizaOutput)
		fmt.Print("\n")
	}//Main function