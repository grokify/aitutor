package aitutor

type Personalization struct {
	CommunicationStyles CommunicationStyles `json:"communication_styles"`
	Depth               Depth               `json:"depth"`
	LearningStyles      LearningStyles      `json:"learning_styles"`
	ReasoningFrameworks ReasoningFrameworks `json:"reasoning_frameworks"`
	ToneStyles          ToneStyles          `json:"tone_styles"`
}

type Depth struct {
	Description string      `json:"description"`
	DepthLevels DepthLevels `json:"depth_levels"`
}

type DepthLevels struct {
	Level1  string `json:"Level_1"`
	Level2  string `json:"Level_2"`
	Level3  string `json:"Level_3"`
	Level4  string `json:"Level_4"`
	Level5  string `json:"Level_5"`
	Level6  string `json:"Level_6"`
	Level7  string `json:"Level_7"`
	Level8  string `json:"Level_8"`
	Level9  string `json:"Level_9"`
	Level10 string `json:"Level_10"`
}

type CommunicationStyles struct {
	Stochastic   string `json:"stochastic"`    // "Incorporates randomness or variability, generating slight variations in responses for a dynamic, less repetitive conversation.",
	Formal       string `json:"Formal"`        // "Follows strict grammatical rules and avoids contractions, slang, or colloquialisms for a structured and polished presentation.",
	Textbook     string `json:"Textbook"`      // "Resembles language in textbooks, using well-structured sentences, rich vocabulary, and focusing on clarity and coherence.",
	Layman       string `json:"Layman"`        // "Simplifies complex concepts, using everyday language and relatable examples for accessible and engaging explanations.",
	StoryTelling string `json:"Story Telling"` // "Presents information through narratives or anecdotes, making ideas engaging and memorable with relatable stories.",
	Socratic     string `json:"Socratic"`      // "Asks thought-provoking questions to stimulate intellectual curiosity, critical thinking, and self-directed learning.",
	Humorous     string `json:"Humorous"`      // "Incorporates wit, jokes, and light-hearted elements for enjoyable, engaging, and memorable content in a relaxed atmosphere."
}

type LearningStyles struct {
	Sensing    string `json:"Sensing"`                   // "Sensing": "Concrete, practical, oriented towards facts and procedures.",
	Visual     string `json:"Visual *REQUIRES PLUGINS*"` // "Visual *REQUIRES PLUGINS*": "Prefer visual representations of presented material - pictures, diagrams, flow charts",
	Inductive  string `json:"Inductive"`                 // "Inductive": "Prefer presentations that proceed from the specific to the general",
	Active     string `json:"Active"`                    // "Active": "Learn by trying things out, experimenting, and doing",
	Sequential string `json:"Sequential"`                // "Sequential": "Linear, orderly learn in small incremental steps",
	Intuitive  string `json:"Intuitive"`                 // "Intuitive": "Conceptual, innovative, oriented toward theories and meanings",
	Verbal     string `json:"Verbal"`                    // "Verbal": "Prefer written and spoken explanations",
	Deductive  string `json:"Deductive"`                 // "Deductive": "Prefer presentations that go from the general to the specific",
	Reflective string `json:"Reflective"`                // "Reflective": "Learn by thinking things through, working alone",
	Global     string `json:"Global"`                    // "Global": "Holistic, system thinkers, learn in large leaps"
}

type ToneStyles struct {
	Debate      string `json:"Debate"`      // "Assertive and competitive, challenges users to think critically and defend their position. Suitable for confident learners.",
	Encouraging string `json:"Encouraging"` // "Supportive and empathetic, provides positive reinforcement. Ideal for sensitive learners preferring collaboration.",
	Neutral     string `json:"Neutral"`     // "Objective and impartial, avoids taking sides or expressing strong opinions. Fits reserved learners valuing neutrality.",
	Informative string `json:"Informative"` // "Clear and precise, focuses on facts and avoids emotional language. Ideal for analytical learners seeking objectivity.",
	Friendly    string `json:"Friendly"`    // "Warm and conversational, establishes connection using friendly language. Best for extroverted learners preferring personal interactions."
}

type ReasoningFrameworks struct {
	Deductive  string `json:"Deductive"`  // "Draws conclusions from general principles, promoting critical thinking and logical problem-solving skills.",
	Inductive  string `json:"Inductive"`  // "Forms general conclusions from specific observations, encouraging pattern recognition and broader theories.",
	Abductive  string `json:"Abductive"`  // "Generates likely explanations based on limited information, supporting plausible hypothesis formation.",
	Analogical string `json:"Analogical"` // "Compares similarities between situations or concepts, fostering deep understanding and creative problem-solving.",
	Causal     string `json:"Causal"`     // "Identifies cause-and-effect relationships, developing critical thinking and understanding of complex systems."

}
