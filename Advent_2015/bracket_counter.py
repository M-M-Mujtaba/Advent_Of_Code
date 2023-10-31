from Solver import Solver



class BracketCounter(Solver):

    bracket_counter = {'(': 1 ,')':-1}



    def count_final_floor(self, input_str: str) -> int:
        total = 0
        for char in input_str:
            total += self.bracket_counter[char]
        return total

    def count_basement_1(self, input_str: str) -> int:
        total = 0
        for i, char in enumerate(input_str):
            total += self.bracket_counter[char]
            if total == -1:
                return i + 1