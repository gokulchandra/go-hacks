package main

import (
	"log"
	"math/rand"
	"bufio"
	"os"
	"regexp"
	"fmt"
	"strings"
)

type ElizaResponses struct {
	patterns map[string][]string
}

type Reflection struct {
	responses map[string]string
}

func (r *Reflection) reflect(in string) string {
	return r.responses[in]
}

func (e *ElizaResponses) respond(s string) string {
	for pattern, replies := range e.patterns {
		r := regexp.MustCompile(pattern)
		match := r.MatchString(s)
		if match {
			return fmt.Sprintf(getElizaResponse(replies[rand.Int()%len(replies)], getMatchedMap(r, s)))
		}
	}
	return "Sorry I dont know what you are saying."
}
func getElizaResponse(s string, matchedMap map[string]string) string {
	return strings.Replace(s, "%s", matchedMap["context"], -1)
}
func getMatchedMap(r *regexp.Regexp, s string) map[string]string {
	paramsMap := make(map[string]string)
	match := r.FindStringSubmatch(s)
	for i, name := range r.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}

var reflections = &Reflection{make(map[string]string)}

func main() {
	println("How are you doing today?")
	s := bufio.NewScanner(os.Stdin)
	eliza := initializeEliza()
	for s.Scan() {
		userInput := s.Text()
		if userInput == "quit" {
			return
		}
		log.Println(eliza.respond(userInput))
	}
}
func initializeEliza() ElizaResponses {
	eliza := ElizaResponses{make(map[string][]string)}
	eliza.patterns = map[string][]string{
		`I need (?P<context>.*)`: {`Why do you need %s?`,
			`Would it really help you to get %s?`,
			`Are you sure you need %s?`,
		},

		`Why don\'?t you (?P<context>[^\?]*)\??`: {`Do you really think I don't %s?`,
			`Perhaps eventually I will %s.`,
			`Do you really want me to %s?`,
		},

		`Why can\'?t I (?P<context>[^\?]*)\??`: {`Do you think you should be able to %s?`,
			`If you could %s, what would you do?`,
			`I don't know -- why can't you %s?`,
			`Have you really tried?`,
		},

		`I can\'?t (?P<context>.*)`: {`How do you know you can't %s?`,
			`Perhaps you could %s if you tried.`,
			`What would it take for you to %s?`,
		},

		`I am (?P<context>.*)`: {`Did you come to me because you are %s?`,
			`How long have you been %s?`,
			`How do you feel about being %s?`,
		},

		`I\'?m (?P<context>.*)`: {`How does being %s make you feel?`,
			`Do you enjoy being %s?`,
			`Why do you tell me you're %s?`,
			`Why do you think you're %s?`,
		},

		`Are you (?P<context>[^\?]*)\??`: {`Why does it matter whether I am %s?`,
			`Would you prefer it if I were not %s?`,
			`Perhaps you believe I am %s.`,
			`I may be %s -- what do you think?`,
		},

		`What (?P<context>.*)`: {`Why do you ask?`,
			`How would an answer to that help you?`,
			`What do you think?`,
		},

		`How (?P<context>.*)`: {`How do you suppose?`,
			`Perhaps you can answer your own question.`,
			`What is it you're really asking?`,
		},

		`Because (?P<context>.*)`: {`Is that the real reason?`,
			`What other reasons come to mind?`,
			`Does that reason apply to anything else?`,
			`If %s, what else must be true?`,
		},

		`(?P<context>.*) sorry (?P<context>.*)`: {`There are many times when no apology is needed.`,
			`What feelings do you have when you apologize?`,
		},

		`Hello(?P<context>.*)`: {`Hello... I'm glad you could drop by today.`,
			`Hi there... how are you today?`,
			`Hello, how are you feeling today?`,
		},

		`I think (?P<context>.*)`: {`Do you doubt %s?`,
			`Do you really think so?`,
			`But you're not sure %s?`,
		},

		`(?P<context>.*) friend (?P<context>.*)`: {`Tell me more about your friends.`,
			`When you think of a friend, what comes to mind?`,
			`Why don't you tell me about a childhood friend?`,
		},

		`Yes`: {`You seem quite sure.`,
			`OK, but can you elaborate a bit?`,
		},

		`(?P<context>.*) computer(?P<context>.*)`: {`Are you really talking about me?`,
			`Does it seem strange to talk to a computer?`,
			`How do computers make you feel?`,
			`Do you feel threatened by computers?`,
		},

		`Is it (?P<context>.*)`: {`Do you think it is %s?`,
			`Perhaps it's %s -- what do you think?`,
			`If it were %s, what would you do?`,
			`It could well be that %s.`,
		},

		`It is (?P<context>.*)`: {`You seem very certain.`,
			`If I told you that it probably isn't %s, what would you feel?`,
		},

		`Can you (?P<context>[^\?]*)\??`: {`What makes you think I can't %s?`,
			`If I could %s, then what?`,
			`Why do you ask if I can %s?`,
		},
	}
	return eliza
}
