import unittest

from cp import xor

class TestSingleByte(unittest.TestCase):
    def test(self):
        ba = bytearray([1, 2, 3, 4, 5])
        xor_byte = 77
        result = xor.single_byte(ba, xor_byte)
        self.assertEqual(result, bytearray([1 ^ 77, 2 ^ 77, 3 ^ 77, 4 ^ 77, 5 ^ 77]))

class TestRepeatingBytes(unittest.TestCase):
    def test(self):
        ba = bytearray([1, 2, 3, 4, 5])
        xor_bytes = bytearray([77, 88, 99])
        result = xor.repeating_bytes(ba, xor_bytes)
        self.assertEqual(result, bytearray([1 ^ 77, 2 ^ 88, 3 ^ 99, 4 ^ 77, 5 ^ 88]))

    def test_long_xor_bytes(self):
        ba = bytearray([1, 2, 3, 4, 5])
        xor_bytes = bytearray([11, 22, 33, 44, 55, 66, 77, 88])
        result = xor.repeating_bytes(ba, xor_bytes)
        self.assertEqual(result, bytearray([1 ^ 11, 2 ^ 22, 3 ^ 33, 4 ^ 44, 5 ^ 55]))
