### Timeline
#### 16-03-2021
- cli                     
    made to take user input of height and width.
- pkg                    
    created a base backbone of the game with a dynamic 2d array(height and width taken from user).
#### 17-03-2021
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