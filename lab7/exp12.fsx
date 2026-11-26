let list1 = [1; 2; 3; 4; 5]
let list2 = [2; 4; 6; 8]

let square = List.map (fun x -> x * x)
let even = List.filter (fun x -> x % 2 = 0)
let sum = List.fold (fun total x -> total + x) 0

printfn "MAP"
printfn "%A" (square list1)
printfn "%A" (square list2)

printfn "\nFILTER"
printfn "%A" (even list1)
printfn "%A" (even list2)

printfn "\nFOLD"
printfn "%d" (sum list1)
printfn "%d" (sum list2)