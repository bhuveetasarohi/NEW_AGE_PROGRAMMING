
fun main() {

    // 1. Hello World
    println("Hello Kotlin!!!")

    // 2. Employee Bonus
    print("Enter department (IT/HR/SALES): ")
    val dept = readLine()!!.uppercase()

    print("Enter years of experience: ")
    val years = readLine()!!.toInt()

    print("Enter salary: ")
    val salary = readLine()!!.toDouble()

    val bonus = when {
        dept == "IT" && years > 5 -> 20
        dept == "HR" && years > 3 -> 15
        dept == "SALES" -> 12
        else -> 10
    }

    val amount = salary * bonus / 100

    println("Bonus Percentage: $bonus%")
    println("Bonus Amount: $amount")

    // 3. Average of Digits
    print("Enter an integer: ")
    val num = readLine()!!.toInt()

    var n = kotlin.math.abs(num)
    var sum = 0
    var count = 0

    while (n > 0) {
        sum += n % 10
        count++
        n /= 10
    }

    if (count == 0) count = 1

    println("Average of digits: ${sum.toDouble() / count}")
}
