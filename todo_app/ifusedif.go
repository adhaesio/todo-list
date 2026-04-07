package todo_app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

func RunCommands() {
	scanner := bufio.NewScanner(os.Stdin)
	todoSlice := []string{}

	for {
		fmt.Print("Введите вашу команду: ")

		if ok := scanner.Scan(); !ok {
			fmt.Println("ERROR")
			return
		}
		text := scanner.Text()
		fields := strings.Fields(text)
		if len(fields) == 0 {
			fmt.Println("Вы ничего не ввели!")
			continue
		}

		cmd := fields[0]
		if cmd == "выйти" || cmd == "Выйти" {
			pp.Println("До скорого!")
			return
		}

		if cmd == "добавить" || cmd == "Добавить" {
			if len(fields) < 2 {
				fmt.Println("Что именно ты хотел добавить?")
				fmt.Println("")
				for {
					fmt.Print("Введите задачу: ")
					if !scanner.Scan() {
						return
					}

					addtext := scanner.Text()
					todoSlice = append(todoSlice, addtext)
					fmt.Println("Задача добавлена: ", addtext)
					fmt.Println("")
					break

				}
				continue

			} else {

				task := strings.TrimSpace(strings.Join(fields[1:], " "))

				todoSlice = append(todoSlice, task)

				fmt.Println("Задача добавлена: ", task)
				fmt.Println("")
				continue

			}
		}

		if cmd == "удалить" || cmd == "Удалить" {
			if len(todoSlice) == 0 {
				pp.Println("Список задач пуст")
				continue
			}
			var deltext int
			var err error
			if len(fields) < 2 {
				fmt.Println("")
				fmt.Println("Вы хотите удалить одну из задач?")
				fmt.Println("Укажите номер")

				if !scanner.Scan() {
					return
				}
				deltext, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
			} else {
				deltext, err = strconv.Atoi(fields[1])
			}

			if err != nil {
				fmt.Println("Номер должен быть числом")
				fmt.Println("")
				continue
			}

			if deltext < 1 || deltext > len(todoSlice) {
				fmt.Println("Задачи с таким номером не существует")
				fmt.Println("")
				continue
			}
			removed := todoSlice[deltext-1]
			todoSlice = append(todoSlice[:deltext-1], todoSlice[deltext:]...)
			fmt.Println("Задача удалена: ", removed)
			continue
		}

		if cmd == "help" {
			fmt.Println("Команда: help")
			fmt.Println("-- это команда выводит список доступных команд")
			fmt.Println("Команда: добавить {что нужно добавить}")

			fmt.Println("-- это команда позволяет добавлять что-то")
			fmt.Println("")
			fmt.Println("Команда: удалить {что нужно удалить}")
			fmt.Println("-- это команда позволяет удалять что-то")
			fmt.Println("")
		}
		if cmd == "показать" {
			if len(todoSlice) == 0 {
				fmt.Println("Список задач пуст")
				fmt.Println("")
				continue
			}
			fmt.Println("Ваши задачи:")
			for i, task := range todoSlice {
				pp.Printf("%d. %s\n", i+1, task)
			}
			fmt.Println("")

		} else {
			fmt.Println("Вы ввели неизвестную команду")
		}
	}

}
