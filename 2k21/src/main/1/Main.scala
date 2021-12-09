import scala.io.Source.fromFile

object Main {
  val InputPath = "src/main/1/input.txt"

  def main(args: Array[String]): Unit = {
    two()
  }

  def one(): Unit = {
    val depths = fromFile(InputPath).mkString.split("\n").map(_.toInt).toList

    val res = {
      for (i <- 1 until depths.length)
        yield if (depths(i - 1) < depths(i)) 1 else 0
    }.sum

    println(res)
  }

  def slidingWindowSum(depths: List[Int], index: Int): Int = {
    if (index < 2)
      sys.error(s"index too low $index")

    {
      (index - 2) to index
    }.map(depths(_)).sum
  }

  def two(): Unit = {

    val depths = fromFile(InputPath).mkString.split("\n").map(_.toInt).toList

    val res = {
      for (i <- 3 until depths.length)
        yield if (slidingWindowSum(depths, i - 1) < slidingWindowSum(depths, i)) 1 else 0
    }.sum

    println(res)
  }
}