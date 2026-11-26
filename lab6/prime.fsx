let rec check n i =
    if n < 2 then
        false
    elif i * i > n then
        true
    elif n % i = 0 then
        false
    else
        check n (i + 1)

printf "Enter a number: "
let n = int (System.Console.ReadLine())

if check n 2 then
    printfn "Prime Number"
else
    printfn "Not a Prime Number"