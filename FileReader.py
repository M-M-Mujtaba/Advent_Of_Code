from abc import ABC, abstractmethod
from os import path

class EmptyFileError(Exception):
    pass

class FileReader(ABC):

    def __init__(self, file_name):
        self.file_name = file_name
        if path.getsize(file_name) == 0:
            raise EmptyFileError(f"The file {file_name} is empty.")
    @abstractmethod
    def read(self):
        pass

class FileFullReader(FileReader):

    def read(self):
        with open(self.file_name) as fp:
            return fp.read()

class FileLineReader(FileReader):

    def read(self):
        with open(self.file_name) as fp:
            for line in fp.readlines():
                yield line.strip()
