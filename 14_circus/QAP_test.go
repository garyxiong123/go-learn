package circus

import (
	"fmt"
	"github.com/arnaucube/go-snark/circuitcompiler"
	"github.com/arnaucube/go-snark/fields"
	"github.com/arnaucube/go-snark/r1csqap"
	"math/big"
	"strings"
	"testing"
)

func Test_QAP(t *testing.T) {
	str := `
       func main(private a, private b,public c):
		 s1 = b + 2
		 c = a * s1
    `

	fmt.Printf("code of the circuit: %s\n", str)
	parser := circuitcompiler.NewParser(strings.NewReader(str))
	circuit, _ := parser.Parse()

	fmt.Println("\ncode to R1CS")
	a, b, c := circuit.GenerateR1CS()
	fmt.Printf("a: %s\n", a)
	fmt.Printf("b: %s\n", b)
	fmt.Printf("c: %d\n\n", c)

	fmt.Println("\nR1CS to QAP")
	r, _ := new(big.Int).SetString("11", 10)
	pf := r1csqap.NewPolynomialField(fields.NewFq(r))
	alphas, betas, gammas, _ := pf.R1CSToQAP(a, b, c)

	fmt.Printf("As: %d\n\n", alphas)
	fmt.Printf("Bs: %d\n\n", betas)
	fmt.Printf("Cs: %d\n\n", gammas)

	// witness
	privateVal := []*big.Int{big.NewInt(int64(3)), big.NewInt(int64(2))}
	publicVal := []*big.Int{big.NewInt(int64(12))}
	w, _ := circuit.CalculateWitness(privateVal, publicVal)
	fmt.Println("\nWitness: ", w)

	ax, bx, cx, px := pf.CombinePolynomials(w, alphas, betas, gammas)
	fmt.Printf("\nax: %d\n", ax)
	fmt.Printf("\nbx: %d\n", bx)
	fmt.Printf("\ncx: %d\n", cx)
	fmt.Printf("\npx: %d\n", px)
}
