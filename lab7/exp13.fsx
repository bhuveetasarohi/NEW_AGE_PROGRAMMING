open System

let task1 = async {
    printfn "Task 1 Started"
    do! Async.Sleep(3000)
    printfn "Task 1 Completed"
    return "Task 1 Result"
}

let task2 = async {
    printfn "Task 2 Started"
    do! Async.Sleep(2000)
    printfn "Task 2 Completed"
    return "Task 2 Result"
}

let task3 = async {
    printfn "Task 3 Started"
    do! Async.Sleep(1000)
    printfn "Task 3 Completed"
    return "Task 3 Result"
}

let tasks = [task1; task2; task3]

let results = Async.Parallel tasks |> Async.RunSynchronously

printfn "\nAll Tasks Completed"

for result in results do
    printfn "%s" result

    