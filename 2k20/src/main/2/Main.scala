
object Main {
  def main(args: Array[String]): Unit = {
    two()
  }
  def one(): Unit = {
    val lines = scala.io.Source.fromFile("input.txt").mkString
    val input = lines.split("\n").map(row => row.toInt).sorted
    var leftPtr = 0
    var rightPtr = input.length - 1

    while (input(leftPtr) + input(rightPtr) != 2020) {
      if (input(leftPtr) + input(rightPtr) < 2020) {
        leftPtr += 1
      } else if (input(leftPtr) + input(rightPtr) > 2020) {
        rightPtr -= 1
      }
    }



    println(input(leftPtr) * input(rightPtr))
  }

  def two(): Unit = {
    val lines = scala.io.Source.fromFile("input.txt").mkString
    val input = lines.split("\n").map(row => row.toInt).sorted

    for (p1 <- input; p2 <- input; p3 <- input) {
      if (p1 + p2 + p3 == 2020){
        println(p1 * p2 * p3)
      }
    }



  }
}