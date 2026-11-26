
fun main() {

    val list = listOf(1, 2, 3)
    println(list)
    println(list.contains(2))

    val mList = mutableListOf(1, 2, 3)
    mList.add(4)
    mList.remove(1)
    println(mList)

    val set = setOf(1, 2, 3)
    println(set)

    val mSet = mutableSetOf(1, 2, 3)
    mSet.add(4)
    mSet.remove(1)
    println(mSet)

    val map = mapOf(1 to "Apple", 2 to "Mango")
    println(map)
    println(map[1])

    val mMap = mutableMapOf(1 to "Apple", 2 to "Mango")
    mMap[3] = "Orange"
    mMap.remove(2)
    println(mMap)

    val numbers = listOf(1, 2, 3, 4, 5)

    println(numbers.map { it * 2 })
    println(numbers.filter { it % 2 == 0 })
    println(numbers.reduce { a, b -> a + b })

    data class Student(val name: String, val marks: Int)

    val students = listOf(
        Student("Amit", 70),
        Student("Riya", 90),
        Student("Rahul", 80)
    )

    println(students.sortedBy { it.marks })
    println(students.groupBy {
        if (it.marks >= 75) "Pass" else "Fail"
    })
}
