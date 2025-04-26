def generate_lorem_ipsum(paragraphs=2):
    """Generate Lorem Ipsum text with specified number of paragraphs."""
    import random
    
    words = [
        "lorem", "ipsum", "dolor", "sit", "amet", "consectetur", 
        "adipiscing", "elit", "sed", "do", "eiusmod"
    ]
    
    result = []
    
    for i in range(paragraphs):
        num_sentences = random.randint(2, 4)
        paragraph = []
        
        for j in range(num_sentences):
            sentence_length = random.randint(5, 8)
            sentence = random.sample(words, sentence_length)
            sentence[0] = sentence[0].capitalize()
            paragraph.append(" ".join(sentence) + ".")
            
        result.append(" ".join(paragraph))
    
    return "\n\n".join(result)


class TextGenerator:
    def __init__(self, style="lorem"):
        self.style = style
        self.count = 0
        
    def generate(self, amount=1):
        self.count += amount
        return generate_lorem_ipsum(amount)

# Example usage
text = generate_lorem_ipsum()
print(text)