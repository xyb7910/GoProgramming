package main

import "fmt"

func main() {
	var str1, str2, str3 string
	fmt.Scanf("%s", &str1)
	fmt.Scanf("%s", &str2)
	fmt.Scanf("%s", &str3)

	if str1 == "vertebrado" {
		if str2 == "ave" {
			if str3 == "carnivoro" {
				fmt.Println("aguia")
			} else if str3 == "onivoro" {
				fmt.Println("pomba")
			}
		} else if str2 == "mamifero" {
			if str3 == "onivoro" {
				fmt.Println("homem")
			} else if str3 == "herbivoro" {
				fmt.Println("vaca")
			}
		}
	} else if str1 == "invertebrado" {
		if str2 == "inseto" {
			if str3 == "hematofago" {
				fmt.Println("pulga")
			} else if str3 == "herbivoro" {
				fmt.Println("lagarta")
			}
		} else if str2 == "anelideo" {
			if str3 == "hematofago" {
				fmt.Println("sanguessuga")
			} else if str3 == "onivoro" {
				fmt.Println("minhoca")
			}
		}
	}
}
