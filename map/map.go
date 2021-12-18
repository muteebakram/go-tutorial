package main

import (
	"strings"
)

func main() {
	stocks := map[string]int{
		"A": 100,
		"B": 200,
		"C": 300,
	}

	println("Number of stocks: ", len(stocks))

	value, ok := stocks["A"]
	if ok {
		println("Stock A: ", value)
	}

	println("\nValues of stocks")
	println("----------------------")
	for k, v := range stocks {
		println("Stock: ", k, "Value: ", v)
	}

	println("\nDeleting stock A")
	delete(stocks, "A")
	println("\nValues of stocks")
	println("----------------------")
	for k, v := range stocks {
		println("Stock: ", k, "Value: ", v)
	}

	println("\nCount words in Poem")
	println("\nPoem: 'Do not go gentle into that good night'")
	println("----------------------")
	poem := `
	Do not go gentle into that good night,
	Old age should burn and rave at close of day;
	Rage, rage against the dying of the light.
	
	Though wise men at their end know dark is right,
	Because their words had forked no lightning they
	Do not go gentle into that good night.
	
	Good men, the last wave by, crying how bright
	Their frail deeds might have danced in a green bay,
	Rage, rage against the dying of the light.
	
	Wild men who caught and sang the sun in flight,
	And learn, too late, they grieved it on its way,
	Do not go gentle into that good night.
	
	Grave men, near death, who see with blinding sight
	Blind eyes could blaze like meteors and be gay,
	Rage, rage against the dying of the light.
	
	And you, my father, there on the sad height,
	Curse, bless, me now with your fierce tears, I pray.
	Do not go gentle into that good night.
	Rage, rage against the dying of the light.
	`
	poem = strings.ToLower(poem)
	words := strings.Fields(poem)
	countWords := map[string]int{}

	for _, word := range words {
		countWords[word] += 1
	}

	for k, v := range countWords {
		println("Count: ", v, " Word: ", k)
	}
}
