open System

let reverse (s: string) =
    new string(Array.rev (s.ToCharArray()))

let palindrome (s: string) =
    if s = reverse s then
        printfn "%s is a palindrome" s
    else
        printfn "%s is not a palindrome" s

let transform (s: string) =
    if s <> "" && Seq.forall Char.IsLetter s then
        s.ToUpper()
    elif s <> "" && Seq.forall Char.IsDigit s then
        s + s
    else
        s

let email (s: string) =
    if System.Text.RegularExpressions.Regex.IsMatch(s, @"^[^@\s]+@[^@\s]+\.[^@\s]+$") then
        printfn "%s is valid" s
    else
        printfn "%s is invalid" s

palindrome "madam"
palindrome "hello"
palindrome "level"

printfn "%s" (transform "hello")
printfn "%s" (transform "123")
printfn "%s" (transform "abc123")

email "abc@gmail.com"
email "abc.com"
email "abc@yahoo.com"