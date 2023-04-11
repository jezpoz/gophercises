# Exercise 1 - Quiz

[Task description](https://github.com/gophercises/quiz)

My code does the following:

1. Parse flags that are specified
   - `-file <path relative to main.go>`
   - `-time <duration>`
   - `-shuffle`
2. Load questions from CSV file
   - Either default `problems.csv` or a csv-file supplied with the `file` flag. (`go run main.go -file other.csv`)
   - Note: If `-shuffe` flag is present/set to true, also randomize the question-array
3. Set up a reader that reads from `os.Stdin`
4. Prompt user to click enter when ready
5. Set up timer for the specified duration (either default 30 seconds, or the duration set with the `-time`-flag)
6. Loop through the questions in the CSV file, and ask the user these questions and compare their answers to what is defined in the CSV
7. When all questions are answered or times runs out, print a score and quit.
