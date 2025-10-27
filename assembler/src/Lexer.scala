import scala.util.Try

object Lexer:
    def lexemes(source : String) : Option[Vector[String]] =
        def loop(s : Vector[Char], result : Vector[String], carry : String = "") : Option[Vector[String]] =
            val head : Option[Char] = s.headOption
            val tail : Option[Vector[Char]] = Try {s.tail}.toOption
            (head, tail) match {
                case (Some(head), Some(tail)) => {
                    println(head)
                    loop(tail, result)
                }
                case _ => Option(result)
            }

        loop(source.trim.toVector, Vector[String]())
