import scala.io.Source
import scala.util.Using
import java.io.File

object Main:
    @main
    def init(args: String*): Unit =
        args.length match
            case 1 => proccess(args(0))
            case _ => help

    def help: Unit =
        println("How to use me")
        println("assembler file.asm")

    def readFile(path: String): Option[String] =
      val file = new File(path)
      if !file.exists() then None
      else Using(Source.fromFile(file))(_.mkString).toOption

    def proccess(file: String): Unit =
        println(file)
        val source : Option[String] = readFile(file)
        source match {
            case Some(s) => Lexer.lexemes(s).filter(_!="").foreach(println)
            case _ => println("Failed to open file")
        }
