/*
# date 2019-08-08 14:31:13
# author calllivecn <calllivecn@outlook.com>
*/


package zxlib

func Clearbuf(buf []byte) (int, error) {
    //block := make([]byte, 1024)

    empty := byte(0)

    l := len(buf)

    var i int = 0

    for ;i<l; i++ {

        if buf[i] == empty {
            break
        }

        buf[i] = empty
    }

    return i, nil

}
