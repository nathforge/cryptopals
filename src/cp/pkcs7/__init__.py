def pad(ba, block_size):
    ba = bytearray(ba)
    ba.extend(get_padding(len(ba), block_size))
    return ba

def get_padding(size, block_size):
    pad_size = block_size - (size % block_size)
    if pad_size > 255:
        raise ValueError("Pad size is too large")
    return bytearray([pad_size] * pad_size)
