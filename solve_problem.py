from Advent_2015.bracket_counter import BracketCounter
from Advent_2015.present_packaging import PresetPackaging
from Advent_2015.house import House
from Advent_2015.md500 import MD5
from Advent_2015.naughty import Naughty
from FileReader import FileFullReader, FileLineReader, EmptyFileError



if __name__ == '__main__':
    problem = Naughty()
    try:
        print(problem.solve('how_many_nice', 'Advent_2015/naughty_data.txt',FileLineReader ))
    except EmptyFileError as exception:
        print(exception)