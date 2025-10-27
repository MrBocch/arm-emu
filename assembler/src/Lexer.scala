import scala.util.Try

object Lexer:
    def lexemes(source : String) : Vector[String] =
        def loop(s : Vector[Char], result : Vector[String], carry : String = "") : Vector[String] =
            val head : Option[Char] = s.headOption
            val tail : Option[Vector[Char]] = Try {s.tail}.toOption
            (head, tail) match {
                case (Some(head), Some(tail)) => {
                    head match {
                        case ' ' | '\n' => loop(tail, result.appended(carry), "")
                        case ',' | ':' => loop(tail, result.appended(carry).appended(head.toString), "")
                        case _ => loop(tail, result, carry+head.toString)
                    }
                }
                case (None, Some(tail)) => {
                    println("head is none")
                    println(s"result: $result carry: $carry")
                    result
                }
                case _ => {
                    println("gest to this case")
                    result
                }
            }

        loop(source.trim.toVector, Vector[String]())
