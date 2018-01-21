import collections
import itertools

class Position(collections.namedtuple("Position", ("start", "end"))):
    def slice(self, iterable):
        return itertools.islice(iterable, self.start, self.end)

def positions(size, block_size):
    for start in range(0, size, block_size):
        yield Position(start, min(size, start + block_size))

def transpose(ba, block_size):
    for offset in range(block_size):
        yield bytearray(ba[i] for i in range(offset, len(ba), block_size))
