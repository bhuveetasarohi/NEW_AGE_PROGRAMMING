
fun main() {

    var name: String = "Amit"
    var city: String? = null

    println(name)
    println(city)

    city = "Delhi"
    println(city?.uppercase())

    city = null
    println(city?.uppercase())

    var username: String? = null
    println(username ?: "Guest")

    try {
        println(city!!.length)
    } catch (e: NullPointerException) {
        println("Null pointer exception caught")
    }
}
