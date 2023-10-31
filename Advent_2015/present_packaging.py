from typing import List
from Solver import Solver
from functools import reduce
from operator import mul


class PresetPackaging(Solver):
    
    def get_formatted_lines(self, input_lines:iter) ->iter:
        for package_lengths in input_lines:
            yield [int(length) for length in package_lengths.split('x')]
    
    def get_present_area(self, package_lengths: List[int]) -> int:
        package_lengths.sort()
        min_surface_area = package_lengths[0] * package_lengths[1]
        return (3 * min_surface_area) + (2 * package_lengths[0] * package_lengths[2]) + (2 * package_lengths[1] * package_lengths[2])

    def get_bow_length(self, package_lengths: List[int]) -> int:
        package_lengths.sort()
        ribon = 2 * package_lengths[0] + 2 * package_lengths[1]
        bow = reduce(mul, package_lengths)
        return ribon + bow
    
    def total_packaging(self, input_lines: iter) -> int:
        total_area = 0
        for package_lengths in self.get_formatted_lines(input_lines):
            total_area += self.get_present_area(package_lengths)
        return total_area
    
    def total_ribon(self, input_lines: iter) -> int:
        total_length = 0
        for package_lengths in self.get_formatted_lines(input_lines):
            total_length += self.get_bow_length(package_lengths)
        return total_length