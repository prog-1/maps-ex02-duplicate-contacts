# Phone Book (Exercise)

In this exercise, you need to write a program that creates a phone book storing names
and phone numbers of your contacts.

Write a program that asks a user to enter names and phone number of your contacts as
a keyboard input, stores them in a map, and prints the map.

Write a function `func createPhoneBook(names []string, numbers []string) map[string]string`,
where `names` is a slice of names entered by the user and `numbers` is a slice of phone numbers entered by the user.

A name may be entered multipled times. In this case, the map must store the first phone number occurrence.

Create tests for the `createPhoneBook` function.

### Example

```
How many contacts do you want to enter?
5
Enter name #1: Alina
Enter phone number #1: +37126017505
Enter name #2: Deniss B
Enter phone number #2: +37127785804
Enter name #3: Antons
Enter phone number #3: +37123622588
Enter name #4: Alina
Enter phone number #4: +37126505719
Enter name #5: Antons
Enter phone number #5: +37128852154
map[Alina:+37126017505 Antons:+37123622588 Deniss B:+37127785804]
```

The input for the `createPhoneBook` function in this example must be

```
names = []string{
  "Alina",
  "Deniss B",
  "Antons",
  "Alina",
  "Antons",
}
numbers = []string{
  "+37126017505",
  "+37127785804",
  "+37123622588",
  "+37126505719",
  "+37128852154",
}
```

The returned map must be

```
map[string]string{
  "Alina": "+37126017505",
  "Antons": "+37123622588",
  "Deniss B": "+37127785804",
}
```

In this example, "Alina" and "Antons" occur twice. The phone numbers of the first occurrences are stored in the map.
