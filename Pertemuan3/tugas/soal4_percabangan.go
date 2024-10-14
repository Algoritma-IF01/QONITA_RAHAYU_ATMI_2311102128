// Sebelum di perbaiki

// package main

// import "fmt"

// func main() {
//     var nam float64
//     var nmk string
//     fmt.Print("Nilai akhir mata kuliah: ")
//     fmt.Scanln(&nam)
    
//     if nam > 80 {
//         nam = "A"
//     }
//     if nam > 72.5 {
//         nam = "AB"
//     }
//     if nam > 65 {
//         nam = "B"
//     }
//     if nam > 57.5 {
//         nam = "BC"
//     }
//     if nam > 50 {
//         nam = "C"
//     }
//     if nam > 40 {
//         nam = "D"
//     } else if nam <= 40 {
//         nam = "E"
//     }
    
//     fmt.Println("Nilai mata kuliah: ", nmk)
// }

// Sesudah di Perbaiki

package main

import "fmt"

func main() {
    var nam float64
    var nmk string
    fmt.Print("Nilai akhir mata kuliah: ")
    fmt.Scanln(&nam)

    if nam > 80 {
        nmk = "A"
    } else if nam > 72.5 {
        nmk = "AB"
    } else if nam > 65 {
        nmk = "B"
    } else if nam > 57.5 {
        nmk = "BC"
    } else if nam > 50 {
        nmk = "C"
    } else if nam > 40 {
        nmk = "D"
    } else {
        nmk = "E"
    }

    fmt.Println("Nilai mata kuliah: ", nmk)
}
