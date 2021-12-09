import scala.annotation.tailrec
import scala.io.Source.fromFile

object Main {
  case class BoardChance(combination: Set[Int], depth: Int = 0, aim: Int = 0)

  val InputPath = "src/main/4/input.txt"

  def main(args: Array[String]): Unit = {
    two()
  }

  // draws, (combination, boardIndex), boards
  def input(): (List[Int], List[(Set[Int], Int)], List[Set[Int]]) = {
    val rows: List[String] = fromFile(InputPath).mkString.split("\n\n").toList
    val draws = rows.head.split(",").map(_.toInt).toList
    val combinationsWithBoards = rows.tail.map(parseBoard)
    val boards = combinationsWithBoards.map(_._2)
    val combinations = combinationsWithBoards.zipWithIndex.flatMap(v => {
      val i = v._2
      val (combs, _) = v._1
      combs.map((_, i))
    })
    (draws, combinations, boards)
  }

  def parseBoard(rawBoard: String): (List[Set[Int]], Set[Int]) = {
    val rows = rawBoard.split("\n").toList
    val cleanedRows = rows.map(row => row.split(" +").filter(_.nonEmpty).map(_.trim))

    val rowCombinations: List[List[Int]] = cleanedRows.map(_.map(_.toInt).toList)
    // poor man's transpose
    val horizontalCombinations = {
      for (i <- rowCombinations.head.indices)
        yield rowCombinations.map(a => a(i))
    }.toList
    val combinations = (rowCombinations ++ horizontalCombinations).map(_.toSet)

    val board = combinations.fold(Set())((acc: Set[Int], curr: Set[Int]) => acc ++ curr)
    (combinations, board)
  }

  def bestCombination(draws: List[Int], combinations: List[(Set[Int], Int)]): (Int, Int) = {
    // (drawsRequired, index)
    var best = (draws.length, -1)
    for (comb <- combinations) {
      val current = combinationToFinish(draws, comb, best._1)
      if (current < best._1) {
        best = (current, comb._2)
      }
    }
    best
  }

  def combinationToFinish(draws: List[Int], combination: (Set[Int], Int), earlyStopAfter: Int = -1): Int = {
    val found = scala.collection.mutable.Set[Int]()
    var i = -1
    while (found.size < combination._1.size && (earlyStopAfter == -1 || i < earlyStopAfter)) {
      i += 1
      val draw = draws(i)
      if (combination._1.contains(draw)) {
        found.add(draw)
      }
    }
    i
  }

  def one(): Unit = {
    // all of the boards as set of ints in list
    // verify the all sets are equally big (no duplicates of ints in board)
    // list of tuple of sets with possible winning combination and index to
    // board in board list
    // calculate how many draws it would require for a given board to finish,
    //  - this is done by iterating through the drawn and construct a found set
    //  - when this set is == winning combination
    //  - save if fewest
    // do so for each board
    // if a given board has exceeded the fewest, stop counting that, it wont win
    // when a board has won
    // we can take the set of numbers that was drawn and get the disjunction
    //    between the whole board and the drawn

    val (draws, combinations, boards) = input()
    val bestComb = bestCombination(draws, combinations)
    val winningBoard = boards(bestComb._2)
    val drawnNumbers = draws.take(bestComb._1 + 1).toSet
    val rest = winningBoard.diff(drawnNumbers)
    val res = rest.sum * draws(bestComb._1)
    println(res)
  }

  def two(): Unit = {
    val (draws, combinations, boards) = input()

    val worstBestBoardComb = combinations.groupBy(comb => comb._2).map(boardToCombs =>
      bestCombination(draws, boardToCombs._2)
    ).maxBy(comb => comb._1)

    val worstBoard = boards(worstBestBoardComb._2)
    val drawnNumbers = draws.take(worstBestBoardComb._1 + 1).toSet
    val rest = worstBoard.diff(drawnNumbers)
    val res = rest.sum * draws(worstBestBoardComb._1)
    println(res)

  }
}
