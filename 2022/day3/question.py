def find_common_letter_three_elf_group(groups):
    elf_group_common_letter = list(set(groups[0]) & set(groups[1]) & set(groups[2]))
    # print(elf_group_common_letter)
    return elf_group_common_letter[0]


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

    three_elf_group = []

    with open("input.txt") as f:
        for i, rucksack in enumerate(f.readlines()):
            # first_comp, second_comp = split_compartments(rucksack)
            # common_letter = find_common_letter(first_comp, second_comp)
            three_elf_group.append(rucksack.strip())

            if len(three_elf_group) == 3:
                common_letter = find_common_letter_three_elf_group(three_elf_group)
                ltr_prt = letter_priority(common_letter)

                total_priorities+= ltr_prt
                three_elf_group.clear()
                
    print(total_priorities)