from sage.all import MatrixSpace, matrix, VectorSpace, vector, \
    is_prime, GF

CHARSET = "abcdefghijklmnopqrstuvwxyz01345678{}_"
p = len(CHARSET)
assert is_prime(p)
Fp = GF(p)

def double_char_to_vector(s):
    assert len(s) == 2
    return vector(Fp, [CHARSET.index(s[0]), CHARSET.index(s[1])])


def vector_to_double_char(v):
    assert len(v) == 2
    return CHARSET[v[0]] + CHARSET[v[1]]


def encrypt(msg, K):
    assert all(c in CHARSET for c in msg)
    assert len(msg) % 2 == 0

    A, v = K
    ciphertext = ""
    for i in range(0, len(msg), 2):
        tmp = A * double_char_to_vector(msg[i:i+2]) + v
        ciphertext += vector_to_double_char(tmp)

    return ciphertext


if __name__ == '__main__':
    from secret import flag
    assert flag.startswith('kmactf')
    A = MatrixSpace(Fp, 2, 2).random_element()
    v = VectorSpace(Fp, 2).random_element()
    assert A.determinant() != 0
    print('ciphertext =', repr(encrypt(flag, (A,v))))

    # Output:
    # ciphertext = 'u_rm3eefa}_7det1znb{sce{qo0h7yf0b}sktse8xtr6'
