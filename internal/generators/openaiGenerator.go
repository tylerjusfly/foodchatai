package generators

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func OpenaiGenerator(prompt string) (string, error) {
	ctx := context.Background()

	OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")

	if OPENAI_API_KEY == "" {
		fmt.Println("OPENAI_API_KEY is not set!")
		err := errors.New("OPENAI_API_KEY is not set")
		return "", err
	}

	//   fmt.Println(OPENAI_API_KEY)

	llm, err := openai.New(openai.WithToken(OPENAI_API_KEY))

	if err != nil {
		return "", err
	}

	AIQuery := fmt.Sprintf(templateStr, prompt)
	//   fmt.Println(AIQuery)

	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, AIQuery)
	if err != nil {
		return "", err
	}

	return completion, nil

}

const templateStr = `
I will ask you a question about food stuff and how to eat healthy.

Get your information from the book How Not to Die by Michael Greger, M.D. FACLM and advice user based on what you read.

If the question does not relates to nutrition and food, Just reply them "other conversation outside of eating healthy is off limits".

Remember is a app for eating healthy, give then nutrition facts.

Sometimes ask if you can setup a timetable for them to fix thier eating habits.

Question:
%s

`
