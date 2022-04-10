// Программа простейшего echo сервера целых чисел с обменом сообщениями через командную строку
// Егор Логинов (GO-11) по заданию SkillFactory в модуле 25 - демонстрация работы с Git

package main
 
importgi (
    "fmt""log"
)
 
func main() {
    n :=0
    fmt.Print("Введите целое число: ")
    _, err := fmt.Scan(&n)
    if err !=nil {
        log.Fatal(err)
    }
    fmt.Printf("Вы ввели число: %d\n", n)
}