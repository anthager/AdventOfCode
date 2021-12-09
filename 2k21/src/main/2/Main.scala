import scala.io.Source.fromFile

object Main {
  case class Position(horizontal: Int = 0, depth: Int = 0, aim: Int = 0)

  val InputPath = "src/main/2/input.txt"

  def main(args: Array[String]): Unit = {
    one()
  }

  def one(): Unit = {
    val rows = fromFile(InputPath).mkString.split("\n").map(parseInputRow).toList

    val endPosition = rows.foldLeft(Position())((state: Position, action: (Int, Int)) => {
      move(state, action)
    })

    println(endPosition.depth * endPosition.horizontal)
  }

  def two(): Unit = {
    val rows = fromFile(InputPath).mkString.split("\n").map(parseInputRow).toList

    val endPosition = rows.foldLeft(Position())((state: Position, action: (Int, Int)) => {
      move(state, action)
    })

    println(endPosition.depth * endPosition.horizontal)
  }

  // returns forward, down
  def parseInputRow(row: String): (Int, Int) =
    row.split(" ").toList match {
      case List("down", value) => (0, value.toInt)
      case List("up", value) => (0, -value.toInt)
      case List("forward", value) => (value.toInt, 0)
      case a => sys.error(a.toString)
    }


  def move(state: Position, action: (Int, Int)): Position = {
    val (forward, down) = action
    val horizontalChange = forward
    val depthChange = forward * state.aim
    val aimChange = down

    val horizontalNew = horizontalChange + state.horizontal
    val depthNew = depthChange + state.depth
    val aimNew = aimChange + state.aim

    Position(horizontalNew, depthNew, aimNew)
  }
}