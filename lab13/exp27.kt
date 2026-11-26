
class Student(
    private var rollNo: Int,
    private var name: String,
    private var course: String,
    private var percentage: Double
) {

    // Function to accept student data
    fun acceptData() {
        print("Enter Roll Number: ")
        rollNo = readLine()!!.toInt()

        print("Enter Name: ")
        name = readLine()!!

        print("Enter Course: ")
        course = readLine()!!

        print("Enter Percentage: ")
        percentage = readLine()!!.toDouble()
    }

    // Function to display student data
    fun displayData() {
        println("\n----- Student Details -----")
        println("Roll Number: $rollNo")
        println("Name: $name")
        println("Course: $course")
        println("Percentage: $percentage%")
    }
}

fun main() {

    // Creating first object
    val s1 = Student(1, "Aman", "CSE", 85.0)
    println("Enter details of Student 1:")
    s1.acceptData()
    s1.displayData()

    // Creating second object
    val s2 = Student(2, "Riya", "IT", 90.0)
    println("\nEnter details of Student 2:")
    s2.acceptData()
    s2.displayData()
}
