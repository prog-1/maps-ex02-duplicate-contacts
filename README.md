# Duplicate Contacts 1 (Exercise)

In this exercise, you need to write two programs that manage contacts and their phone numbers in your phone. 

## Part 1: Phone Book

Write a program that asks a user to enter names and phone number of your contacts.
Create a map `map[string]string` that maps name to the corresponding phone numbers.
Print the map.

> **NOTE** Multiple names may have the same phone number.

### Example

```
How many contacts do you want to enter?
9
Enter name #1: Alina
Enter phone number #1: +37126017505
Enter name #2: Deniss
Enter phone number #2: +37127785804
Enter name #3: Andrejs
Enter phone number #3: +37124540266
Enter name #4: Polina
Enter phone number #4: +37121256444
Enter name #5: Deniss B
Enter phone number #5: +37127785804
Enter name #6: Polina J
Enter phone number #6: +37121256444
Enter name #7: Antons
Enter phone number #7: +37123622588
Enter name #8: Valerija
Enter phone number #8: +37127754705
Enter name #9: Deniss M
Enter phone number #9: +37123228388
map[Alina:+37126017505 Andrejs:+37124540266 Antons:+37123622588 Deniss:+37127785804 Deniss B:+37127785804 Deniss M:+37123228388 Polina:+37121256444 Polina J:+37121256444 Valerija:+37127754705]
```

Note that the phone book contains duplicate contacts `Deniss` and `Deniss B` (`Deniss M` is not a duplicate, because it corresponds to a different phone number. `Polina` and `Polina J` are duplicates as well.

## Part 2: Duplicate Contacts

Write a function `removeDuplicateContacts(contacts map[string]string) map[string]string`
that merges contacts with the same number in the phone book.

### Example

```
phoneBook := map[string]string{
		"Alina":    "+37126017505",
		"Andrejs":  "+37124540266",
		"Antons":   "+37123622588",
		"Deniss":   "+37127785804",
		"Deniss B": "+37127785804",
		"Deniss M": "+37123228388",
		"Polina":   "+37121256444",
		"Polina J": "+37121256444",
		"Valerija": "+37127754705",
	}
removeDuplicateContacts(phoneBook)
```

must return

```
phoneBook := map[string]string{
		"Alina":    "+37126017505",
		"Andrejs":  "+37124540266",
		"Antons":   "+37123622588",
		"Deniss":   "+37127785804",
		"Deniss M": "+37123228388",
		"Polina J": "+37121256444",
		"Valerija": "+37127754705",
	}
```

`Deniss B` may occur instead of `Denis`; `Polina` may occur instead of `Polina J`.
