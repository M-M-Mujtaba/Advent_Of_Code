from FileReader import FileReader



class Solver():

    def solve(self, type: str, file_name: str, file_reader: FileReader ) -> int:
        method = getattr(self, type, None)
        if method is not None and callable(method):
            return method(file_reader(file_name).read())
        else:
            raise ValueError(f"No method {type} found in {self.__class__.__name__}")