package main
import "fmt"

func main() {
	var a int = 10	/* actual variable declaration */
	var ip *int		/* pointer variable declaration */

	ip = &a			/* store address of a in pointer variable */

	fmt.Println("Alamat dari variable a : ",&a)
	fmt.Println("Alamat yang disimpan di dalam variable ip : ", ip)
	fmt.Println("Nilai dari variable ip : ",*ip)

	// nilai a berubah
	a = 30
	fmt.Println("Nilai dari variable ip setelah value berubah : ",*ip)



}