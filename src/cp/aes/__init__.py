import oscrypto.symmetric

from cp import blocks

# OSCrypto doesn't provide ECB mode (for obvious reasons),
# so we'll simulate by encrypting blocks in isolation with
# a null IV.

def encrypt_ecb(key, data):
    key = bytes(key)
    null_iv = bytes([0] * 16)
    result = bytearray()
    for position in blocks.positions(len(data), 16):
        block = bytes(position.slice(data))
        _, enc_block = oscrypto.symmetric.aes_cbc_no_padding_encrypt(key, block, null_iv)
        result.extend(enc_block)
    return result

def decrypt_ecb(key, data):
    key = bytes(key)
    null_iv = bytes([0] * 16)
    result = bytearray()
    for position in blocks.positions(len(data), 16):
        block = bytes(position.slice(data))
        dec_block = oscrypto.symmetric.aes_cbc_no_padding_decrypt(key, block, null_iv)
        result.extend(dec_block)
    return result
