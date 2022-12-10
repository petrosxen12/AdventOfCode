from typing import Union


def definetely_visible_outside_trees(trees: list[list[int]]):
    """
    Returns all trees visible from edge of grid
    """

    top_row = len(trees[0])
    side = len(trees)

    total_outside_visible = (2*top_row)+(2*side) - 4

    return total_outside_visible

def is_tree_visible(position: tuple[int,int] , tree_heights_table: list[list[int]]):

    height  = len(tree_heights_table)
    width = len(tree_heights_table[0])

    visible_up = look_up(position, tree_heights_table, height, width)
    visible_down = look_down(position, tree_heights_table, height, width)
    visible_left = look_left(position, tree_heights_table, height , width)
    visible_right = look_right(position, tree_heights_table, height, width)

    # print([visible_up, visible_down, visible_left, visible_right])

    is_visibile = any([visible_up, visible_down, visible_left, visible_right])
    return is_visibile


def look_up(position, tree_heights_table, height, width):
    y = position[0]
    x = position[1]

    current_tree_height = tree_heights_table[y][x]

    for i in range(1, y+1):
        if tree_heights_table[y-i][x] >= current_tree_height:
            return False

    return True

def look_left(position, tree_heights_table, height, width):
    y = position[0]
    x = position[1]

    current_tree_height = tree_heights_table[y][x]

    for i in range(1, x+1):
        if tree_heights_table[y][x-i] >= current_tree_height:
            return False

    return True

def look_right(position, tree_heights_table, height, width):
    y = position[0]
    x = position[1]

    current_tree_height = tree_heights_table[y][x]

    for i in range(1, (width-x)):
        if tree_heights_table[y][x+i] >= current_tree_height:
            return False

    return True

def look_down(position, tree_heights_table, height, width):
    y = position[0]
    x = position[1]

    current_tree_height = tree_heights_table[y][x]

    for i in range(1, (height-y)):
        if tree_heights_table[y+i][x] >= current_tree_height:
            return False

    return True


if __name__ == "__main__":
    filename = "input.txt"
    test_filename = "test_input.txt"

    trees = []


    with open(filename, encoding="utf-8") as f:
        trees = [x.strip() for x in f.readlines()]
        tree_heights = []

        for tr in trees:
            tr_h = []
            for tree in tr:
                tr_h.append(int(tree))

            tree_heights.append(tr_h)

        # print(tree_heights)
        
        visible_trees = 0

        for y in range(1, (len(tree_heights)-1)):
            for x in range(1, len(tree_heights[0])-1):
                
                # print(f"Position: x: {x}, y: {y}")
                # print(f"Tree: {tree_heights[y][x]}")
                # print()

                if is_tree_visible((y,x), tree_heights):
                    # print()
                    # print(f"Position: x: {x}, y: {y}")
                    # print(f"Tree: {tree_heights[y][x]}")
                    # # print()
                    # print(f"Visible at: x: {x}, y: {y}")
                    # print()
                    visible_trees+=1

    outside_visibility = definetely_visible_outside_trees(tree_heights)
    print(visible_trees)
    print(outside_visibility)

    print(f"Total visible trees: {visible_trees+outside_visibility}")