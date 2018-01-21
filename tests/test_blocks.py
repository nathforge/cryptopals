import unittest

from cp import blocks

class TestPositions(unittest.TestCase):
    def test(self):
        self.assertEqual(list(blocks.positions(64, 16)), [
            blocks.Position(start=0, end=16),
            blocks.Position(start=16, end=32),
            blocks.Position(start=32, end=48),
            blocks.Position(start=48, end=64)
        ])

class TestTranspose(unittest.TestCase):
    def test(self):
        # Transp block: 0  1  2  0  1  2  0  1  2  0
        ba = bytearray([0, 1, 2, 3, 4, 5, 6, 7, 8, 9])
        self.assertEqual(list(blocks.transpose(ba, 3)), [
            bytearray([0, 3, 6, 9]),
            bytearray([1, 4, 7]),
            bytearray([2, 5, 8])
        ])
