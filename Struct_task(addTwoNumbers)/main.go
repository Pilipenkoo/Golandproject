package main

import "fmt"

// Определение структуры узла
type ListNode struct {
	Val  int
	Next *ListNode
}

// Основная функция, которая складывает два числа, представленных в виде связных списков
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{} // Создание фиктивного узла для удобства
	current := dummyHead     // Указатель на текущий узел в результирующем списке
	carry := 0               // Переменная для хранения переноса

	for l1 != nil || l2 != nil {
		var x, y int
		if l1 != nil {
			x = l1.Val // Значение текущего узла первого списка
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val // Значение текущего узла второго списка
			l2 = l2.Next
		}

		sum := carry + x + y                    // Сумма значений двух узлов и переноса
		carry = sum / 10                        // Вычисление нового переноса
		current.Next = &ListNode{Val: sum % 10} // Создание нового узла с результатом
		current = current.Next                  // Переход к следующему узлу
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry} // Если после последней итерации есть перенос, добавляем новый узел
	}

	return dummyHead.Next // Возвращаем следующий узел после фиктивного
}

func main() {
	// Пример использования функции
	l1 := &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
}
