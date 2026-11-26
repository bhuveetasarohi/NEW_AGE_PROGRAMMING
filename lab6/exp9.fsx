printf "Enter marks of subject 1: "
let a = int (System.Console.ReadLine())

printf "Enter marks of subject 2: "
let b = int (System.Console.ReadLine())

printf "Enter marks of subject 3: "
let c = int (System.Console.ReadLine())

printf "Enter marks of subject 4: "
let d = int (System.Console.ReadLine())

let per = float (a + b + c + d) / 4.0

let result =
    match per with
    | x when x >= 60.0 -> "First Division"
    | x when x >= 50.0 -> "Second Division"
    | x when x >= 40.0 -> "Third Division"
    | _ -> "Fail"

printfn "Percentage = %.2f" per
printfn "Division = %s" result