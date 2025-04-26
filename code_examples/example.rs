use std::collections::HashMap;
use rand::{Rng, thread_rng};
use rand::seq::SliceRandom;

/// Structure for Lorem Ipsum text generation
#[derive(Debug)]
struct LoremGenerator {
    words: Vec<String>,
    stats: HashMap<String, usize>,
}

impl LoremGenerator {
    /// Creates a new generator instance
    pub fn new() -> Self {
        let words = vec![
            "lorem", "ipsum", "dolor", "sit", "amet", "consectetur", 
            "adipiscing", "elit", "sed", "do", "eiusmod"
        ]
        .iter()
        .map(|s| s.to_string())
        .collect();

        Self {
            words,
            stats: HashMap::new(),
        }
    }

    /// Generates Lorem Ipsum text with specified number of paragraphs
    pub fn generate(&mut self, paragraphs: usize) -> String {
        let mut result = Vec::with_capacity(paragraphs);
        let mut rng = thread_rng();

        for _ in 0..paragraphs {
            let num_sentences = rng.gen_range(2..5);
            let paragraph = self.generate_paragraph(num_sentences);
            result.push(paragraph);
        }

        // Update statistics
        self.update_stats(&result);

        result.join("\n\n")
    }

    fn generate_paragraph(&self, num_sentences: usize) -> String {
        let mut result = Vec::with_capacity(num_sentences);
        let mut rng = thread_rng();

        for _ in 0..num_sentences {
            let length = rng.gen_range(5..9);
            result.push(self.generate_sentence(length));
        }

        result.join(" ")
    }

    fn generate_sentence(&self, word_count: usize) -> String {
        let mut rng = thread_rng();
        let mut sentence = Vec::with_capacity(word_count);

        for i in 0..word_count {
            let word = self.words.choose(&mut rng).unwrap().clone();
            
            if i == 0 {
                // Capitalize first word
                let mut chars = word.chars();
                match chars.next() {
                    None => sentence.push(word),
                    Some(first) => {
                        let capitalized = first.to_uppercase().collect::<String>() + chars.as_str();
                        sentence.push(capitalized);
                    }
                }
            } else {
                sentence.push(word);
            }
        }

        sentence.join(" ") + "."
    }

    fn update_stats(&mut self, texts: &[String]) {
        for text in texts {
            for word in text.split_whitespace() {
                let clean_word = word.trim_matches(|c: char| !c.is_alphanumeric()).to_lowercase();
                if !clean_word.is_empty() {
                    *self.stats.entry(clean_word).or_insert(0) += 1;
                }
            }
        }
    }
}

// Example enum
#[derive(Debug)]
enum LoremType {
    Short,
    Medium,
    Long,
}

fn main() {
    // Basic usage
    let mut generator = LoremGenerator::new();
    let text = generator.generate(1);
    println!("{}", text);
    
    // Using enum
    let lorem_type = LoremType::Short;
    let paragraphs = match lorem_type {
        LoremType::Short => 1,
        LoremType::Medium => 2,
        LoremType::Long => 3,
    };
    
    let more_text = generator.generate(paragraphs);
    println!("\n{}", more_text);
}