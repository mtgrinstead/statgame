# statgame


Welcome to the **Stat Game**! The goal is simple: raise one of your stats (Strength, Health, Stamina, or Intelligence) to 10 while balancing the others. Every stat increase causes the other stats to decrease, and you cannot increase the same stat twice in a row. Can you reach 10 in any stat and win?

## How to Play

1. **Objective**: The goal of the game is to raise one stat to 10 while balancing others. The four stats you can manipulate are:
   - Strength
   - Health
   - Stamina
   - Intelligence

2. **Gameplay**:
   - Each time you choose to increase a stat, the others will decrease.
   - You cannot increase the same stat twice in a row.
   - Each stat increase is a random amount between 1 and 5.
   - Each stat decrease is a random amount between 0 and 2.
   - The game ends when one of the stats reaches 10 (you win!) or you quit.

3. **Commands**:
   - On each turn, you'll be prompted to choose a number corresponding to the stat you want to raise:
     - **1**: Increase Strength
     - **2**: Increase Health
     - **3**: Increase Stamina
     - **4**: Increase Intelligence
     - **5**: View your current stats
     - **6**: Quit the game

4. **Win Condition**:
   - You win the game if any stat reaches 10.

## How to Run the Game

### Prerequisites

- You need to have Go installed on your machine.

### Steps to Play

1. Clone or download the project to your local machine.

2. Navigate to the project directory in the terminal.

3. Run the game by executing the following command:

   go run *.go
