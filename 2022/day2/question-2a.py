def round_outcome(my_choice, opponent_choice):
    """
    1. Rock defeats Scissors
    2. Scissors defeats Paper
    3. Paper defeats Rock. 
    - If both players choose the same shape, the round instead ends in a draw.

    my_choice is: X for Rock, Y for Paper, and Z for Scissors.
    opponent_choice is: A for Rock, B for Paper, and C for Scissors
    """ 

    if (my_choice == "X" and opponent_choice == "A") or (my_choice == "Y" and opponent_choice == "B") or (my_choice == "Z" and opponent_choice == "C"):
        return "d"

    # winning scenarios
    if my_choice == "X" and opponent_choice == "C":
        return "w"

    if my_choice == "Z" and opponent_choice == "B":
        return "w"

    if my_choice == "Y" and opponent_choice == "A":
        return "w"

    # lose scenarios
    if my_choice == "Z" and opponent_choice == "A":
        return "l"

    if my_choice == "Y" and opponent_choice == "C":
        return "l"

    if my_choice == "X" and opponent_choice == "B":
        return "l"

def choice_points(choice):
    if choice == "X":
        return 1
    
    if choice == "Y":
        return 2

    if choice == "Z":
        return 3

def win_lose_points(outcome):
    if outcome == "l":
        return 0
    if outcome == "w":
        return 6
    if outcome == "d":
        return 3

if __name__ == "__main__":

    total_points = 0

    with open("input.txt") as f:
        for line in f.readlines():
            opponent_mychoice = [x.strip() for x in line.split(" ")]
            print(opponent_mychoice)

            opponent_choice = opponent_mychoice[0]
            my_choice = opponent_mychoice[1]
       
            round_score = win_lose_points(round_outcome(my_choice=my_choice, opponent_choice=opponent_choice)) + choice_points(choice=my_choice)

            total_points+=int(round_score)

    print(total_points)
