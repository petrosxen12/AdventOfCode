import re


def parse_expression(expression):
   
    regex = r"move (\d+) from (\d+) to (\d+)"

    matches = re.findall(regex, expression)

    # print(matches)
    matches_ints = [int(x) for x in matches[0]]

    return matches_ints


def move_crates(soc, number_of_crates, origin_stack_id, destination_stack_id):

    origin_stack = soc[origin_stack_id]
    origin_stack_cut = origin_stack[:-number_of_crates]

    # reversal happens in place to simulate movement of one crate at a time
    origin_stack_cut.reverse()

    for crate in origin_stack_cut:
        soc[destination_stack_id].append(crate)

    new_origin_stack = soc[origin_stack_id][:(len(origin_stack)-number_of_crates)]

    soc[origin_stack_id] = new_origin_stack

    return soc


if __name__ == "__main__":
    stack_of_crates = {
        1: ['Z','J','N','W','P','S'],
        2: ['G','S','T'],
        3: ['V','Q','R','L','H'],
        4: ['V','S','T','D'],
        5: ['Q','Z','T','D','B','M','J'],
        6: ['M','W','T','J','D','C','Z','L'],
        7: ['L','P','M','W','G','T','J'],
        8: ['N','G','M','T','B','F','Q','H'],
        9: ['R','D','G','C','P','B','Q','M'],
    }

    # print(stack_of_crates)
    read_flag = 0
    counter_flag = 0

    with open("input.txt") as f:
        all_lines = f.readlines()
        length_lines = len(all_lines)

        for index,line in enumerate(all_lines[10:]):
            # if read_flag == 0:
            #     counter_flag+=1
            # if read_flag == 1:
            print((index/length_lines)*100)
            movements = parse_expression(line.strip())
            stack_of_crates = move_crates(soc=stack_of_crates, number_of_crates=movements[0], origin_stack_id=movements[1],destination_stack_id= movements[2])

            print(line)
            # print(stack_of_crates)

            # if line == '\n':
            #     read_flag = 1        
    
    for x in stack_of_crates.items():
        print(f"{x[0]}, {x[1]}")