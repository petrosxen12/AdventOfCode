from typing import List, Set, Union
from dataclasses import dataclass

def command(cmd_input: str):
    if cmd_input[0] == "$":
        return cmd_input[1:].strip()

    return None

def file(cmd_input: str) -> Union[Union[int, int], None] :
    # print(cmd_input[0:3])
    if cmd_input[0] != "$" and cmd_input[0:3] != 'dir':
        file_details = cmd_input.split(sep=" ")

        size = file_details[0]
        name = file_details[1].strip()

        return size, name

    return None

def directory(cmd_input: str) -> Union[str , None]:
    if cmd_input[0:3] == 'dir':
        return cmd_input[3:].strip()

    return None

def get_files_of_directory(directory_to_cd_to: str, ls_command: str):
    """
    To list files of a directory we first cd in the directory and then execute ls
    then iterate until we find another command.
    """
    directory = ""

    cmd = command(directory_to_cd_to)
    cmd_size = len(cmd.split())
    
    ls = command(ls_command)
    ls_cmd_size = len(ls)
    
    if cmd and ls:
        if cmd_size== 1:
            return None
        if cmd_size == 2 and cmd.split()[1] != "..":
            directory = cmd.split()[1]

            if ls_cmd_size == 1:
                pass


@dataclass
class File:
    """Holds a file"""
    size: int
    name: str


@dataclass
class Directory:
    """Holds a directory"""
    files: List[str]
    directories: List[str]
    size: int



if __name__ == "__main__":

    filename: str = "test_input.txt"

    directories: List[str] = []
    files: Set[str] = []


    with open(filename, encoding="utf-8") as f:
        for line in f.readlines():
            # print(command(line) if command(line) is None else "")
            # print(f"{file(line)}" if file(line) is None else "")
            # print(f"{directory(line)}" if directory(line) is None else "")
            cmd = command(line)
            dirct = directory(line)
            fle = file(line)

            if dirct:
                directories.append(dirct)

            if fle:
                files.append(fle[1])


        print(f"Directories: {directories}")
        print(f"Files: {files}")

