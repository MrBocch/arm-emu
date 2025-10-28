import scala.util.Try

object Lexer:
    // need to handle comments
    // also handle "strings"
    def lexemes(source : String) : Vector[String] =
        def loop(s : Vector[Char], result : Vector[String], carry : String = "", comment : Boolean = false) : Vector[String] =
            val head : Option[Char] = s.headOption
            val tail : Option[Vector[Char]] = Try {s.tail}.toOption
            (head, tail) match {
                case (Some(head), Some(tail)) => {
                    println(s"Head: $head")
                    head match {
                        case ' ' | '\n' | '\t' => loop(tail, result.appended(carry), "")
                        case ',' | ':' | '[' | ']' | '-' | '+' | '#' => loop(tail, result.appended(carry).appended(head.toString), "")
                        case _ => loop(tail, result, carry+head.toString)
                    }
                }
                case (None, Some(tail)) => {
                    println("head is none")
                    println(s"result: $result carry: $carry")
                    result
                }
                case _ => result.appended(carry.toString)
            }

        loop(source.trim.toVector, Vector[String]())
