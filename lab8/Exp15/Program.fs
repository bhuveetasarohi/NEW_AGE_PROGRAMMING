
open Microsoft.Data.Sqlite

let con = new SqliteConnection("Data Source=student.db")
con.Open()

let execute sql =
    let cmd = con.CreateCommand()
    cmd.CommandText <- sql
    cmd.ExecuteNonQuery() |> ignore

let read() =
    let cmd = con.CreateCommand()
    cmd.CommandText <- "SELECT * FROM Student"
    let reader = cmd.ExecuteReader()

    while reader.Read() do
        printfn "%d %s" (reader.GetInt32(0)) (reader.GetString(1))

execute "CREATE TABLE IF NOT EXISTS Student (Id INTEGER PRIMARY KEY, Name TEXT)"

execute "INSERT OR REPLACE INTO Student VALUES (1, 'Rahul')"
printfn "After Create:"
read()

execute "UPDATE Student SET Name='Aman' WHERE Id=1"
printfn "After Update:"
read()

execute "DELETE FROM Student WHERE Id=1"
printfn "After Delete:"
read()

con.Close()
