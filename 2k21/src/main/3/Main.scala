import scala.annotation.tailrec
import scala.io.Source.fromFile

object Main {
  case class Position(horizontal: Int = 0, depth: Int = 0, aim: Int = 0)

  val InputPath = "src/main/3/input.txt"

  def main(args: Array[String]): Unit = {
    two()
  }

  def one(): Unit = {
    val rows: List[List[Char]] = fromFile(InputPath).mkString.split("\n").map(_.toList).toList
    val epsilon = getEpsilon(rows)
    val epsilonDec = BigInt(toDecimal(epsilon))
    val gamma = getGamma(rows)
    val gammaDec = BigInt(toDecimal(gamma))
    println(epsilonDec * gammaDec)
  }

  def two(): Unit = {
    // go through all to find most common
    // then filter on it
    // take is oxy param to determine both common/uncommon filter and equal
    val rows: List[List[String]] = fromFile(InputPath).mkString.split("\n").map(_.toList.map(_.toString)).toList
    val oxy = calculateValue(rows, isOxy = true)
    val co2 = calculateValue(rows, isOxy = false)
    println(oxy * co2)
  }

  def calculateValue(numbers: List[List[String]], isOxy: Boolean): Int = {
    val binNum: String = _calculateValue(numbers, 0, isOxy).mkString
    toDecimal(binNum)
  }

  @tailrec
  def _calculateValue(numbers: List[List[String]], i: Int, isOxy: Boolean): List[String] = {
    numbers.length match {
      case 0 => sys.error("length 0 reached")
      case 1 => numbers.head
      case _ => {
        val mostCommon = clipInt(findMostCommonForBit(numbers, i))
        val newNumbers = numbers.filter(num => {
          (num(i), mostCommon) match {
            case ("0", -1) => isOxy
            case ("1", 1) => isOxy
            case ("0", 0) => !isOxy
            case ("1", 0) => isOxy
            case ("0", 1) => !isOxy
            case ("1", -1) => !isOxy
            case (_, _) => sys.error("input doesnt match any")
          }
        })
        _calculateValue(newNumbers, i + 1, isOxy)
      }
    }

  }

  def clipInt(num: Int): Int = num match {
    case n if n < 0 => -1
    case n if n == 0 => 0
    case n if n > 0 => 1
  }

  def filter(numbers: List[List[String]], index: Int, bit: String, oneOnEqual: Boolean): List[List[String]] = {
    numbers.filter(num => num(index) == bit)
  }

  // will return < 0, 0, or 0 < for 0 most common, 0 and 1 equally
  // common and 1 most common respectively
  def findMostCommonForBit(numbers: List[List[String]], i: Int): Int =
    numbers.foldLeft(0)((acc: Int, chars: List[String]) => acc + (if (chars(i) == "1") 1 else -1))

  def getEpsilon(numbers: List[List[Char]]): String = {
    numbers.head.indices.map(i => {
      val charSummer = numbers.foldLeft(0)((acc: Int, chars: List[Char]) => acc + (if (chars(i) == '1') 1 else -1))
      if (charSummer < 0) "0" else "1"
    })
  }.mkString

  def getGamma(numbers: List[List[Char]]): String =
    getEpsilon(numbers).map(n => if (n == '0') '1' else '0')

  def toDecimal(bin: String): Int = {
    val indices = bin.indices.reverse
    bin.zip(indices).map(charAndIndex => {
      val binNum = charAndIndex._1.toString.toInt
      binNum * Math.pow(2, charAndIndex._2)
    }).sum.toInt
  }
}
