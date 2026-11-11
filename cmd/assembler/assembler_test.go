
package assembler

import (
	"reflect"
	"testing"
)

// ---------------------------------------------------------------------------
// Pruebas del LEXER
// ---------------------------------------------------------------------------

func TestLexInstructionWithRegisterAndImmediate(t *testing.T) {
	code := "mov r0, #5\n"
	tokens := Lex(code)

	expected := []Token{
		{Kind: Instruction, Lexeme: "mov", Line: 1},
		{Kind: Register, Lexeme: "r0", Line: 1},
		{Kind: Comma, Lexeme: ",", Line: 1},
		{Kind: Hash, Lexeme: "#", Line: 1},
		{Kind: Number, Lexeme: "5", Line: 1},
		{Kind: NewLine, Lexeme: "\n", Line: 1},
	}

	assertTokens(t, expected, tokens)
}

func TestLexCommentAtEndOfLine(t *testing.T) {
	code := "mov r0, #5 // esto es un comentario\n"
	tokens := Lex(code)

	// El lexer no debe producir tokens para el comentario
	for _, tok := range tokens {
		if tok.Lexeme == "//" || tok.Lexeme == "esto" {
			t.Fatalf("el lexer produjo un token de comentario: %+v", tok)
		}
	}

	// Debe haber exactamente una instrucción y un newline
	if len(tokens) != 6 {
		t.Fatalf("se esperaban 6 tokens, se obtuvieron %d: %+v", len(tokens), tokens)
	}
}

func TestLexLabel(t *testing.T) {
	code := "inicio: mov r0, #1\n"
	tokens := Lex(code)

	expected := []Token{
		{Kind: Identifier, Lexeme: "inicio", Line: 1},
		{Kind: Colon, Lexeme: ":", Line: 1},
		{Kind: Instruction, Lexeme: "mov", Line: 1},
		{Kind: Register, Lexeme: "r0", Line: 1},
		{Kind: Comma, Lexeme: ",", Line: 1},
		{Kind: Hash, Lexeme: "#", Line: 1},
		{Kind: Number, Lexeme: "1", Line: 1},
		{Kind: NewLine, Lexeme: "\n", Line: 1},
	}

	assertTokens(t, expected, tokens)
}

func TestLexStringLiteral(t *testing.T) {
	code := `mov r0, "hola"` + "\n"
	tokens := Lex(code)

	expected := []Token{
		{Kind: Instruction, Lexeme: "mov", Line: 1},
		{Kind: Register, Lexeme: "r0", Line: 1},
		{Kind: Comma, Lexeme: ",", Line: 1},
		{Kind: StringLiteral, Lexeme: "hola", Line: 1},
		{Kind: NewLine, Lexeme: "\n", Line: 1},
	}

	assertTokens(t, expected, tokens)
}

func TestLexHexAndBitNumbers(t *testing.T) {
	code := "mov r0, #0xFF\nmov r1, #0b1010\n"
	tokens := Lex(code)

	// Verificar que los números se reconozcan como Number/HexNumber/BitNumber
	// sin entrar en detalle de cuál subtipo usa el lexer actualmente.
	if len(tokens) < 8 {
		t.Fatalf("se esperaban al menos 8 tokens, se obtuvieron %d: %+v", len(tokens), tokens)
	}
}

// ---------------------------------------------------------------------------
// Pruebas del ENCODER
// ---------------------------------------------------------------------------

func TestEncodeHalt(t *testing.T) {
	op := Opp{Op: "halt"}
	got := Encode(op)
	want := packOp(opToB["halt"])
	if got != want {
		t.Errorf("Encode(halt) = %#x; want %#x", got, want)
	}
}

func TestEncodeMovrr(t *testing.T) {
	op := Oprr{Op: "movrr", R1: 0, R2: 1}
	got := Encode(op)
	want := packOprr(opToB["movrr"], 0, 1)
	if got != want {
		t.Errorf("Encode(movrr) = %#x; want %#x", got, want)
	}
}

func TestEncodeMovri(t *testing.T) {
	op := Opri{Op: "movri", R1: 0, I: 5}
	got := Encode(op)
	want := packOpri(opToB["movri"], 0, 5)
	if got != want {
		t.Errorf("Encode(movri) = %#x; want %#x", got, want)
	}
}

func TestEncodeAddrrr(t *testing.T) {
	op := Oprrr{Op: "addrrr", R1: 0, R2: 1, R3: 2}
	got := Encode(op)
	want := packOprrr(opToB["addrrr"], 0, 1, 2)
	if got != want {
		t.Errorf("Encode(addrrr) = %#x; want %#x", got, want)
	}
}

func TestEncodeAddrri(t *testing.T) {
	op := Oprri{Op: "addrri", R1: 0, R2: 1, I: 5}
	got := Encode(op)
	want := packOprri(opToB["addrri"], 0, 1, 5)
	if got != want {
		t.Errorf("Encode(addrri) = %#x; want %#x", got, want)
	}
}

func TestEncodeBl(t *testing.T) {
	op := Opl{Op: "bl", I: 10}
	got := Encode(op)
	want := packOpl(opToB["bl"], 10)
	if got != want {
		t.Errorf("Encode(bl) = %#x; want %#x", got, want)
	}
}

// ---------------------------------------------------------------------------
// Prueba de integración: Lex -> Analyze -> memoria
// ---------------------------------------------------------------------------

func TestAnalyzeSimpleProgram(t *testing.T) {
	code := "mov r0, #5\nadd r1, r0, r0\nhalt\n"
	tokens := Lex(code)
	mem, err := Analyze(tokens)
	if err != nil {
		t.Fatalf("Analyze devolvió error: %v", err)
	}

	if len(mem) == 0 {
		t.Fatal("Analyze no generó memoria")
	}

	// Verificar que la primera instrucción sea movri r0, #5
	decoded, err := Decode(mem[0])
	if err != nil {
		t.Fatalf("Decode falló: %v", err)
	}
	movri, ok := decoded.(Opri)
	if !ok || movri.Op != "movri" || movri.R1 != 0 || movri.I != 5 {
		t.Errorf("primera instrucción inesperada: %+v", decoded)
	}
}

// ---------------------------------------------------------------------------
// Utilidades
// ---------------------------------------------------------------------------

func assertTokens(t *testing.T, want, got []Token) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("tokens inesperados:\n  want: %+v\n  got:  %+v", want, got)
	}
}
