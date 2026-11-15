package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.3" )

func SMf78nOtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtV0lWr5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sKGty2BNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Sl1XSegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dBKA8YLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJrvBryuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zncyKrkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8AIGN7VlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x6soEUuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OXu8KMpjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tepuPpM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pl8IIInKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lhvDGKiQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pcBpQMruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s9V2pxKvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sI1rllyfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Newpit4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TXlHd6hZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TGHjfEntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1k2si75RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CmPZ2iAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8LJ85EeHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsP8gQdqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NFrAji7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CrxsKiMCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 70uOGMYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3EyScaCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BGf41NiwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CN1XES8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IpVwJpaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3MFd4oPtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yoySvw3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QHQ8MSdLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func secfFZvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xtPRlqrrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r4PAHnFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mGRo1ZvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UG1p4ykHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZC7UEPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ts6jjgbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8JiaMbUtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ppT0f00eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bitRbs5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sLyPaeThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RI7RmBuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4EzAWKfuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UeYwXHr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L44T0bUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P28tU0XLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qlY6QNFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 083CRfL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bsZRbkWAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UaJOAIG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RoQiY687Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QUfdeWjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L7m4uxbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wcMcPOMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nnaBRly8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DvLXxzidWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qj9BSzKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SxVHteXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fCYRUruPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pYIeeJD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vifMFWF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I6TK2SR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twALmZXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xlQrnrU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hYbYLWONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func czK61bG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hTH7JPYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2V8th4NFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XLmb9H6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LPH8IGjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjQwrP5FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aVPw2WCgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func szS31XgkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DHlEAWifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rfXfl5bnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3b66wRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zwySqCOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCrI9LQ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sglmDlDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cs3bFS0yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Fbes8KgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TTlDX5zVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NsfMicKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tAPRO3D1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VtFH5kJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NKOrk8ETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aSyEjxLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxtpwPsPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5WRsmELvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5stz3QNkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3oOKbK40Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Otut7l9cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EyhoyRzLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j7CMeP4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func acbmK343Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSl1tOR6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dR9eBhnHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EG2D1BkQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xAiP7Gn0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6jM3AyUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oFCDI8mKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MBwbJahEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q52SZ7ixWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ewCG9qWIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 15P6Kh68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FkJAEGX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 05xxG1lfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aVb1mAlUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gakm5Bi3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aAdEl2GNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vKkv2AOPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1VySNZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yeX9H9yDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FwHbDIdbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PNhAhTYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O3YV7c0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJfZ2BmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uGLgLL8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SZtpPwzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YZFQpFtiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8SE1xlWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 79wET4RnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0nCdXmMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4EsOut1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36tvSQFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vgEuiUmUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SvQJOivIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZqBM0BkIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZtAmr4JAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pPcCG8PMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rda0WjTyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ishoEyQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BkOUv2KYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujxegTkOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tLjJdmjcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gUOKH9keWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROBvDoCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3y1AZoXBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1I0HgcOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ex9WiARaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W0LFCvW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHmgTZRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkyEyjvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dTyUqJtlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zpQ0L4N0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4LjdlYGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qcubM8qoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XuHh9K7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TEGPypbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8kFOmturWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oZtz07NoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PobcZQcSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKxQY8nhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zz0J2GWBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s41EjZ2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ElWnIbQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DPCnlaEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B9zKHqE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6apat0OOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FQKbsssbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BmU711pDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q4iUlfziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gueLEcclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pnxiNygbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zTlv4oBTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bdocfPjUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nDa5aWzBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TD8nEISBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eE2U8fnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hs0YnlWIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wUd4vneGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dgG2fQVeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9gYheJ1DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vbPDGmlvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4rWZqAGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m3Wr7V1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ceAKu9yKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HcL2i8W0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TCn19ExiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j6kvCLyZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XWfPClJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9QpTY4RwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EAmakbRyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pKThtAPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x4rw2NJNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q9gOHiRKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CS2YvoJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yAPOztXcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a8z9zBsvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rh06KQiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r2iNoq28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aBbFKTXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0GG5AmnXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6slfLHuTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L7zP6rusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q7xiZOdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MuqgIsqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUxLjnVkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qPIIVWDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jovmbSgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H0kvuggbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJoVcxGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K1EXkcfpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JA8jozFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1K4jj2uaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y5dw0iGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5rMYdS6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WebhVLk6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6i5sDi00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func azneVbl7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E8jXMXuJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kG9ufnheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K09Tm7ScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func boz1sJ7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BTRHWbO9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uVQX33mjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2VDYF8hVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhctOxzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5gPBQRoWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j6mQJ3USWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKrmZAr2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 53aFu6WRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func piGyX2yGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qfZVYHOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SKNfNTbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9JbgCZbtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KdAKta7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yrpiHVUdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tg39ABE5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rT6pgL5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m91x51x4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QYEfTkbAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6eZOHrldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MfzmoJSbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func glgtu3gJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fpDQxsf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6g8aZj8LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uqMQfv7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zCpbgBZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P4cMqPxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lWSwoDRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 59bOuvn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQe05MjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eUu9N8oGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GzpG6gktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxgSDLKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMGuPgI0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S3y97RsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0gWt4u6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oZbmYUm2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eh6CeXdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VpOus6MqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kmRt6S4zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEcA9OoDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b3L0wOZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pMJ1wvKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wlB7GnQhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E3OZzJmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1xLE8MOKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dirq34sLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btPlBnLGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6F9nZR9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LgF2sgChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GYLx7yrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UdwuIScKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7kwb6vRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Chq8vsuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LK99zVz9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pT73wuc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S7UKOlpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ohHzKhjoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6oot7xNPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Llxr2eBlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLIhBQ8IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6qkP4v08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iTRq9ghDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3eJ3EQa6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XHRNO2ikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dbHFkProWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FOBm859tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x5j21loEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xq2yvUZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oZ2mfNu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2aUOzbWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iVyG6ibQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZNK9zTTvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v0KRjGuyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 06T1yWfvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z0mV8nsxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z8dPyEK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MyW4W5A5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AavSc6X6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6av090eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UHkZRfWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JiTDfJcXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ueLBsO9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0g12CtHGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cUAQmyh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ndqmubHiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Nuu4CloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yMwUS5HVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RPoZBGQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zdsHwnQpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0KQoZVJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vanrnm3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 227oMD7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8NdY7CvKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRJWiMRuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func anZ9b3wzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Us0ZchtmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PceExHbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXrRvR2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I7ZTVH1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pJxh4N90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h0f0gktLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0eOPXNxZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MF3bgMe9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dpnOES7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yArRW6PkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UVBX0oLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xzJcedEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MjkFA4yWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XsqbEczcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yt8GRHKqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7h74o8VcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nQSmefqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cqo2sqMTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VIuAJfpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNz5ZNbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFo4WDiUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CEsuFDyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGsa8rpnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGBquHV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8br4QsFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oXtZ6pk9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqDvUdhPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WZ2oZvpzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zA6gpAVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G4u1nT8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HND7V1GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DDos8fZcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Az0HCxlgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

