package main
import "fmt"

// 課題1.
// 2つのfloat64型を引数に対して 「引き算・掛け算・割り算」の計算をし,
// その値を戻り値として返す関数をそれぞれ作ってください。

func tasizan(x float64,y float64) float64 {
	return x + y
}
func hikizan(x float64,y float64) float64 {
	return x - y
}
func kakezan(x float64,y float64) float64 {
	return x * y
}
func warizan(x float64,y float64) float64 {
	return x / y
}

// 課題2.
// 課題1を発展させ, 第2, 3引数にfloat64型の値を入れ,
// その値に対して第1引数に指定した符号(例: +,-,*,/)に合わせた計算をし,
// 結果を戻り値として返す関数を作ってください。

func calc(kigou string,x float64,y float64) float64 {
	if kigou == "+" {
		return tasizan(x,y)
	} else if kigou == "-" {
		return hikizan(x,y)
	} else if kigou == "*" {
		return kakezan(x,y)
	} else if kigou == "/" {
		return warizan(x,y)
	} else{
		return 0
	}
}

	/*
	課題3.
	第2回の資料において, Structの2番に書いてあるコードを
	発展させ以下の要件を満たしたメソッドをそれぞれ作成してください。

	Turtle構造体のaの値を, 引数をもとに足し算するメソッド
	Turtle構造体のnameの値を, 引数をもとに変更するメソッド
	Turtle構造体のnameの値を利用して, 引数をもとに
	「Turtleのname : こんにちは！」のようなメッセージを出力するメソッド
	*/

	type Turtle struct {
		name string
		x float64
		y float64
		a float64
	}

	func (t *Turtle) plus(i float64){
		t.a += i
	}

	func (t *Turtle) nameChange(s string){
		t.name = s
	}

	func (t *Turtle) hello(){
		fmt.Println(t.name + " : こんにちは!")
	}

func main() {
	fmt.Println("足し算")
	fmt.Println(calc("+",10,20))
	fmt.Println("-------------")
	fmt.Println("引き算")
	fmt.Println(calc("-",50,10))
	fmt.Println("-------------")
	fmt.Println("かけ算")
	fmt.Println(calc("*",2,10))
	fmt.Println("-------------")
	fmt.Println("わり算")
	fmt.Println(calc("/",20,2))
	fmt.Println("-------------")

	var t1 Turtle = Turtle{"師匠",1000,500,100}
	var t2 Turtle = Turtle{"弟子",100,200,300}

	fmt.Println(t1)
	fmt.Println(t2)

	t1.plus(1)
	t1.nameChange("sample")
	t1.hello()
	fmt.Println(t1)

}
