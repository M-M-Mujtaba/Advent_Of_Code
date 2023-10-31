from Solver import Solver


class Coords:
    x,y = 0, 0
    def __repr__(self):
        return str(self.x) + str(self.y)

class House(Solver):
    
    visited_houses = set()
    total_unique_houses_visited = 0
    santa_coords = Coords()
    robot_coords = Coords()

    def move_up(self, coords: Coords):
        coords.y+=1
    def move_down(self, coords: Coords):
        coords.y-=1
    def move_left(self, coords: Coords):
        coords.x-=1
    def move_right(self, coords: Coords):
        coords.x+=1
    direction_to_coords = {'>': move_right, '<': move_left, '^': move_up, 'v': move_down}

    def add_to_visited_houses(self, coords: Coords):
        if str(coords) not in self.visited_houses:
            self.total_unique_houses_visited +=1
            self.visited_houses.add(str(coords))
    def move(self, char, coords: Coords):
        self.add_to_visited_houses(coords)
        self.direction_to_coords[char](self, coords)

    
    def unique_houses_visited(self, input_str: str) -> int:
        for char in input_str:
            self.move(char, self.santa_coords)
        return self.total_unique_houses_visited
    def unique_houses_visited_with_robot(self, input_str: str) -> int:
        santa_move = True
        for char in input_str:
            if santa_move:
                self.move(char, self.santa_coords)
            else:
                self.move(char, self.robot_coords)
            santa_move = not santa_move
        return self.total_unique_houses_visited