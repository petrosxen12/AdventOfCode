def part_two_match_outcome(opponent_choice, match_outcome):
    """
     1. X means you need to lose.
     2. Y means you need to end the round in a draw.
     3. Z means you need to win. 
    """

    if match_outcome == "X":
        if opponent_choice == "A":
            return choice_points("Z")

        if opponent_choice == "C":
            return choice_points("Y")

        if opponent_choice == "B":
            return choice_points("X")

    if match_outcome == "Y":
        if opponent_choice == "A":
            return choice_points("X")

        if opponent_choice == "C":
            return choice_points("Z")

        if opponent_choice == "B":
            return choice_points("Y")

    if match_outcome == "Z":
        if opponent_choice == "C":
            return choice_points("X")

        if opponent_choice == "B":
            return choice_points("Z")

        if opponent_choice == "A":
            return choice_points("Y")


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

def win_lose_points_part_two(outcome):
    if outcome == "X":
        return 0
    if outcome == "Z":
        return 6
    if outcome == "Y":
        return 3

if __name__ == "__main__":

    total_points = 0

    with open("input.txt") as f:
        for line in f.readlines():
            opponent_mychoice = [x.strip() for x in line.split(" ")]
            print(opponent_mychoice)

            opponent_choice = opponent_mychoice[0]
            my_choice = opponent_mychoice[1]
            
            outcome = opponent_mychoice[1]
       
            # round_score = win_lose_points(round_outcome(my_choice=my_choice, 
            # opponent_choice=opponent_choice)) + choice_points(choice=my_choice)

            round_score = part_two_match_outcome(opponent_choice, outcome) + win_lose_points_part_two(outcome)


            total_points+=int(round_score)

    print(total_points)
