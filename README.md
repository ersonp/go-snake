### Timeline
#### 16-03-2021 (30 mins)
- cli                     
    made to take user input of height and width.
- pkg                    
    created a base backbone of the game with a dynamic 2d array(height and width taken from user).
#### 17-03-2021 (2-3 hrs)
- cli
    - clear.go                    
    added code to clear the cmd to render the sake.
- pkg
    - board.go                    
        added func to move the sake.
    - game-state.go                    
        made a loop so the position of snake will be updated and also used CallClear() from cli/clear.go to clear the cmd.
    - render.go                    
        rendered the snake.
    - snake.go                    
        set up movement parameters for the snake.
#### 18-03-2021 (6-7 hrs)
- pkg
    - game-state.go                    
        added keyboard events.
    - render.go                    
        added additional info.
    - snake.go                    
        derectionalised movements of snake and collusion detection.
    - keyboard.go                    
        to get real-time keyboard inputs with [This Library](github.com/eiannone/keyboard).
    - food.go                    
        added food logic.
- vendor                    
    added vendor directory
- Makefile                    
    created makefile with run and format