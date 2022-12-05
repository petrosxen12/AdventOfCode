def split_compartments(rucksack_contents):
    contents_length = len(rucksack_contents)
    
    halfway = int(contents_length/2)

    first = rucksack_contents[:halfway]
    second = rucksack_contents[halfway:]

    return first, second

def find_common_letter(first, second):
    f = set(first)
    s = set(second)

    common_letter = list(f & s)

    return common_letter[0]

def letter_priority(letter):
    """
    - Lowercase item types a through z have priorities 1 through 26.
    - Uppercase item types A through Z have priorities 27 through 52.
    """
    # means lower letter
    if ord(letter) > 96:
        return ord(letter) - 96

    else:
        return ord(letter) - 38


if __name__ == "__main__":
    total_priorities = 0

    with open("input.txt") as f:
        for rucksack in f.readlines():
            first_comp, second_comp = split_compartments(rucksack)
            common_letter = find_common_letter(first_comp, second_comp)
    
            ltr_prt = letter_priority(common_letter)

            total_priorities+= ltr_prt

    print(total_priorities)