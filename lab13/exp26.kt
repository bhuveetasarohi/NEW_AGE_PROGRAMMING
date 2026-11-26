
// Function with default parameters
fun student(name: String, age: Int = 18, city: String = "Delhi") {
    println("Name: $name, Age: $age, City: $city")
}

// Function with vararg
fun numbers(vararg n: Int) {
    println("Numbers:")
    for (i in n) {
        println(i)
    }
}

// Higher-order function
fun calculate(list: List<Int>, operation: (Int) -> Int) {
    for (i in list) {
        println(operation(i))
    }
}

fun main() {

    // 1. Default parameters
    student("Aman")
    student("Riya", 21, "Mumbai")

    // 2. Named arguments (out of order)
    student(name = "Rahul", city = "Pune", age = 22)

    // 3. Vararg
    numbers(10, 20, 30, 40, 50)

    // 4. Higher-order function
    val list = listOf(1, 2, 3, 4, 5)

    println("Squares:")
    calculate(list) { it * it }
}
