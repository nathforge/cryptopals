import itertools

def single_byte(bs, xor_byte):
    bs = bytearray(bs)
    return bytearray(
        byte ^ xor_byte
        for byte in bs
    )

def repeating_bytes(bs, xor_bytes):
    bs = bytearray(bs)
    xor_bytes = bytearray(xor_bytes)
    return bytearray(
        byte ^ xor_byte
        for byte, xor_byte in itertools.islice(zip(
            bs, itertools.cycle(xor_bytes)
        ), len(bs))
    )
