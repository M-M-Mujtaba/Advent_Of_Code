import multiprocessing
from hashlib import md5
import time

class MD5:

    @staticmethod
    def get_md5(data: bytes) -> str:
        return md5(data).hexdigest()

    def worker(self, input_bytes, target, number_range, result):
        start, end = number_range
        for number in range(start, end):
            if result.value != -1:
                return  # Exit if a solution has been found
            check_bytes = input_bytes + str(number).encode()
            if self.get_md5(check_bytes)[0:6] == target:
                with result.get_lock():
                    if result.value == -1:  # Double-check the result hasn't been set by another process
                        result.value = number
                return  # Exit after finding a solution

if __name__ == "__main__":
    solver = MD5()
    input_bytes = b"iwrupvqb"
    target = '0' * 6
    num_processes = 4
    step = 2500000  # 2.5 million

    # Calculate the ranges for each process
    ranges = [(i * step, (i + 1) * step) for i in range(num_processes)]

    start_time = time.time()

    with multiprocessing.Manager() as manager:
        result = manager.Value('i', -1)  # Shared result variable
        with multiprocessing.Pool(processes=num_processes) as pool:
            pool.starmap(solver.worker, [(input_bytes, target, number_range, result) for number_range in ranges])

    end_time = time.time()

    elapsed_time = end_time - start_time
    print(f"Result: {result.value}")
    print(f"Time taken: {elapsed_time:.4f} seconds")
