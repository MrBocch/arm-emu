import scala.io.Source

object Main {
    @main
    def init(args: String*): Unit =
        args.length match
            case 0 => help
            case 1 => proccess(args(0))
            case _ => println("multiple choices?")

    def help: Unit =
        println("How to use me")
        println("assembler file.asm")


    def proccess(file: String): Unit =
        // use Try, dont want to write any try {} catch {} code
        println(s"open contents of $file")

        // Option(file_string) match
        println("taking that, turning it into string")
        // do you remember crafting interpreters?
        println("feed that to a 'assembler'\n[ASSEMBLER]")
        println("- Lexer")
        println("- Tokenizer")
        println("- Serializer")
}
