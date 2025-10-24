object Lexer:

    def lexemes(source : Option[String]) : Vector[String] =
        def loop(source : String, carry : Vector[String]) : Vector[String] =
            Vector("hi")

        source match:
            case Some(s) => loop(s)
            case _ => None
