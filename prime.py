#!/usr/bin/env python3
#coding=utf-8


import math

import numba

@numba.jit(nopython=True)
def prime(n):
    if n == 1 :
        return True
    
    for i in range(2,int(math.sqrt(n)) + 1):
        if 0 == n % i:
            return False

    return True

@numba.jit(nopython=True)
def main(m):
    for i in range(1,m+1):
        if prime(i):
            #pass
            print(i,"是素数")



if __name__ == "__main__":
    main(500)

