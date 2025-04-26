interface LoremConfig {
    paragraphs: number;
    minWords?: number;
    maxWords?: number;
  }
  
  /**
   * Class for generating Lorem Ipsum text
   */
  class LoremGenerator {
    private words: string[];
    private static instance: LoremGenerator;
  
    private constructor() {
      this.words = [
        "lorem", "ipsum", "dolor", "sit", "amet", "consectetur", 
        "adipiscing", "elit", "sed", "do", "eiusmod"
      ];
    }
  
    public static getInstance(): LoremGenerator {
      if (!LoremGenerator.instance) {
        LoremGenerator.instance = new LoremGenerator();
      }
      return LoremGenerator.instance;
    }
  
    /**
     * Generates Lorem Ipsum text based on provided configuration
     */
    public generateText(config: LoremConfig): string {
      const { paragraphs, minWords = 5, maxWords = 8 } = config;
      const result: string[] = [];
  
      for (let i = 0; i < paragraphs; i++) {
        const numSentences = this.randomNumber(2, 4);
        const paragraph: string[] = [];
        
        for (let j = 0; j < numSentences; j++) {
          paragraph.push(this.generateSentence(minWords, maxWords));
        }
  
        result.push(paragraph.join(" "));
      }
  
      return result.join("\n\n");
    }
  
    private generateSentence(min: number, max: number): string {
      const length = this.randomNumber(min, max);
      const sentence = Array(length).fill("")
        .map(() => this.words[Math.floor(Math.random() * this.words.length)]);
      
      sentence[0] = sentence[0].charAt(0).toUpperCase() + sentence[0].slice(1);
      
      return sentence.join(" ") + ".";
    }
  
    private randomNumber(min: number, max: number): number {
      return Math.floor(Math.random() * (max - min + 1)) + min;
    }
  }
  
  // Example usage
  const generator = LoremGenerator.getInstance();
  const text = generator.generateText({ paragraphs: 1 });
  console.log(text);