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
        prev = input_string[0]
        pairs.add(input_string[0] + input_string[1])
        for idx, char in enumerate(input_string[1:-1]):
            nxt = input_string[idx + 2]
            if not (prev == char and char == nxt) and not pair_repeat:
                pair = char + nxt
                if not (pair_repeat := pair in pairs):
                    pairs.add(pair)
            in_between = prev == nxt if not in_between else in_between
            if pair_repeat and in_between:
                return True
            prev = char
        return False
                    
    

    def how_many_nice(self, input_lines: iter) -> int:
        nice_strings = 0
        
        for input_string in input_lines:
            if self.is_nice_v2(input_string):
                nice_strings += 1
                print(input_string)
            
        return nice_strings
        