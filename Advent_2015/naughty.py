from Solver import Solver



class Naughty(Solver):

    vowels = {
        'a', 'e' , 'i', 'o', 'u'
    }
    bad_pairs = {
        'ab', 'cd', 'pq', 'xy'
    }
    
    def is_nice(self, input_string: str) -> bool:
        three_vowels = 0
        repeat = False
        bad_pair = False
        prev = input_string[0]
        three_vowels += int(prev in self.vowels)
        for char in input_string[1:]:
            if three_vowels < 3:
                three_vowels += int(char in self.vowels)
            if not repeat:
                repeat = char == prev
            if not bad_pair:
                if bad_pair := (prev + char in self.bad_pairs):
                    break
            prev = char
        return not bad_pair and  three_vowels == 3 and repeat

    def is_nice_v2(self, input_string: str) -> bool:
        pair_repeat = False
        pairs = set()
        in_between = False
        idx = 1
        while idx < len(input_string)-1:
            pair = input_string[idx - 1] + input_string[idx]
            if not ((input_string[idx - 1] == input_string[idx]) and (input_string[idx] == input_string[idx + 1])):
                if not (pair_repeat := pair  in pairs):
                    pairs.add(pair)
            in_between = input_string[idx - 1] == input_string[idx + 1]
            if in_between and pair_repeat:
                return True
            idx += 1
        return in_between and input_string[idx - 1] + input_string[idx] in pairs


    def how_many_nice(self, input_lines: iter) -> int:
        nice_strings = 0
        
        for input_string in input_lines:
            nice_strings += int(self.is_nice_v2(input_string))
        return nice_strings
        