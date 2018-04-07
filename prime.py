#!/usr/bin/env python3
#coding=utf-8


import math

import numba

@numba.jit(nopython=True)
def prime(n):
    if n <= 2 :
        return True
    
    for i in range(3,int(math.sqrt(n))):
    #for i in range(3,n):
        if 0 == n % i:
            return False

    return True

@numba.jit(nopython=True)
def main(m):
    for i in range(1,m+1):
        if prime(i):
            pass
            #print(i,"是素数")



if __name__ == "__main__":
    main(500000)

