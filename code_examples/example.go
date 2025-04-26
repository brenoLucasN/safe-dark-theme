package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// LoremConfig defines configuration for Lorem Ipsum generation
type LoremConfig struct {
	Paragraphs int
	MinWords   int
	MaxWords   int
	MinSent    int
	MaxSent    int
}

// LoremGenerator handles Lorem Ipsum text generation
type LoremGenerator struct {
	words  []string
	config LoremConfig
	rng    *rand.Rand
}

// NewLoremGenerator creates a new generator with default settings
func NewLoremGenerator() *LoremGenerator {
	words := []string{
		"lorem", "ipsum", "dolor", "sit", "amet", "consectetur",
		"adipiscing", "elit", "sed", "do", "eiusmod",
	}

	config := LoremConfig{
		Paragraphs: 2,
		MinWords:   5,
		MaxWords:   8,
		MinSent:    2,
		MaxSent:    4,
	}

	seed := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(seed)

	return &LoremGenerator{
		words:  words,
		config: config,
		rng:    rng,
	}
}

// SetConfig updates generator configuration
func (lg *LoremGenerator) SetConfig(config LoremConfig) {
	lg.config = config
}

// GenerateText creates Lorem Ipsum text based on configuration
func (lg *LoremGenerator) GenerateText() string {
	paragraphs := make([]string, 0, lg.config.Paragraphs)

	for i := 0; i < lg.config.Paragraphs; i++ {
		paragraphs = append(paragraphs, lg.generateParagraph())
	}

	return strings.Join(paragraphs, "\n\n")
}

func (lg *LoremGenerator) generateParagraph() string {
	numSent := lg.rng.Intn(lg.config.MaxSent-lg.config.MinSent+1) + lg.config.MinSent
	sentences := make([]string, 0, numSent)

	for i := 0; i < numSent; i++ {
		sentences = append(sentences, lg.generateSentence())
	}

	return strings.Join(sentences, " ")
}

func (lg *LoremGenerator) generateSentence() string {
	numWords := lg.rng.Intn(lg.config.MaxWords-lg.config.MinWords+1) + lg.config.MinWords
	words := make([]string, 0, numWords)

	for i := 0; i < numWords; i++ {
		wordIdx := lg.rng.Intn(len(lg.words))
		word := lg.words[wordIdx]

		if i == 0 {
			// Capitalize first word
			word = strings.ToUpper(word[:1]) + word[1:]
		}

		words = append(words, word)
	}

	return strings.Join(words, " ") + "."
}

// Statistics contains information about text generation
type Statistics struct {
	WordCount   int
	SentCount   int
	ParaCount   int
}

func main() {
	// Basic example
	generator := NewLoremGenerator()
	text := generator.GenerateText()
	fmt.Println(text)
	fmt.Println()

	// Custom configuration
	generator.SetConfig(LoremConfig{
		Paragraphs: 1,
		MinWords:   6,
		MaxWords:   10,
		MinSent:    3,
		MaxSent:    5,
	})

	customText := generator.GenerateText()
	fmt.Println(customText)
}