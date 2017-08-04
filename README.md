# countdown
Countdown solver as a service in Go


# TODO
- Basic page for / and /letters
- Unit tests   
- Move dictionary into .go file to make binary completely portable
- Command-line flags
- Improve numbers round efficiency
- Generate numbers based on rules (new route too?)
- Benchmark tests using fuzzing
- Numbers round scoring
- Numbers round answer using fewest numbers?
- Improve efficiency of numbers
- Refactor perms and combi packages to use more consistent APIs 
- ~~Numbers round~~
- ~~Error when no words are found~~
- ~~Return JSON response~~
- ~~JSON content-type header~~
- ~~Better error handling~~
- ~~Simple logging~~
- ~~Upper limit on word length - long words can be very slow~~
- ~~Time limit per search~~ (not needed)~~
- ~~Combi comments and example~~


# Useful tests for numbers round
- numbers/952,25,50,75,100,3,6 currently takes up to 384ms
- numbers/821,25,100,75,50,6,4
- numbers/556,50,8,3,7,2,10